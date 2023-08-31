// Copyright © 2023 The Knative Authors
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

package plugin

import (
	"bytes"
	"encoding/json"
	"os"
	"os/exec"
	"strings"
)

//--TYPES--
//TODO: move types into its own file

type ContextData map[string]string

type Manifest struct {
	// Path to external plugin binary. Always empty for inlined plugins.
	Path string `json:"path,omitempty"`

	// Plugin declares its own manifest to be included in Context Sharing feature
	HasManifest bool `json:"hasManifest"`

	// ProducesContextDataKeys is a list of keys for the ContextData that
	// a plugin can produce. Nil or an empty list declares that this
	// plugin is not ContextDataProducer
	ProducesContextDataKeys []string `json:"producesKeys,omitempty"`

	// ConsumesContextDataKeys is a list of keys from a ContextData that a
	// plugin is interested in to consume. Nil or an empty list declares
	// that this plugin is not a ContextDataConsumer
	ConsumesContextDataKeys []string `json:"consumesKeys,omitempty"`
}

type ContextDataConsumer interface {
	// ExecuteWithContextData executes the plugin with the given args much like
	// Execute() but with an additional argument that holds the ContextData
	ExecuteWithContextData(args []string, data ContextData) error
}

//--TYPES--

var ctxManager *ContextDataManager

type ContextDataManager struct {
	ContextData map[string]ContextData `json:"contextData"`
	Manifests   map[string]Manifest    `json:"manifests"`
}

func NewContextManager() (*ContextDataManager, error) {
	if ctxManager == nil {
		var err error
		//file, err := os.Open("~/.config/kn/context.json")
		//if err != nil {
		//	return nil, err
		//}
		//TODO: mock data to test things out
		mockData := `
{
	"contextData": {
		"default": {
			"service": "hello",
			"namespace": "default"
		}
	},
	"manifests": {
		"kn-service-log": {
			"path": "/usr/local/bin/kn-service/log",
			"hasManifest": true,
			"description": "Show logs of a service",
			"consumesKeys": ["service", "namespace"]
		}
	}
}
`
		mockReader := strings.NewReader(mockData)
		decoder := json.NewDecoder(mockReader)
		contextData := &ContextDataManager{}
		err = decoder.Decode(&contextData)
		if err != nil {
			return nil, err
		}
		ctxManager = contextData
	}
	return ctxManager, nil
}

// GetContext returns context data by key
func (c *ContextDataManager) GetContext(key string) ContextData {
	return c.ContextData[key]
}

// GetDefault returns default context data
func (c *ContextDataManager) GetDefault() ContextData {
	return c.GetContext("default")
}

// GetConsumesKeys returns array of keys consumed by plugin
func (c *ContextDataManager) GetConsumesKeys(pluginName string) []string {
	return c.Manifests[pluginName].ConsumesContextDataKeys
}

// GetProducesKeys returns array of keys produced by plugin
func (c *ContextDataManager) GetProducesKeys(pluginName string) []string {
	return c.Manifests[pluginName].ProducesContextDataKeys
}

// FetchManifests it tries to retrieve manifest from both inlined and external plugins
func (c *ContextDataManager) FetchManifests(pluginManager *Manager) error {
	plugins, err := pluginManager.ListPlugins()
	if err != nil {
		return err
	}
	for _, p := range plugins {
		// Add new plugins only
		if _, exists := c.Manifests[p.Name()]; !exists {
			manifest := &Manifest{}
			if p.Path() != "" {
				// Fetch from external plugin
				if m := fetchExternalManifest(p); m != nil {
					manifest = m
				}
			} else {
				// Fetch from internal plugin
				if pwm, ok := p.(PluginWithManifest); ok {
					manifest = pwm.GetManifest()
				}
			}
			// Add manifest to map
			c.Manifests[p.Name()] = *manifest
		}
	}
	return nil
}

// TODO: We should cautiously execute external binaries
// fetchExternalManifest returns Manifest from external plugin by exec `$plugin manifest get`
func fetchExternalManifest(p Plugin) *Manifest {
	cmd := exec.Command(p.Path(), "manifest", "get") //nolint:gosec
	stdOut := new(bytes.Buffer)
	cmd.Stdout = stdOut
	manifest := &Manifest{
		Path:        p.Path(),
		HasManifest: false,
	}
	if err := cmd.Run(); err != nil {
		//TODO: debug log
		println("No manifest cmd found")
		return manifest
	}
	d := json.NewDecoder(stdOut)
	if err := d.Decode(manifest); err != nil {
		//TODO: debug log
		println("Error reading manifest")
		return manifest
	}
	manifest.HasManifest = true
	return manifest
}

// TODO: store to file actually
// WriteCache store data back to cache file
func (c *ContextDataManager) WriteCache() error {
	println("\n====\nContext Data to be stored:")
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	return enc.Encode(c)
}
