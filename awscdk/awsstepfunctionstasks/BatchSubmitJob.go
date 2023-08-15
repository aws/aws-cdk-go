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

// Task to submits an AWS Batch job from a job definition.
//
// Example:
//   import batch "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//   var batchJobDefinition ecsJobDefinition
//   var batchQueue jobQueue
//
//
//   task := tasks.NewBatchSubmitJob(this, jsii.String("Submit Job"), &BatchSubmitJobProps{
//   	JobDefinitionArn: batchJobDefinition.JobDefinitionArn,
//   	JobName: jsii.String("MyJob"),
//   	JobQueueArn: batchQueue.JobQueueArn,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-batch.html
//
type BatchSubmitJob interface {
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

// The jsii proxy struct for BatchSubmitJob
type jsiiProxy_BatchSubmitJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_BatchSubmitJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BatchSubmitJob) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewBatchSubmitJob(scope constructs.Construct, id *string, props *BatchSubmitJobProps) BatchSubmitJob {
	_init_.Initialize()

	if err := validateNewBatchSubmitJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BatchSubmitJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBatchSubmitJob_Override(b BatchSubmitJob, scope constructs.Construct, id *string, props *BatchSubmitJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_BatchSubmitJob)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_BatchSubmitJob)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func BatchSubmitJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateBatchSubmitJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func BatchSubmitJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateBatchSubmitJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func BatchSubmitJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateBatchSubmitJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
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
func BatchSubmitJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBatchSubmitJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func BatchSubmitJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateBatchSubmitJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.BatchSubmitJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (b *jsiiProxy_BatchSubmitJob) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := b.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addBranch",
		[]interface{}{branch},
	)
}

func (b *jsiiProxy_BatchSubmitJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := b.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		b,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := b.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (b *jsiiProxy_BatchSubmitJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := b.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (b *jsiiProxy_BatchSubmitJob) AddPrefix(x *string) {
	if err := b.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addPrefix",
		[]interface{}{x},
	)
}

func (b *jsiiProxy_BatchSubmitJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := b.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		b,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := b.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (b *jsiiProxy_BatchSubmitJob) MakeDefault(def awsstepfunctions.State) {
	if err := b.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"makeDefault",
		[]interface{}{def},
	)
}

func (b *jsiiProxy_BatchSubmitJob) MakeNext(next awsstepfunctions.State) {
	if err := b.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"makeNext",
		[]interface{}{next},
	)
}

func (b *jsiiProxy_BatchSubmitJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := b.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		b,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BatchSubmitJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := b.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

