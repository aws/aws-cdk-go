package previewawswisdommixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogs",
		reflect.TypeOf((*CfnAssistantEventLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAssistantEventLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsDestProps",
		reflect.TypeOf((*CfnAssistantEventLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsFirehoseProps",
		reflect.TypeOf((*CfnAssistantEventLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsLogGroupProps",
		reflect.TypeOf((*CfnAssistantEventLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat",
		reflect.TypeOf((*CfnAssistantEventLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAssistantEventLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAssistantEventLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAssistantEventLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAssistantEventLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAssistantEventLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAssistantEventLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAssistantEventLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAssistantEventLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAssistantEventLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAssistantEventLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAssistantEventLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAssistantEventLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAssistantEventLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsRecordFields",
		reflect.TypeOf((*CfnAssistantEventLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"ASSISTANT_ID": CfnAssistantEventLogsRecordFields_ASSISTANT_ID,
			"EVENT_TIMESTAMP": CfnAssistantEventLogsRecordFields_EVENT_TIMESTAMP,
			"EVENT_TYPE": CfnAssistantEventLogsRecordFields_EVENT_TYPE,
			"SESSION_ID": CfnAssistantEventLogsRecordFields_SESSION_ID,
			"SESSION_NAME": CfnAssistantEventLogsRecordFields_SESSION_NAME,
			"RECOMMENDATION_ID": CfnAssistantEventLogsRecordFields_RECOMMENDATION_ID,
			"RECOMMENDATION": CfnAssistantEventLogsRecordFields_RECOMMENDATION,
			"IS_RECOMMENDATION_USEFUL": CfnAssistantEventLogsRecordFields_IS_RECOMMENDATION_USEFUL,
			"RELEVANCE_SCORE": CfnAssistantEventLogsRecordFields_RELEVANCE_SCORE,
			"RECOMMENDATION_TITLE": CfnAssistantEventLogsRecordFields_RECOMMENDATION_TITLE,
			"RECOMMENDATION_SOURCES": CfnAssistantEventLogsRecordFields_RECOMMENDATION_SOURCES,
			"INTENT_ID": CfnAssistantEventLogsRecordFields_INTENT_ID,
			"INTENT": CfnAssistantEventLogsRecordFields_INTENT,
			"INTENT_CLICKED": CfnAssistantEventLogsRecordFields_INTENT_CLICKED,
			"UTTERANCE": CfnAssistantEventLogsRecordFields_UTTERANCE,
			"PROMPT": CfnAssistantEventLogsRecordFields_PROMPT,
			"RESPONSE": CfnAssistantEventLogsRecordFields_RESPONSE,
			"SESSION_EVENT_ID": CfnAssistantEventLogsRecordFields_SESSION_EVENT_ID,
			"SESSION_EVENT_IDS": CfnAssistantEventLogsRecordFields_SESSION_EVENT_IDS,
			"ISSUE_PROBABILITY": CfnAssistantEventLogsRecordFields_ISSUE_PROBABILITY,
			"IS_VALID_TRIGGER": CfnAssistantEventLogsRecordFields_IS_VALID_TRIGGER,
			"PROMPT_TYPE": CfnAssistantEventLogsRecordFields_PROMPT_TYPE,
			"COMPLETION": CfnAssistantEventLogsRecordFields_COMPLETION,
			"MODEL_ID": CfnAssistantEventLogsRecordFields_MODEL_ID,
			"CONNECT_USER_ARN": CfnAssistantEventLogsRecordFields_CONNECT_USER_ARN,
			"CONVERSATION_SESSION_DATA": CfnAssistantEventLogsRecordFields_CONVERSATION_SESSION_DATA,
			"SESSION_MESSAGE_ID": CfnAssistantEventLogsRecordFields_SESSION_MESSAGE_ID,
			"PARSED_RESPONSE": CfnAssistantEventLogsRecordFields_PARSED_RESPONSE,
			"ANSWER_ID": CfnAssistantEventLogsRecordFields_ANSWER_ID,
			"GENERATION_ID": CfnAssistantEventLogsRecordFields_GENERATION_ID,
			"AI_AGENT_ID": CfnAssistantEventLogsRecordFields_AI_AGENT_ID,
			"AI_AGENT_NAME": CfnAssistantEventLogsRecordFields_AI_AGENT_NAME,
			"AI_GUARDRAIL_ID": CfnAssistantEventLogsRecordFields_AI_GUARDRAIL_ID,
			"NAME": CfnAssistantEventLogsRecordFields_NAME,
			"AI_GUARDRAIL": CfnAssistantEventLogsRecordFields_AI_GUARDRAIL,
			"CONTENT": CfnAssistantEventLogsRecordFields_CONTENT,
			"ACTION": CfnAssistantEventLogsRecordFields_ACTION,
			"ACTION_REASON": CfnAssistantEventLogsRecordFields_ACTION_REASON,
			"OUTPUTS": CfnAssistantEventLogsRecordFields_OUTPUTS,
			"ASSESSMENTS": CfnAssistantEventLogsRecordFields_ASSESSMENTS,
			"USAGE": CfnAssistantEventLogsRecordFields_USAGE,
			"GUARDRAIL_COVERAGE": CfnAssistantEventLogsRecordFields_GUARDRAIL_COVERAGE,
			"AI_AGENT_VERSION": CfnAssistantEventLogsRecordFields_AI_AGENT_VERSION,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantEventLogsS3Props",
		reflect.TypeOf((*CfnAssistantEventLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_wisdom.mixins.CfnAssistantLogsMixin",
		reflect.TypeOf((*CfnAssistantLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAssistantLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
