package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/chscz/test_project/otel_sender/config"
	"github.com/chscz/test_project/otel_sender/service"
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

	ctx := context.Background()
	cfg, err := config.LoadFromEnv()
	if err != nil {
		panic(err)
	}
	container := service.NewContainer(cfg)
	defer container.Close()

	senders := cfg.GetSenderList()

	var wg sync.WaitGroup
	for _, sender := range senders {
		wg.Add(1)
		s := container.InitSender(ctx, sender)
		success, failed := s.Start(ctx)
		fmt.Println("type::", sender, "success::", success, "failed::", failed)
		wg.Done()
	}
	wg.Wait()
}
