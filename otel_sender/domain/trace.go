package domain

import tracepb "go.opentelemetry.io/proto/otlp/trace/v1"

type Trace struct {
	tracepb.TracesData
}
