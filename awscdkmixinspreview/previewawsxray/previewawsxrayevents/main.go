package previewawsxrayevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate",
		reflect.TypeOf((*AWSXRayInsightUpdate)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSXRayInsightUpdate{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.AWSXRayInsightUpdateProps",
		reflect.TypeOf((*AWSXRayInsightUpdate_AWSXRayInsightUpdateProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.AwsxRayInsightUpdateItem",
		reflect.TypeOf((*AWSXRayInsightUpdate_AwsxRayInsightUpdateItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.ClientRequestImpactStatistics",
		reflect.TypeOf((*AWSXRayInsightUpdate_ClientRequestImpactStatistics)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.Event",
		reflect.TypeOf((*AWSXRayInsightUpdate_Event)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.RootCauseServiceId",
		reflect.TypeOf((*AWSXRayInsightUpdate_RootCauseServiceId)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_xray.events.AWSXRayInsightUpdate.ServiceId",
		reflect.TypeOf((*AWSXRayInsightUpdate_ServiceId)(nil)).Elem(),
	)
}
