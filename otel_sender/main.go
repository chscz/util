package main

import (
	"github.com/chscz/test_project/otel_sender/config"
	collogspb "go.opentelemetry.io/proto/otlp/collector/logs/v1"
	colmetricspb "go.opentelemetry.io/proto/otlp/collector/metrics/v1"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	logspb "go.opentelemetry.io/proto/otlp/logs/v1"
	metricspb "go.opentelemetry.io/proto/otlp/metrics/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

func main() {
	_ = collogspb.ExportLogsServiceRequest{}
	_ = colmetricspb.ExportMetricsServiceRequest{}
	_ = coltracepb.ExportTraceServiceRequest{}
	_ = commonpb.AnyValue{}
	_ = logspb.LogsData{}
	_ = metricspb.MetricsData{}
	_ = resourcepb.Resource{}
	_ = tracepb.TracesData{}

	cfg, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}
_ cfg
}
