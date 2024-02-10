package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctionstasks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Get an Athena Query Results as a Task.
//
// Example:
//   getQueryResultsJob := tasks.NewAthenaGetQueryResults(this, jsii.String("Get Query Results"), &AthenaGetQueryResultsProps{
//   	QueryExecutionId: sfn.JsonPath_StringAt(jsii.String("$.QueryExecutionId")),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
type AthenaGetQueryResults interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	// Continuable states of this Chainable.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	Processor() awsstepfunctions.StateGraph
	SetProcessor(val awsstepfunctions.StateGraph)
	ProcessorConfig() *awsstepfunctions.ProcessorConfig
	SetProcessorConfig(val *awsstepfunctions.ProcessorConfig)
	ProcessorMode() awsstepfunctions.ProcessorMode
	SetProcessorMode(val awsstepfunctions.ProcessorMode)
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	StateName() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a parallel branch to this state.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	// Add a choice branch to this state.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions)
	// Add a item processor to this state.
	AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig)
	// Add a map iterator to this state.
	AddIterator(iteration awsstepfunctions.StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph awsstepfunctions.StateGraph)
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def awsstepfunctions.State)
	// Make the indicated state the default transition of this state.
	MakeNext(next awsstepfunctions.State)
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
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
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
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaGetQueryResults
type jsiiProxy_AthenaGetQueryResults struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaGetQueryResults) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewAthenaGetQueryResults(scope constructs.Construct, id *string, props *AthenaGetQueryResultsProps) AthenaGetQueryResults {
	_init_.Initialize()

	if err := validateNewAthenaGetQueryResultsParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AthenaGetQueryResults{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewAthenaGetQueryResults_Override(a AthenaGetQueryResults, scope constructs.Construct, id *string, props *AthenaGetQueryResultsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func AthenaGetQueryResults_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateAthenaGetQueryResults_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func AthenaGetQueryResults_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAthenaGetQueryResults_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func AthenaGetQueryResults_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateAthenaGetQueryResults_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
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
func AthenaGetQueryResults_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAthenaGetQueryResults_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func AthenaGetQueryResults_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateAthenaGetQueryResults_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := a.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := a.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := a.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := a.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := a.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddPrefix(x *string) {
	if err := a.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := a.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) MakeDefault(def awsstepfunctions.State) {
	if err := a.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) MakeNext(next awsstepfunctions.State) {
	if err := a.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

func (a *jsiiProxy_AthenaGetQueryResults) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := a.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := a.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaGetQueryResults) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := a.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

