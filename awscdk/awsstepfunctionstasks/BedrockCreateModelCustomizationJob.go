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

// A Step Functions Task to create model customization in Bedrock.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//   var outputBucket iBucket
//   var trainingBucket iBucket
//   var validationBucket iBucket
//   var kmsKey iKey
//   var vpc iVpc
//
//
//   model := bedrock.FoundationModel_FromFoundationModelId(this, jsii.String("Model"), bedrock.FoundationModelIdentifier_AMAZON_TITAN_TEXT_G1_EXPRESS_V1())
//
//   task := tasks.NewBedrockCreateModelCustomizationJob(this, jsii.String("CreateModelCustomizationJob"), &BedrockCreateModelCustomizationJobProps{
//   	BaseModel: model,
//   	ClientRequestToken: jsii.String("MyToken"),
//   	CustomizationType: tasks.CustomizationType_FINE_TUNING,
//   	CustomModelKmsKey: kmsKey,
//   	CustomModelName: jsii.String("MyCustomModel"),
//   	 // required
//   	CustomModelTags: []customModelTag{
//   		&customModelTag{
//   			Key: jsii.String("key1"),
//   			Value: jsii.String("value1"),
//   		},
//   	},
//   	HyperParameters: map[string]*string{
//   		"batchSize": jsii.String("10"),
//   	},
//   	JobName: jsii.String("MyCustomizationJob"),
//   	 // required
//   	JobTags: []*customModelTag{
//   		&customModelTag{
//   			Key: jsii.String("key2"),
//   			Value: jsii.String("value2"),
//   		},
//   	},
//   	OutputData: &OutputBucketConfiguration{
//   		Bucket: outputBucket,
//   		 // required
//   		Path: jsii.String("output-data/"),
//   	},
//   	TrainingData: &TrainingBucketConfiguration{
//   		Bucket: trainingBucket,
//   		Path: jsii.String("training-data/data.json"),
//   	},
//   	 // required
//   	// If you don't provide validation data, you have to specify `Evaluation percentage` hyperparameter.
//   	ValidationData: []validationBucketConfiguration{
//   		&validationBucketConfiguration{
//   			Bucket: validationBucket,
//   			Path: jsii.String("validation-data/data.json"),
//   		},
//   	},
//   	VpcConfig: map[string][]iSecurityGroup{
//   		"securityGroups": []*iSecurityGroup{
//   			ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   				"vpc": vpc,
//   			}),
//   		},
//   		"subnets": vpc.privateSubnets,
//   	},
//   })
//
type BedrockCreateModelCustomizationJob interface {
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
	// The IAM role for the bedrock create model customization job.
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

// The jsii proxy struct for BedrockCreateModelCustomizationJob
type jsiiProxy_BedrockCreateModelCustomizationJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewBedrockCreateModelCustomizationJob(scope constructs.Construct, id *string, props *BedrockCreateModelCustomizationJobProps) BedrockCreateModelCustomizationJob {
	_init_.Initialize()

	if err := validateNewBedrockCreateModelCustomizationJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockCreateModelCustomizationJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewBedrockCreateModelCustomizationJob_Override(b BedrockCreateModelCustomizationJob, scope constructs.Construct, id *string, props *BedrockCreateModelCustomizationJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_BedrockCreateModelCustomizationJob)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func BedrockCreateModelCustomizationJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateBedrockCreateModelCustomizationJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func BedrockCreateModelCustomizationJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateBedrockCreateModelCustomizationJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func BedrockCreateModelCustomizationJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateBedrockCreateModelCustomizationJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
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
func BedrockCreateModelCustomizationJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockCreateModelCustomizationJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func BedrockCreateModelCustomizationJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateBedrockCreateModelCustomizationJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.BedrockCreateModelCustomizationJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := b.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addBranch",
		[]interface{}{branch},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := b.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := b.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := b.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddPrefix(x *string) {
	if err := b.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addPrefix",
		[]interface{}{x},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := b.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MakeDefault(def awsstepfunctions.State) {
	if err := b.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"makeDefault",
		[]interface{}{def},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MakeNext(next awsstepfunctions.State) {
	if err := b.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"makeNext",
		[]interface{}{next},
	)
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
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

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) ToStateJson(topLevelQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"toStateJson",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockCreateModelCustomizationJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := b.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

