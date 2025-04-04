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

// A StepFunctions Task to send messages to SQS queue.
//
// Example:
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   // Use a field from the execution data as message.
//   task1 := tasks.NewSqsSendMessage(this, jsii.String("Send1"), &SqsSendMessageProps{
//   	Queue: Queue,
//   	MessageBody: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
//   })
//
//   // Combine a field from the execution data with
//   // a literal object.
//   task2 := tasks.NewSqsSendMessage(this, jsii.String("Send2"), &SqsSendMessageProps{
//   	Queue: Queue,
//   	MessageBody: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"field1": jsii.String("somedata"),
//   		"field2": sfn.JsonPath_stringAt(jsii.String("$.field2")),
//   	}),
//   })
//
type SqsSendMessage interface {
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
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
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
	ToStateJson(topLevelQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SqsSendMessage
type jsiiProxy_SqsSendMessage struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SqsSendMessage) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsSendMessage) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewSqsSendMessage(scope constructs.Construct, id *string, props *SqsSendMessageProps) SqsSendMessage {
	_init_.Initialize()

	if err := validateNewSqsSendMessageParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsSendMessage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSqsSendMessage_Override(s SqsSendMessage, scope constructs.Construct, id *string, props *SqsSendMessageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SqsSendMessage)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SqsSendMessage)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_SqsSendMessage)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_SqsSendMessage)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_SqsSendMessage)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func SqsSendMessage_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateSqsSendMessage_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func SqsSendMessage_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSqsSendMessage_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func SqsSendMessage_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSqsSendMessage_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
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
func SqsSendMessage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSqsSendMessage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// A StepFunctions Task to send messages to SQS queue using JSONata.
func SqsSendMessage_Jsonata(scope constructs.Construct, id *string, props *SqsSendMessageJsonataProps) SqsSendMessage {
	_init_.Initialize()

	if err := validateSqsSendMessage_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns SqsSendMessage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// A StepFunctions Task to send messages to SQS queue using JSONPath.
func SqsSendMessage_JsonPath(scope constructs.Construct, id *string, props *SqsSendMessageJsonPathProps) SqsSendMessage {
	_init_.Initialize()

	if err := validateSqsSendMessage_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns SqsSendMessage

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func SqsSendMessage_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateSqsSendMessage_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.SqsSendMessage",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := s.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := s.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := s.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := s.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := s.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddPrefix(x *string) {
	if err := s.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

func (s *jsiiProxy_SqsSendMessage) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := s.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := s.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (s *jsiiProxy_SqsSendMessage) MakeDefault(def awsstepfunctions.State) {
	if err := s.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

func (s *jsiiProxy_SqsSendMessage) MakeNext(next awsstepfunctions.State) {
	if err := s.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

func (s *jsiiProxy_SqsSendMessage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := s.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := s.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) ToStateJson(topLevelQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := s.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

