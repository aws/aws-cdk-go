package previewawswafv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogs",
		reflect.TypeOf((*CfnWebACLAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWebACLAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnWebACLAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnWebACLAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat",
		reflect.TypeOf((*CfnWebACLAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnWebACLAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnWebACLAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWebACLAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnWebACLAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnWebACLAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnWebACLAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnWebACLAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnWebACLAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnWebACLAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWebACLAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnWebACLAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnWebACLAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnWebACLAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLAccessLogsS3Props",
		reflect.TypeOf((*CfnWebACLAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLLogsMixin",
		reflect.TypeOf((*CfnWebACLLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWebACLLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogs",
		reflect.TypeOf((*CfnWebACLTokenLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWebACLTokenLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsFirehoseProps",
		reflect.TypeOf((*CfnWebACLTokenLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsLogGroupProps",
		reflect.TypeOf((*CfnWebACLTokenLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat",
		reflect.TypeOf((*CfnWebACLTokenLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnWebACLTokenLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnWebACLTokenLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWebACLTokenLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnWebACLTokenLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnWebACLTokenLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnWebACLTokenLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnWebACLTokenLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnWebACLTokenLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsOutputFormat.S3",
		reflect.TypeOf((*CfnWebACLTokenLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWebACLTokenLogsOutputFormat_S3_JSON,
			"PLAIN": CfnWebACLTokenLogsOutputFormat_S3_PLAIN,
			"W3C": CfnWebACLTokenLogsOutputFormat_S3_W3C,
			"PARQUET": CfnWebACLTokenLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wafv2.mixins.CfnWebACLTokenLogsS3Props",
		reflect.TypeOf((*CfnWebACLTokenLogsS3Props)(nil)).Elem(),
	)
}
