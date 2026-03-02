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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomMixinProps",
		reflect.TypeOf((*CfnBrowserCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBrowserCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.BrowserNetworkConfigurationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.BrowserSigningProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_BrowserSigningProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.RecordingConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_RecordingConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnBrowserCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnBrowserCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomMixinProps",
		reflect.TypeOf((*CfnCodeInterpreterCustomMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCodeInterpreterCustomPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin.CodeInterpreterNetworkConfigurationProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_CodeInterpreterNetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnCodeInterpreterCustomPropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnCodeInterpreterCustomPropsMixin_VpcConfigProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayMixinProps",
		reflect.TypeOf((*CfnGatewayMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin",
		reflect.TypeOf((*CfnGatewayPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.GatewayInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.GatewayProtocolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_GatewayProtocolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.InterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.InterceptorInputConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_InterceptorInputConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.LambdaInterceptorConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_LambdaInterceptorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.MCPGatewayConfigurationProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_MCPGatewayConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayPropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnGatewayPropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetMixinProps",
		reflect.TypeOf((*CfnGatewayTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnGatewayTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiGatewayTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiGatewayToolConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiGatewayToolFilterProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiGatewayToolOverrideProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiGatewayToolOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiKeyCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiKeyCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ApiSchemaConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ApiSchemaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.CredentialProviderConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.CredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_CredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpLambdaTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpLambdaTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpServerTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpServerTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.McpTargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_McpTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.MetadataConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_MetadataConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.OAuthCredentialProviderProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_OAuthCredentialProviderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.S3ConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_S3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.SchemaDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_SchemaDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.TargetConfigurationProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_TargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ToolDefinitionProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnGatewayTargetPropsMixin.ToolSchemaProperty",
		reflect.TypeOf((*CfnGatewayTargetPropsMixin_ToolSchemaProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryMixinProps",
		reflect.TypeOf((*CfnMemoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin",
		reflect.TypeOf((*CfnMemoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMemoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.CustomConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.CustomMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_CustomMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicOverrideReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicOverrideReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.EpisodicReflectionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_EpisodicReflectionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.InvocationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_InvocationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.MemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.MessageBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_MessageBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SelfManagedConfigurationProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SelfManagedConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SemanticOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SemanticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.SummaryOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_SummaryOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TimeBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TimeBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TokenBasedTriggerInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TokenBasedTriggerInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.TriggerConditionInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_TriggerConditionInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceMemoryStrategyProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceMemoryStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideConsolidationConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideConsolidationConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideExtractionConfigurationInputProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideExtractionConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnMemoryPropsMixin.UserPreferenceOverrideProperty",
		reflect.TypeOf((*CfnMemoryPropsMixin_UserPreferenceOverrideProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeEndpointMixinProps",
		reflect.TypeOf((*CfnRuntimeEndpointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeEndpointPropsMixin",
		reflect.TypeOf((*CfnRuntimeEndpointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimeEndpointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimeMixinProps",
		reflect.TypeOf((*CfnRuntimeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin",
		reflect.TypeOf((*CfnRuntimePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRuntimePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.AgentRuntimeArtifactProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AgentRuntimeArtifactProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.AuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.AuthorizingClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_AuthorizingClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.ClaimMatchValueTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_ClaimMatchValueTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CodeConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CodeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.ContainerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_ContainerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CustomClaimValidationTypeProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CustomClaimValidationTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.CustomJWTAuthorizerConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_CustomJWTAuthorizerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.LifecycleConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_LifecycleConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.RequestHeaderConfigurationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_RequestHeaderConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.VpcConfigProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_VpcConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnRuntimePropsMixin.WorkloadIdentityDetailsProperty",
		reflect.TypeOf((*CfnRuntimePropsMixin_WorkloadIdentityDetailsProperty)(nil)).Elem(),
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
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityMixinProps",
		reflect.TypeOf((*CfnWorkloadIdentityMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrockagentcore.mixins.CfnWorkloadIdentityPropsMixin",
		reflect.TypeOf((*CfnWorkloadIdentityPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkloadIdentityPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
