package previewawsiotfleetwisemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogs",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCampaignIotFleetwiseLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsFirehoseProps",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsLogGroupProps",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCampaignIotFleetwiseLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCampaignIotFleetwiseLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCampaignIotFleetwiseLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCampaignIotFleetwiseLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCampaignIotFleetwiseLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCampaignIotFleetwiseLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCampaignIotFleetwiseLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCampaignIotFleetwiseLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCampaignIotFleetwiseLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCampaignIotFleetwiseLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignIotFleetwiseLogsS3Props",
		reflect.TypeOf((*CfnCampaignIotFleetwiseLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnCampaignLogsMixin",
		reflect.TypeOf((*CfnCampaignLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaignLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogs",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnVehicleIotFleetwiseLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsFirehoseProps",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsLogGroupProps",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnVehicleIotFleetwiseLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVehicleIotFleetwiseLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnVehicleIotFleetwiseLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnVehicleIotFleetwiseLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnVehicleIotFleetwiseLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnVehicleIotFleetwiseLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsOutputFormat.S3",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnVehicleIotFleetwiseLogsOutputFormat_S3_JSON,
			"PLAIN": CfnVehicleIotFleetwiseLogsOutputFormat_S3_PLAIN,
			"W3C": CfnVehicleIotFleetwiseLogsOutputFormat_S3_W3C,
			"PARQUET": CfnVehicleIotFleetwiseLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleIotFleetwiseLogsS3Props",
		reflect.TypeOf((*CfnVehicleIotFleetwiseLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnVehicleLogsMixin",
		reflect.TypeOf((*CfnVehicleLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVehicleLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
