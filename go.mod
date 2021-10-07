module knative.dev/client

go 1.15

require (
	github.com/google/go-cmp v0.5.6
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	golang.org/x/term v0.0.0-20210503060354-a79de5458b56
	gotest.tools/v3 v3.0.3
	k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/cli-runtime v0.20.7
	k8s.io/client-go v0.20.7
	k8s.io/code-generator v0.20.7
	knative.dev/eventing v0.24.0
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/kn-plugin-func v0.18.0
	knative.dev/kn-plugin-source-kafka v0.24.0
	knative.dev/networking v0.0.0-20210622182128-53f45d6d2cfa
	knative.dev/pkg v0.0.0-20210622173328-dd0db4b05c80
	knative.dev/serving v0.24.0
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3

replace knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v0.18.1-0.20211007090909-a33974383b9e
