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
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/cli-runtime v0.19.7
	k8s.io/client-go v0.19.7
	k8s.io/code-generator v0.19.7
	knative.dev/eventing v0.23.3
	knative.dev/hack v0.0.0-20210602212444-509255f29a24
	knative.dev/kn-plugin-func v0.17.0
	knative.dev/kn-plugin-source-kafka v0.23.0
	knative.dev/networking v0.0.0-20210608114541-4b1712c029b7
	knative.dev/pkg v0.0.0-20210510175900-4564797bf3b7
	knative.dev/serving v0.23.1
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/go-openapi/spec => github.com/go-openapi/spec v0.19.3

//Using main@9db1a3d902016d59e60b732de43bdf4be198334f for the release
replace knative.dev/kn-plugin-func => knative.dev/kn-plugin-func v0.16.1-0.20210803132815-9db1a3d90201
