package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a Task state in the state machine.
//
// Reaching a Task state causes some work to be executed, represented by the
// Task's resource property. Task constructs represent a generic Amazon
// States Language Task.
//
// For some resource types, more specific subclasses of Task may be available
// which are more convenient to use.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var parameters interface{}
//   var stepFunctionsTask iStepFunctionsTask
//
//   task := awscdk.Aws_stepfunctions.NewTask(this, jsii.String("MyTask"), &taskProps{
//   	task: stepFunctionsTask,
//
//   	// the properties below are optional
//   	comment: jsii.String("comment"),
//   	inputPath: jsii.String("inputPath"),
//   	outputPath: jsii.String("outputPath"),
//   	parameters: map[string]interface{}{
//   		"parametersKey": parameters,
//   	},
//   	resultPath: jsii.String("resultPath"),
//   	timeout: duration,
//   })
//
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
type Task interface {
	State
	INextable
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Branches() *[]StateGraph
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Comment() *string
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	DefaultChoice() State
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	SetDefaultChoice(val State)
	// Continuable states of this Chainable.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Id() *string
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	InputPath() *string
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Iteration() StateGraph
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	SetIteration(val StateGraph)
	// The construct tree node associated with this construct.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Node() awscdk.ConstructNode
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	OutputPath() *string
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Parameters() *map[string]interface{}
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	ResultPath() *string
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	StateId() *string
	// Add a paralle branch to this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddBranch(branch StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddCatch(handler IChainable, props *CatchProps) Task
	// Add a choice branch to this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddChoice(condition Condition, next State)
	// Add a map iterator to this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	AddRetry(props *RetryProps) Task
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	BindToGraph(graph StateGraph)
	// Make the indicated state the default choice transition of this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MakeDefault(def State)
	// Make the indicated state the default transition of this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MakeNext(next State)
	// Return the given named metric for this Task.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity fails.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times the heartbeat times out for this activity.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the Task starts and the time it closes.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is scheduled.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, for which the activity stays in the schedule state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity is started.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity succeeds.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of times this activity times out.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Continue normal execution with the given state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Next(next IChainable) Chain
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Prepare()
	// Render parallel branches in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderChoices() interface{}
	// Render InputPath/Parameters/OutputPath in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderInputOutput() interface{}
	// Render map iterator in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderNextEnd() interface{}
	// Render ResultSelector in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	RenderRetryCatch() interface{}
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Synthesize(session awscdk.ISynthesisSession)
	// Return the Amazon States Language object for this state.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	ToStateJson() *map[string]interface{}
	// Returns a string representation of this construct.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	Validate() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Task
type jsiiProxy_Task struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_Task) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Task) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func NewTask(scope constructs.Construct, id *string, props *TaskProps) Task {
	_init_.Initialize()

	if err := validateNewTaskParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Task{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Task",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func NewTask_Override(t Task, scope constructs.Construct, id *string, props *TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Task",
		[]interface{}{scope, id, props},
		t,
	)
}

func (j *jsiiProxy_Task)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Task)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func Task_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateTask_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Task",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func Task_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateTask_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Task",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func Task_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateTask_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Task",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func Task_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTask_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Task",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Deprecated: - replaced by service integration specific classes (i.e. LambdaInvoke, SnsPublish)
func Task_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateTask_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions.Task",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (t *jsiiProxy_Task) AddBranch(branch StateGraph) {
	if err := t.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addBranch",
		[]interface{}{branch},
	)
}

func (t *jsiiProxy_Task) AddCatch(handler IChainable, props *CatchProps) Task {
	if err := t.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		t,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) AddChoice(condition Condition, next State) {
	if err := t.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (t *jsiiProxy_Task) AddIterator(iteration StateGraph) {
	if err := t.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (t *jsiiProxy_Task) AddPrefix(x *string) {
	if err := t.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addPrefix",
		[]interface{}{x},
	)
}

func (t *jsiiProxy_Task) AddRetry(props *RetryProps) Task {
	if err := t.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns Task

	_jsii_.Invoke(
		t,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) BindToGraph(graph StateGraph) {
	if err := t.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (t *jsiiProxy_Task) MakeDefault(def State) {
	if err := t.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"makeDefault",
		[]interface{}{def},
	)
}

func (t *jsiiProxy_Task) MakeNext(next State) {
	if err := t.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"makeNext",
		[]interface{}{next},
	)
}

func (t *jsiiProxy_Task) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := t.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		t,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) Next(next IChainable) Chain {
	if err := t.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		t,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) OnPrepare() {
	_jsii_.InvokeVoid(
		t,
		"onPrepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) OnSynthesize(session constructs.ISynthesisSession) {
	if err := t.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Task) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) Prepare() {
	_jsii_.InvokeVoid(
		t,
		"prepare",
		nil, // no parameters
	)
}

func (t *jsiiProxy_Task) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) Synthesize(session awscdk.ISynthesisSession) {
	if err := t.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"synthesize",
		[]interface{}{session},
	)
}

func (t *jsiiProxy_Task) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Task) WhenBoundToGraph(graph StateGraph) {
	if err := t.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

