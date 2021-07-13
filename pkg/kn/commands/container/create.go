// Copyright © 2021 The Knative Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"knative.dev/client/pkg/kn/commands"
	knflags "knative.dev/client/pkg/kn/flags"
	sigyaml "sigs.k8s.io/yaml"
)

// NewContainerCreateCommand to create event channels
func NewContainerCreateCommand(p *commands.KnParams) *cobra.Command {
	var podSpecFlags knflags.PodSpecFlags
	machineReadablePrintFlags := genericclioptions.NewPrintFlags("")

	cmd := &cobra.Command{
		Use:   "create NAME",
		Short: "Create a container",
		Example: `
  The 'container create' represents utility command that prints YAML container spec to standard output. It's useful for 
  multi-container use cases to create definition with help of standard 'kn' option flags. The command can be chained through
  Unix pipes to create multiple containers at once.

  # Create a container 'sidecart' from image 'docker.io/example/sidecart' 
  kn container create sidecart --image docker.io/example/sidecart

  # Create command can be chained by standard Unix pipe symbol '|'
  kn container create sidecart --image docker.io/example/sidecart | \
  kn container create second --image docker.io/example/sidecart:second`,
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			//TODO: require name for every container?
			name := ""
			if len(args) == 1 {
				name = args[0]
			}
			// Detect pipe input from previous container command
			if podSpecFlags.Containers == "" && detectPipeInput(os.Stdin) {
				podSpecFlags.Containers = "-"
			}
			podSpec := &corev1.PodSpec{}
			if err = podSpecFlags.ResolvePodSpec(podSpec, cmd.Flags(), os.Args); err != nil {
				return err
			}
			// Add container's name to current one
			if name != "" {
				podSpec.Containers[0].Name = name
			}
			b, err := sigyaml.Marshal(podSpec)
			if err != nil {
				return err
			}
			fmt.Fprintf(cmd.OutOrStdout(), "%s", b)
			return nil
		},
	}
	machineReadablePrintFlags.AddFlags(cmd)
	podSpecFlags.AddFlags(cmd.Flags())
	// Volume is not part of ContainerSpec
	cmd.Flag("volume").Hidden = true

	return cmd
}

func detectPipeInput(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice == 0
}
