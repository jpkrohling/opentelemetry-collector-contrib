module github.com/open-telemetry/opentelemetry-collector-contrib/examples/demo/server

go 1.18

require (
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.37.0
	go.opentelemetry.io/otel v1.11.2
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric v0.31.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v0.31.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.11.1
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.1
	go.opentelemetry.io/otel/metric v0.34.0
	go.opentelemetry.io/otel/sdk v1.11.1
	go.opentelemetry.io/otel/sdk/metric v0.31.0
	go.opentelemetry.io/otel/trace v1.11.2
	google.golang.org/grpc v1.51.0
)

require (
	github.com/cenkalti/backoff/v4 v4.1.3 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.14.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.11.1 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20221114212237-e4508ebdbee1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

retract v0.65.0
