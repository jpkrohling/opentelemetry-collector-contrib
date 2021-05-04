// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package policysamplingprocessor

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/configtest"
)

func TestLoadConfig(t *testing.T) {
	factories, err := componenttest.NopFactories()
	assert.NoError(t, err)

	factory := NewFactory()
	factories.Processors[factory.Type()] = factory

	cfg, err := configtest.LoadConfigFile(t, path.Join(".", "testdata", "config.yaml"), factories)
	require.Nil(t, err)
	require.NotNil(t, cfg)

	assert.Equal(t, cfg.Processors["policysampling"],
		&Config{
			ProcessorSettings: config.ProcessorSettings{
				TypeVal: "policysampling",
				NameVal: "policysampling",
			},
			PolicyCfgs: []PolicyCfg{
				{
					Name: "test-policy-1",
					Type: AlwaysSample,
				},
				{
					Name:                "test-policy-2",
					Type:                NumericAttribute,
					NumericAttributeCfg: NumericAttributeCfg{Key: "key1", MinValue: 50, MaxValue: 100},
				},
				{
					Name:               "test-policy-3",
					Type:               StringAttribute,
					StringAttributeCfg: StringAttributeCfg{Key: "key2", Values: []string{"value1", "value2"}},
				},
				{
					Name:            "test-policy-4",
					Type:            RateLimiting,
					RateLimitingCfg: RateLimitingCfg{SpansPerSecond: 35},
				},
			},
		})
}
