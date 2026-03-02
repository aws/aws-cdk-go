package previewawsbackupgatewaymixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogs",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnHypervisorBgwHypervisorLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsFirehoseProps",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsLogGroupProps",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnHypervisorBgwHypervisorLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnHypervisorBgwHypervisorLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnHypervisorBgwHypervisorLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnHypervisorBgwHypervisorLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnHypervisorBgwHypervisorLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnHypervisorBgwHypervisorLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsOutputFormat.S3",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnHypervisorBgwHypervisorLogsOutputFormat_S3_JSON,
			"PLAIN": CfnHypervisorBgwHypervisorLogsOutputFormat_S3_PLAIN,
			"W3C": CfnHypervisorBgwHypervisorLogsOutputFormat_S3_W3C,
			"PARQUET": CfnHypervisorBgwHypervisorLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorBgwHypervisorLogsS3Props",
		reflect.TypeOf((*CfnHypervisorBgwHypervisorLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogs",
		reflect.TypeOf((*CfnHypervisorDataAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnHypervisorDataAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnHypervisorDataAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnHypervisorDataAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnHypervisorDataAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnHypervisorDataAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnHypervisorDataAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnHypervisorDataAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnHypervisorDataAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnHypervisorDataAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnHypervisorDataAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnHypervisorDataAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorDataAccessLogsS3Props",
		reflect.TypeOf((*CfnHypervisorDataAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorLogsMixin",
		reflect.TypeOf((*CfnHypervisorLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHypervisorLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorMixinProps",
		reflect.TypeOf((*CfnHypervisorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_backupgateway.mixins.CfnHypervisorPropsMixin",
		reflect.TypeOf((*CfnHypervisorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHypervisorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
