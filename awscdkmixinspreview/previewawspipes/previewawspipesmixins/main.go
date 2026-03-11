package previewawspipesmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogs",
		reflect.TypeOf((*CfnPipeExecutionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnPipeExecutionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsFirehoseProps",
		reflect.TypeOf((*CfnPipeExecutionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsLogGroupProps",
		reflect.TypeOf((*CfnPipeExecutionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat",
		reflect.TypeOf((*CfnPipeExecutionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnPipeExecutionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnPipeExecutionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPipeExecutionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnPipeExecutionLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnPipeExecutionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnPipeExecutionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnPipeExecutionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnPipeExecutionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnPipeExecutionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnPipeExecutionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnPipeExecutionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnPipeExecutionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnPipeExecutionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeExecutionLogsS3Props",
		reflect.TypeOf((*CfnPipeExecutionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_pipes.mixins.CfnPipeLogsMixin",
		reflect.TypeOf((*CfnPipeLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnPipeLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
