// Package metrics sets up and handles our prometheus collectors.
package metrics

import (
	"context"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	api "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric"
)

var (
	requestsReceived     api.Int64Counter
	timeToProcessRequest api.Float64Histogram
)

// RegisterPrometheusCollectors tells prometheus to set up collectors.
func RegisterPrometheusCollectors() {
	// The exporter embeds a default OpenTelemetry Reader and
	// implements prometheus.Collector, allowing it to be used as
	// both a Reader and Collector.
	exporter, err := prometheus.New()
	if err != nil {
		panic(err)
	}

	provider := metric.NewMeterProvider(metric.WithReader(exporter))
	meter := provider.Meter("api-gateway")

	requestsReceived, _ = meter.Int64Counter("http_request_status_code",
		api.WithDescription("Status codes returned by the API"),
		api.WithUnit("{call}"),
	)

	timeToProcessRequest, _ = meter.Float64Histogram("http_request_duration",
		api.WithDescription("Time spent processing requests"),
		api.WithUnit("s"),
	)

}

// ObserveTimeToProcess records the time spent processing an operation.
func ObserveTimeToProcess(ctx context.Context, operation string, t float64) {
	opt := api.WithAttributes(
		attribute.Key("endpoint").String(operation),
	)

	timeToProcessRequest.Record(ctx, t, opt)
}

// ReceivedRequest records the status code returned for each request.
func ReceivedRequest(ctx context.Context, statusCode int, operationName string) {
	opt := api.WithAttributes(
		attribute.Key(operationName).Int(statusCode),
	)

	requestsReceived.Add(ctx, 1, opt)
}
