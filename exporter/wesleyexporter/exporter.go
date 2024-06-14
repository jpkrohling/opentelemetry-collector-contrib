// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter"

import (
	"context"

	"go.opentelemetry.io/collector/pdata/ptrace"
)

type wesleyExporter struct {
}

func newWesleyExporter(*Config) (*wesleyExporter, error) {
	return &wesleyExporter{}, nil
}

func (we *wesleyExporter) consumeTraces(context.Context, ptrace.Traces) error {
	return nil
}
