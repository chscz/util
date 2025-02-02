package domain

import (
	"io"

	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
	"google.golang.org/protobuf/proto"
)

type Trace struct {
	tracepb.TracesData
}

func SetData() io.Reader {
	trace := &Trace{
		TracesData: tracepb.TracesData{
			ResourceSpans: []*tracepb.ResourceSpans{
				{
					Resource: &resourcepb.Resource{
						Attributes:             []*commonpb.KeyValue{},
						DroppedAttributesCount: 0,
					},
					ScopeSpans: nil,
					SchemaUrl:  "",
				},
				{
					Resource:   nil,
					ScopeSpans: nil,
					SchemaUrl:  "",
				},
				{
					Resource:   nil,
					ScopeSpans: nil,
					SchemaUrl:  "",
				},
			},
		},
	}

	b, err := proto.Marshal(trace)
	if err != nil {
		//
	}
	_ = b
	return nil
}
