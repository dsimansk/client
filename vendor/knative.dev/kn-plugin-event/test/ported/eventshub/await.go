package eventshub

import (
	"context"

	"knative.dev/kn-plugin-event/test/ported/k8s"
	"knative.dev/reconciler-test/pkg/feature"
)

// Deprecated: not required after 1.3+.
func WaitForSinkReady(sinkName string) feature.StepFn {
	return func(ctx context.Context, t feature.T) {
		k8s.WaitForServiceReadyOrFail(ctx, t, sinkName, k8s.ReadinessEndpoint{
			Path: "/",
			Head: true,
		})
	}
}
