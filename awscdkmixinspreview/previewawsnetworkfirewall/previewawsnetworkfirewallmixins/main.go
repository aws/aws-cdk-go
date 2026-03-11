package previewawsnetworkfirewallmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogs",
		reflect.TypeOf((*CfnFirewallAlertLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallAlertLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallAlertLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallAlertLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallAlertLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallAlertLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallAlertLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallAlertLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallAlertLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallAlertLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallAlertLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallAlertLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallAlertLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallAlertLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallAlertLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsS3Props",
		reflect.TypeOf((*CfnFirewallAlertLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogs",
		reflect.TypeOf((*CfnFirewallFlowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallFlowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallFlowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallFlowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallFlowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallFlowLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallFlowLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallFlowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallFlowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallFlowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallFlowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallFlowLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallFlowLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallFlowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallFlowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallFlowLogsS3Props",
		reflect.TypeOf((*CfnFirewallFlowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallLogsMixin",
		reflect.TypeOf((*CfnFirewallLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFirewallLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogs",
		reflect.TypeOf((*CfnFirewallTlsLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallTlsLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallTlsLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallTlsLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallTlsLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallTlsLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFirewallTlsLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFirewallTlsLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallTlsLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallTlsLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallTlsLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFirewallTlsLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFirewallTlsLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFirewallTlsLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallTlsLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallTlsLogsS3Props",
		reflect.TypeOf((*CfnFirewallTlsLogsS3Props)(nil)).Elem(),
	)
}
