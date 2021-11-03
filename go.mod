module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/cli-runtime v0.20.7
	k8s.io/client-go v0.20.7
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.25.2
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/kn-plugin-func v0.19.0
	knative.dev/kn-plugin-source-kafka v0.25.0
	knative.dev/networking v0.0.0-20210803181815-acdfd41c575c
	knative.dev/pkg v0.0.0-20210902173607-844a6bc45596
	knative.dev/serving v0.25.1
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3

replace knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v0.18.1-0.20211103194307-d8d721a8334a
