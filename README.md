# K8s Monitoring Playground on kind

[![yamllint](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/yamllint.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/yamllint.yaml)
[![golangci-lint](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/golangci-lint.yml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/golangci-lint.yml)
[![kind e2e](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/kind-e2e.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/k8s-monitoring-playground/actions/workflows/kind-e2e.yaml)
[![Renovate](https://img.shields.io/badge/renovate-enabled-brightgreen.svg)](https://renovatebot.com)

## Features

### Dashboard

Dashboard is provided by [Grafana](https://grafana.com).<br>
Following resources are added as data sources.

### Traces

Traces are measured in OpenTelemetry format, and sent to [Jaeger](https://www.jaegertracing.io).

### Metrics

Metrics are collected by [Prometheus](https://prometheus.io).

### Logs

Logs are collected by [Promtail](https://grafana.com/docs/loki/latest/clients/promtail), and aggregated by [Loki](https://grafana.com/oss/loki).
