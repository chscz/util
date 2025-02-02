package mock

import (
	"math/rand"

	"github.com/chscz/test_project/otel_sender/domain"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

const (
	ResourceAttributeKeyServiceName    = "service.name"
	ResourceAttributeKeyServiceVersion = "service.version"
	ResourceAttributeKeySessionID      = "session.id"
)

func GenerateResource() *domain.Resource {
	n := rand.Intn(len(ServiceSession))
	serviceSession := ServiceSession[n]
	r := &domain.Resource{
		Resource: resourcepb.Resource{
			Attributes: []*commonpb.KeyValue{
				{
					Key: ResourceAttributeKeyServiceName,
					Value: &commonpb.AnyValue{
						Value: &commonpb.AnyValue_StringValue{
							StringValue: serviceSession.ServiceName,
						},
					},
				},
				{
					Key: ResourceAttributeKeyServiceVersion,
					Value: &commonpb.AnyValue{
						Value: &commonpb.AnyValue_StringValue{
							StringValue: serviceSession.ServiceVersion,
						},
					},
				},
				{
					Key: ResourceAttributeKeySessionID,
					Value: &commonpb.AnyValue{
						Value: &commonpb.AnyValue_StringValue{
							StringValue: serviceSession.SessionID,
						},
					},
				},
			},
			DroppedAttributesCount: 0,
		},
	}
	return r
}
