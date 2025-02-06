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

// A Step Functions Task to create a job in MediaConvert.
//
// The JobConfiguration/Request Syntax is defined in the Parameters in the Task State.
//
// Example:
//   tasks.NewMediaConvertCreateJob(this, jsii.String("CreateJob"), &MediaConvertCreateJobProps{
//   	CreateJobRequest: map[string]interface{}{
//   		"Role": jsii.String("arn:aws:iam::123456789012:role/MediaConvertRole"),
//   		"Settings": map[string][]map[string]interface{}{
//   			"OutputGroups": []map[string]interface{}{
//   				map[string]interface{}{
//   					"Outputs": []map[string]interface{}{
//   						map[string]interface{}{
//   							"ContainerSettings": map[string]*string{
//   								"Container": jsii.String("MP4"),
//   							},
//   							"VideoDescription": map[string]map[string]interface{}{
//   								"CodecSettings": map[string]interface{}{
//   									"Codec": jsii.String("H_264"),
//   									"H264Settings": map[string]interface{}{
//   										"MaxBitrate": jsii.Number(1000),
//   										"RateControlMode": jsii.String("QVBR"),
//   										"SceneChangeDetect": jsii.String("TRANSITION_DETECTION"),
//   									},
//   								},
//   							},
//   							"AudioDescriptions": []map[string]map[string]interface{}{
//   								map[string]map[string]interface{}{
//   									"CodecSettings": map[string]interface{}{
//   										"Codec": jsii.String("AAC"),
//   										"AacSettings": map[string]interface{}{
//   											"Bitrate": jsii.Number(96000),
//   											"CodingMode": jsii.String("CODING_MODE_2_0"),
//   											"SampleRate": jsii.Number(48000),
//   										},
//   									},
//   								},
//   							},
//   						},
//   					},
//   					"OutputGroupSettings": map[string]interface{}{
//   						"Type": jsii.String("FILE_GROUP_SETTINGS"),
//   						"FileGroupSettings": map[string]*string{
//   							"Destination": jsii.String("s3://EXAMPLE-DESTINATION-BUCKET/"),
//   						},
//   					},
//   				},
//   			},
//   			"Inputs": []map[string]interface{}{
//   				map[string]interface{}{
//   					"AudioSelectors": map[string]map[string]*string{
//   						"Audio Selector 1": map[string]*string{
//   							"DefaultSelection": jsii.String("DEFAULT"),
//   						},
//   					},
//   					"FileInput": jsii.String("s3://EXAMPLE-SOURCE-BUCKET/EXAMPLE-SOURCE_FILE"),
//   				},
//   			},
//   		},
//   	},
//   	IntegrationPattern: sfn.IntegrationPattern_RUN_JOB,
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-mediaconvert.html
//
// Response syntax: see CreateJobResponse schema
// https://docs.aws.amazon.com/mediaconvert/latest/apireference/jobs.html#jobs-response-examples
//
type MediaConvertCreateJob interface {
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

// The jsii proxy struct for MediaConvertCreateJob
type jsiiProxy_MediaConvertCreateJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_MediaConvertCreateJob) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) Processor() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) ProcessorConfig() *awsstepfunctions.ProcessorConfig {
	var returns *awsstepfunctions.ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) ProcessorMode() awsstepfunctions.ProcessorMode {
	var returns awsstepfunctions.ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) QueryLanguage() awsstepfunctions.QueryLanguage {
	var returns awsstepfunctions.QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MediaConvertCreateJob) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


func NewMediaConvertCreateJob(scope constructs.Construct, id *string, props *MediaConvertCreateJobProps) MediaConvertCreateJob {
	_init_.Initialize()

	if err := validateNewMediaConvertCreateJobParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MediaConvertCreateJob{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMediaConvertCreateJob_Override(m MediaConvertCreateJob, scope constructs.Construct, id *string, props *MediaConvertCreateJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		[]interface{}{scope, id, props},
		m,
	)
}

func (j *jsiiProxy_MediaConvertCreateJob)SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_MediaConvertCreateJob)SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_MediaConvertCreateJob)SetProcessor(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_MediaConvertCreateJob)SetProcessorConfig(val *awsstepfunctions.ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_MediaConvertCreateJob)SetProcessorMode(val awsstepfunctions.ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func MediaConvertCreateJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func MediaConvertCreateJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func MediaConvertCreateJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
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
func MediaConvertCreateJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// A Step Functions Task to create a job in MediaConvert using JSONata.
func MediaConvertCreateJob_Jsonata(scope constructs.Construct, id *string, props *MediaConvertCreateJobJsonataProps) MediaConvertCreateJob {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_JsonataParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns MediaConvertCreateJob

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"jsonata",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// A Step Functions Task to create a job in MediaConvert using JSONPath.
func MediaConvertCreateJob_JsonPath(scope constructs.Construct, id *string, props *MediaConvertCreateJobJsonPathProps) MediaConvertCreateJob {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_JsonPathParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns MediaConvertCreateJob

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"jsonPath",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func MediaConvertCreateJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateMediaConvertCreateJob_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions_tasks.MediaConvertCreateJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddBranch(branch awsstepfunctions.StateGraph) {
	if err := m.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addBranch",
		[]interface{}{branch},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	if err := m.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		m,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State, options *awsstepfunctions.ChoiceTransitionOptions) {
	if err := m.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddItemProcessor(processor awsstepfunctions.StateGraph, config *awsstepfunctions.ProcessorConfig) {
	if err := m.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	if err := m.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddPrefix(x *string) {
	if err := m.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addPrefix",
		[]interface{}{x},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	if err := m.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		m,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	if err := m.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) MakeDefault(def awsstepfunctions.State) {
	if err := m.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeDefault",
		[]interface{}{def},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) MakeNext(next awsstepfunctions.State) {
	if err := m.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeNext",
		[]interface{}{next},
	)
}

func (m *jsiiProxy_MediaConvertCreateJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricFailedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricHeartbeatTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricRunTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricScheduledParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricScheduleTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricStartedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricSucceededParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricTimeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := m.validateMetricTimedOutParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		m,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	if err := m.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		m,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderAssign(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderChoices(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderQueryLanguage(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) RenderRetryCatch(topLevelQueryLanguage awsstepfunctions.QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) ToStateJson(topLevelQueryLanguage awsstepfunctions.QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"toStateJson",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MediaConvertCreateJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	if err := m.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

