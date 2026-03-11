package previewawselasticloadbalancingv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogs",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsDestProps",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLoadBalancerAlbAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLoadBalancerAlbAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLoadBalancerAlbAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLoadBalancerAlbAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLoadBalancerAlbAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLoadBalancerAlbAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLoadBalancerAlbAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsRecordFields",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TYPE": CfnLoadBalancerAlbAccessLogsRecordFields_TYPE,
			"TIME": CfnLoadBalancerAlbAccessLogsRecordFields_TIME,
			"ELB": CfnLoadBalancerAlbAccessLogsRecordFields_ELB,
			"CLIENT_PORT": CfnLoadBalancerAlbAccessLogsRecordFields_CLIENT_PORT,
			"TARGET_PORT": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_PORT,
			"REQUEST_PROCESSING_TIME": CfnLoadBalancerAlbAccessLogsRecordFields_REQUEST_PROCESSING_TIME,
			"TARGET_PROCESSING_TIME": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_PROCESSING_TIME,
			"RESPONSE_PROCESSING_TIME": CfnLoadBalancerAlbAccessLogsRecordFields_RESPONSE_PROCESSING_TIME,
			"ELB_STATUS_CODE": CfnLoadBalancerAlbAccessLogsRecordFields_ELB_STATUS_CODE,
			"TARGET_STATUS_CODE": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_STATUS_CODE,
			"RECEIVED_BYTES": CfnLoadBalancerAlbAccessLogsRecordFields_RECEIVED_BYTES,
			"SENT_BYTES": CfnLoadBalancerAlbAccessLogsRecordFields_SENT_BYTES,
			"REQUEST_LINE": CfnLoadBalancerAlbAccessLogsRecordFields_REQUEST_LINE,
			"USER_AGENT": CfnLoadBalancerAlbAccessLogsRecordFields_USER_AGENT,
			"SSL_CIPHER": CfnLoadBalancerAlbAccessLogsRecordFields_SSL_CIPHER,
			"SSL_PROTOCOL": CfnLoadBalancerAlbAccessLogsRecordFields_SSL_PROTOCOL,
			"TARGET_GROUP_ARN": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_GROUP_ARN,
			"TRACE_ID": CfnLoadBalancerAlbAccessLogsRecordFields_TRACE_ID,
			"DOMAIN_NAME": CfnLoadBalancerAlbAccessLogsRecordFields_DOMAIN_NAME,
			"CHOSEN_CERT_ARN": CfnLoadBalancerAlbAccessLogsRecordFields_CHOSEN_CERT_ARN,
			"MATCHED_RULE_PRIORITY": CfnLoadBalancerAlbAccessLogsRecordFields_MATCHED_RULE_PRIORITY,
			"REQUEST_CREATION_TIME": CfnLoadBalancerAlbAccessLogsRecordFields_REQUEST_CREATION_TIME,
			"ACTIONS_EXECUTED": CfnLoadBalancerAlbAccessLogsRecordFields_ACTIONS_EXECUTED,
			"REDIRECT_URL": CfnLoadBalancerAlbAccessLogsRecordFields_REDIRECT_URL,
			"ERROR_REASON": CfnLoadBalancerAlbAccessLogsRecordFields_ERROR_REASON,
			"TARGET_PORT_LIST": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_PORT_LIST,
			"TARGET_STATUS_CODE_LIST": CfnLoadBalancerAlbAccessLogsRecordFields_TARGET_STATUS_CODE_LIST,
			"CLASSIFICATION": CfnLoadBalancerAlbAccessLogsRecordFields_CLASSIFICATION,
			"CLASSIFICATION_REASON": CfnLoadBalancerAlbAccessLogsRecordFields_CLASSIFICATION_REASON,
			"CONN_TRACE_ID": CfnLoadBalancerAlbAccessLogsRecordFields_CONN_TRACE_ID,
			"TRANSFORMED_HOST": CfnLoadBalancerAlbAccessLogsRecordFields_TRANSFORMED_HOST,
			"TRANSFORMED_URI": CfnLoadBalancerAlbAccessLogsRecordFields_TRANSFORMED_URI,
			"REQUEST_TRANSFORM_STATUS": CfnLoadBalancerAlbAccessLogsRecordFields_REQUEST_TRANSFORM_STATUS,
			"IP_ADDRESS": CfnLoadBalancerAlbAccessLogsRecordFields_IP_ADDRESS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbAccessLogsS3Props",
		reflect.TypeOf((*CfnLoadBalancerAlbAccessLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogs",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbConnectionLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsDestProps",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsFirehoseProps",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsLogGroupProps",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbConnectionLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbConnectionLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLoadBalancerAlbConnectionLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLoadBalancerAlbConnectionLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLoadBalancerAlbConnectionLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLoadBalancerAlbConnectionLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbConnectionLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLoadBalancerAlbConnectionLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLoadBalancerAlbConnectionLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLoadBalancerAlbConnectionLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsRecordFields",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIME": CfnLoadBalancerAlbConnectionLogsRecordFields_TIME,
			"CLIENT_IP": CfnLoadBalancerAlbConnectionLogsRecordFields_CLIENT_IP,
			"CLIENT_PORT": CfnLoadBalancerAlbConnectionLogsRecordFields_CLIENT_PORT,
			"LISTENER_PORT": CfnLoadBalancerAlbConnectionLogsRecordFields_LISTENER_PORT,
			"TLS_PROTOCOL": CfnLoadBalancerAlbConnectionLogsRecordFields_TLS_PROTOCOL,
			"TLS_CIPHER": CfnLoadBalancerAlbConnectionLogsRecordFields_TLS_CIPHER,
			"TLS_HANDSHAKE_LATENCY": CfnLoadBalancerAlbConnectionLogsRecordFields_TLS_HANDSHAKE_LATENCY,
			"LEAF_CLIENT_CERT_SUBJECT": CfnLoadBalancerAlbConnectionLogsRecordFields_LEAF_CLIENT_CERT_SUBJECT,
			"LEAF_CLIENT_CERT_VALIDITY": CfnLoadBalancerAlbConnectionLogsRecordFields_LEAF_CLIENT_CERT_VALIDITY,
			"LEAF_CLIENT_CERT_SERIAL_NUMBER": CfnLoadBalancerAlbConnectionLogsRecordFields_LEAF_CLIENT_CERT_SERIAL_NUMBER,
			"TLS_VERIFY_STATUS": CfnLoadBalancerAlbConnectionLogsRecordFields_TLS_VERIFY_STATUS,
			"CONN_TRACE_ID": CfnLoadBalancerAlbConnectionLogsRecordFields_CONN_TRACE_ID,
			"TLS_KEYEXCHANGE": CfnLoadBalancerAlbConnectionLogsRecordFields_TLS_KEYEXCHANGE,
			"ELB": CfnLoadBalancerAlbConnectionLogsRecordFields_ELB,
			"IP_ADDRESS": CfnLoadBalancerAlbConnectionLogsRecordFields_IP_ADDRESS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbConnectionLogsS3Props",
		reflect.TypeOf((*CfnLoadBalancerAlbConnectionLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogs",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbHealthCheckLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsDestProps",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsFirehoseProps",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsLogGroupProps",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerAlbHealthCheckLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLoadBalancerAlbHealthCheckLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsRecordFields",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TYPE": CfnLoadBalancerAlbHealthCheckLogsRecordFields_TYPE,
			"TIME": CfnLoadBalancerAlbHealthCheckLogsRecordFields_TIME,
			"LATENCY": CfnLoadBalancerAlbHealthCheckLogsRecordFields_LATENCY,
			"TARGET_ADDR": CfnLoadBalancerAlbHealthCheckLogsRecordFields_TARGET_ADDR,
			"TARGET_GROUP_ID": CfnLoadBalancerAlbHealthCheckLogsRecordFields_TARGET_GROUP_ID,
			"STATUS": CfnLoadBalancerAlbHealthCheckLogsRecordFields_STATUS,
			"STATUS_CODE": CfnLoadBalancerAlbHealthCheckLogsRecordFields_STATUS_CODE,
			"REASON_CODE": CfnLoadBalancerAlbHealthCheckLogsRecordFields_REASON_CODE,
			"ELB": CfnLoadBalancerAlbHealthCheckLogsRecordFields_ELB,
			"IP_ADDRESS": CfnLoadBalancerAlbHealthCheckLogsRecordFields_IP_ADDRESS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerAlbHealthCheckLogsS3Props",
		reflect.TypeOf((*CfnLoadBalancerAlbHealthCheckLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerLogsMixin",
		reflect.TypeOf((*CfnLoadBalancerLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLoadBalancerLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogs",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerNlbAccessLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsDestProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsFirehoseProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsLogGroupProps",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnLoadBalancerNlbAccessLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnLoadBalancerNlbAccessLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsOutputFormat.S3",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_JSON,
			"PLAIN": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_PLAIN,
			"W3C": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_W3C,
			"PARQUET": CfnLoadBalancerNlbAccessLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsRecordFields",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TYPE": CfnLoadBalancerNlbAccessLogsRecordFields_TYPE,
			"VERSION": CfnLoadBalancerNlbAccessLogsRecordFields_VERSION,
			"TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TIME,
			"ELB": CfnLoadBalancerNlbAccessLogsRecordFields_ELB,
			"LISTENER": CfnLoadBalancerNlbAccessLogsRecordFields_LISTENER,
			"CLIENT_PORT": CfnLoadBalancerNlbAccessLogsRecordFields_CLIENT_PORT,
			"DESTINATION_PORT": CfnLoadBalancerNlbAccessLogsRecordFields_DESTINATION_PORT,
			"CONNECTION_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_CONNECTION_TIME,
			"TLS_HANDSHAKE_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_HANDSHAKE_TIME,
			"RECEIVED_BYTES": CfnLoadBalancerNlbAccessLogsRecordFields_RECEIVED_BYTES,
			"SENT_BYTES": CfnLoadBalancerNlbAccessLogsRecordFields_SENT_BYTES,
			"INCOMING_TLS_ALERT": CfnLoadBalancerNlbAccessLogsRecordFields_INCOMING_TLS_ALERT,
			"CHOSEN_CERT_ARN": CfnLoadBalancerNlbAccessLogsRecordFields_CHOSEN_CERT_ARN,
			"CHOSEN_CERT_SERIAL": CfnLoadBalancerNlbAccessLogsRecordFields_CHOSEN_CERT_SERIAL,
			"TLS_CIPHER": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_CIPHER,
			"TLS_PROTOCOL_VERSION": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_PROTOCOL_VERSION,
			"TLS_KEYEXCHANGE": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_KEYEXCHANGE,
			"DOMAIN_NAME": CfnLoadBalancerNlbAccessLogsRecordFields_DOMAIN_NAME,
			"ALPN_FE_PROTOCOL": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_FE_PROTOCOL,
			"ALPN_BE_PROTOCOL": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_BE_PROTOCOL,
			"ALPN_CLIENT_PREFERENCE_LIST": CfnLoadBalancerNlbAccessLogsRecordFields_ALPN_CLIENT_PREFERENCE_LIST,
			"TLS_CONNECTION_CREATION_TIME": CfnLoadBalancerNlbAccessLogsRecordFields_TLS_CONNECTION_CREATION_TIME,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_elasticloadbalancingv2.mixins.CfnLoadBalancerNlbAccessLogsS3Props",
		reflect.TypeOf((*CfnLoadBalancerNlbAccessLogsS3Props)(nil)).Elem(),
	)
}
