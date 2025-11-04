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

// Start a Job run as a Task.
//
// Example:
//   tasks.NewGlueDataBrewStartJobRun(this, jsii.String("Task"), &GlueDataBrewStartJobRunProps{
//   	Name: jsii.String("databrew-job"),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html
//
type GlueDataBrewStartJobRun interface {
	awsstepfunctions.TaskStateBase
	Arguments() *map[string]interface{}
	Assign() *map[string]interface{}
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
	Outputs() *map[string]interface{}
	Parameters() *map[string]interface{}
	Processor() awsstepfunctions.StateGraph
	SetProcessor(val awsstepfunctions.StateGraph)
	ProcessorConfig() *awsstepfunctions.ProcessorConfig
	SetProcessorConfig(val *awsstepfunctions.ProcessorConfig)
	ProcessorMode() awsstepfunctions.ProcessorMode
	SetProcessorMode(val awsstepfunctions.ProcessorMode)
	QueryLanguage() awsstepfunctions.QueryLanguage
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
	Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain
	// Render the assign in ASL JSON format.
	RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{}
	// Render parallel branches in ASL JSON format.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{}
	// Render InputPath/Parameters/OutputPath/Arguments/Output in ASL JSON format.
	RenderInputOutput() interface{}
	// Render ItemProcessor in ASL JSON format.
	RenderItemProcessor() interface{}
	// Render map iterator in ASL JSON format.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	RenderNextEnd() interface{}
	// Render QueryLanguage in ASL JSON format if needed.
	RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{}
	// Render ResultSelector in ASL JSON format.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{}
	// Return the Amazon States Language object for this state.
	ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for GlueDataBrewStartJobRun
type jsiiProxy_GlueDataBrewStartJobRun struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewGlueDataBrewStartJobRun(scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) GlueDataBrewStartJobRun {
	_init_.Initialize()

	if err := validateNewGlueDataBrewStartJobRunParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueDataBrewStartJobRun{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewGlueDataBrewStartJobRun_Override(g GlueDataBrewStartJobRun, scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func GlueDataBrewStartJobRun_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func GlueDataBrewStartJobRun_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func GlueDataBrewStartJobRun_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
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
func GlueDataBrewStartJobRun_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Start a Job run as a Task using JSONata.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html
//
func GlueDataBrewStartJobRun_Jsonata(scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunJsonataProps) GlueDataBrewStartJobRun {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns GlueDataBrewStartJobRun

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Start a Job run as a Task using JSONPath.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html
//
func GlueDataBrewStartJobRun_JsonPath(scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunJsonPathProps) GlueDataBrewStartJobRun {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns GlueDataBrewStartJobRun

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func GlueDataBrewStartJobRun_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateGlueDataBrewStartJobRun_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := g.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addBranch",
		[]interface{}{branch},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := g.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := g.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := g.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddPrefix(x *string) {
	if err := g.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addPrefix",
		[]interface{}{x},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeDefault(def awsstepfunctions.State) {
	if err := g.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeDefault",
		[]interface{}{def},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeNext(next awsstepfunctions.State) {
	if err := g.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"makeNext",
		[]interface{}{next},
	)
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (g *jsiiProxy_GlueDataBrewStartJobRun) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := g.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		g,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toStateJson",
		[]interface{}{stateMachineQueryLanguage},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueDataBrewStartJobRun) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := g.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

