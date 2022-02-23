module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.6
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.9.0
	golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.22.5
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/cli-runtime v0.21.4
	k8s.io/client-go v0.22.5
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.27.2
	knative.dev/hack v0.0.0-20220216040439-0456e8bf6547
	knative.dev/kn-plugin-event v0.27.1
	knative.dev/kn-plugin-func v0.21.0
	knative.dev/kn-plugin-source-kafka v0.27.0
	knative.dev/networking v0.0.0-20211101215640-8c71a2708e7d
	knative.dev/pkg v0.0.0-20220215153400-3c00bb0157b9
	knative.dev/serving v0.27.1
	sigs.k8s.io/yaml v1.3.0
)

replace (
	k8s.io/api => k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.4
	k8s.io/client-go => k8s.io/client-go v0.21.4
	k8s.io/code-generator => k8s.io/code-generator v0.21.4
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.27.1-0.20220223114256-af13ecf492aa
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v0.22.1-0.20220223003546-b1c8c38b7608
)
