package previewawsvpclatticemixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationLogsMixin",
		reflect.TypeOf((*CfnResourceConfigurationLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResourceConfigurationLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogs",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnResourceConfigurationResourceAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnResourceConfigurationResourceAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnResourceConfigurationResourceAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnResourceConfigurationResourceAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnResourceConfigurationResourceAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnResourceConfigurationResourceAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnResourceConfigurationResourceAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnResourceConfigurationResourceAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnResourceConfigurationResourceAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnResourceConfigurationResourceAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnResourceConfigurationResourceAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnResourceConfigurationResourceAccessLogsS3Props",
		reflect.TypeOf((*CfnResourceConfigurationResourceAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogs",
		reflect.TypeOf((*CfnServiceAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnServiceAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnServiceAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnServiceAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat",
		reflect.TypeOf((*CfnServiceAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnServiceAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnServiceAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnServiceAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnServiceAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnServiceAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnServiceAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnServiceAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnServiceAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnServiceAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnServiceAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnServiceAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnServiceAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnServiceAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceAccessLogsS3Props",
		reflect.TypeOf((*CfnServiceAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_vpclattice.mixins.CfnServiceLogsMixin",
		reflect.TypeOf((*CfnServiceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServiceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
