package previewawsapigatewaymixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogs",
		reflect.TypeOf((*CfnRestApiAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRestApiAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsDestProps",
		reflect.TypeOf((*CfnRestApiAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnRestApiAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnRestApiAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat",
		reflect.TypeOf((*CfnRestApiAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRestApiAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRestApiAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRestApiAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRestApiAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRestApiAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRestApiAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRestApiAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRestApiAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRestApiAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRestApiAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRestApiAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRestApiAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRestApiAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsRecordFields",
		reflect.TypeOf((*CfnRestApiAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnRestApiAccessLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRestApiAccessLogsRecordFields_EVENT_TIMESTAMP,
			"STAGE": CfnRestApiAccessLogsRecordFields_STAGE,
			"PAYLOAD": CfnRestApiAccessLogsRecordFields_PAYLOAD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiAccessLogsS3Props",
		reflect.TypeOf((*CfnRestApiAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogs",
		reflect.TypeOf((*CfnRestApiExecutionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRestApiExecutionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsDestProps",
		reflect.TypeOf((*CfnRestApiExecutionLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsFirehoseProps",
		reflect.TypeOf((*CfnRestApiExecutionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsLogGroupProps",
		reflect.TypeOf((*CfnRestApiExecutionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat",
		reflect.TypeOf((*CfnRestApiExecutionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRestApiExecutionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRestApiExecutionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRestApiExecutionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRestApiExecutionLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRestApiExecutionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRestApiExecutionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRestApiExecutionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRestApiExecutionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRestApiExecutionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRestApiExecutionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRestApiExecutionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRestApiExecutionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRestApiExecutionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsRecordFields",
		reflect.TypeOf((*CfnRestApiExecutionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnRestApiExecutionLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRestApiExecutionLogsRecordFields_EVENT_TIMESTAMP,
			"STAGE": CfnRestApiExecutionLogsRecordFields_STAGE,
			"PAYLOAD": CfnRestApiExecutionLogsRecordFields_PAYLOAD,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiExecutionLogsS3Props",
		reflect.TypeOf((*CfnRestApiExecutionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apigateway.mixins.CfnRestApiLogsMixin",
		reflect.TypeOf((*CfnRestApiLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRestApiLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
