package previewawsbedrockmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogs",
		reflect.TypeOf((*CfnAgentAliasEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAgentAliasEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsDestProps",
		reflect.TypeOf((*CfnAgentAliasEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsFirehoseProps",
		reflect.TypeOf((*CfnAgentAliasEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsLogGroupProps",
		reflect.TypeOf((*CfnAgentAliasEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat",
		reflect.TypeOf((*CfnAgentAliasEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAgentAliasEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAgentAliasEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentAliasEventLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAgentAliasEventLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAgentAliasEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAgentAliasEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentAliasEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAgentAliasEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAgentAliasEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentAliasEventLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAgentAliasEventLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAgentAliasEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentAliasEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsRecordFields",
		reflect.TypeOf((*CfnAgentAliasEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnAgentAliasEventLogsRecordFields_TIMESTAMP,
			"RESOURCEID": CfnAgentAliasEventLogsRecordFields_RESOURCEID,
			"TRACEID": CfnAgentAliasEventLogsRecordFields_TRACEID,
			"SPANID": CfnAgentAliasEventLogsRecordFields_SPANID,
			"SESSIONID": CfnAgentAliasEventLogsRecordFields_SESSIONID,
			"REQUESTID": CfnAgentAliasEventLogsRecordFields_REQUESTID,
			"OPERATION": CfnAgentAliasEventLogsRecordFields_OPERATION,
			"ATTRIBUTES": CfnAgentAliasEventLogsRecordFields_ATTRIBUTES,
			"BODY": CfnAgentAliasEventLogsRecordFields_BODY,
			"EVENTTYPE": CfnAgentAliasEventLogsRecordFields_EVENTTYPE,
			"EVENTVERSION": CfnAgentAliasEventLogsRecordFields_EVENTVERSION,
			"EVENTNAME": CfnAgentAliasEventLogsRecordFields_EVENTNAME,
			"LEVEL": CfnAgentAliasEventLogsRecordFields_LEVEL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsS3Props",
		reflect.TypeOf((*CfnAgentAliasEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasLogsMixin",
		reflect.TypeOf((*CfnAgentAliasLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentAliasLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogs",
		reflect.TypeOf((*CfnAgentApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAgentApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsDestProps",
		reflect.TypeOf((*CfnAgentApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnAgentApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnAgentApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnAgentApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAgentApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAgentApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAgentApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAgentApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAgentApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAgentApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAgentApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAgentApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAgentApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsRecordFields",
		reflect.TypeOf((*CfnAgentApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"AGENT_ALIAS_ARN": CfnAgentApplicationLogsRecordFields_AGENT_ALIAS_ARN,
			"EVENT_TIMESTAMP": CfnAgentApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"EVENT_VERSION": CfnAgentApplicationLogsRecordFields_EVENT_VERSION,
			"OPERATION": CfnAgentApplicationLogsRecordFields_OPERATION,
			"EVENT_TYPE": CfnAgentApplicationLogsRecordFields_EVENT_TYPE,
			"WORKFLOW": CfnAgentApplicationLogsRecordFields_WORKFLOW,
			"LEVEL": CfnAgentApplicationLogsRecordFields_LEVEL,
			"EVENT": CfnAgentApplicationLogsRecordFields_EVENT,
			"AGENT_ARN": CfnAgentApplicationLogsRecordFields_AGENT_ARN,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentApplicationLogsS3Props",
		reflect.TypeOf((*CfnAgentApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentLogsMixin",
		reflect.TypeOf((*CfnAgentLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogs",
		reflect.TypeOf((*CfnFlowApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnFlowApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsDestProps",
		reflect.TypeOf((*CfnFlowApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnFlowApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnFlowApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnFlowApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnFlowApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnFlowApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFlowApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnFlowApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnFlowApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnFlowApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnFlowApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnFlowApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnFlowApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnFlowApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnFlowApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnFlowApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnFlowApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsRecordFields",
		reflect.TypeOf((*CfnFlowApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENT_TIMESTAMP": CfnFlowApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE_ARN": CfnFlowApplicationLogsRecordFields_RESOURCE_ARN,
			"SCHEMA_VERSION": CfnFlowApplicationLogsRecordFields_SCHEMA_VERSION,
			"FLOW_ID": CfnFlowApplicationLogsRecordFields_FLOW_ID,
			"FLOW_ALIAS_ID": CfnFlowApplicationLogsRecordFields_FLOW_ALIAS_ID,
			"FLOW_VERSION": CfnFlowApplicationLogsRecordFields_FLOW_VERSION,
			"REQUEST_ID": CfnFlowApplicationLogsRecordFields_REQUEST_ID,
			"EXECUTION_ID": CfnFlowApplicationLogsRecordFields_EXECUTION_ID,
			"EVENT_TYPE": CfnFlowApplicationLogsRecordFields_EVENT_TYPE,
			"EVENT": CfnFlowApplicationLogsRecordFields_EVENT,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowApplicationLogsS3Props",
		reflect.TypeOf((*CfnFlowApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnFlowLogsMixin",
		reflect.TypeOf((*CfnFlowLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFlowLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogs",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsDestProps",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnKnowledgeBaseApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnKnowledgeBaseApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_W3C,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsRecordFields",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENT_TIMESTAMP": CfnKnowledgeBaseApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"EVENT": CfnKnowledgeBaseApplicationLogsRecordFields_EVENT,
			"EVENT_VERSION": CfnKnowledgeBaseApplicationLogsRecordFields_EVENT_VERSION,
			"EVENT_TYPE": CfnKnowledgeBaseApplicationLogsRecordFields_EVENT_TYPE,
			"LEVEL": CfnKnowledgeBaseApplicationLogsRecordFields_LEVEL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsS3Props",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseLogsMixin",
		reflect.TypeOf((*CfnKnowledgeBaseLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnKnowledgeBaseLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogs",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseRuntimeLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsDestProps",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsFirehoseProps",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsLogGroupProps",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseRuntimeLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnKnowledgeBaseRuntimeLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnKnowledgeBaseRuntimeLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnKnowledgeBaseRuntimeLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnKnowledgeBaseRuntimeLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnKnowledgeBaseRuntimeLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsOutputFormat.S3",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnKnowledgeBaseRuntimeLogsOutputFormat_S3_JSON,
			"PLAIN": CfnKnowledgeBaseRuntimeLogsOutputFormat_S3_PLAIN,
			"W3C": CfnKnowledgeBaseRuntimeLogsOutputFormat_S3_W3C,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsRecordFields",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENT_TIMESTAMP": CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_TIMESTAMP,
			"EVENT": CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT,
			"EVENT_VERSION": CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_VERSION,
			"EVENT_TYPE": CfnKnowledgeBaseRuntimeLogsRecordFields_EVENT_TYPE,
			"LEVEL": CfnKnowledgeBaseRuntimeLogsRecordFields_LEVEL,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseRuntimeLogsS3Props",
		reflect.TypeOf((*CfnKnowledgeBaseRuntimeLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTraces",
		reflect.TypeOf((*CfnKnowledgeBaseTraces)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toXRay", GoMethod: "ToXRay"},
		},
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseTraces{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesDestProps",
		reflect.TypeOf((*CfnKnowledgeBaseTracesDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesOutputFormat",
		reflect.TypeOf((*CfnKnowledgeBaseTracesOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnKnowledgeBaseTracesOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesRecordFields",
		reflect.TypeOf((*CfnKnowledgeBaseTracesRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TRACE": CfnKnowledgeBaseTracesRecordFields_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesXRayProps",
		reflect.TypeOf((*CfnKnowledgeBaseTracesXRayProps)(nil)).Elem(),
	)
}
