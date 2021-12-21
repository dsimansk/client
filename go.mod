module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery v0.22.3
	k8s.io/cli-runtime v0.21.4
	k8s.io/client-go v0.21.4
	k8s.io/code-generator v0.21.4
	knative.dev/eventing v0.26.1
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/kn-plugin-event v0.26.0
	knative.dev/kn-plugin-func v0.20.0
	knative.dev/kn-plugin-source-kafka v0.26.0
	knative.dev/networking v0.0.0-20210916065741-5e884aff221e
	knative.dev/pkg v0.0.0-20210919202233-5ae482141474
	knative.dev/serving v0.26.0
	sigs.k8s.io/yaml v1.3.0
)

replace (
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.4
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.26.2-0.20211209202740-89c860ca5062
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v0.20.0
)
