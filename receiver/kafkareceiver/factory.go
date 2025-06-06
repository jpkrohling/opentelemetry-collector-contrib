// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkareceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/kafka/configkafka"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver/internal/metadata"
)

const (
	defaultTracesTopic  = "otlp_spans"
	defaultMetricsTopic = "otlp_metrics"
	defaultLogsTopic    = "otlp_logs"
	defaultEncoding     = "otlp_proto"
)

// NewFactory creates Kafka receiver factory.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		metadata.Type,
		createDefaultConfig,
		receiver.WithTraces(createTracesReceiver, metadata.TracesStability),
		receiver.WithMetrics(createMetricsReceiver, metadata.MetricsStability),
		receiver.WithLogs(createLogsReceiver, metadata.LogsStability),
	)
}

func createDefaultConfig() component.Config {
	return &Config{
		ClientConfig:   configkafka.NewDefaultClientConfig(),
		ConsumerConfig: configkafka.NewDefaultConsumerConfig(),
		Encoding:       defaultEncoding,
		MessageMarking: MessageMarking{
			After:   false,
			OnError: false,
		},
		HeaderExtraction: HeaderExtraction{
			ExtractHeaders: false,
		},
	}
}

func createTracesReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Traces,
) (receiver.Traces, error) {
	oCfg := *(cfg.(*Config))
	if oCfg.Topic == "" {
		oCfg.Topic = defaultTracesTopic
	}

	r, err := newTracesReceiver(oCfg, set, nextConsumer)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func createMetricsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Metrics,
) (receiver.Metrics, error) {
	oCfg := *(cfg.(*Config))
	if oCfg.Topic == "" {
		oCfg.Topic = defaultMetricsTopic
	}

	r, err := newMetricsReceiver(oCfg, set, nextConsumer)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func createLogsReceiver(
	_ context.Context,
	set receiver.Settings,
	cfg component.Config,
	nextConsumer consumer.Logs,
) (receiver.Logs, error) {
	oCfg := *(cfg.(*Config))
	if oCfg.Topic == "" {
		oCfg.Topic = defaultLogsTopic
	}

	r, err := newLogsReceiver(oCfg, set, nextConsumer)
	if err != nil {
		return nil, err
	}
	return r, nil
}
