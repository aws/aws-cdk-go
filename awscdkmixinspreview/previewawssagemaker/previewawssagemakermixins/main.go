package previewawssagemakermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogs",
		reflect.TypeOf((*CfnWorkteamActivityLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWorkteamActivityLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsLogGroupProps",
		reflect.TypeOf((*CfnWorkteamActivityLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsOutputFormat",
		reflect.TypeOf((*CfnWorkteamActivityLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnWorkteamActivityLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnWorkteamActivityLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnWorkteamActivityLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnWorkteamActivityLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsOutputFormat.S3",
		reflect.TypeOf((*CfnWorkteamActivityLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWorkteamActivityLogsOutputFormat_S3_JSON,
			"PLAIN": CfnWorkteamActivityLogsOutputFormat_S3_PLAIN,
			"W3C": CfnWorkteamActivityLogsOutputFormat_S3_W3C,
			"PARQUET": CfnWorkteamActivityLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamActivityLogsS3Props",
		reflect.TypeOf((*CfnWorkteamActivityLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnWorkteamLogsMixin",
		reflect.TypeOf((*CfnWorkteamLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkteamLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
