package domain

import commonpb "go.opentelemetry.io/proto/otlp/common/v1"

type Common struct {
	commonpb.KeyValue
}
