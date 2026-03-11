package previewawsdevopsagentmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogs",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toFirehose", GoMethod: "ToFirehose"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnAgentSpaceApplicationLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsDestProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsFirehoseProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsFirehoseProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsLogGroupProps",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnAgentSpaceApplicationLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.Firehose",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_Firehose)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_JSON,
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_PLAIN,
			"RAW": CfnAgentSpaceApplicationLogsOutputFormat_Firehose_RAW,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsOutputFormat.S3",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"JSON": CfnAgentSpaceApplicationLogsOutputFormat_S3_JSON,
			"PLAIN": CfnAgentSpaceApplicationLogsOutputFormat_S3_PLAIN,
			"W3C": CfnAgentSpaceApplicationLogsOutputFormat_S3_W3C,
			"PARQUET": CfnAgentSpaceApplicationLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsRecordFields",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"OPTIONAL_ACCOUNT_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ACCOUNT_ID,
			"OPTIONAL_AGENT_SPACE_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_AGENT_SPACE_ID,
			"OPTIONAL_LEVEL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_LEVEL,
			"OPTIONAL_ASSOCIATION_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ASSOCIATION_ID,
			"OPTIONAL_STATUS": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_STATUS,
			"OPTIONAL_WEBHOOK_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_WEBHOOK_ID,
			"OPTIONAL_MCP_ENDPOINT_URL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_MCP_ENDPOINT_URL,
			"OPTIONAL_SERVICE_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_TYPE,
			"OPTIONAL_SERVICE_ENDPOINT_URL": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_ENDPOINT_URL,
			"OPTIONAL_SERVICE_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_SERVICE_ID,
			"OPTIONAL_REQUEST_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_REQUEST_ID,
			"OPTIONAL_OPERATION": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_OPERATION,
			"OPTIONAL_TASK_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_TASK_TYPE,
			"OPTIONAL_TASK_ID": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_TASK_ID,
			"OPTIONAL_REFERENCE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_REFERENCE,
			"OPTIONAL_ERROR_TYPE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ERROR_TYPE,
			"OPTIONAL_ERROR_MESSAGE": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_ERROR_MESSAGE,
			"OPTIONAL_DETAILS": CfnAgentSpaceApplicationLogsRecordFields_OPTIONAL_DETAILS,
			"RESOURCE_ARN": CfnAgentSpaceApplicationLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnAgentSpaceApplicationLogsRecordFields_EVENT_TIMESTAMP,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceApplicationLogsS3Props",
		reflect.TypeOf((*CfnAgentSpaceApplicationLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_devopsagent.mixins.CfnAgentSpaceLogsMixin",
		reflect.TypeOf((*CfnAgentSpaceLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentSpaceLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}
