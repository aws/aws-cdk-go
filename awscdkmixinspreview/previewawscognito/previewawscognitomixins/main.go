package previewawscognitomixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogs",
		reflect.TypeOf((*CfnUserPoolApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnUserPoolApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnUserPoolApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnUserPoolApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnUserPoolApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnUserPoolApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnUserPoolApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnUserPoolApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnUserPoolApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnUserPoolApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnUserPoolApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnUserPoolApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnUserPoolApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnUserPoolApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnUserPoolApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnUserPoolApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnUserPoolApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnUserPoolApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolApplicationLogsS3Props",
		reflect.TypeOf((*CfnUserPoolApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cognito.mixins.CfnUserPoolLogsMixin",
		reflect.TypeOf((*CfnUserPoolLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnUserPoolLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
