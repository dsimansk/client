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
	"strings"

	"k8s.io/apimachinery/pkg/util/yaml"
)

var ctxManager *ContextDataManager

type ContextDataManager struct {
	ContextData ContextData
}

func NewContextManager() (*ContextDataManager, error) {
	if ctxManager == nil {
		var err error
		//file, err := os.Open("~/.config/kn/context-cache.json")
		//if err != nil {
		//	return nil, err
		//}

		mockData := `
{  
  "service":"hello",
  "namespace":"default"
}
`
		mockReader := strings.NewReader(mockData)
		contextData := map[string]string{}
		decoder := yaml.NewYAMLOrJSONDecoder(mockReader, 512)
		err = decoder.Decode(&contextData)
		if err != nil {
			return nil, err
		}
		ctxManager = &ContextDataManager{
			ContextData: contextData,
		}
	}
	return ctxManager, nil
}

func (c *ContextDataManager) Find(key string) string {
	return c.ContextData[key]
}
