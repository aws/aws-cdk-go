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
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsDestProps",
		reflect.TypeOf((*CfnFirewallAlertLogsDestProps)(nil)).Elem(),
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
			"PLAIN": CfnFirewallAlertLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnFirewallAlertLogsOutputFormat_Firehose_JSON,
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
			"PLAIN": CfnFirewallAlertLogsOutputFormat_S3_PLAIN,
			"JSON": CfnFirewallAlertLogsOutputFormat_S3_JSON,
			"W3C": CfnFirewallAlertLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallAlertLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsRecordFields",
		reflect.TypeOf((*CfnFirewallAlertLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnFirewallAlertLogsRecordFields_TIMESTAMP,
			"RESOURCE_ARN": CfnFirewallAlertLogsRecordFields_RESOURCE_ARN,
			"RESOURCEID": CfnFirewallAlertLogsRecordFields_RESOURCEID,
			"CLIENT_SRC_IP": CfnFirewallAlertLogsRecordFields_CLIENT_SRC_IP,
			"CLIENT_SRC_PORT": CfnFirewallAlertLogsRecordFields_CLIENT_SRC_PORT,
			"SRC_VPC": CfnFirewallAlertLogsRecordFields_SRC_VPC,
			"SRC_VPCE": CfnFirewallAlertLogsRecordFields_SRC_VPCE,
			"VPCE_ACCOUNT_ID": CfnFirewallAlertLogsRecordFields_VPCE_ACCOUNT_ID,
			"DEST_DOMAIN": CfnFirewallAlertLogsRecordFields_DEST_DOMAIN,
			"URL": CfnFirewallAlertLogsRecordFields_URL,
			"HTTP_METHOD": CfnFirewallAlertLogsRecordFields_HTTP_METHOD,
			"DEST_IP": CfnFirewallAlertLogsRecordFields_DEST_IP,
			"DEST_PORT": CfnFirewallAlertLogsRecordFields_DEST_PORT,
			"HTTP_STATUS_CODE": CfnFirewallAlertLogsRecordFields_HTTP_STATUS_CODE,
			"FIRST_ALERT_MATCH": CfnFirewallAlertLogsRecordFields_FIRST_ALERT_MATCH,
			"ALL_MATCHES": CfnFirewallAlertLogsRecordFields_ALL_MATCHES,
			"FINAL_RULE_NAME": CfnFirewallAlertLogsRecordFields_FINAL_RULE_NAME,
			"FINAL_RULE_GROUP_NAME": CfnFirewallAlertLogsRecordFields_FINAL_RULE_GROUP_NAME,
			"REQUEST_SIZE": CfnFirewallAlertLogsRecordFields_REQUEST_SIZE,
			"RESPONSE_SIZE": CfnFirewallAlertLogsRecordFields_RESPONSE_SIZE,
			"EVENT_TIMESTAMP": CfnFirewallAlertLogsRecordFields_EVENT_TIMESTAMP,
			"PROXY_NAME": CfnFirewallAlertLogsRecordFields_PROXY_NAME,
			"FINAL_ACTION": CfnFirewallAlertLogsRecordFields_FINAL_ACTION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAlertLogsS3Props",
		reflect.TypeOf((*CfnFirewallAlertLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogs",
		reflect.TypeOf((*CfnFirewallAllowLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallAllowLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsDestProps",
		reflect.TypeOf((*CfnFirewallAllowLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallAllowLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallAllowLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallAllowLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallAllowLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallAllowLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallAllowLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnFirewallAllowLogsOutputFormat_Firehose_JSON,
			"RAW": CfnFirewallAllowLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallAllowLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallAllowLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallAllowLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallAllowLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallAllowLogsOutputFormat_S3_PLAIN,
			"JSON": CfnFirewallAllowLogsOutputFormat_S3_JSON,
			"W3C": CfnFirewallAllowLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallAllowLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsRecordFields",
		reflect.TypeOf((*CfnFirewallAllowLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnFirewallAllowLogsRecordFields_TIMESTAMP,
			"RESOURCE_ARN": CfnFirewallAllowLogsRecordFields_RESOURCE_ARN,
			"RESOURCEID": CfnFirewallAllowLogsRecordFields_RESOURCEID,
			"CLIENT_SRC_IP": CfnFirewallAllowLogsRecordFields_CLIENT_SRC_IP,
			"CLIENT_SRC_PORT": CfnFirewallAllowLogsRecordFields_CLIENT_SRC_PORT,
			"SRC_VPC": CfnFirewallAllowLogsRecordFields_SRC_VPC,
			"SRC_VPCE": CfnFirewallAllowLogsRecordFields_SRC_VPCE,
			"VPCE_ACCOUNT_ID": CfnFirewallAllowLogsRecordFields_VPCE_ACCOUNT_ID,
			"DEST_DOMAIN": CfnFirewallAllowLogsRecordFields_DEST_DOMAIN,
			"URL": CfnFirewallAllowLogsRecordFields_URL,
			"HTTP_METHOD": CfnFirewallAllowLogsRecordFields_HTTP_METHOD,
			"DEST_IP": CfnFirewallAllowLogsRecordFields_DEST_IP,
			"DEST_PORT": CfnFirewallAllowLogsRecordFields_DEST_PORT,
			"HTTP_STATUS_CODE": CfnFirewallAllowLogsRecordFields_HTTP_STATUS_CODE,
			"FIRST_ALERT_MATCH": CfnFirewallAllowLogsRecordFields_FIRST_ALERT_MATCH,
			"ALL_MATCHES": CfnFirewallAllowLogsRecordFields_ALL_MATCHES,
			"FINAL_RULE_NAME": CfnFirewallAllowLogsRecordFields_FINAL_RULE_NAME,
			"FINAL_RULE_GROUP_NAME": CfnFirewallAllowLogsRecordFields_FINAL_RULE_GROUP_NAME,
			"REQUEST_SIZE": CfnFirewallAllowLogsRecordFields_REQUEST_SIZE,
			"RESPONSE_SIZE": CfnFirewallAllowLogsRecordFields_RESPONSE_SIZE,
			"EVENT_TIMESTAMP": CfnFirewallAllowLogsRecordFields_EVENT_TIMESTAMP,
			"PROXY_NAME": CfnFirewallAllowLogsRecordFields_PROXY_NAME,
			"FINAL_ACTION": CfnFirewallAllowLogsRecordFields_FINAL_ACTION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallAllowLogsS3Props",
		reflect.TypeOf((*CfnFirewallAllowLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogs",
		reflect.TypeOf((*CfnFirewallDenyLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFirewallDenyLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsDestProps",
		reflect.TypeOf((*CfnFirewallDenyLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsFirehoseProps",
		reflect.TypeOf((*CfnFirewallDenyLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsLogGroupProps",
		reflect.TypeOf((*CfnFirewallDenyLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat",
		reflect.TypeOf((*CfnFirewallDenyLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFirewallDenyLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFirewallDenyLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallDenyLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnFirewallDenyLogsOutputFormat_Firehose_JSON,
			"RAW": CfnFirewallDenyLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFirewallDenyLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallDenyLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFirewallDenyLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFirewallDenyLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFirewallDenyLogsOutputFormat_S3_PLAIN,
			"JSON": CfnFirewallDenyLogsOutputFormat_S3_JSON,
			"W3C": CfnFirewallDenyLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFirewallDenyLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsRecordFields",
		reflect.TypeOf((*CfnFirewallDenyLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnFirewallDenyLogsRecordFields_TIMESTAMP,
			"RESOURCE_ARN": CfnFirewallDenyLogsRecordFields_RESOURCE_ARN,
			"RESOURCEID": CfnFirewallDenyLogsRecordFields_RESOURCEID,
			"CLIENT_SRC_IP": CfnFirewallDenyLogsRecordFields_CLIENT_SRC_IP,
			"CLIENT_SRC_PORT": CfnFirewallDenyLogsRecordFields_CLIENT_SRC_PORT,
			"SRC_VPC": CfnFirewallDenyLogsRecordFields_SRC_VPC,
			"SRC_VPCE": CfnFirewallDenyLogsRecordFields_SRC_VPCE,
			"VPCE_ACCOUNT_ID": CfnFirewallDenyLogsRecordFields_VPCE_ACCOUNT_ID,
			"DEST_DOMAIN": CfnFirewallDenyLogsRecordFields_DEST_DOMAIN,
			"URL": CfnFirewallDenyLogsRecordFields_URL,
			"HTTP_METHOD": CfnFirewallDenyLogsRecordFields_HTTP_METHOD,
			"DEST_IP": CfnFirewallDenyLogsRecordFields_DEST_IP,
			"DEST_PORT": CfnFirewallDenyLogsRecordFields_DEST_PORT,
			"HTTP_STATUS_CODE": CfnFirewallDenyLogsRecordFields_HTTP_STATUS_CODE,
			"FIRST_ALERT_MATCH": CfnFirewallDenyLogsRecordFields_FIRST_ALERT_MATCH,
			"ALL_MATCHES": CfnFirewallDenyLogsRecordFields_ALL_MATCHES,
			"FINAL_RULE_NAME": CfnFirewallDenyLogsRecordFields_FINAL_RULE_NAME,
			"FINAL_RULE_GROUP_NAME": CfnFirewallDenyLogsRecordFields_FINAL_RULE_GROUP_NAME,
			"REQUEST_SIZE": CfnFirewallDenyLogsRecordFields_REQUEST_SIZE,
			"RESPONSE_SIZE": CfnFirewallDenyLogsRecordFields_RESPONSE_SIZE,
			"EVENT_TIMESTAMP": CfnFirewallDenyLogsRecordFields_EVENT_TIMESTAMP,
			"PROXY_NAME": CfnFirewallDenyLogsRecordFields_PROXY_NAME,
			"FINAL_ACTION": CfnFirewallDenyLogsRecordFields_FINAL_ACTION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_networkfirewall.mixins.CfnFirewallDenyLogsS3Props",
		reflect.TypeOf((*CfnFirewallDenyLogsS3Props)(nil)).Elem(),
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
}
