package awsstepfunctionstasks

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AcceleratorClass",
		reflect.TypeOf((*AcceleratorClass)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			return &jsiiProxy_AcceleratorClass{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AcceleratorType",
		reflect.TypeOf((*AcceleratorType)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			return &jsiiProxy_AcceleratorType{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.ActionOnFailure",
		reflect.TypeOf((*ActionOnFailure)(nil)).Elem(),
		map[string]interface{}{
			"TERMINATE_CLUSTER": ActionOnFailure_TERMINATE_CLUSTER,
			"CANCEL_AND_WAIT": ActionOnFailure_CANCEL_AND_WAIT,
			"CONTINUE": ActionOnFailure_CONTINUE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.AlgorithmSpecification",
		reflect.TypeOf((*AlgorithmSpecification)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ApplicationConfiguration",
		reflect.TypeOf((*ApplicationConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.AssembleWith",
		reflect.TypeOf((*AssembleWith)(nil)).Elem(),
		map[string]interface{}{
			"NONE": AssembleWith_NONE,
			"LINE": AssembleWith_LINE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		reflect.TypeOf((*AthenaGetQueryExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AthenaGetQueryExecution{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryExecutionProps",
		reflect.TypeOf((*AthenaGetQueryExecutionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		reflect.TypeOf((*AthenaGetQueryResults)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AthenaGetQueryResults{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResultsProps",
		reflect.TypeOf((*AthenaGetQueryResultsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		reflect.TypeOf((*AthenaStartQueryExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AthenaStartQueryExecution{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaStartQueryExecutionProps",
		reflect.TypeOf((*AthenaStartQueryExecutionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		reflect.TypeOf((*AthenaStopQueryExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_AthenaStopQueryExecution{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaStopQueryExecutionProps",
		reflect.TypeOf((*AthenaStopQueryExecutionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.AuthType",
		reflect.TypeOf((*AuthType)(nil)).Elem(),
		map[string]interface{}{
			"NO_AUTH": AuthType_NO_AUTH,
			"IAM_ROLE": AuthType_IAM_ROLE,
			"RESOURCE_POLICY": AuthType_RESOURCE_POLICY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchContainerOverrides",
		reflect.TypeOf((*BatchContainerOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchJobDependency",
		reflect.TypeOf((*BatchJobDependency)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchStrategy",
		reflect.TypeOf((*BatchStrategy)(nil)).Elem(),
		map[string]interface{}{
			"MULTI_RECORD": BatchStrategy_MULTI_RECORD,
			"SINGLE_RECORD": BatchStrategy_SINGLE_RECORD,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		reflect.TypeOf((*BatchSubmitJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_BatchSubmitJob{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJobProps",
		reflect.TypeOf((*BatchSubmitJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallApiGatewayEndpointBaseProps",
		reflect.TypeOf((*CallApiGatewayEndpointBaseProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		reflect.TypeOf((*CallApiGatewayHttpApiEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "arnForExecuteApi", GoGetter: "ArnForExecuteApi"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberMethod{JsiiMethod: "createPolicyStatements", GoMethod: "CreatePolicyStatements"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_CallApiGatewayHttpApiEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpointProps",
		reflect.TypeOf((*CallApiGatewayHttpApiEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		reflect.TypeOf((*CallApiGatewayRestApiEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberProperty{JsiiProperty: "apiEndpoint", GoGetter: "ApiEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "arnForExecuteApi", GoGetter: "ArnForExecuteApi"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberMethod{JsiiMethod: "createPolicyStatements", GoMethod: "CreatePolicyStatements"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "stageName", GoGetter: "StageName"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_CallApiGatewayRestApiEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpointProps",
		reflect.TypeOf((*CallApiGatewayRestApiEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallAwsService",
		reflect.TypeOf((*CallAwsService)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_CallAwsService{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CallAwsServiceProps",
		reflect.TypeOf((*CallAwsServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.Channel",
		reflect.TypeOf((*Channel)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.Classification",
		reflect.TypeOf((*Classification)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "classificationStatement", GoGetter: "ClassificationStatement"},
		},
		func() interface{} {
			return &jsiiProxy_Classification{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		reflect.TypeOf((*CodeBuildStartBuild)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_CodeBuildStartBuild{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuildProps",
		reflect.TypeOf((*CodeBuildStartBuildProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.CommonEcsRunTaskProps",
		reflect.TypeOf((*CommonEcsRunTaskProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.CompressionType",
		reflect.TypeOf((*CompressionType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": CompressionType_NONE,
			"GZIP": CompressionType_GZIP,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.ContainerDefinition",
		reflect.TypeOf((*ContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_ContainerDefinition{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IContainerDefinition)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ContainerDefinitionConfig",
		reflect.TypeOf((*ContainerDefinitionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ContainerDefinitionOptions",
		reflect.TypeOf((*ContainerDefinitionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ContainerOverride",
		reflect.TypeOf((*ContainerOverride)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ContainerOverrides",
		reflect.TypeOf((*ContainerOverrides)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DataSource",
		reflect.TypeOf((*DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImage",
		reflect.TypeOf((*DockerImage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_DockerImage{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DockerImageConfig",
		reflect.TypeOf((*DockerImageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoAttributeValue",
		reflect.TypeOf((*DynamoAttributeValue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "attributeValue", GoGetter: "AttributeValue"},
			_jsii_.MemberMethod{JsiiMethod: "toObject", GoMethod: "ToObject"},
		},
		func() interface{} {
			return &jsiiProxy_DynamoAttributeValue{}
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoConsumedCapacity",
		reflect.TypeOf((*DynamoConsumedCapacity)(nil)).Elem(),
		map[string]interface{}{
			"INDEXES": DynamoConsumedCapacity_INDEXES,
			"TOTAL": DynamoConsumedCapacity_TOTAL,
			"NONE": DynamoConsumedCapacity_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoDeleteItem",
		reflect.TypeOf((*DynamoDeleteItem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoDeleteItem{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoDeleteItemProps",
		reflect.TypeOf((*DynamoDeleteItemProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		reflect.TypeOf((*DynamoGetItem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoGetItem{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItemProps",
		reflect.TypeOf((*DynamoGetItemProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoItemCollectionMetrics",
		reflect.TypeOf((*DynamoItemCollectionMetrics)(nil)).Elem(),
		map[string]interface{}{
			"SIZE": DynamoItemCollectionMetrics_SIZE,
			"NONE": DynamoItemCollectionMetrics_NONE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoProjectionExpression",
		reflect.TypeOf((*DynamoProjectionExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "atIndex", GoMethod: "AtIndex"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "withAttribute", GoMethod: "WithAttribute"},
		},
		func() interface{} {
			return &jsiiProxy_DynamoProjectionExpression{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoPutItem",
		reflect.TypeOf((*DynamoPutItem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoPutItem{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoPutItemProps",
		reflect.TypeOf((*DynamoPutItemProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoReturnValues",
		reflect.TypeOf((*DynamoReturnValues)(nil)).Elem(),
		map[string]interface{}{
			"NONE": DynamoReturnValues_NONE,
			"ALL_OLD": DynamoReturnValues_ALL_OLD,
			"UPDATED_OLD": DynamoReturnValues_UPDATED_OLD,
			"ALL_NEW": DynamoReturnValues_ALL_NEW,
			"UPDATED_NEW": DynamoReturnValues_UPDATED_NEW,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoUpdateItem",
		reflect.TypeOf((*DynamoUpdateItem)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_DynamoUpdateItem{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoUpdateItemProps",
		reflect.TypeOf((*DynamoUpdateItemProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		reflect.TypeOf((*EcsEc2LaunchTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EcsEc2LaunchTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsLaunchTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsEc2LaunchTargetOptions",
		reflect.TypeOf((*EcsEc2LaunchTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		reflect.TypeOf((*EcsFargateLaunchTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_EcsFargateLaunchTarget{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEcsLaunchTarget)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsFargateLaunchTargetOptions",
		reflect.TypeOf((*EcsFargateLaunchTargetOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsLaunchTargetConfig",
		reflect.TypeOf((*EcsLaunchTargetConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsRunTask",
		reflect.TypeOf((*EcsRunTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EcsRunTask{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EcsRunTaskProps",
		reflect.TypeOf((*EcsRunTaskProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EksCall",
		reflect.TypeOf((*EksCall)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EksCall{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EksCallProps",
		reflect.TypeOf((*EksCallProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EksClusterInput",
		reflect.TypeOf((*EksClusterInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
		},
		func() interface{} {
			return &jsiiProxy_EksClusterInput{}
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStep",
		reflect.TypeOf((*EmrAddStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrAddStep{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrAddStepProps",
		reflect.TypeOf((*EmrAddStepProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCancelStep",
		reflect.TypeOf((*EmrCancelStep)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrCancelStep{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCancelStepProps",
		reflect.TypeOf((*EmrCancelStepProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersCreateVirtualCluster",
		reflect.TypeOf((*EmrContainersCreateVirtualCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrContainersCreateVirtualCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersCreateVirtualClusterProps",
		reflect.TypeOf((*EmrContainersCreateVirtualClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersDeleteVirtualCluster",
		reflect.TypeOf((*EmrContainersDeleteVirtualCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrContainersDeleteVirtualCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersDeleteVirtualClusterProps",
		reflect.TypeOf((*EmrContainersDeleteVirtualClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersStartJobRun",
		reflect.TypeOf((*EmrContainersStartJobRun)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrContainersStartJobRun{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrContainersStartJobRunProps",
		reflect.TypeOf((*EmrContainersStartJobRunProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster",
		reflect.TypeOf((*EmrCreateCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberProperty{JsiiProperty: "autoScalingRole", GoGetter: "AutoScalingRole"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "clusterRole", GoGetter: "ClusterRole"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "serviceRole", GoGetter: "ServiceRole"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrCreateCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ApplicationConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_ApplicationConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.AutoScalingPolicyProperty",
		reflect.TypeOf((*EmrCreateCluster_AutoScalingPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.BootstrapActionConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_BootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmComparisonOperator",
		reflect.TypeOf((*EmrCreateCluster_CloudWatchAlarmComparisonOperator)(nil)).Elem(),
		map[string]interface{}{
			"GREATER_THAN_OR_EQUAL": EmrCreateCluster_CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL,
			"GREATER_THAN": EmrCreateCluster_CloudWatchAlarmComparisonOperator_GREATER_THAN,
			"LESS_THAN": EmrCreateCluster_CloudWatchAlarmComparisonOperator_LESS_THAN,
			"LESS_THAN_OR_EQUAL": EmrCreateCluster_CloudWatchAlarmComparisonOperator_LESS_THAN_OR_EQUAL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmDefinitionProperty",
		reflect.TypeOf((*EmrCreateCluster_CloudWatchAlarmDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmStatistic",
		reflect.TypeOf((*EmrCreateCluster_CloudWatchAlarmStatistic)(nil)).Elem(),
		map[string]interface{}{
			"SAMPLE_COUNT": EmrCreateCluster_CloudWatchAlarmStatistic_SAMPLE_COUNT,
			"AVERAGE": EmrCreateCluster_CloudWatchAlarmStatistic_AVERAGE,
			"SUM": EmrCreateCluster_CloudWatchAlarmStatistic_SUM,
			"MINIMUM": EmrCreateCluster_CloudWatchAlarmStatistic_MINIMUM,
			"MAXIMUM": EmrCreateCluster_CloudWatchAlarmStatistic_MAXIMUM,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.CloudWatchAlarmUnit",
		reflect.TypeOf((*EmrCreateCluster_CloudWatchAlarmUnit)(nil)).Elem(),
		map[string]interface{}{
			"NONE": EmrCreateCluster_CloudWatchAlarmUnit_NONE,
			"SECONDS": EmrCreateCluster_CloudWatchAlarmUnit_SECONDS,
			"MICRO_SECONDS": EmrCreateCluster_CloudWatchAlarmUnit_MICRO_SECONDS,
			"MILLI_SECONDS": EmrCreateCluster_CloudWatchAlarmUnit_MILLI_SECONDS,
			"BYTES": EmrCreateCluster_CloudWatchAlarmUnit_BYTES,
			"KILO_BYTES": EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES,
			"MEGA_BYTES": EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES,
			"GIGA_BYTES": EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES,
			"TERA_BYTES": EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES,
			"BITS": EmrCreateCluster_CloudWatchAlarmUnit_BITS,
			"KILO_BITS": EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS,
			"MEGA_BITS": EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS,
			"GIGA_BITS": EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS,
			"TERA_BITS": EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS,
			"PERCENT": EmrCreateCluster_CloudWatchAlarmUnit_PERCENT,
			"COUNT": EmrCreateCluster_CloudWatchAlarmUnit_COUNT,
			"BYTES_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_BYTES_PER_SECOND,
			"KILO_BYTES_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES_PER_SECOND,
			"MEGA_BYTES_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES_PER_SECOND,
			"GIGA_BYTES_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES_PER_SECOND,
			"TERA_BYTES_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES_PER_SECOND,
			"BITS_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_BITS_PER_SECOND,
			"KILO_BITS_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS_PER_SECOND,
			"MEGA_BITS_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS_PER_SECOND,
			"GIGA_BITS_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS_PER_SECOND,
			"TERA_BITS_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS_PER_SECOND,
			"COUNT_PER_SECOND": EmrCreateCluster_CloudWatchAlarmUnit_COUNT_PER_SECOND,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ConfigurationProperty",
		reflect.TypeOf((*EmrCreateCluster_ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_EbsBlockDeviceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.EbsBlockDeviceVolumeType",
		reflect.TypeOf((*EmrCreateCluster_EbsBlockDeviceVolumeType)(nil)).Elem(),
		map[string]interface{}{
			"GP2": EmrCreateCluster_EbsBlockDeviceVolumeType_GP2,
			"IO1": EmrCreateCluster_EbsBlockDeviceVolumeType_IO1,
			"STANDARD": EmrCreateCluster_EbsBlockDeviceVolumeType_STANDARD,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.EbsConfigurationProperty",
		reflect.TypeOf((*EmrCreateCluster_EbsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.EmrClusterScaleDownBehavior",
		reflect.TypeOf((*EmrCreateCluster_EmrClusterScaleDownBehavior)(nil)).Elem(),
		map[string]interface{}{
			"TERMINATE_AT_INSTANCE_HOUR": EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_INSTANCE_HOUR,
			"TERMINATE_AT_TASK_COMPLETION": EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_TASK_COMPLETION,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceFleetConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_InstanceFleetConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceFleetProvisioningSpecificationsProperty",
		reflect.TypeOf((*EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceGroupConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_InstanceGroupConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceMarket",
		reflect.TypeOf((*EmrCreateCluster_InstanceMarket)(nil)).Elem(),
		map[string]interface{}{
			"ON_DEMAND": EmrCreateCluster_InstanceMarket_ON_DEMAND,
			"SPOT": EmrCreateCluster_InstanceMarket_SPOT,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceRoleType",
		reflect.TypeOf((*EmrCreateCluster_InstanceRoleType)(nil)).Elem(),
		map[string]interface{}{
			"MASTER": EmrCreateCluster_InstanceRoleType_MASTER,
			"CORE": EmrCreateCluster_InstanceRoleType_CORE,
			"TASK": EmrCreateCluster_InstanceRoleType_TASK,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstanceTypeConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_InstanceTypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.InstancesConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_InstancesConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.KerberosAttributesProperty",
		reflect.TypeOf((*EmrCreateCluster_KerberosAttributesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.MetricDimensionProperty",
		reflect.TypeOf((*EmrCreateCluster_MetricDimensionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.PlacementTypeProperty",
		reflect.TypeOf((*EmrCreateCluster_PlacementTypeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScalingActionProperty",
		reflect.TypeOf((*EmrCreateCluster_ScalingActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScalingAdjustmentType",
		reflect.TypeOf((*EmrCreateCluster_ScalingAdjustmentType)(nil)).Elem(),
		map[string]interface{}{
			"CHANGE_IN_CAPACITY": EmrCreateCluster_ScalingAdjustmentType_CHANGE_IN_CAPACITY,
			"PERCENT_CHANGE_IN_CAPACITY": EmrCreateCluster_ScalingAdjustmentType_PERCENT_CHANGE_IN_CAPACITY,
			"EXACT_CAPACITY": EmrCreateCluster_ScalingAdjustmentType_EXACT_CAPACITY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScalingConstraintsProperty",
		reflect.TypeOf((*EmrCreateCluster_ScalingConstraintsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScalingRuleProperty",
		reflect.TypeOf((*EmrCreateCluster_ScalingRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScalingTriggerProperty",
		reflect.TypeOf((*EmrCreateCluster_ScalingTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.ScriptBootstrapActionConfigProperty",
		reflect.TypeOf((*EmrCreateCluster_ScriptBootstrapActionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.SimpleScalingPolicyConfigurationProperty",
		reflect.TypeOf((*EmrCreateCluster_SimpleScalingPolicyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.SpotAllocationStrategy",
		reflect.TypeOf((*EmrCreateCluster_SpotAllocationStrategy)(nil)).Elem(),
		map[string]interface{}{
			"CAPACITY_OPTIMIZED": EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.SpotProvisioningSpecificationProperty",
		reflect.TypeOf((*EmrCreateCluster_SpotProvisioningSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.SpotTimeoutAction",
		reflect.TypeOf((*EmrCreateCluster_SpotTimeoutAction)(nil)).Elem(),
		map[string]interface{}{
			"SWITCH_TO_ON_DEMAND": EmrCreateCluster_SpotTimeoutAction_SWITCH_TO_ON_DEMAND,
			"TERMINATE_CLUSTER": EmrCreateCluster_SpotTimeoutAction_TERMINATE_CLUSTER,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateCluster.VolumeSpecificationProperty",
		reflect.TypeOf((*EmrCreateCluster_VolumeSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrCreateClusterProps",
		reflect.TypeOf((*EmrCreateClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		reflect.TypeOf((*EmrModifyInstanceFleetByName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrModifyInstanceFleetByName{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceFleetByNameProps",
		reflect.TypeOf((*EmrModifyInstanceFleetByNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		reflect.TypeOf((*EmrModifyInstanceGroupByName)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrModifyInstanceGroupByName{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName.InstanceGroupModifyConfigProperty",
		reflect.TypeOf((*EmrModifyInstanceGroupByName_InstanceGroupModifyConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName.InstanceResizePolicyProperty",
		reflect.TypeOf((*EmrModifyInstanceGroupByName_InstanceResizePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName.ShrinkPolicyProperty",
		reflect.TypeOf((*EmrModifyInstanceGroupByName_ShrinkPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrModifyInstanceGroupByNameProps",
		reflect.TypeOf((*EmrModifyInstanceGroupByNameProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		reflect.TypeOf((*EmrSetClusterTerminationProtection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrSetClusterTerminationProtection{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrSetClusterTerminationProtectionProps",
		reflect.TypeOf((*EmrSetClusterTerminationProtectionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrTerminateCluster",
		reflect.TypeOf((*EmrTerminateCluster)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EmrTerminateCluster{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EmrTerminateClusterProps",
		reflect.TypeOf((*EmrTerminateClusterProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EncryptionConfiguration",
		reflect.TypeOf((*EncryptionConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.EncryptionOption",
		reflect.TypeOf((*EncryptionOption)(nil)).Elem(),
		map[string]interface{}{
			"S3_MANAGED": EncryptionOption_S3_MANAGED,
			"KMS": EncryptionOption_KMS,
			"CLIENT_SIDE_KMS": EncryptionOption_CLIENT_SIDE_KMS,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EvaluateExpression",
		reflect.TypeOf((*EvaluateExpression)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EvaluateExpression{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EvaluateExpressionProps",
		reflect.TypeOf((*EvaluateExpressionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.EventBridgePutEvents",
		reflect.TypeOf((*EventBridgePutEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_EventBridgePutEvents{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EventBridgePutEventsEntry",
		reflect.TypeOf((*EventBridgePutEventsEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.EventBridgePutEventsProps",
		reflect.TypeOf((*EventBridgePutEventsProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		reflect.TypeOf((*GlueDataBrewStartJobRun)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_GlueDataBrewStartJobRun{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRunProps",
		reflect.TypeOf((*GlueDataBrewStartJobRunProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		reflect.TypeOf((*GlueStartJobRun)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_GlueStartJobRun{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRunProps",
		reflect.TypeOf((*GlueStartJobRunProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpMethod",
		reflect.TypeOf((*HttpMethod)(nil)).Elem(),
		map[string]interface{}{
			"GET": HttpMethod_GET,
			"POST": HttpMethod_POST,
			"PUT": HttpMethod_PUT,
			"DELETE": HttpMethod_DELETE,
			"PATCH": HttpMethod_PATCH,
			"HEAD": HttpMethod_HEAD,
			"OPTIONS": HttpMethod_OPTIONS,
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpMethods",
		reflect.TypeOf((*HttpMethods)(nil)).Elem(),
		map[string]interface{}{
			"GET": HttpMethods_GET,
			"POST": HttpMethods_POST,
			"PUT": HttpMethods_PUT,
			"DELETE": HttpMethods_DELETE,
			"PATCH": HttpMethods_PATCH,
			"HEAD": HttpMethods_HEAD,
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_stepfunctions_tasks.IContainerDefinition",
		reflect.TypeOf((*IContainerDefinition)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IContainerDefinition{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_stepfunctions_tasks.IEcsLaunchTarget",
		reflect.TypeOf((*IEcsLaunchTarget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_IEcsLaunchTarget{}
		},
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_stepfunctions_tasks.ISageMakerTask",
		reflect.TypeOf((*ISageMakerTask)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
		},
		func() interface{} {
			j := jsiiProxy_ISageMakerTask{}
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.InputMode",
		reflect.TypeOf((*InputMode)(nil)).Elem(),
		map[string]interface{}{
			"PIPE": InputMode_PIPE,
			"FILE": InputMode_FILE,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.JobDependency",
		reflect.TypeOf((*JobDependency)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.JobDriver",
		reflect.TypeOf((*JobDriver)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvocationType",
		reflect.TypeOf((*LambdaInvocationType)(nil)).Elem(),
		map[string]interface{}{
			"REQUEST_RESPONSE": LambdaInvocationType_REQUEST_RESPONSE,
			"EVENT": LambdaInvocationType_EVENT,
			"DRY_RUN": LambdaInvocationType_DRY_RUN,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
		reflect.TypeOf((*LambdaInvoke)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaInvoke{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvokeProps",
		reflect.TypeOf((*LambdaInvokeProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.LaunchTargetBindOptions",
		reflect.TypeOf((*LaunchTargetBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.MessageAttribute",
		reflect.TypeOf((*MessageAttribute)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.MessageAttributeDataType",
		reflect.TypeOf((*MessageAttributeDataType)(nil)).Elem(),
		map[string]interface{}{
			"STRING": MessageAttributeDataType_STRING,
			"STRING_ARRAY": MessageAttributeDataType_STRING_ARRAY,
			"NUMBER": MessageAttributeDataType_NUMBER,
			"BINARY": MessageAttributeDataType_BINARY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.MetricDefinition",
		reflect.TypeOf((*MetricDefinition)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.Mode",
		reflect.TypeOf((*Mode)(nil)).Elem(),
		map[string]interface{}{
			"SINGLE_MODEL": Mode_SINGLE_MODEL,
			"MULTI_MODEL": Mode_MULTI_MODEL,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ModelClientOptions",
		reflect.TypeOf((*ModelClientOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.Monitoring",
		reflect.TypeOf((*Monitoring)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.OutputDataConfig",
		reflect.TypeOf((*OutputDataConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ProductionVariant",
		reflect.TypeOf((*ProductionVariant)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.QueryExecutionContext",
		reflect.TypeOf((*QueryExecutionContext)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.RecordWrapperType",
		reflect.TypeOf((*RecordWrapperType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": RecordWrapperType_NONE,
			"RECORD_IO": RecordWrapperType_RECORD_IO,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.ReleaseLabel",
		reflect.TypeOf((*ReleaseLabel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
		},
		func() interface{} {
			return &jsiiProxy_ReleaseLabel{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ResourceConfig",
		reflect.TypeOf((*ResourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ResultConfiguration",
		reflect.TypeOf((*ResultConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3DataDistributionType",
		reflect.TypeOf((*S3DataDistributionType)(nil)).Elem(),
		map[string]interface{}{
			"FULLY_REPLICATED": S3DataDistributionType_FULLY_REPLICATED,
			"SHARDED_BY_S3_KEY": S3DataDistributionType_SHARDED_BY_S3_KEY,
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3DataSource",
		reflect.TypeOf((*S3DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3DataType",
		reflect.TypeOf((*S3DataType)(nil)).Elem(),
		map[string]interface{}{
			"MANIFEST_FILE": S3DataType_MANIFEST_FILE,
			"S3_PREFIX": S3DataType_S3_PREFIX,
			"AUGMENTED_MANIFEST_FILE": S3DataType_AUGMENTED_MANIFEST_FILE,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3Location",
		reflect.TypeOf((*S3Location)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			return &jsiiProxy_S3Location{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3LocationBindOptions",
		reflect.TypeOf((*S3LocationBindOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.S3LocationConfig",
		reflect.TypeOf((*S3LocationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		reflect.TypeOf((*SageMakerCreateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerCreateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		reflect.TypeOf((*SageMakerCreateEndpointConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerCreateEndpointConfig{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateEndpointConfigProps",
		reflect.TypeOf((*SageMakerCreateEndpointConfigProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateEndpointProps",
		reflect.TypeOf((*SageMakerCreateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateModel",
		reflect.TypeOf((*SageMakerCreateModel)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerCreateModel{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateModelProps",
		reflect.TypeOf((*SageMakerCreateModelProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		reflect.TypeOf((*SageMakerCreateTrainingJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "addSecurityGroup", GoMethod: "AddSecurityGroup"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "connections", GoGetter: "Connections"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "grantPrincipal", GoGetter: "GrantPrincipal"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerCreateTrainingJob{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			_jsii_.InitJsiiProxy(&j.Type__awsec2IConnectable)
			_jsii_.InitJsiiProxy(&j.Type__awsiamIGrantable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTrainingJobProps",
		reflect.TypeOf((*SageMakerCreateTrainingJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		reflect.TypeOf((*SageMakerCreateTransformJob)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "role", GoGetter: "Role"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerCreateTransformJob{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJobProps",
		reflect.TypeOf((*SageMakerCreateTransformJobProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		reflect.TypeOf((*SageMakerUpdateEndpoint)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SageMakerUpdateEndpoint{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerUpdateEndpointProps",
		reflect.TypeOf((*SageMakerUpdateEndpointProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.ShuffleConfig",
		reflect.TypeOf((*ShuffleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SnsPublish",
		reflect.TypeOf((*SnsPublish)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SnsPublish{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SnsPublishProps",
		reflect.TypeOf((*SnsPublishProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SparkSubmitJobDriver",
		reflect.TypeOf((*SparkSubmitJobDriver)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"aws-cdk-lib.aws_stepfunctions_tasks.SplitType",
		reflect.TypeOf((*SplitType)(nil)).Elem(),
		map[string]interface{}{
			"NONE": SplitType_NONE,
			"LINE": SplitType_LINE,
			"RECORD_IO": SplitType_RECORD_IO,
			"TF_RECORD": SplitType_TF_RECORD,
		},
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		reflect.TypeOf((*SqsSendMessage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_SqsSendMessage{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessageProps",
		reflect.TypeOf((*SqsSendMessageProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		reflect.TypeOf((*StepFunctionsInvokeActivity)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_StepFunctionsInvokeActivity{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.StepFunctionsInvokeActivityProps",
		reflect.TypeOf((*StepFunctionsInvokeActivityProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		reflect.TypeOf((*StepFunctionsStartExecution)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addBranch", GoMethod: "AddBranch"},
			_jsii_.MemberMethod{JsiiMethod: "addCatch", GoMethod: "AddCatch"},
			_jsii_.MemberMethod{JsiiMethod: "addChoice", GoMethod: "AddChoice"},
			_jsii_.MemberMethod{JsiiMethod: "addIterator", GoMethod: "AddIterator"},
			_jsii_.MemberMethod{JsiiMethod: "addPrefix", GoMethod: "AddPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "addRetry", GoMethod: "AddRetry"},
			_jsii_.MemberMethod{JsiiMethod: "bindToGraph", GoMethod: "BindToGraph"},
			_jsii_.MemberProperty{JsiiProperty: "branches", GoGetter: "Branches"},
			_jsii_.MemberProperty{JsiiProperty: "comment", GoGetter: "Comment"},
			_jsii_.MemberProperty{JsiiProperty: "defaultChoice", GoGetter: "DefaultChoice"},
			_jsii_.MemberProperty{JsiiProperty: "endStates", GoGetter: "EndStates"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "inputPath", GoGetter: "InputPath"},
			_jsii_.MemberProperty{JsiiProperty: "iteration", GoGetter: "Iteration"},
			_jsii_.MemberMethod{JsiiMethod: "makeDefault", GoMethod: "MakeDefault"},
			_jsii_.MemberMethod{JsiiMethod: "makeNext", GoMethod: "MakeNext"},
			_jsii_.MemberMethod{JsiiMethod: "metric", GoMethod: "Metric"},
			_jsii_.MemberMethod{JsiiMethod: "metricFailed", GoMethod: "MetricFailed"},
			_jsii_.MemberMethod{JsiiMethod: "metricHeartbeatTimedOut", GoMethod: "MetricHeartbeatTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "metricRunTime", GoMethod: "MetricRunTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduled", GoMethod: "MetricScheduled"},
			_jsii_.MemberMethod{JsiiMethod: "metricScheduleTime", GoMethod: "MetricScheduleTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricStarted", GoMethod: "MetricStarted"},
			_jsii_.MemberMethod{JsiiMethod: "metricSucceeded", GoMethod: "MetricSucceeded"},
			_jsii_.MemberMethod{JsiiMethod: "metricTime", GoMethod: "MetricTime"},
			_jsii_.MemberMethod{JsiiMethod: "metricTimedOut", GoMethod: "MetricTimedOut"},
			_jsii_.MemberMethod{JsiiMethod: "next", GoMethod: "Next"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "outputPath", GoGetter: "OutputPath"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberMethod{JsiiMethod: "renderBranches", GoMethod: "RenderBranches"},
			_jsii_.MemberMethod{JsiiMethod: "renderChoices", GoMethod: "RenderChoices"},
			_jsii_.MemberMethod{JsiiMethod: "renderInputOutput", GoMethod: "RenderInputOutput"},
			_jsii_.MemberMethod{JsiiMethod: "renderIterator", GoMethod: "RenderIterator"},
			_jsii_.MemberMethod{JsiiMethod: "renderNextEnd", GoMethod: "RenderNextEnd"},
			_jsii_.MemberMethod{JsiiMethod: "renderResultSelector", GoMethod: "RenderResultSelector"},
			_jsii_.MemberMethod{JsiiMethod: "renderRetryCatch", GoMethod: "RenderRetryCatch"},
			_jsii_.MemberProperty{JsiiProperty: "resultPath", GoGetter: "ResultPath"},
			_jsii_.MemberProperty{JsiiProperty: "resultSelector", GoGetter: "ResultSelector"},
			_jsii_.MemberProperty{JsiiProperty: "startState", GoGetter: "StartState"},
			_jsii_.MemberProperty{JsiiProperty: "stateId", GoGetter: "StateId"},
			_jsii_.MemberProperty{JsiiProperty: "taskMetrics", GoGetter: "TaskMetrics"},
			_jsii_.MemberProperty{JsiiProperty: "taskPolicies", GoGetter: "TaskPolicies"},
			_jsii_.MemberMethod{JsiiMethod: "toStateJson", GoMethod: "ToStateJson"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "validateState", GoMethod: "ValidateState"},
			_jsii_.MemberMethod{JsiiMethod: "whenBoundToGraph", GoMethod: "WhenBoundToGraph"},
		},
		func() interface{} {
			j := jsiiProxy_StepFunctionsStartExecution{}
			_jsii_.InitJsiiProxy(&j.Type__awsstepfunctionsTaskStateBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.StepFunctionsStartExecutionProps",
		reflect.TypeOf((*StepFunctionsStartExecutionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.StoppingCondition",
		reflect.TypeOf((*StoppingCondition)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TaskEnvironmentVariable",
		reflect.TypeOf((*TaskEnvironmentVariable)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TransformDataSource",
		reflect.TypeOf((*TransformDataSource)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TransformInput",
		reflect.TypeOf((*TransformInput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TransformOutput",
		reflect.TypeOf((*TransformOutput)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TransformResources",
		reflect.TypeOf((*TransformResources)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.TransformS3DataSource",
		reflect.TypeOf((*TransformS3DataSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_stepfunctions_tasks.VirtualClusterInput",
		reflect.TypeOf((*VirtualClusterInput)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
		},
		func() interface{} {
			return &jsiiProxy_VirtualClusterInput{}
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_stepfunctions_tasks.VpcConfig",
		reflect.TypeOf((*VpcConfig)(nil)).Elem(),
	)
}
