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

// Start a CodeBuild Build as a task.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   codebuildProject := codebuild.NewProject(this, jsii.String("Project"), &ProjectProps{
//   	ProjectName: jsii.String("MyTestProject"),
//   	BuildSpec: codebuild.BuildSpec_FromObject(map[string]interface{}{
//   		"version": jsii.String("0.2"),
//   		"phases": map[string]map[string][]*string{
//   			"build": map[string][]*string{
//   				"commands": []*string{
//   					jsii.String("echo \"Hello, CodeBuild!\""),
//   				},
//   			},
//   		},
//   	}),
//   })
//
//   task := tasks.NewCodeBuildStartBuild(this, jsii.String("Task"), &CodeBuildStartBuildProps{
//   	Project: codebuildProject,
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   	EnvironmentVariablesOverride: map[string]BuildEnvironmentVariable{
//   		"ZONE": &BuildEnvironmentVariable{
//   			"type": codebuild.BuildEnvironmentVariableType_PLAINTEXT,
//   			"value": sfn.JsonPath_stringAt(jsii.String("$.envVariables.zone")),
//   		},
//   	},
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-codebuild.html
//
type CodeBuildStartBuild interface {
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
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for CodeBuildStartBuild
type jsiiProxy_CodeBuildStartBuild struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_CodeBuildStartBuild) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildStartBuild) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewCodeBuildStartBuild(scope constructs.Construct, id *string, props *CodeBuildStartBuildProps) CodeBuildStartBuild {
	_init_.Initialize()

	if err := validateNewCodeBuildStartBuildParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildStartBuild{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCodeBuildStartBuild_Override(c CodeBuildStartBuild, scope constructs.Construct, id *string, props *CodeBuildStartBuildProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func CodeBuildStartBuild_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func CodeBuildStartBuild_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func CodeBuildStartBuild_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
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
func CodeBuildStartBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Start a CodeBuild Build as a task using JSONata.
func CodeBuildStartBuild_Jsonata(scope constructs.Construct, id *string, props *CodeBuildStartBuildJsonataProps) CodeBuildStartBuild {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns CodeBuildStartBuild

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Start a CodeBuild Build as a task using JSONPath.
func CodeBuildStartBuild_JsonPath(scope constructs.Construct, id *string, props *CodeBuildStartBuildJsonPathProps) CodeBuildStartBuild {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns CodeBuildStartBuild

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func CodeBuildStartBuild_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateCodeBuildStartBuild_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := c.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := c.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := c.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := c.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := c.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddPrefix(x *string) {
	if err := c.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := c.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := c.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) MakeDefault(def awsstepfunctions.State) {
	if err := c.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) MakeNext(next awsstepfunctions.State) {
	if err := c.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := c.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := c.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		[]interface{}{stateMachineQueryLanguage},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildStartBuild) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := c.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

func (c *jsiiProxy_CodeBuildStartBuild) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"with",
		args,
		&returns,
	)

	return returns
}

