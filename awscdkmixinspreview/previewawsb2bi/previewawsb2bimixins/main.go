package previewawsb2bimixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogs",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnTransformerB2biExecutionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsFirehoseProps",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsLogGroupProps",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnTransformerB2biExecutionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnTransformerB2biExecutionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnTransformerB2biExecutionLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnTransformerB2biExecutionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnTransformerB2biExecutionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnTransformerB2biExecutionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnTransformerB2biExecutionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnTransformerB2biExecutionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnTransformerB2biExecutionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnTransformerB2biExecutionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerB2biExecutionLogsS3Props",
		reflect.TypeOf((*CfnTransformerB2biExecutionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_b2bi.mixins.CfnTransformerLogsMixin",
		reflect.TypeOf((*CfnTransformerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTransformerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
