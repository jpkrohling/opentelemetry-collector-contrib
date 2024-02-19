module github.com/open-telemetry/opentelemetry-collector-contrib/internal/splunk

go 1.21

require (
	github.com/stretchr/testify v1.8.4
	go.opentelemetry.io/collector/consumer v0.94.2-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/exporter v0.94.2-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/pdata v1.1.1-0.20240216233521-0ab8f4443937
	go.opentelemetry.io/collector/semconv v0.94.2-0.20240216233521-0ab8f4443937
	go.uber.org/goleak v1.3.0
	go.uber.org/multierr v1.11.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.0.0-alpha.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.1.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20231216201459-8508981c8b6c // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/component v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/config/configretry v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/confmap v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/collector/extension v0.94.2-0.20240216233521-0ab8f4443937 // indirect
	go.opentelemetry.io/otel v1.23.1 // indirect
	go.opentelemetry.io/otel/metric v1.23.1 // indirect
	go.opentelemetry.io/otel/trace v1.23.1 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
	google.golang.org/grpc v1.61.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
