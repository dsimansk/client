## Crafting a new release

### Create a release branch for openshift/knative-client repo:

* Check that a remote reference to openshift and upstream exists
```bash
$ git remote -v | grep -e 'openshift\|upstream'
openshift	git@github.com:openshift/knative-client.git (fetch)
openshift	git@github.com:openshift/knative-client.git (push)
upstream	https://github.com/knative/client.git (fetch)
upstream	https://github.com/knative/client.git (push)
```

* Create a new release branch which points to upstream release branch + OpenShift specific files:
```bash
# Create a new release branch. Parameters are the upstream release tag
# and the name of the branch to create
# Usage: ./create-release-branch.sh <upstream-tag> <downstream-release-branch>
# <upstream-tag>: The tag referring the upstream release
# <downstream-release-branch>: Name of the release branch to create
$ ./create-release-branch.sh v0.14.0 release-v0.14.0

# Push release branch to openshift/knative-client repo
$ git push openshift release-v0.14.0

# Note: Make sure the golang version being used for CI operator build image
# and SPEC file, the paths of the mentioned files are
# openshift/ci-operator/build-image/Dockerfile
# openshift-serverless-clients.spec
```

### Create a ci-operator configuration, prow job configurations and image mirroring config:

* Create a fork and clone of https://github.com/openshift/release
* Create a new ci-operator configuration:
```bash
# Jump into the knative client config directory in the openshift/release
$ cd ci-operator/config/openshift/knative-client

# Copy over the last release's config to a release specific config with
# the name of the yaml file ends with the new release branch name (e.g. release-v0.14.0)
$ cp openshift-knative-client-release-v0.13.0.yaml openshift-knative-client-release-v0.14.0.yaml

# Adapt the configuration for the kn new image name
# - Change .promotion.name to a release specific name (knative-v0.14.0)
# - Change .binary_build_commands to a release tag as below:
   TAG=v0.14.0 make install
   TAG=v0.14.0 make build-cross
$ vi openshift-knative-client-release-v0.14.0.yaml
```

* Create prow job configurations:
```bash
# Jump to top-level repo directory
$ cd ../../../../

# Call Prow job generators using 'make jobs' (you need a local Docker daemon installed)
# This will generate necessary presubmit and postsubmit prow jobs config YAML files
# ref: [doc](https://docs.google.com/document/d/1SQ_qlkcplqhe8h6ONXdgBr7YUVbs4oRSj4ISl3gpLW4/edit#heading=h.8w7nj9363nsd)
$ make jobs

# Update generated metadata `zz_generated_metadata`
$ make ci-operator-config
```

* Create image mirroring config:
```bash
# Add the image mirroring settings (create an empty file if not present)
$ vi core-services/image-mirroring/knative/mapping_knative_v0_14_quay

# Add following lines for the kn image like below
registry.ci.openshift.org/openshift/knative-v0.14.0:kn quay.io/openshift-knative/kn:v0.14.0
registry.ci.openshift.org/openshift/knative-v0.14.0:kn-cli-artifacts quay.io/openshift-knative/kn-cli-artifacts:v0.14.0
```

### Create a PR against openshift/release repo for CI setup of release branch using configs generated above:
```bash
# Verify the changes
$ git status
On branch master
Your branch is ahead of 'origin/master' by 180 commits.
  (use "git push" to publish your local commits)

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   core-services/image-mirroring/knative/mapping_knative_v0_12_quay

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	ci-operator/config/openshift/knative-client/openshift-knative-client-release-v0.14.0.yaml
	ci-operator/jobs/openshift/knative-client/openshift-knative-client-release-v0.14.0-postsubmits.yaml
	ci-operator/jobs/openshift/knative-client/openshift-knative-client-release-v0.14.0-presubmits.yaml

# Add & Commit all and push to your repo
$ git add .
$ git commit -a -m "knative-client release v0.14.0 setup"
$ git push

# Create pull request on https://github.com/openshift/release with your changes
# Once PR against openshift/release repo is merged, the CI is setup for release-branch
```

### Update RPM SPEC file in release branch with correct version:
```bash
# Update RPM SPEC file, update the version and add changelog entry
$ vi openshift-serverless-clients.spec

# Verify the changes and raise a PR against release branch
$ git status
$ git add .
$ git commit -m "Update SPEC file for release v0.14.0"
```
Note: After CI is setup for release branch, we'll need to raise at least one PR against
target release branch, have CI run and merge of PR, this ensures image gets mirrored at quay as postsubmit job.

* For further changes which are specific to OpenShift, raise PR against release branch.

### Update RPM SPEC file in main branch with latest version:
Once updated SPEC file is merged into release branch it should be synced with `main` branch file to ensure that log will be stacked correctly with the future releases.

```bash
# Alternatively you can cherry-pick spec file commit from release branch to main
git checkout openshift/main
git checkout openshift/release-v0.14.0 openshift-serverless-clients.spec

git add . 
git commit -m "Update SPEC file for latest release"
git push 
```
Create a new pull request to update `main` branch.

### Plugins

Currently, we ship 3 plugin within productized `kn` binary. It's achieved by mechanism called "plugin inlining", for futher reference see upstream's docs.
In a nutshell it's plugin defined as a dependency by empty import in `plugin_register.go` [file](https://github.com/openshift/knative-client/blob/release-v1.1.0/pkg/kn/root/plugin_register.go) and then vendored through standard Go modules.

Midstream plugins:
* kn-plugin-event
* kn-plugin-func

Upstream plugins:
* kn-plugin-source-kafka


Similar to other midstream repositories, a midstream plugin is the potentially modified source of its upstream counterpart.
An upstream plugin is taken directly from Knative release without any further modification.

There's a tool called [kn builder](https://github.com/knative/client/tree/main/tools/knb) (`knb`) to assist adding and generating necessary sources for plugins.

The following steps should be performed for each plugin.

* Create a `knb` config file in `./openshift/release/kn.yaml`
```yaml
plugins:
  - name: kn-plugin-source-kafka
    module: knative.dev/kn-plugin-source-kafka
    pluginImportPath: knative.dev/kn-plugin-source-kafka/plugin
    version: v0.28.0
  - name: kn-plugin-func
    module: knative.dev/kn-plugin-func
    pluginImportPath: knative.dev/kn-plugin-func
    version: v0.23.1
  - name: kn-plugin-event
    module: knative.dev/kn-plugin-event
    pluginImportPath: knative.dev/kn-plugin-event/pkg/plugin
    version: v0.28.0
```
NOTE: knb is lacking midstream replace support. That's solved in further steps.

* Use `knb` to populate registration file and go modules
```console
$ knb plugin distro -c openshift/release/kn.yaml
```

* Go modules `require` 
```go
// All the plugins should be already added, 
// but modules versions are used as display version for func and event plugin.
// Event plugin uses corresponding upstream's 1.x scheme
knative.dev/kn-plugin-event v1.1.1
// Func plugin uses corresponding upstream's 0.x.y
knative.dev/kn-plugin-func v0.23.1
// No modification needed
knative.dev/kn-plugin-source-kafka v0.28.0
```

* Go modules `replace`
```go
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
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event release-1.1
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func openshift-v0.23.1
)
```
Depending on the release, there're usually several other replacements to align with correct `k8s.io` version or to solve conflicts. Mileage may hugely vary here. As a rule of thumb less is definitely better. Additional comments with explanation are greatly appreciated.

* Update Go modules (repeat the previous and this step until success :))
```bash
$ ./hack/update-deps.sh
```

* The final look of updated `go.mod` replace directive 
```go
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
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.27.1-0.20220412125940-01b7cf265a69
	knative.dev/kn-plugin-func => github.com/openshift-knative/kn-plugin-func v1.1.3-0.20220413141425-a3b388f3a2d8
)
```

* Create a new PR with added plugins


### Once the changes to release branch is finalized, and we are ready for QA, create tag and push:
```bash
$ git tag openshift-v0.14.0
$ git push openshift openshift-v0.14.0
```

Note: Notify any changes required for this release, for e.g.: new commands, commands output update, etc. to docs team.
