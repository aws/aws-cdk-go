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
type SageMakerCreateTransformJob interface {
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
	// The execution role for the Sagemaker transform job.
	//
	// Only available after task has been added to a state machine.
	Role() awsiam.IRole
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

// The jsii proxy struct for SageMakerCreateTransformJob
type jsiiProxy_SageMakerCreateTransformJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_SageMakerCreateTransformJob) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_SageMakerCreateTransformJob) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
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

func (j *jsiiProxy_SageMakerCreateTransformJob) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTransformJob) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
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

func (j *jsiiProxy_SageMakerCreateTransformJob) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
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


func NewSageMakerCreateTransformJob(scope constructs.Construct, id *string, props *SageMakerCreateTransformJobProps) SageMakerCreateTransformJob {
	_init_.Initialize()

	if err := validateNewSageMakerCreateTransformJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SageMakerCreateTransformJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSageMakerCreateTransformJob_Override(s SageMakerCreateTransformJob, scope constructs.Construct, id *string, props *SageMakerCreateTransformJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
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

func (j *jsiiProxy_SageMakerCreateTransformJob)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateTransformJob)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateTransformJob)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func SageMakerCreateTransformJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func SageMakerCreateTransformJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func SageMakerCreateTransformJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
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
func SageMakerCreateTransformJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Class representing the SageMaker Create Transform Job task using JSONata.
func SageMakerCreateTransformJob_Jsonata(scope constructs.Construct, id *string, props *SageMakerCreateTransformJobJsonataProps) SageMakerCreateTransformJob {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns SageMakerCreateTransformJob

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Class representing the SageMaker Create Transform Job task using JSONPath.
func SageMakerCreateTransformJob_JsonPath(scope constructs.Construct, id *string, props *SageMakerCreateTransformJobJsonPathProps) SageMakerCreateTransformJob {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns SageMakerCreateTransformJob

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func SageMakerCreateTransformJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateSageMakerCreateTransformJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
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

func (s *jsiiProxy_SageMakerCreateTransformJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := s.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (s *jsiiProxy_SageMakerCreateTransformJob) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := s.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addItemProcessor",
		[]interface{}{processor, config},
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

func (s *jsiiProxy_SageMakerCreateTransformJob) Next(state awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := s.validateNextParameters(state); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{state},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
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

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
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

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderItemProcessor",
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

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
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

func (s *jsiiProxy_SageMakerCreateTransformJob) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SageMakerCreateTransformJob) ToStateJson(stateMachineQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		[]interface{}{stateMachineQueryLanguage},
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

func (s *jsiiProxy_SageMakerCreateTransformJob) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validateState",
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

