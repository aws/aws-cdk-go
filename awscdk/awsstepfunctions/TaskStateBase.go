package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Task state in the state machine.
//
// Reaching a Task state causes some work to be executed, represented by the
// Task's resource property. Task constructs represent a generic Amazon
// States Language Task.
//
// For some resource types, more specific subclasses of Task may be available
// which are more convenient to use.
type TaskStateBase interface {
	State
	INextable
	Branches() *[]StateGraph
	Comment() *string
	DefaultChoice() State
	SetDefaultChoice(val State)
	// Continuable states of this Chainable.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Iteration() StateGraph
	SetIteration(val StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	Processor() StateGraph
	SetProcessor(val StateGraph)
	ProcessorConfig() *ProcessorConfig
	SetProcessorConfig(val *ProcessorConfig)
	ProcessorMode() ProcessorMode
	SetProcessorMode(val ProcessorMode)
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	StateName() *string
	TaskMetrics() *TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a parallel branch to this state.
	AddBranch(branch StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler IChainable, props *CatchProps) TaskStateBase
	// Add a choice branch to this state.
	AddChoice(condition Condition, next State, options *ChoiceTransitionOptions)
	// Add a item processor to this state.
	AddItemProcessor(processor StateGraph, config *ProcessorConfig)
	// Add a map iterator to this state.
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	AddRetry(props *RetryProps) TaskStateBase
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph StateGraph)
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def State)
	// Make the indicated state the default transition of this state.
	MakeNext(next State)
	// Return the given named metric for this Task.
	// Default: - sum over 5 minutes.
	//
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	// Default: - sum over 5 minutes.
	//
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	// Default: - sum over 5 minutes.
	//
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	// Default: - average over 5 minutes.
	//
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	// Default: - sum over 5 minutes.
	//
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	// Default: - average over 5 minutes.
	//
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	// Default: - sum over 5 minutes.
	//
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	// Default: - sum over 5 minutes.
	//
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	// Default: - average over 5 minutes.
	//
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	// Default: - sum over 5 minutes.
	//
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Continue normal execution with the given state.
	Next(next IChainable) Chain
	// Render parallel branches in ASL JSON format.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	RenderInputOutput() interface{}
	// Render ItemProcessor in ASL JSON format.
	RenderItemProcessor() interface{}
	// Render map iterator in ASL JSON format.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	RenderNextEnd() interface{}
	// Render ResultSelector in ASL JSON format.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	RenderRetryCatch() interface{}
	// Return the Amazon States Language object for this state.
	ToStateJson() *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for TaskStateBase
type jsiiProxy_TaskStateBase struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_TaskStateBase) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) TaskMetrics() *TaskMetricsConfig {
	var returns *TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TaskStateBase) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewTaskStateBase_Override(t TaskStateBase, scope constructs.Construct, id *string, props *TaskStateBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_TaskStateBase)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_TaskStateBase)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_TaskStateBase)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_TaskStateBase)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_TaskStateBase)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func TaskStateBase_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateTaskStateBase_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func TaskStateBase_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateTaskStateBase_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func TaskStateBase_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateTaskStateBase_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func TaskStateBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTaskStateBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func TaskStateBase_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateTaskStateBase_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.TaskStateBase",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (t *jsiiProxy_TaskStateBase) AddBranch(branch StateGraph) {
	if err := t.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addBranch",
		[]interface{}{branch},
	)
}

func (t *jsiiProxy_TaskStateBase) AddCatch(handler IChainable, props *CatchProps) TaskStateBase {
	if err := t.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns TaskStateBase

	_jsii_.Invoke(
		t,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := t.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (t *jsiiProxy_TaskStateBase) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := t.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (t *jsiiProxy_TaskStateBase) AddIterator(iteration StateGraph) {
	if err := t.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (t *jsiiProxy_TaskStateBase) AddPrefix(x *string) {
	if err := t.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addPrefix",
		[]interface{}{x},
	)
}

func (t *jsiiProxy_TaskStateBase) AddRetry(props *RetryProps) TaskStateBase {
	if err := t.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns TaskStateBase

	_jsii_.Invoke(
		t,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) BindToGraph(graph StateGraph) {
	if err := t.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (t *jsiiProxy_TaskStateBase) MakeDefault(def State) {
	if err := t.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"makeDefault",
		[]interface{}{def},
	)
}

func (t *jsiiProxy_TaskStateBase) MakeNext(next State) {
	if err := t.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"makeNext",
		[]interface{}{next},
	)
}

func (t *jsiiProxy_TaskStateBase) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) Next(next IChainable) Chain {
	if err := t.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		t,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskStateBase) WhenBoundToGraph(graph StateGraph) {
	if err := t.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

