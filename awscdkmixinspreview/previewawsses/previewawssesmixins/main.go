package previewawssesmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogs",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsDestProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerIngressPointApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerIngressPointApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"INGRESS_POINT_TYPE": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_TYPE,
			"INGRESS_POINT_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_ID,
			"SENDER_IP_ADDRESS": CfnMailManagerIngressPointApplicationLogsRecordFields_SENDER_IP_ADDRESS,
			"SMTP_MAIL_FROM": CfnMailManagerIngressPointApplicationLogsRecordFields_SMTP_MAIL_FROM,
			"SMTP_HELO": CfnMailManagerIngressPointApplicationLogsRecordFields_SMTP_HELO,
			"TLS_PROTOCOL": CfnMailManagerIngressPointApplicationLogsRecordFields_TLS_PROTOCOL,
			"RECIPIENTS": CfnMailManagerIngressPointApplicationLogsRecordFields_RECIPIENTS,
			"INGRESS_POINT_METADATA": CfnMailManagerIngressPointApplicationLogsRecordFields_INGRESS_POINT_METADATA,
			"TLS_CIPHER_SUITE": CfnMailManagerIngressPointApplicationLogsRecordFields_TLS_CIPHER_SUITE,
			"MESSAGE_SIZE_BYTES": CfnMailManagerIngressPointApplicationLogsRecordFields_MESSAGE_SIZE_BYTES,
			"RULE_SET_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_RULE_SET_ID,
			"RESOURCE_ARN": CfnMailManagerIngressPointApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerIngressPointApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE_ID": CfnMailManagerIngressPointApplicationLogsRecordFields_MESSAGE_ID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointApplicationLogsS3Props",
		reflect.TypeOf((*CfnMailManagerIngressPointApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointLogsMixin",
		reflect.TypeOf((*CfnMailManagerIngressPointLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerIngressPointLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogs",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerIngressPointTrafficPolicyDebugLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"INGRESS_POINT_TYPE": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_TYPE,
			"INGRESS_POINT_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_ID,
			"INGRESS_POINT_SESSION_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_INGRESS_POINT_SESSION_ID,
			"TRAFFIC_POLICY_ID": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_ID,
			"TRAFFIC_POLICY_EVALUATION": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_EVALUATION,
			"TRAFFIC_POLICY_VERDICT": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TRAFFIC_POLICY_VERDICT,
			"SENDER_IP_ADDRESS": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SENDER_IP_ADDRESS,
			"SMTP_MAIL_FROM": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SMTP_MAIL_FROM,
			"SMTP_HELO": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_SMTP_HELO,
			"TLS_PROTOCOL": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TLS_PROTOCOL,
			"RECIPIENT": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_RECIPIENT,
			"TLS_CIPHER_SUITE": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_TLS_CIPHER_SUITE,
			"RESOURCE_ARN": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerIngressPointTrafficPolicyDebugLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props",
		reflect.TypeOf((*CfnMailManagerIngressPointTrafficPolicyDebugLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogs",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMailManagerRuleSetApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsDestProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMailManagerRuleSetApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMailManagerRuleSetApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMailManagerRuleSetApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsRecordFields",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RULE_NAME": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_NAME,
			"RULE_INDEX": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_INDEX,
			"RECIPIENTS_MATCHED": CfnMailManagerRuleSetApplicationLogsRecordFields_RECIPIENTS_MATCHED,
			"ACTION_METADATA": CfnMailManagerRuleSetApplicationLogsRecordFields_ACTION_METADATA,
			"RESOURCE_ARN": CfnMailManagerRuleSetApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMailManagerRuleSetApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MESSAGE_ID": CfnMailManagerRuleSetApplicationLogsRecordFields_MESSAGE_ID,
			"RULE_SET_NAME": CfnMailManagerRuleSetApplicationLogsRecordFields_RULE_SET_NAME,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetApplicationLogsS3Props",
		reflect.TypeOf((*CfnMailManagerRuleSetApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ses.mixins.CfnMailManagerRuleSetLogsMixin",
		reflect.TypeOf((*CfnMailManagerRuleSetLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMailManagerRuleSetLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
