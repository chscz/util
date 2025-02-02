package domain

import (
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
)

type Resource struct {
	resourcepb.Resource
}

var TestResource1 = &Resource{
	resourcepb.Resource{
		Attributes: []*commonpb.KeyValue{
			{
				Key:   "",
				Value: nil,
			},
		},

		DroppedAttributesCount: 0,
	},
}

/*

{'telemetry.sdk.name': 'opentelemetry', 'telemetry.sdk.version': '1.40.0', 'device.manufacturer': 'Google', 'imqa.sdk.version': '0.0.50', 'service.name': 'io.imqa.examples.basic', 'service.version': '1.0.0', 'telemetry.sdk.language': 'java', 'device.brand': 'google', 'device.id': 'ed95805728546c21', 'device.model.identifier': 'sdk_gphone64_arm64', 'os.name': 'android', 'os.version': '15'}
{'imqa.agent.version': '0.0.11', 'session.id': 'ae20a7852acf5b8d7a8c8da6db659a27', 'service.name': 'react-sample-web', 'telemetry.sdk.language': 'webjs', 'telemetry.sdk.version': '0.0.11', 'service.version': '0.0.1', 'service.key': 'DwRjafKFLVMskhHy43DzohNy66EifcoDyAKVWZdfZN6zzVxm0DucKAdSCDHqL46drRgVTslsAKQHY6_kiljJINfY-qHtr6rhnh_yBippTg.6JQgbbWgxC6jvMyjGqMCTg', 'rum.version': '0.0.11', 'rum.scriptInstance': '39ae9485f90f752d', 'telemetry.sdk.name': '@imqa/web-agent', 'process.runtime.name': 'browser', 'os.name': 'macOS', 'os.version': '10.15.7'}
*/
