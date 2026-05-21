package previewawsbedrockagentcoremixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomLogsMixin",
		reflect.TypeOf((*CfnBrowserCustomLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustomLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogs",
		reflect.TypeOf((*CfnBrowserCustomUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnBrowserCustomUsageLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsDestProps",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsFirehoseProps",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsLogGroupProps",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnBrowserCustomUsageLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBrowserCustomUsageLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnBrowserCustomUsageLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnBrowserCustomUsageLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnBrowserCustomUsageLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnBrowserCustomUsageLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsOutputFormat.S3",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBrowserCustomUsageLogsOutputFormat_S3_JSON,
			"PLAIN": CfnBrowserCustomUsageLogsOutputFormat_S3_PLAIN,
			"W3C": CfnBrowserCustomUsageLogsOutputFormat_S3_W3C,
			"PARQUET": CfnBrowserCustomUsageLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsRecordFields",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnBrowserCustomUsageLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnBrowserCustomUsageLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE": CfnBrowserCustomUsageLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnBrowserCustomUsageLogsRecordFields_ATTRIBUTES,
			"METRICS": CfnBrowserCustomUsageLogsRecordFields_METRICS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomUsageLogsS3Props",
		reflect.TypeOf((*CfnBrowserCustomUsageLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserLogsMixin",
		reflect.TypeOf((*CfnBrowserLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogs",
		reflect.TypeOf((*CfnBrowserUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnBrowserUsageLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsDestProps",
		reflect.TypeOf((*CfnBrowserUsageLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsFirehoseProps",
		reflect.TypeOf((*CfnBrowserUsageLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsLogGroupProps",
		reflect.TypeOf((*CfnBrowserUsageLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat",
		reflect.TypeOf((*CfnBrowserUsageLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnBrowserUsageLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnBrowserUsageLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBrowserUsageLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnBrowserUsageLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnBrowserUsageLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnBrowserUsageLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnBrowserUsageLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnBrowserUsageLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsOutputFormat.S3",
		reflect.TypeOf((*CfnBrowserUsageLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnBrowserUsageLogsOutputFormat_S3_JSON,
			"PLAIN": CfnBrowserUsageLogsOutputFormat_S3_PLAIN,
			"W3C": CfnBrowserUsageLogsOutputFormat_S3_W3C,
			"PARQUET": CfnBrowserUsageLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsRecordFields",
		reflect.TypeOf((*CfnBrowserUsageLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnBrowserUsageLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnBrowserUsageLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE": CfnBrowserUsageLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnBrowserUsageLogsRecordFields_ATTRIBUTES,
			"METRICS": CfnBrowserUsageLogsRecordFields_METRICS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserUsageLogsS3Props",
		reflect.TypeOf((*CfnBrowserUsageLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogs",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsDestProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCodeInterpreterCustomApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCodeInterpreterCustomApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCodeInterpreterCustomApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCodeInterpreterCustomApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCodeInterpreterCustomApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCodeInterpreterCustomApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCodeInterpreterCustomApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCodeInterpreterCustomApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCodeInterpreterCustomApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsRecordFields",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"ACCOUNT_ID": CfnCodeInterpreterCustomApplicationLogsRecordFields_ACCOUNT_ID,
			"REQUEST_ID": CfnCodeInterpreterCustomApplicationLogsRecordFields_REQUEST_ID,
			"TOOL_SESSION_ID": CfnCodeInterpreterCustomApplicationLogsRecordFields_TOOL_SESSION_ID,
			"SPAN_ID": CfnCodeInterpreterCustomApplicationLogsRecordFields_SPAN_ID,
			"TRACE_ID": CfnCodeInterpreterCustomApplicationLogsRecordFields_TRACE_ID,
			"SERVICE_NAME": CfnCodeInterpreterCustomApplicationLogsRecordFields_SERVICE_NAME,
			"OPERATION": CfnCodeInterpreterCustomApplicationLogsRecordFields_OPERATION,
			"REQUEST_PAYLOAD": CfnCodeInterpreterCustomApplicationLogsRecordFields_REQUEST_PAYLOAD,
			"RESPONSE_PAYLOAD": CfnCodeInterpreterCustomApplicationLogsRecordFields_RESPONSE_PAYLOAD,
			"RESOURCE": CfnCodeInterpreterCustomApplicationLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnCodeInterpreterCustomApplicationLogsRecordFields_ATTRIBUTES,
			"TIMEUNIXNANO": CfnCodeInterpreterCustomApplicationLogsRecordFields_TIMEUNIXNANO,
			"SEVERITYNUMBER": CfnCodeInterpreterCustomApplicationLogsRecordFields_SEVERITYNUMBER,
			"SEVERITYTEXT": CfnCodeInterpreterCustomApplicationLogsRecordFields_SEVERITYTEXT,
			"BODY": CfnCodeInterpreterCustomApplicationLogsRecordFields_BODY,
			"TRACEID": CfnCodeInterpreterCustomApplicationLogsRecordFields_TRACEID,
			"SPANID": CfnCodeInterpreterCustomApplicationLogsRecordFields_SPANID,
			"RESOURCE_ARN": CfnCodeInterpreterCustomApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCodeInterpreterCustomApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomApplicationLogsS3Props",
		reflect.TypeOf((*CfnCodeInterpreterCustomApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomLogsMixin",
		reflect.TypeOf((*CfnCodeInterpreterCustomLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustomLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogs",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomUsageLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsDestProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsFirehoseProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsLogGroupProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnCodeInterpreterCustomUsageLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCodeInterpreterCustomUsageLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnCodeInterpreterCustomUsageLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnCodeInterpreterCustomUsageLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnCodeInterpreterCustomUsageLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnCodeInterpreterCustomUsageLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsOutputFormat.S3",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnCodeInterpreterCustomUsageLogsOutputFormat_S3_JSON,
			"PLAIN": CfnCodeInterpreterCustomUsageLogsOutputFormat_S3_PLAIN,
			"W3C": CfnCodeInterpreterCustomUsageLogsOutputFormat_S3_W3C,
			"PARQUET": CfnCodeInterpreterCustomUsageLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsRecordFields",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnCodeInterpreterCustomUsageLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnCodeInterpreterCustomUsageLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE": CfnCodeInterpreterCustomUsageLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnCodeInterpreterCustomUsageLogsRecordFields_ATTRIBUTES,
			"METRICS": CfnCodeInterpreterCustomUsageLogsRecordFields_METRICS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomUsageLogsS3Props",
		reflect.TypeOf((*CfnCodeInterpreterCustomUsageLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogs",
		reflect.TypeOf((*CfnGatewayApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGatewayApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsDestProps",
		reflect.TypeOf((*CfnGatewayApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnGatewayApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnGatewayApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnGatewayApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnGatewayApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnGatewayApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnGatewayApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnGatewayApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnGatewayApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnGatewayApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnGatewayApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnGatewayApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnGatewayApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnGatewayApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnGatewayApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnGatewayApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnGatewayApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsRecordFields",
		reflect.TypeOf((*CfnGatewayApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"BODY": CfnGatewayApplicationLogsRecordFields_BODY,
			"ACCOUNT_ID": CfnGatewayApplicationLogsRecordFields_ACCOUNT_ID,
			"REQUEST_ID": CfnGatewayApplicationLogsRecordFields_REQUEST_ID,
			"TRACE_ID": CfnGatewayApplicationLogsRecordFields_TRACE_ID,
			"SPAN_ID": CfnGatewayApplicationLogsRecordFields_SPAN_ID,
			"RESOURCE": CfnGatewayApplicationLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnGatewayApplicationLogsRecordFields_ATTRIBUTES,
			"TIMEUNIXNANO": CfnGatewayApplicationLogsRecordFields_TIMEUNIXNANO,
			"SEVERITYNUMBER": CfnGatewayApplicationLogsRecordFields_SEVERITYNUMBER,
			"SEVERITYTEXT": CfnGatewayApplicationLogsRecordFields_SEVERITYTEXT,
			"TRACEID": CfnGatewayApplicationLogsRecordFields_TRACEID,
			"SPANID": CfnGatewayApplicationLogsRecordFields_SPANID,
			"RESOURCE_ARN": CfnGatewayApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnGatewayApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayApplicationLogsS3Props",
		reflect.TypeOf((*CfnGatewayApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayLogsMixin",
		reflect.TypeOf((*CfnGatewayLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTraces",
		reflect.TypeOf((*CfnGatewayTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnGatewayTraces{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesDestProps",
		reflect.TypeOf((*CfnGatewayTracesDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesOutputFormat",
		reflect.TypeOf((*CfnGatewayTracesOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnGatewayTracesOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesRecordFields",
		reflect.TypeOf((*CfnGatewayTracesRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TRACE": CfnGatewayTracesRecordFields_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTracesXRayProps",
		reflect.TypeOf((*CfnGatewayTracesXRayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogs",
		reflect.TypeOf((*CfnMemoryApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMemoryApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsDestProps",
		reflect.TypeOf((*CfnMemoryApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnMemoryApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnMemoryApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnMemoryApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMemoryApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnMemoryApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMemoryApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnMemoryApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnMemoryApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnMemoryApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnMemoryApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnMemoryApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnMemoryApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnMemoryApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnMemoryApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnMemoryApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnMemoryApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsRecordFields",
		reflect.TypeOf((*CfnMemoryApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnMemoryApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnMemoryApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"MEMORY_STRATEGY_ID": CfnMemoryApplicationLogsRecordFields_MEMORY_STRATEGY_ID,
			"NAMESPACE": CfnMemoryApplicationLogsRecordFields_NAMESPACE,
			"ACTOR_ID": CfnMemoryApplicationLogsRecordFields_ACTOR_ID,
			"SESSION_ID": CfnMemoryApplicationLogsRecordFields_SESSION_ID,
			"EVENT_ID": CfnMemoryApplicationLogsRecordFields_EVENT_ID,
			"BODY": CfnMemoryApplicationLogsRecordFields_BODY,
			"RESOURCE": CfnMemoryApplicationLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnMemoryApplicationLogsRecordFields_ATTRIBUTES,
			"TIMEUNIXNANO": CfnMemoryApplicationLogsRecordFields_TIMEUNIXNANO,
			"SEVERITYNUMBER": CfnMemoryApplicationLogsRecordFields_SEVERITYNUMBER,
			"SEVERITYTEXT": CfnMemoryApplicationLogsRecordFields_SEVERITYTEXT,
			"TRACEID": CfnMemoryApplicationLogsRecordFields_TRACEID,
			"SPANID": CfnMemoryApplicationLogsRecordFields_SPANID,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryApplicationLogsS3Props",
		reflect.TypeOf((*CfnMemoryApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryLogsMixin",
		reflect.TypeOf((*CfnMemoryLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemoryLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTraces",
		reflect.TypeOf((*CfnMemoryTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnMemoryTraces{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesDestProps",
		reflect.TypeOf((*CfnMemoryTracesDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesOutputFormat",
		reflect.TypeOf((*CfnMemoryTracesOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnMemoryTracesOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesRecordFields",
		reflect.TypeOf((*CfnMemoryTracesRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TRACE": CfnMemoryTracesRecordFields_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryTracesXRayProps",
		reflect.TypeOf((*CfnMemoryTracesXRayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogs",
		reflect.TypeOf((*CfnRuntimeApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsDestProps",
		reflect.TypeOf((*CfnRuntimeApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnRuntimeApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnRuntimeApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnRuntimeApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRuntimeApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRuntimeApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRuntimeApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRuntimeApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRuntimeApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRuntimeApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRuntimeApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRuntimeApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRuntimeApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRuntimeApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRuntimeApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRuntimeApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRuntimeApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsRecordFields",
		reflect.TypeOf((*CfnRuntimeApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"ACCOUNT_ID": CfnRuntimeApplicationLogsRecordFields_ACCOUNT_ID,
			"REQUEST_ID": CfnRuntimeApplicationLogsRecordFields_REQUEST_ID,
			"SESSION_ID": CfnRuntimeApplicationLogsRecordFields_SESSION_ID,
			"SPAN_ID": CfnRuntimeApplicationLogsRecordFields_SPAN_ID,
			"TRACE_ID": CfnRuntimeApplicationLogsRecordFields_TRACE_ID,
			"SERVICE_NAME": CfnRuntimeApplicationLogsRecordFields_SERVICE_NAME,
			"OPERATION": CfnRuntimeApplicationLogsRecordFields_OPERATION,
			"REQUEST_PAYLOAD": CfnRuntimeApplicationLogsRecordFields_REQUEST_PAYLOAD,
			"RESPONSE_PAYLOAD": CfnRuntimeApplicationLogsRecordFields_RESPONSE_PAYLOAD,
			"RESOURCE": CfnRuntimeApplicationLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnRuntimeApplicationLogsRecordFields_ATTRIBUTES,
			"TIMEUNIXNANO": CfnRuntimeApplicationLogsRecordFields_TIMEUNIXNANO,
			"SEVERITYNUMBER": CfnRuntimeApplicationLogsRecordFields_SEVERITYNUMBER,
			"SEVERITYTEXT": CfnRuntimeApplicationLogsRecordFields_SEVERITYTEXT,
			"BODY": CfnRuntimeApplicationLogsRecordFields_BODY,
			"TRACEID": CfnRuntimeApplicationLogsRecordFields_TRACEID,
			"SPANID": CfnRuntimeApplicationLogsRecordFields_SPANID,
			"RESOURCE_ARN": CfnRuntimeApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRuntimeApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeApplicationLogsS3Props",
		reflect.TypeOf((*CfnRuntimeApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeLogsMixin",
		reflect.TypeOf((*CfnRuntimeLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTraces",
		reflect.TypeOf((*CfnRuntimeTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeTraces{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesDestProps",
		reflect.TypeOf((*CfnRuntimeTracesDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesOutputFormat",
		reflect.TypeOf((*CfnRuntimeTracesOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRuntimeTracesOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesRecordFields",
		reflect.TypeOf((*CfnRuntimeTracesRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TRACE": CfnRuntimeTracesRecordFields_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeTracesXRayProps",
		reflect.TypeOf((*CfnRuntimeTracesXRayProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogs",
		reflect.TypeOf((*CfnRuntimeUsageLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnRuntimeUsageLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsDestProps",
		reflect.TypeOf((*CfnRuntimeUsageLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsFirehoseProps",
		reflect.TypeOf((*CfnRuntimeUsageLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsLogGroupProps",
		reflect.TypeOf((*CfnRuntimeUsageLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsOutputFormat",
		reflect.TypeOf((*CfnRuntimeUsageLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnRuntimeUsageLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnRuntimeUsageLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRuntimeUsageLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnRuntimeUsageLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnRuntimeUsageLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnRuntimeUsageLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnRuntimeUsageLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnRuntimeUsageLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsOutputFormat.S3",
		reflect.TypeOf((*CfnRuntimeUsageLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnRuntimeUsageLogsOutputFormat_S3_JSON,
			"PLAIN": CfnRuntimeUsageLogsOutputFormat_S3_PLAIN,
			"W3C": CfnRuntimeUsageLogsOutputFormat_S3_W3C,
			"PARQUET": CfnRuntimeUsageLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsRecordFields",
		reflect.TypeOf((*CfnRuntimeUsageLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ARN": CfnRuntimeUsageLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnRuntimeUsageLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE": CfnRuntimeUsageLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnRuntimeUsageLogsRecordFields_ATTRIBUTES,
			"METRICS": CfnRuntimeUsageLogsRecordFields_METRICS,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeUsageLogsS3Props",
		reflect.TypeOf((*CfnRuntimeUsageLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogs",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnWorkloadIdentityApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsDestProps",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnWorkloadIdentityApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWorkloadIdentityApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnWorkloadIdentityApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnWorkloadIdentityApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnWorkloadIdentityApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnWorkloadIdentityApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnWorkloadIdentityApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnWorkloadIdentityApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnWorkloadIdentityApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnWorkloadIdentityApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsRecordFields",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"REQUEST_ID": CfnWorkloadIdentityApplicationLogsRecordFields_REQUEST_ID,
			"OPERATION_NAME": CfnWorkloadIdentityApplicationLogsRecordFields_OPERATION_NAME,
			"OPERATION_TYPE": CfnWorkloadIdentityApplicationLogsRecordFields_OPERATION_TYPE,
			"START_TIME": CfnWorkloadIdentityApplicationLogsRecordFields_START_TIME,
			"END_TIME": CfnWorkloadIdentityApplicationLogsRecordFields_END_TIME,
			"DURATION_MS": CfnWorkloadIdentityApplicationLogsRecordFields_DURATION_MS,
			"ACCOUNT_ID": CfnWorkloadIdentityApplicationLogsRecordFields_ACCOUNT_ID,
			"REQUEST": CfnWorkloadIdentityApplicationLogsRecordFields_REQUEST,
			"RESPONSE": CfnWorkloadIdentityApplicationLogsRecordFields_RESPONSE,
			"TRACE_ID": CfnWorkloadIdentityApplicationLogsRecordFields_TRACE_ID,
			"SPAN_ID": CfnWorkloadIdentityApplicationLogsRecordFields_SPAN_ID,
			"RESOURCE": CfnWorkloadIdentityApplicationLogsRecordFields_RESOURCE,
			"ATTRIBUTES": CfnWorkloadIdentityApplicationLogsRecordFields_ATTRIBUTES,
			"TIMEUNIXNANO": CfnWorkloadIdentityApplicationLogsRecordFields_TIMEUNIXNANO,
			"SEVERITYNUMBER": CfnWorkloadIdentityApplicationLogsRecordFields_SEVERITYNUMBER,
			"SEVERITYTEXT": CfnWorkloadIdentityApplicationLogsRecordFields_SEVERITYTEXT,
			"BODY": CfnWorkloadIdentityApplicationLogsRecordFields_BODY,
			"TRACEID": CfnWorkloadIdentityApplicationLogsRecordFields_TRACEID,
			"SPANID": CfnWorkloadIdentityApplicationLogsRecordFields_SPANID,
			"RESOURCE_ARN": CfnWorkloadIdentityApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnWorkloadIdentityApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityApplicationLogsS3Props",
		reflect.TypeOf((*CfnWorkloadIdentityApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityLogsMixin",
		reflect.TypeOf((*CfnWorkloadIdentityLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentityLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
