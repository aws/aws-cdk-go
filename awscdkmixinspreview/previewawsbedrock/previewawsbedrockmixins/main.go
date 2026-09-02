package previewawsbedrockmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogs",
		reflect.TypeOf((*CfnAgentAliasApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAgentAliasApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsDestProps",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAgentAliasApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentAliasApplicationLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnAgentAliasApplicationLogsOutputFormat_Firehose_JSON,
			"RAW": CfnAgentAliasApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentAliasApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAgentAliasApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentAliasApplicationLogsOutputFormat_S3_PLAIN,
			"JSON": CfnAgentAliasApplicationLogsOutputFormat_S3_JSON,
			"W3C": CfnAgentAliasApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentAliasApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsRecordFields",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnAgentAliasApplicationLogsRecordFields_TIMESTAMP,
			"AGENT_ALIAS_ARN": CfnAgentAliasApplicationLogsRecordFields_AGENT_ALIAS_ARN,
			"RESOURCE_ID": CfnAgentAliasApplicationLogsRecordFields_RESOURCE_ID,
			"EVENT_TIMESTAMP": CfnAgentAliasApplicationLogsRecordFields_EVENT_TIMESTAMP,
			"EVENT_VERSION": CfnAgentAliasApplicationLogsRecordFields_EVENT_VERSION,
			"OPERATION": CfnAgentAliasApplicationLogsRecordFields_OPERATION,
			"EVENT_TYPE": CfnAgentAliasApplicationLogsRecordFields_EVENT_TYPE,
			"WORKFLOW": CfnAgentAliasApplicationLogsRecordFields_WORKFLOW,
			"LEVEL": CfnAgentAliasApplicationLogsRecordFields_LEVEL,
			"EVENT": CfnAgentAliasApplicationLogsRecordFields_EVENT,
			"AGENT_ID": CfnAgentAliasApplicationLogsRecordFields_AGENT_ID,
			"AGENT_ALIAS_ID": CfnAgentAliasApplicationLogsRecordFields_AGENT_ALIAS_ID,
			"AGENT_ARN": CfnAgentAliasApplicationLogsRecordFields_AGENT_ARN,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasApplicationLogsS3Props",
		reflect.TypeOf((*CfnAgentAliasApplicationLogsS3Props)(nil)).Elem(),
	)
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
			"PLAIN": CfnAgentAliasEventLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnAgentAliasEventLogsOutputFormat_Firehose_JSON,
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
			"PLAIN": CfnAgentAliasEventLogsOutputFormat_S3_PLAIN,
			"JSON": CfnAgentAliasEventLogsOutputFormat_S3_JSON,
			"W3C": CfnAgentAliasEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentAliasEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnAgentAliasEventLogsRecordFields",
		reflect.TypeOf((*CfnAgentAliasEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"RESOURCE_ID": CfnAgentAliasEventLogsRecordFields_RESOURCE_ID,
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
			"PLAIN": CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose_PLAIN,
			"JSON": CfnKnowledgeBaseApplicationLogsOutputFormat_Firehose_JSON,
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
			"PLAIN": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_PLAIN,
			"JSON": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_JSON,
			"W3C": CfnKnowledgeBaseApplicationLogsOutputFormat_S3_W3C,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseApplicationLogsRecordFields",
		reflect.TypeOf((*CfnKnowledgeBaseApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnKnowledgeBaseApplicationLogsRecordFields_TIMESTAMP,
			"RESOURCE_ID": CfnKnowledgeBaseApplicationLogsRecordFields_RESOURCE_ID,
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
			"TIMESTAMP": CfnKnowledgeBaseTracesRecordFields_TIMESTAMP,
			"RESOURCEARN": CfnKnowledgeBaseTracesRecordFields_RESOURCEARN,
			"TRACE": CfnKnowledgeBaseTracesRecordFields_TRACE,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_bedrock.mixins.CfnKnowledgeBaseTracesXRayProps",
		reflect.TypeOf((*CfnKnowledgeBaseTracesXRayProps)(nil)).Elem(),
	)
}
