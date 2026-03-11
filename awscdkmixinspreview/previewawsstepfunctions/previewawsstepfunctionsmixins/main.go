package previewawsstepfunctionsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogs",
		reflect.TypeOf((*CfnStateMachineExpressLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_CfnStateMachineExpressLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsDestProps",
		reflect.TypeOf((*CfnStateMachineExpressLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsLogGroupProps",
		reflect.TypeOf((*CfnStateMachineExpressLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsOutputFormat",
		reflect.TypeOf((*CfnStateMachineExpressLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnStateMachineExpressLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnStateMachineExpressLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnStateMachineExpressLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnStateMachineExpressLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineExpressLogsRecordFields",
		reflect.TypeOf((*CfnStateMachineExpressLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"DETAILS": CfnStateMachineExpressLogsRecordFields_DETAILS,
			"MAP_RUN_ARN": CfnStateMachineExpressLogsRecordFields_MAP_RUN_ARN,
			"PARENT_EXECUTION_ARN": CfnStateMachineExpressLogsRecordFields_PARENT_EXECUTION_ARN,
			"REDRIVE_COUNT": CfnStateMachineExpressLogsRecordFields_REDRIVE_COUNT,
			"ID": CfnStateMachineExpressLogsRecordFields_ID,
			"TYPE": CfnStateMachineExpressLogsRecordFields_TYPE,
			"PREVIOUS_EVENT_ID": CfnStateMachineExpressLogsRecordFields_PREVIOUS_EVENT_ID,
			"EVENT_TIMESTAMP": CfnStateMachineExpressLogsRecordFields_EVENT_TIMESTAMP,
			"EXECUTION_ARN": CfnStateMachineExpressLogsRecordFields_EXECUTION_ARN,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineLogsMixin",
		reflect.TypeOf((*CfnStateMachineLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStateMachineLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogs",
		reflect.TypeOf((*CfnStateMachineStandardLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_CfnStateMachineStandardLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsDestProps",
		reflect.TypeOf((*CfnStateMachineStandardLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsLogGroupProps",
		reflect.TypeOf((*CfnStateMachineStandardLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsOutputFormat",
		reflect.TypeOf((*CfnStateMachineStandardLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnStateMachineStandardLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnStateMachineStandardLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnStateMachineStandardLogsOutputFormat_LogGroup_PLAIN,
			"JSON": CfnStateMachineStandardLogsOutputFormat_LogGroup_JSON,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_stepfunctions.mixins.CfnStateMachineStandardLogsRecordFields",
		reflect.TypeOf((*CfnStateMachineStandardLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"DETAILS": CfnStateMachineStandardLogsRecordFields_DETAILS,
			"MAP_RUN_ARN": CfnStateMachineStandardLogsRecordFields_MAP_RUN_ARN,
			"PARENT_EXECUTION_ARN": CfnStateMachineStandardLogsRecordFields_PARENT_EXECUTION_ARN,
			"REDRIVE_COUNT": CfnStateMachineStandardLogsRecordFields_REDRIVE_COUNT,
			"ID": CfnStateMachineStandardLogsRecordFields_ID,
			"TYPE": CfnStateMachineStandardLogsRecordFields_TYPE,
			"PREVIOUS_EVENT_ID": CfnStateMachineStandardLogsRecordFields_PREVIOUS_EVENT_ID,
			"EVENT_TIMESTAMP": CfnStateMachineStandardLogsRecordFields_EVENT_TIMESTAMP,
			"EXECUTION_ARN": CfnStateMachineStandardLogsRecordFields_EXECUTION_ARN,
		},
	)
}
