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

// Invoke a Lambda function as a Task.
//
// Example:
//   var fn function
//
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke with empty object as payload"), &LambdaInvokeProps{
//   	LambdaFunction: fn,
//   	Payload: sfn.TaskInput_FromObject(map[string]interface{}{
//   	}),
//   })
//
//   // use the output of fn as input
//   // use the output of fn as input
//   tasks.NewLambdaInvoke(this, jsii.String("Invoke with payload field in the state input"), &LambdaInvokeProps{
//   	LambdaFunction: fn,
//   	Payload: sfn.TaskInput_FromJsonPathAt(jsii.String("$.Payload")),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html
//
type LambdaInvoke interface {
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
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() awsstepfunctions.State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	// Add a paralle branch to this state.
	AddBranch(branch awsstepfunctions.StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	// Add a choice branch to this state.
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
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

func (j *jsiiProxy_LambdaInvoke) Node() constructs.Node {
	var returns constructs.Node
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


func NewLambdaInvoke(scope constructs.Construct, id *string, props *LambdaInvokeProps) LambdaInvoke {
	_init_.Initialize()

	if err := validateNewLambdaInvokeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvoke{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewLambdaInvoke_Override(l LambdaInvoke, scope constructs.Construct, id *string, props *LambdaInvokeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
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
func LambdaInvoke_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateLambdaInvoke_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func LambdaInvoke_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateLambdaInvoke_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func LambdaInvoke_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateLambdaInvoke_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
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
func LambdaInvoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaInvoke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func LambdaInvoke_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateLambdaInvoke_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.LambdaInvoke",
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

func (l *jsiiProxy_LambdaInvoke) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validateState",
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

