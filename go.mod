module knative.dev/client

go 1.16

require (
	github.com/google/go-cmp v0.5.7
	github.com/maximilien/kn-source-pkg v0.6.3
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	gotest.tools/v3 v3.1.0
	k8s.io/api v0.23.4
	k8s.io/apiextensions-apiserver v0.22.5
	k8s.io/apimachinery v0.23.4
	k8s.io/cli-runtime v0.21.4
	k8s.io/client-go v1.5.2
	k8s.io/code-generator v0.22.5
	knative.dev/eventing v0.28.4
	knative.dev/hack v0.0.0-20220128200847-51a42b2eb63e
	knative.dev/kn-plugin-event v1.1.1
	knative.dev/kn-plugin-func v0.23.1
	knative.dev/kn-plugin-source-kafka v0.28.0
	knative.dev/networking v0.0.0-20220120045035-ec849677a316
	knative.dev/pkg v0.0.0-20220222214539-0b8a9403de7e
	knative.dev/reconciler-test v0.0.0-20211207070557-0d138a88867b
	knative.dev/serving v0.28.4
	sigs.k8s.io/yaml v1.3.0
)

replace (
	// Required by kn-plugin-func to use newer Docker version
	github.com/openshift/source-to-image => github.com/boson-project/source-to-image v1.3.2

	// Aligned k8s.io version with Knative release
	k8s.io/api => k8s.io/api v0.21.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.21.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.21.4
	k8s.io/client-go => k8s.io/client-go v0.21.4
	k8s.io/code-generator => k8s.io/code-generator v0.21.4

	// Inlined plugins
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.27.1-0.20220331134246-2c22e325c977
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v1.1.3-0.20220413141425-a3b388f3a2d8
)
