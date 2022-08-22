// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tenant // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter/internal/tenant"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

var _ Source = (*AttributeTenantSource)(nil)

type AttributeTenantSource struct {
	Value  string
	Logger *zap.Logger
}

func (ts *AttributeTenantSource) GetTenant(_ context.Context, logs plog.Logs) (string, error) {
	ret := ""
	for i := 0; i < logs.ResourceLogs().Len(); i++ {
		rl := logs.ResourceLogs().At(i)
		if v, found := rl.Resource().Attributes().Get(ts.Value); found {
			tenant := v.StringVal()
			if len(ret) > 0 && ret != tenant {
				ts.Logger.Info("found a different tenant in resource attribute", zap.String("tenant", ret), zap.Int("index", i))
			}
			ret = tenant
		}
	}
	return ret, nil
}
