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

// A Step Functions Task to call a public third-party API.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//   var connection Connection
//
//
//   getIssue := tasks.HttpInvoke_Jsonata(this, jsii.String("Get Issue"), &HttpInvokeJsonataProps{
//   	Connection: Connection,
//   	ApiRoot: jsii.String("{% 'https://' & $states.input.hostname %}"),
//   	ApiEndpoint: sfn.TaskInput_FromText(jsii.String("{% 'issues/' & $states.input.issue.id %}")),
//   	Method: sfn.TaskInput_*FromText(jsii.String("GET")),
//   	// Parse the API call result to object and set to the variables
//   	Assign: map[string]interface{}{
//   		"hostname": jsii.String("{% $states.input.hostname %}"),
//   		"issue": jsii.String("{% $parse($states.result.ResponseBody) %}"),
//   	},
//   })
//
//   updateLabels := tasks.HttpInvoke_Jsonata(this, jsii.String("Update Issue Labels"), &HttpInvokeJsonataProps{
//   	Connection: Connection,
//   	ApiRoot: jsii.String("{% 'https://' & $states.input.hostname %}"),
//   	ApiEndpoint: sfn.TaskInput_*FromText(jsii.String("{% 'issues/' & $states.input.issue.id & 'labels' %}")),
//   	Method: sfn.TaskInput_*FromText(jsii.String("POST")),
//   	Body: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"labels": jsii.String("{% [$type, $component] %}"),
//   	}),
//   })
//   notMatchTitleTemplate := sfn.Pass_Jsonata(this, jsii.String("Not Match Title Template"))
//
//   definition := getIssue.Next(sfn.Choice_Jsonata(this, jsii.String("Match Title Template?")).When(sfn.Condition_Jsonata(jsii.String("{% $contains($issue.title, /(feat)|(fix)|(chore)(w*):.*/) %}")), updateLabels, &ChoiceTransitionOptions{
//   	Assign: map[string]interface{}{
//   		"type": jsii.String("{% $match($states.input.title, /(w*)((.*))/).groups[0] %}"),
//   		"component": jsii.String("{% $match($states.input.title, /(w*)((.*))/).groups[1] %}"),
//   	},
//   }).Otherwise(notMatchTitleTemplate))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	DefinitionBody: sfn.DefinitionBody_FromChainable(definition),
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	Comment: jsii.String("automate issue labeling state machine"),
//   })
//
type HttpInvoke interface {
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
	BuildTaskPolicyStatements() *[]awsiam.PolicyStatement
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

// The jsii proxy struct for HttpInvoke
type jsiiProxy_HttpInvoke struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_HttpInvoke) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpInvoke) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewHttpInvoke(scope constructs.Construct, id *string, props *HttpInvokeProps) HttpInvoke {
	_init_.Initialize()

	if err := validateNewHttpInvokeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpInvoke{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewHttpInvoke_Override(h HttpInvoke, scope constructs.Construct, id *string, props *HttpInvokeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HttpInvoke)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_HttpInvoke)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_HttpInvoke)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_HttpInvoke)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_HttpInvoke)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func HttpInvoke_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateHttpInvoke_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func HttpInvoke_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateHttpInvoke_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func HttpInvoke_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateHttpInvoke_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
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
func HttpInvoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpInvoke_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// A Step Functions Task to call a public third-party API using JSONata.
func HttpInvoke_Jsonata(scope constructs.Construct, id *string, props *HttpInvokeJsonataProps) HttpInvoke {
	_init_.Initialize()

	if err := validateHttpInvoke_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns HttpInvoke

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// A Step Functions Task to call a public third-party API using JSONPath.
func HttpInvoke_JsonPath(scope constructs.Construct, id *string, props *HttpInvokeJsonPathProps) HttpInvoke {
	_init_.Initialize()

	if err := validateHttpInvoke_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns HttpInvoke

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func HttpInvoke_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateHttpInvoke_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.HttpInvoke",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (h *jsiiProxy_HttpInvoke) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := h.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addBranch",
		[]interface{}{branch},
	)
}

func (h *jsiiProxy_HttpInvoke) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := h.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		h,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := h.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (h *jsiiProxy_HttpInvoke) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := h.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (h *jsiiProxy_HttpInvoke) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := h.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (h *jsiiProxy_HttpInvoke) AddPrefix(x *string) {
	if err := h.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"addPrefix",
		[]interface{}{x},
	)
}

func (h *jsiiProxy_HttpInvoke) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := h.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		h,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := h.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (h *jsiiProxy_HttpInvoke) BuildTaskPolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		h,
		"buildTaskPolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MakeDefault(def awsstepfunctions.State) {
	if err := h.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"makeDefault",
		[]interface{}{def},
	)
}

func (h *jsiiProxy_HttpInvoke) MakeNext(next awsstepfunctions.State) {
	if err := h.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"makeNext",
		[]interface{}{next},
	)
}

func (h *jsiiProxy_HttpInvoke) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := h.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		h,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := h.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		h,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		h,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"toStateJson",
		[]interface{}{stateMachineQueryLanguage},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpInvoke) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := h.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

func (h *jsiiProxy_HttpInvoke) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		h,
		"with",
		args,
		&returns,
	)

	return returns
}

