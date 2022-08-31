package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Stop an Athena Query Execution as a Task.
//
// Example:
//   stopQueryExecutionJob := tasks.NewAthenaStopQueryExecution(this, jsii.String("Stop Query Execution"), &athenaStopQueryExecutionProps{
//   	queryExecutionId: sfn.jsonPath.stringAt(jsii.String("$.QueryExecutionId")),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
// Experimental.
type AthenaStopQueryExecution interface {
	awsstepfunctions.TaskStateBase
	// Experimental.
	Branches() *[]awsstepfunctions.StateGraph
	// Experimental.
	Comment() *string
	// Experimental.
	DefaultChoice() awsstepfunctions.State
	// Experimental.
	SetDefaultChoice(val awsstepfunctions.State)
	// Continuable states of this Chainable.
	// Experimental.
	EndStates() *[]awsstepfunctions.INextable
	// Descriptive identifier for this chainable.
	// Experimental.
	Id() *string
	// Experimental.
	InputPath() *string
	// Experimental.
	Iteration() awsstepfunctions.StateGraph
	// Experimental.
	SetIteration(val awsstepfunctions.StateGraph)
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	OutputPath() *string
	// Experimental.
	Parameters() *map[string]interface{}
	// Experimental.
	ResultPath() *string
	// Experimental.
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	// Experimental.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	// Experimental.
	StateId() *string
	// Experimental.
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	// Experimental.
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a paralle branch to this state.
	// Experimental.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	// Experimental.
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	// Add a choice branch to this state.
	// Experimental.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	// Add a map iterator to this state.
	// Experimental.
	AddIterator(iteration awsstepfunctions.StateGraph)
	// Add a prefix to the stateId of this state.
	// Experimental.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	// Experimental.
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	// Experimental.
	BindToGraph(graph awsstepfunctions.StateGraph)
	// Make the indicated state the default choice transition of this state.
	// Experimental.
	MakeDefault(def awsstepfunctions.State)
	// Make the indicated state the default transition of this state.
	// Experimental.
	MakeNext(next awsstepfunctions.State)
	// Return the given named metric for this Task.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	// Experimental.
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	// Experimental.
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	// Experimental.
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	// Experimental.
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	// Experimental.
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	// Experimental.
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	// Experimental.
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	// Experimental.
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	// Experimental.
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Continue normal execution with the given state.
	// Experimental.
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Render parallel branches in ASL JSON format.
	// Experimental.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	// Experimental.
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	// Experimental.
	RenderInputOutput() interface{}
	// Render map iterator in ASL JSON format.
	// Experimental.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	// Experimental.
	RenderNextEnd() interface{}
	// Render ResultSelector in ASL JSON format.
	// Experimental.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	// Experimental.
	RenderRetryCatch() interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Return the Amazon States Language object for this state.
	// Experimental.
	ToStateJson() *map[string]interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	// Experimental.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaStopQueryExecution
type jsiiProxy_AthenaStopQueryExecution struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaStopQueryExecution) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewAthenaStopQueryExecution(scope constructs.Construct, id *string, props *AthenaStopQueryExecutionProps) AthenaStopQueryExecution {
	_init_.Initialize()

	j := jsiiProxy_AthenaStopQueryExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaStopQueryExecution_Override(a AthenaStopQueryExecution, scope constructs.Construct, id *string, props *AthenaStopQueryExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaStopQueryExecution) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaStopQueryExecution) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func AthenaStopQueryExecution_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func AthenaStopQueryExecution_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
// Experimental.
func AthenaStopQueryExecution_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AthenaStopQueryExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func AthenaStopQueryExecution_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AthenaStopQueryExecution) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AthenaStopQueryExecution) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

