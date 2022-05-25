module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.22.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/cli-runtime v0.22.5
	k8s.io/client-go v0.22.5
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.29.0
	knative.dev/hack v0.0.0-20220118141833-9b2ed8471e30
	knative.dev/kn-plugin-source-kafka v0.29.0
	knative.dev/networking v0.0.0-20220120043934-ec785540a732
	knative.dev/pkg v0.0.0-20220118160532-77555ea48cd4
	knative.dev/serving v0.29.0
	sigs.k8s.io/yaml v1.3.0
)
