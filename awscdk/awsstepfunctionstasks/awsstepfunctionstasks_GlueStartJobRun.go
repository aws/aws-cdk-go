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

// Starts an AWS Glue job in a Task state.
//
// OUTPUT: the output of this task is a JobRun structure, for details consult
// https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-JobRun
//
// Example:
//   tasks.NewGlueStartJobRun(this, jsii.String("Task"), &glueStartJobRunProps{
//   	glueJobName: jsii.String("my-glue-job"),
//   	arguments: sfn.taskInput.fromObject(map[string]interface{}{
//   		"key": jsii.String("value"),
//   	}),
//   	taskTimeout: sfn.timeout.duration(awscdk.Duration.minutes(jsii.Number(30))),
//   	notifyDelayAfter: awscdk.Duration.minutes(jsii.Number(5)),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html
//
type GlueStartJobRun interface {
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
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
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

// The jsii proxy struct for GlueStartJobRun
type jsiiProxy_GlueStartJobRun struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_GlueStartJobRun) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewGlueStartJobRun(scope constructs.Construct, id *string, props *GlueStartJobRunProps) GlueStartJobRun {
	_init_.Initialize()

	if err := validateNewGlueStartJobRunParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueStartJobRun{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGlueStartJobRun_Override(g GlueStartJobRun, scope constructs.Construct, id *string, props *GlueStartJobRunProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GlueStartJobRun)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_GlueStartJobRun)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func GlueStartJobRun_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateGlueStartJobRun_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func GlueStartJobRun_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueStartJobRun_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func GlueStartJobRun_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueStartJobRun_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
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
func GlueStartJobRun_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueStartJobRun_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func GlueStartJobRun_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateGlueStartJobRun_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueStartJobRun",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (g *jsiiProxy_GlueStartJobRun) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := g.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addBranch",
		[]interface{}{branch},
	)
}

func (g *jsiiProxy_GlueStartJobRun) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := g.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := g.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (g *jsiiProxy_GlueStartJobRun) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := g.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (g *jsiiProxy_GlueStartJobRun) AddPrefix(x *string) {
	if err := g.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPrefix",
		[]interface{}{x},
	)
}

func (g *jsiiProxy_GlueStartJobRun) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := g.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (g *jsiiProxy_GlueStartJobRun) MakeDefault(def awsstepfunctions.State) {
	if err := g.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeDefault",
		[]interface{}{def},
	)
}

func (g *jsiiProxy_GlueStartJobRun) MakeNext(next awsstepfunctions.State) {
	if err := g.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeNext",
		[]interface{}{next},
	)
}

func (g *jsiiProxy_GlueStartJobRun) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := g.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := g.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		g,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueStartJobRun) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

