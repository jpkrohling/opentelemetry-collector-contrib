// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package wesleyexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/wesleyexporter/internal/metadata"
)

// NewFactory creates a factory for the legacy Loki exporter.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		metadata.Type,
		createDefaultConfig,
		exporter.WithTraces(createTracesExporter, metadata.TracesStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createTracesExporter(ctx context.Context, set exporter.Settings, cfg component.Config) (exporter.Traces, error) {
	return nil, nil
}
