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

// Invoke a Lambda function as a Task.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var orderFn function
//
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("InvokeOrderProcessor"), &LambdaInvokeProps{
//   	LambdaFunction: orderFn,
//   	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"OrderId": sfn.JsonPath_stringAt(jsii.String("$.OrderId")),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html
//
// Experimental.
type LambdaInvoke interface {
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

// The jsii proxy struct for LambdaInvoke
type jsiiProxy_LambdaInvoke struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_LambdaInvoke) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaInvoke(scope constructs.Construct, id *string, props *LambdaInvokeProps) LambdaInvoke {
	_init_.Initialize()

	if err := validateNewLambdaInvokeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvoke{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvoke_Override(l LambdaInvoke, scope constructs.Construct, id *string, props *LambdaInvokeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaInvoke)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_LambdaInvoke)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func LambdaInvoke_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateLambdaInvoke_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func LambdaInvoke_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateLambdaInvoke_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
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
func LambdaInvoke_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateLambdaInvoke_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LambdaInvoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaInvoke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func LambdaInvoke_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateLambdaInvoke_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (l *jsiiProxy_LambdaInvoke) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := l.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addBranch",
		[]interface{}{branch},
	)
}

func (l *jsiiProxy_LambdaInvoke) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := l.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		l,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := l.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (l *jsiiProxy_LambdaInvoke) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := l.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (l *jsiiProxy_LambdaInvoke) AddPrefix(x *string) {
	if err := l.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addPrefix",
		[]interface{}{x},
	)
}

func (l *jsiiProxy_LambdaInvoke) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := l.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		l,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := l.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (l *jsiiProxy_LambdaInvoke) MakeDefault(def awsstepfunctions.State) {
	if err := l.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"makeDefault",
		[]interface{}{def},
	)
}

func (l *jsiiProxy_LambdaInvoke) MakeNext(next awsstepfunctions.State) {
	if err := l.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"makeNext",
		[]interface{}{next},
	)
}

func (l *jsiiProxy_LambdaInvoke) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := l.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := l.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		l,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvoke) OnSynthesize(session constructs.ISynthesisSession) {
	if err := l.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LambdaInvoke) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvoke) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) Synthesize(session awscdk.ISynthesisSession) {
	if err := l.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LambdaInvoke) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvoke) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := l.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

