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

// Class representing the SageMaker Create Transform Job task.
//
// Example:
//   tasks.NewSageMakerCreateTransformJob(this, jsii.String("Batch Inference"), &SageMakerCreateTransformJobProps{
//   	TransformJobName: jsii.String("MyTransformJob"),
//   	ModelName: jsii.String("MyModelName"),
//   	ModelClientOptions: &ModelClientOptions{
//   		InvocationsMaxRetries: jsii.Number(3),
//   		 // default is 0
//   		InvocationsTimeout: awscdk.Duration_Minutes(jsii.Number(5)),
//   	},
//   	TransformInput: &TransformInput{
//   		TransformDataSource: &TransformDataSource{
//   			S3DataSource: &TransformS3DataSource{
//   				S3Uri: jsii.String("s3://inputbucket/train"),
//   				S3DataType: tasks.S3DataType_S3_PREFIX,
//   			},
//   		},
//   	},
//   	TransformOutput: &TransformOutput{
//   		S3OutputPath: jsii.String("s3://outputbucket/TransformJobOutputPath"),
//   	},
//   	TransformResources: &TransformResources{
//   		InstanceCount: jsii.Number(1),
//   		InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_M4, ec2.InstanceSize_XLARGE),
//   	},
//   })
//
// Experimental.
type SageMakerCreateTransformJob interface {
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
	// The execution role for the Sagemaker transform job.
	//
	// Only available after task has been added to a state machine.
	// Experimental.
	Role() awsiam.IRole
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

// The jsii proxy struct for SageMakerCreateTransformJob
type jsiiProxy_SageMakerCreateTransformJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerCreateTransformJob(scope constructs.Construct, id *string, props *SageMakerCreateTransformJobProps) SageMakerCreateTransformJob {
	_init_.Initialize()

	if err := validateNewSageMakerCreateTransformJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerCreateTransformJob{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerCreateTransformJob_Override(s SageMakerCreateTransformJob, scope constructs.Construct, id *string, props *SageMakerCreateTransformJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerCreateTransformJob)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateTransformJob)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerCreateTransformJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerCreateTransformJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
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
func SageMakerCreateTransformJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerCreateTransformJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerCreateTransformJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := s.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	if err := s.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := s.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddPrefix(x *string) {
	if err := s.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := s.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) MakeDefault(def awsstepfunctions.State) {
	if err := s.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) MakeNext(next awsstepfunctions.State) {
	if err := s.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
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

func (s *jsiiProxy_SageMakerCreateTransformJob) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := s.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

