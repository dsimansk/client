package k8s

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	kubeclient "knative.dev/pkg/client/injection/kube/client"
	"knative.dev/pkg/kmeta"
	"knative.dev/reconciler-test/pkg/environment"
	"knative.dev/reconciler-test/pkg/feature"
	"knative.dev/reconciler-test/pkg/k8s"
)

const ubi8Image = "registry.access.redhat.com/ubi8/ubi"

// ReadinessEndpoint defines of what to check for readiness.
type ReadinessEndpoint struct {
	Path string
	Head bool
}

// Deprecated: use reconciler-test/pkg/k8s.WaitForServiceReadyOrFail when available (1.3+).
func WaitForServiceReadyOrFail(ctx context.Context, t feature.T, name string, readiness ReadinessEndpoint) {
	if err := WaitForServiceReady(ctx, t, name, readiness); err != nil {
		t.Fatalf("Failed while %+v", errors.WithStack(err))
	}
}

// ErrWaitingForServiceReady if waiting for service ready failed.
var ErrWaitingForServiceReady = errors.New("waiting for service ready")

// Deprecated: use reconciler-test/pkg/k8s.WaitForServiceReady when available (1.3+).
func WaitForServiceReady(ctx context.Context, t feature.T, name string, readiness ReadinessEndpoint) error {
	env := environment.FromContext(ctx)
	ns := env.Namespace()
	kube := kubeclient.Get(ctx)
	jobs := kube.BatchV1().Jobs(ns)
	label := "readiness-check"
	jobName := feature.MakeRandomK8sName(name + "-" + label)
	sinkURI := apis.HTTP(fmt.Sprintf("%s.%s.svc", name, ns))
	sinkURI.Path = readiness.Path
	head := ""
	if readiness.Head {
		head = "--head"
	}
	curl := fmt.Sprintf("curl %s --max-time 2 "+
		"--trace-ascii %% --trace-time "+
		"--retry 6 --retry-connrefused %s", head, sinkURI)
	var one int32 = 1
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{Name: jobName, Namespace: ns},
		Spec: batchv1.JobSpec{
			Completions: &one,
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: corev1.RestartPolicyOnFailure,
					Containers: []corev1.Container{{
						Name:    jobName,
						Image:   ubi8Image,
						Command: []string{"/bin/sh"},
						Args:    []string{"-c", curl},
					}},
				},
			},
		},
	}
	created, err := jobs.Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	env.Reference(kmeta.ObjectReference(created))
	if err = k8s.WaitUntilJobDone(ctx, t, kube, ns, jobName); err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	job, err = jobs.Get(ctx, jobName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	if !k8s.IsJobSucceeded(job) {
		return dumpPodLogsInError(ctx, jobName, job)
	}

	return nil
}

func dumpPodLogsInError(ctx context.Context, jobName string, job *batchv1.Job) error {
	env := environment.FromContext(ctx)
	ns := env.Namespace()
	kube := kubeclient.Get(ctx)
	pod, err := k8s.GetJobPodByJobName(ctx, kube, ns, jobName)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	logs, err := k8s.PodLogs(ctx, pod.Name, jobName, ns)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	status, err := json.MarshalIndent(job.Status, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: %v", ErrWaitingForServiceReady, err)
	}
	return fmt.Errorf("%w: job failed, status: \n%s\n---\nlogs:\n%s",
		ErrWaitingForServiceReady, status, logs)
}
