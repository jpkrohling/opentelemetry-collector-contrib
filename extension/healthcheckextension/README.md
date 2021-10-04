# Health Check

Health Check extension enables an HTTP url that can be probed to check the
status of the OpenTelemetry Collector. This extension can be used as a
liveness and/or readiness probe on Kubernetes.

There is an optional configuration `metrics_health_check` which allows 
customers to enable metrics functional health check. This feature can monitor the
number of times that components failed send data to the destination. It only 
supports monitoring exporter failures and will support receivers and processors in 
the future.

The following settings are required:

- `endpoint` (default = 0.0.0.0:13133): Address to publish the health check status to
- `port` (default = 13133): [deprecated] What port to expose HTTP health information.
- `metrics_health_check:` (optional): Settings of metrics health check
  - `enabled` (default = false): Whether enable metrics health check or not
  - `interval` (default = "5m"): Time interval to check the number of failures
  - `exporter_failure_threshold` (default = 5): The failure number threshold to mark 
  containers as healthy.

Example:

```yaml
extensions:
  health_check:
  health_check/1:
    endpoint: "localhost:13"
    metrics_health_check:
      enabled: false
      interval: "5m"
      exporter_failure_threshold: 5
```

The full list of settings exposed for this exporter is documented [here](./config.go)
with detailed sample configurations [here](./testdata/config.yaml).
