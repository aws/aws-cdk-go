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

// A StepFunctions task to call DynamoGetItem.
//
// Example:
//   var myTable Table
//
//   tasks.NewDynamoGetItem(this, jsii.String("Get Item"), &DynamoGetItemProps{
//   	Key: map[string]DynamoAttributeValue{
//   		"messageId": tasks.DynamoAttributeValue_fromString(jsii.String("message-007")),
//   	},
//   	Table: myTable,
//   })
//
type DynamoGetItem interface {
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

// The jsii proxy struct for DynamoGetItem
type jsiiProxy_DynamoGetItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_DynamoGetItem) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoGetItem) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewDynamoGetItem(scope constructs.Construct, id *string, props *DynamoGetItemProps) DynamoGetItem {
	_init_.Initialize()

	if err := validateNewDynamoGetItemParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoGetItem{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDynamoGetItem_Override(d DynamoGetItem, scope constructs.Construct, id *string, props *DynamoGetItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoGetItem)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoGetItem)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_DynamoGetItem)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_DynamoGetItem)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_DynamoGetItem)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func DynamoGetItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateDynamoGetItem_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func DynamoGetItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateDynamoGetItem_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func DynamoGetItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateDynamoGetItem_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
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
func DynamoGetItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamoGetItem_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// A StepFunctions task using JSONata to call DynamoGetItem.
func DynamoGetItem_Jsonata(scope constructs.Construct, id *string, props *DynamoGetItemJsonataProps) DynamoGetItem {
	_init_.Initialize()

	if err := validateDynamoGetItem_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns DynamoGetItem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// A StepFunctions task using JSONPath to call DynamoGetItem.
func DynamoGetItem_JsonPath(scope constructs.Construct, id *string, props *DynamoGetItemJsonPathProps) DynamoGetItem {
	_init_.Initialize()

	if err := validateDynamoGetItem_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns DynamoGetItem

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func DynamoGetItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateDynamoGetItem_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.DynamoGetItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := d.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := d.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := d.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := d.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := d.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddPrefix(x *string) {
	if err := d.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

func (d *jsiiProxy_DynamoGetItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := d.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := d.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (d *jsiiProxy_DynamoGetItem) MakeDefault(def awsstepfunctions.State) {
	if err := d.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

func (d *jsiiProxy_DynamoGetItem) MakeNext(next awsstepfunctions.State) {
	if err := d.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

func (d *jsiiProxy_DynamoGetItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := d.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := d.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		[]interface{}{stateMachineQueryLanguage},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoGetItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := d.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

func (d *jsiiProxy_DynamoGetItem) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

