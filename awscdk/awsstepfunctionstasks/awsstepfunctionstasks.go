package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/awscodebuild"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/awsecr"
	"github.com/aws/aws-cdk-go/awscdk/awsecrassets"
	"github.com/aws/aws-cdk-go/awscdk/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/awseks"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awskms"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// The generation of Elastic Inference (EI) instance.
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html
//
// Experimental.
type AcceleratorClass interface {
	Version() *string
}

// The jsii proxy struct for AcceleratorClass
type jsiiProxy_AcceleratorClass struct {
	_ byte // padding
}

func (j *jsiiProxy_AcceleratorClass) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom AcceleratorType.
// Experimental.
func AcceleratorClass_Of(version *string) AcceleratorClass {
	_init_.Initialize()

	var returns AcceleratorClass

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AcceleratorClass",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func AcceleratorClass_EIA1() AcceleratorClass {
	_init_.Initialize()
	var returns AcceleratorClass
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.AcceleratorClass",
		"EIA1",
		&returns,
	)
	return returns
}

func AcceleratorClass_EIA2() AcceleratorClass {
	_init_.Initialize()
	var returns AcceleratorClass
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.AcceleratorClass",
		"EIA2",
		&returns,
	)
	return returns
}

// The size of the Elastic Inference (EI) instance to use for the production variant.
//
// EI instances provide on-demand GPU computing for inference
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html
//
// Experimental.
type AcceleratorType interface {
	ToString() *string
}

// The jsii proxy struct for AcceleratorType
type jsiiProxy_AcceleratorType struct {
	_ byte // padding
}

// Experimental.
func NewAcceleratorType(instanceTypeIdentifier *string) AcceleratorType {
	_init_.Initialize()

	j := jsiiProxy_AcceleratorType{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		[]interface{}{instanceTypeIdentifier},
		&j,
	)

	return &j
}

// Experimental.
func NewAcceleratorType_Override(a AcceleratorType, instanceTypeIdentifier *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		[]interface{}{instanceTypeIdentifier},
		a,
	)
}

// AcceleratorType.
//
// This class takes a combination of a class and size.
// Experimental.
func AcceleratorType_Of(acceleratorClass AcceleratorClass, instanceSize awsec2.InstanceSize) AcceleratorType {
	_init_.Initialize()

	var returns AcceleratorType

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		"of",
		[]interface{}{acceleratorClass, instanceSize},
		&returns,
	)

	return returns
}

// Return the accelerator type as a dotted string.
// Experimental.
func (a *jsiiProxy_AcceleratorType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The action to take when the cluster step fails.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
//
// Here, they are named as TERMINATE_JOB_FLOW, TERMINATE_CLUSTER, CANCEL_AND_WAIT, and CONTINUE respectively.
//
// Experimental.
type ActionOnFailure string

const (
	ActionOnFailure_TERMINATE_CLUSTER ActionOnFailure = "TERMINATE_CLUSTER"
	ActionOnFailure_CANCEL_AND_WAIT ActionOnFailure = "CANCEL_AND_WAIT"
	ActionOnFailure_CONTINUE ActionOnFailure = "CONTINUE"
)

// Specify the training algorithm and algorithm-specific metadata.
// Experimental.
type AlgorithmSpecification struct {
	// Name of the algorithm resource to use for the training job.
	//
	// This must be an algorithm resource that you created or subscribe to on AWS Marketplace.
	// If you specify a value for this parameter, you can't specify a value for TrainingImage.
	// Experimental.
	AlgorithmName *string `json:"algorithmName"`
	// List of metric definition objects.
	//
	// Each object specifies the metric name and regular expressions used to parse algorithm logs.
	// Experimental.
	MetricDefinitions *[]*MetricDefinition `json:"metricDefinitions"`
	// Registry path of the Docker image that contains the training algorithm.
	// Experimental.
	TrainingImage DockerImage `json:"trainingImage"`
	// Input mode that the algorithm supports.
	// Experimental.
	TrainingInputMode InputMode `json:"trainingInputMode"`
}

// How to assemble the results of the transform job as a single S3 object.
// Experimental.
type AssembleWith string

const (
	AssembleWith_NONE AssembleWith = "NONE"
	AssembleWith_LINE AssembleWith = "LINE"
)

// Get an Athena Query Execution as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
// Experimental.
type AthenaGetQueryExecution interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaGetQueryExecution
type jsiiProxy_AthenaGetQueryExecution struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaGetQueryExecution) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryExecution) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewAthenaGetQueryExecution(scope constructs.Construct, id *string, props *AthenaGetQueryExecutionProps) AthenaGetQueryExecution {
	_init_.Initialize()

	j := jsiiProxy_AthenaGetQueryExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaGetQueryExecution_Override(a AthenaGetQueryExecution, scope constructs.Construct, id *string, props *AthenaGetQueryExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaGetQueryExecution) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryExecution) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func AthenaGetQueryExecution_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func AthenaGetQueryExecution_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
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
func AthenaGetQueryExecution_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AthenaGetQueryExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func AthenaGetQueryExecution_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryExecution",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryExecution) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for getting a Query Execution.
// Experimental.
type AthenaGetQueryExecutionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Query that will be retrieved.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	QueryExecutionId *string `json:"queryExecutionId"`
}

// Get an Athena Query Results as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
// Experimental.
type AthenaGetQueryResults interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaGetQueryResults
type jsiiProxy_AthenaGetQueryResults struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaGetQueryResults) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaGetQueryResults) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewAthenaGetQueryResults(scope constructs.Construct, id *string, props *AthenaGetQueryResultsProps) AthenaGetQueryResults {
	_init_.Initialize()

	j := jsiiProxy_AthenaGetQueryResults{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaGetQueryResults_Override(a AthenaGetQueryResults, scope constructs.Construct, id *string, props *AthenaGetQueryResultsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaGetQueryResults) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func AthenaGetQueryResults_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func AthenaGetQueryResults_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
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
func AthenaGetQueryResults_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AthenaGetQueryResults_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func AthenaGetQueryResults_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.AthenaGetQueryResults",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (a *jsiiProxy_AthenaGetQueryResults) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for getting a Query Results.
// Experimental.
type AthenaGetQueryResultsProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Query that will be retrieved.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	QueryExecutionId *string `json:"queryExecutionId"`
	// Max number of results.
	// Experimental.
	MaxResults *float64 `json:"maxResults"`
	// Pagination token.
	// Experimental.
	NextToken *string `json:"nextToken"`
}

// Start an Athena Query as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
// Experimental.
type AthenaStartQueryExecution interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaStartQueryExecution
type jsiiProxy_AthenaStartQueryExecution struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaStartQueryExecution) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStartQueryExecution) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewAthenaStartQueryExecution(scope constructs.Construct, id *string, props *AthenaStartQueryExecutionProps) AthenaStartQueryExecution {
	_init_.Initialize()

	j := jsiiProxy_AthenaStartQueryExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaStartQueryExecution_Override(a AthenaStartQueryExecution, scope constructs.Construct, id *string, props *AthenaStartQueryExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaStartQueryExecution) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaStartQueryExecution) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func AthenaStartQueryExecution_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func AthenaStartQueryExecution_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
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
func AthenaStartQueryExecution_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AthenaStartQueryExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func AthenaStartQueryExecution_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.AthenaStartQueryExecution",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (a *jsiiProxy_AthenaStartQueryExecution) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for starting a Query Execution.
// Experimental.
type AthenaStartQueryExecutionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Query that will be started.
	// Experimental.
	QueryString *string `json:"queryString"`
	// Unique string string to ensure idempotence.
	// Experimental.
	ClientRequestToken *string `json:"clientRequestToken"`
	// Database within which query executes.
	// Experimental.
	QueryExecutionContext *QueryExecutionContext `json:"queryExecutionContext"`
	// Configuration on how and where to save query.
	// Experimental.
	ResultConfiguration *ResultConfiguration `json:"resultConfiguration"`
	// Configuration on how and where to save query.
	// Experimental.
	WorkGroup *string `json:"workGroup"`
}

// Stop an Athena Query Execution as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-athena.html
//
// Experimental.
type AthenaStopQueryExecution interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for AthenaStopQueryExecution
type jsiiProxy_AthenaStopQueryExecution struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_AthenaStopQueryExecution) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AthenaStopQueryExecution) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewAthenaStopQueryExecution(scope constructs.Construct, id *string, props *AthenaStopQueryExecutionProps) AthenaStopQueryExecution {
	_init_.Initialize()

	j := jsiiProxy_AthenaStopQueryExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAthenaStopQueryExecution_Override(a AthenaStopQueryExecution, scope constructs.Construct, id *string, props *AthenaStopQueryExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AthenaStopQueryExecution) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_AthenaStopQueryExecution) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func AthenaStopQueryExecution_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func AthenaStopQueryExecution_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
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
func AthenaStopQueryExecution_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func AthenaStopQueryExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func AthenaStopQueryExecution_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.AthenaStopQueryExecution",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		a,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		a,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		a,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		a,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (a *jsiiProxy_AthenaStopQueryExecution) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		a,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for stoping a Query Execution.
// Experimental.
type AthenaStopQueryExecutionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Query that will be stopped.
	// Experimental.
	QueryExecutionId *string `json:"queryExecutionId"`
}

// The authentication method used to call the endpoint.
// Experimental.
type AuthType string

const (
	AuthType_NO_AUTH AuthType = "NO_AUTH"
	AuthType_IAM_ROLE AuthType = "IAM_ROLE"
	AuthType_RESOURCE_POLICY AuthType = "RESOURCE_POLICY"
)

// The overrides that should be sent to a container.
// Experimental.
type BatchContainerOverrides struct {
	// The command to send to the container that overrides the default command from the Docker image or the job definition.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container
	// at launch, or you can override the existing environment variables from
	// the Docker image or the job definition.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all containers in a job
	// should not exceed the number of available GPUs on the compute
	// resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `json:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// This parameter is not valid for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// Memory reserved for the job.
	// Experimental.
	Memory awscdk.Size `json:"memory"`
	// The number of vCPUs to reserve for the container.
	//
	// This value overrides the value set in the job definition.
	// Experimental.
	Vcpus *float64 `json:"vcpus"`
}

// An object representing an AWS Batch job dependency.
// Experimental.
type BatchJobDependency struct {
	// The job ID of the AWS Batch job associated with this dependency.
	// Experimental.
	JobId *string `json:"jobId"`
	// The type of the job dependency.
	// Experimental.
	Type *string `json:"type"`
}

// Specifies the number of records to include in a mini-batch for an HTTP inference request.
// Experimental.
type BatchStrategy string

const (
	BatchStrategy_MULTI_RECORD BatchStrategy = "MULTI_RECORD"
	BatchStrategy_SINGLE_RECORD BatchStrategy = "SINGLE_RECORD"
)

// Task to submits an AWS Batch job from a job definition.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-batch.html
//
// Experimental.
type BatchSubmitJob interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_BatchSubmitJob) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewBatchSubmitJob(scope constructs.Construct, id *string, props *BatchSubmitJobProps) BatchSubmitJob {
	_init_.Initialize()

	j := jsiiProxy_BatchSubmitJob{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBatchSubmitJob_Override(b BatchSubmitJob, scope constructs.Construct, id *string, props *BatchSubmitJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		[]interface{}{scope, id, props},
		b,
	)
}

func (j *jsiiProxy_BatchSubmitJob) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_BatchSubmitJob) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func BatchSubmitJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func BatchSubmitJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
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
func BatchSubmitJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func BatchSubmitJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func BatchSubmitJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.BatchSubmitJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		b,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		b,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		b,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		b,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		b,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		b,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		b,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		b,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		b,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		b,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) OnPrepare() {
	_jsii_.InvokeVoid(
		b,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) Prepare() {
	_jsii_.InvokeVoid(
		b,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
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

// Render the choices in ASL JSON format.
// Experimental.
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

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
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

// Render map iterator in ASL JSON format.
// Experimental.
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

// Render the default next state in ASL JSON format.
// Experimental.
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

// Render ResultSelector in ASL JSON format.
// Experimental.
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

// Render error recovery options in ASL JSON format.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		b,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (b *jsiiProxy_BatchSubmitJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		b,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for RunBatchJob.
// Experimental.
type BatchSubmitJobProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The arn of the job definition used by this job.
	// Experimental.
	JobDefinitionArn *string `json:"jobDefinitionArn"`
	// The name of the job.
	//
	// The first character must be alphanumeric, and up to 128 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobName *string `json:"jobName"`
	// The arn of the job queue into which the job is submitted.
	// Experimental.
	JobQueueArn *string `json:"jobQueueArn"`
	// The array size can be between 2 and 10,000.
	//
	// If you specify array properties for a job, it becomes an array job.
	// For more information, see Array Jobs in the AWS Batch User Guide.
	// Experimental.
	ArraySize *float64 `json:"arraySize"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and 10 attempts.
	// If the value of attempts is greater than one,
	// the job is retried on failure the same number of attempts as the value.
	// Experimental.
	Attempts *float64 `json:"attempts"`
	// A list of container overrides in JSON format that specify the name of a container in the specified job definition and the overrides it should receive.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-containerOverrides
	//
	// Experimental.
	ContainerOverrides *BatchContainerOverrides `json:"containerOverrides"`
	// A list of dependencies for the job.
	//
	// A job can depend upon a maximum of 20 jobs.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-dependsOn
	//
	// Experimental.
	DependsOn *[]*BatchJobDependency `json:"dependsOn"`
	// The payload to be passed as parameters to the batch job.
	// Experimental.
	Payload awsstepfunctions.TaskInput `json:"payload"`
}

// Base CallApiGatewayEdnpoint Task Props.
// Experimental.
type CallApiGatewayEndpointBaseProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Http method for the API.
	// Experimental.
	Method HttpMethod `json:"method"`
	// Path parameters appended after API endpoint.
	// Experimental.
	ApiPath *string `json:"apiPath"`
	// Authentication methods.
	// Experimental.
	AuthType AuthType `json:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Experimental.
	Headers awsstepfunctions.TaskInput `json:"headers"`
	// Query strings attatched to end of request.
	// Experimental.
	QueryParameters awsstepfunctions.TaskInput `json:"queryParameters"`
	// HTTP Request body.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `json:"requestBody"`
}

// Call HTTP API endpoint as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-api-gateway.html
//
// Experimental.
type CallApiGatewayHttpApiEndpoint interface {
	awsstepfunctions.TaskStateBase
	ApiEndpoint() *string
	ArnForExecuteApi() *string
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StageName() *string
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	CreatePolicyStatements() *[]awsiam.PolicyStatement
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for CallApiGatewayHttpApiEndpoint
type jsiiProxy_CallApiGatewayHttpApiEndpoint struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) ArnForExecuteApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForExecuteApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewCallApiGatewayHttpApiEndpoint(scope constructs.Construct, id *string, props *CallApiGatewayHttpApiEndpointProps) CallApiGatewayHttpApiEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CallApiGatewayHttpApiEndpoint{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCallApiGatewayHttpApiEndpoint_Override(c CallApiGatewayHttpApiEndpoint, scope constructs.Construct, id *string, props *CallApiGatewayHttpApiEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CallApiGatewayHttpApiEndpoint) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func CallApiGatewayHttpApiEndpoint_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func CallApiGatewayHttpApiEndpoint_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
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
func CallApiGatewayHttpApiEndpoint_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CallApiGatewayHttpApiEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func CallApiGatewayHttpApiEndpoint_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayHttpApiEndpoint",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) CreatePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		c,
		"createPolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (c *jsiiProxy_CallApiGatewayHttpApiEndpoint) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for calling an HTTP API Endpoint.
// Experimental.
type CallApiGatewayHttpApiEndpointProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Http method for the API.
	// Experimental.
	Method HttpMethod `json:"method"`
	// Path parameters appended after API endpoint.
	// Experimental.
	ApiPath *string `json:"apiPath"`
	// Authentication methods.
	// Experimental.
	AuthType AuthType `json:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Experimental.
	Headers awsstepfunctions.TaskInput `json:"headers"`
	// Query strings attatched to end of request.
	// Experimental.
	QueryParameters awsstepfunctions.TaskInput `json:"queryParameters"`
	// HTTP Request body.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `json:"requestBody"`
	// The Id of the API to call.
	// Experimental.
	ApiId *string `json:"apiId"`
	// The Stack in which the API is defined.
	// Experimental.
	ApiStack awscdk.Stack `json:"apiStack"`
	// Name of the stage where the API is deployed to in API Gateway.
	// Experimental.
	StageName *string `json:"stageName"`
}

// Call REST API endpoint as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-api-gateway.html
//
// Experimental.
type CallApiGatewayRestApiEndpoint interface {
	awsstepfunctions.TaskStateBase
	ApiEndpoint() *string
	ArnForExecuteApi() *string
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StageName() *string
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	CreatePolicyStatements() *[]awsiam.PolicyStatement
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for CallApiGatewayRestApiEndpoint
type jsiiProxy_CallApiGatewayRestApiEndpoint struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) ApiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) ArnForExecuteApi() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arnForExecuteApi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewCallApiGatewayRestApiEndpoint(scope constructs.Construct, id *string, props *CallApiGatewayRestApiEndpointProps) CallApiGatewayRestApiEndpoint {
	_init_.Initialize()

	j := jsiiProxy_CallApiGatewayRestApiEndpoint{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCallApiGatewayRestApiEndpoint_Override(c CallApiGatewayRestApiEndpoint, scope constructs.Construct, id *string, props *CallApiGatewayRestApiEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CallApiGatewayRestApiEndpoint) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func CallApiGatewayRestApiEndpoint_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func CallApiGatewayRestApiEndpoint_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
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
func CallApiGatewayRestApiEndpoint_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CallApiGatewayRestApiEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func CallApiGatewayRestApiEndpoint_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.CallApiGatewayRestApiEndpoint",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) CreatePolicyStatements() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement

	_jsii_.Invoke(
		c,
		"createPolicyStatements",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for calling an REST API Endpoint.
// Experimental.
type CallApiGatewayRestApiEndpointProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Http method for the API.
	// Experimental.
	Method HttpMethod `json:"method"`
	// Path parameters appended after API endpoint.
	// Experimental.
	ApiPath *string `json:"apiPath"`
	// Authentication methods.
	// Experimental.
	AuthType AuthType `json:"authType"`
	// HTTP request information that does not relate to contents of the request.
	// Experimental.
	Headers awsstepfunctions.TaskInput `json:"headers"`
	// Query strings attatched to end of request.
	// Experimental.
	QueryParameters awsstepfunctions.TaskInput `json:"queryParameters"`
	// HTTP Request body.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `json:"requestBody"`
	// API to call.
	// Experimental.
	Api awsapigateway.IRestApi `json:"api"`
	// Name of the stage where the API is deployed to in API Gateway.
	// Experimental.
	StageName *string `json:"stageName"`
}

// A StepFunctions task to call an AWS service API.
// Experimental.
type CallAwsService interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for CallAwsService
type jsiiProxy_CallAwsService struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_CallAwsService) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CallAwsService) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewCallAwsService(scope constructs.Construct, id *string, props *CallAwsServiceProps) CallAwsService {
	_init_.Initialize()

	j := jsiiProxy_CallAwsService{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCallAwsService_Override(c CallAwsService, scope constructs.Construct, id *string, props *CallAwsServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CallAwsService) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CallAwsService) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func CallAwsService_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func CallAwsService_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
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
func CallAwsService_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CallAwsService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func CallAwsService_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.CallAwsService",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (c *jsiiProxy_CallAwsService) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (c *jsiiProxy_CallAwsService) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (c *jsiiProxy_CallAwsService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (c *jsiiProxy_CallAwsService) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (c *jsiiProxy_CallAwsService) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallAwsService) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallAwsService) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallAwsService) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CallAwsService) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CallAwsService) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CallAwsService) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (c *jsiiProxy_CallAwsService) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (c *jsiiProxy_CallAwsService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CallAwsService) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (c *jsiiProxy_CallAwsService) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for calling an AWS service's API action from your state machine.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
//
// Experimental.
type CallAwsServiceProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The API action to call.
	//
	// Use camelCase.
	// Experimental.
	Action *string `json:"action"`
	// The resources for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// By default the action for this IAM statement will be `service:action`.
	// Experimental.
	IamResources *[]*string `json:"iamResources"`
	// The AWS service to call.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/supported-services-awssdk.html
	//
	// Experimental.
	Service *string `json:"service"`
	// The action for the IAM statement that will be added to the state machine role's policy to allow the state machine to make the API call.
	//
	// Use in the case where the IAM action name does not match with the
	// API service/action name, e.g. `s3:ListBuckets` requires `s3:ListAllMyBuckets`.
	// Experimental.
	IamAction *string `json:"iamAction"`
	// Parameters for the API action call.
	//
	// Use PascalCase for the parameter names.
	// Experimental.
	Parameters *map[string]interface{} `json:"parameters"`
}

// Describes the training, validation or test dataset and the Amazon S3 location where it is stored.
// Experimental.
type Channel struct {
	// Name of the channel.
	// Experimental.
	ChannelName *string `json:"channelName"`
	// Location of the channel data.
	// Experimental.
	DataSource *DataSource `json:"dataSource"`
	// Compression type if training data is compressed.
	// Experimental.
	CompressionType CompressionType `json:"compressionType"`
	// The MIME type of the data.
	// Experimental.
	ContentType *string `json:"contentType"`
	// Input mode to use for the data channel in a training job.
	// Experimental.
	InputMode InputMode `json:"inputMode"`
	// Specify RecordIO as the value when input data is in raw format but the training algorithm requires the RecordIO format.
	//
	// In this case, Amazon SageMaker wraps each individual S3 object in a RecordIO record.
	// If the input data is already in RecordIO format, you don't need to set this attribute.
	// Experimental.
	RecordWrapperType RecordWrapperType `json:"recordWrapperType"`
	// Shuffle config option for input data in a channel.
	// Experimental.
	ShuffleConfig *ShuffleConfig `json:"shuffleConfig"`
}

// Start a CodeBuild Build as a task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-codebuild.html
//
// Experimental.
type CodeBuildStartBuild interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for CodeBuildStartBuild
type jsiiProxy_CodeBuildStartBuild struct {
	internal.Type__awsstepfunctionsTaskStateBase
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

func (j *jsiiProxy_CodeBuildStartBuild) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_CodeBuildStartBuild) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
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


// Experimental.
func NewCodeBuildStartBuild(scope constructs.Construct, id *string, props *CodeBuildStartBuildProps) CodeBuildStartBuild {
	_init_.Initialize()

	j := jsiiProxy_CodeBuildStartBuild{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildStartBuild_Override(c CodeBuildStartBuild, scope constructs.Construct, id *string, props *CodeBuildStartBuildProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CodeBuildStartBuild) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func CodeBuildStartBuild_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func CodeBuildStartBuild_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
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
func CodeBuildStartBuild_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CodeBuildStartBuild_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func CodeBuildStartBuild_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.CodeBuildStartBuild",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		c,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		c,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
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

// Render the choices in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
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

// Render map iterator in ASL JSON format.
// Experimental.
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

// Render the default next state in ASL JSON format.
// Experimental.
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

// Render ResultSelector in ASL JSON format.
// Experimental.
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

// Render error recovery options in ASL JSON format.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (c *jsiiProxy_CodeBuildStartBuild) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for CodeBuildStartBuild.
// Experimental.
type CodeBuildStartBuildProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// CodeBuild project to start.
	// Experimental.
	Project awscodebuild.IProject `json:"project"`
	// A set of environment variables to be used for this build only.
	// Experimental.
	EnvironmentVariablesOverride *map[string]*awscodebuild.BuildEnvironmentVariable `json:"environmentVariablesOverride"`
}

// Basic properties for ECS Tasks.
// Experimental.
type CommonEcsRunTaskProps struct {
	// The topic to run the task on.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// Experimental.
	TaskDefinition awsecs.TaskDefinition `json:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Experimental.
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Experimental.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
}

// Compression type of the data.
// Experimental.
type CompressionType string

const (
	CompressionType_NONE CompressionType = "NONE"
	CompressionType_GZIP CompressionType = "GZIP"
)

// Describes the container, as part of model definition.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ContainerDefinition.html
//
// Experimental.
type ContainerDefinition interface {
	IContainerDefinition
	Bind(task ISageMakerTask) *ContainerDefinitionConfig
}

// The jsii proxy struct for ContainerDefinition
type jsiiProxy_ContainerDefinition struct {
	jsiiProxy_IContainerDefinition
}

// Experimental.
func NewContainerDefinition(options *ContainerDefinitionOptions) ContainerDefinition {
	_init_.Initialize()

	j := jsiiProxy_ContainerDefinition{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ContainerDefinition",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewContainerDefinition_Override(c ContainerDefinition, options *ContainerDefinitionOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.ContainerDefinition",
		[]interface{}{options},
		c,
	)
}

// Called when the ContainerDefinition type configured on Sagemaker Task.
// Experimental.
func (c *jsiiProxy_ContainerDefinition) Bind(task ISageMakerTask) *ContainerDefinitionConfig {
	var returns *ContainerDefinitionConfig

	_jsii_.Invoke(
		c,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Configuration options for the ContainerDefinition.
// Experimental.
type ContainerDefinitionConfig struct {
	// Additional parameters to pass to the base task.
	// Experimental.
	Parameters *map[string]interface{} `json:"parameters"`
}

// Properties to define a ContainerDefinition.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ContainerDefinition.html
//
// Experimental.
type ContainerDefinitionOptions struct {
	// This parameter is ignored for models that contain only a PrimaryContainer.
	//
	// When a ContainerDefinition is part of an inference pipeline,
	// the value of the parameter uniquely identifies the container for the purposes of logging and metrics.
	// Experimental.
	ContainerHostName *string `json:"containerHostName"`
	// The environment variables to set in the Docker container.
	// Experimental.
	EnvironmentVariables awsstepfunctions.TaskInput `json:"environmentVariables"`
	// The Amazon EC2 Container Registry (Amazon ECR) path where inference code is stored.
	// Experimental.
	Image DockerImage `json:"image"`
	// Defines how many models the container hosts.
	// Experimental.
	Mode Mode `json:"mode"`
	// The name or Amazon Resource Name (ARN) of the model package to use to create the model.
	// Experimental.
	ModelPackageName *string `json:"modelPackageName"`
	// The S3 path where the model artifacts, which result from model training, are stored.
	//
	// This path must point to a single gzip compressed tar archive (.tar.gz suffix).
	// The S3 path is required for Amazon SageMaker built-in algorithms, but not if you use your own algorithms.
	// Experimental.
	ModelS3Location S3Location `json:"modelS3Location"`
}

// A list of container overrides that specify the name of a container and the overrides it should receive.
// Experimental.
type ContainerOverride struct {
	// Name of the container inside the task definition.
	// Experimental.
	ContainerDefinition awsecs.ContainerDefinition `json:"containerDefinition"`
	// Command to run inside the container.
	// Experimental.
	Command *[]*string `json:"command"`
	// The number of cpu units reserved for the container.
	// Experimental.
	Cpu *float64 `json:"cpu"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container at launch,
	// or you can override the existing environment variables from the Docker image or the task definition.
	// Experimental.
	Environment *[]*TaskEnvironmentVariable `json:"environment"`
	// The hard limit (in MiB) of memory to present to the container.
	// Experimental.
	MemoryLimit *float64 `json:"memoryLimit"`
	// The soft limit (in MiB) of memory to reserve for the container.
	// Experimental.
	MemoryReservation *float64 `json:"memoryReservation"`
}

// The overrides that should be sent to a container.
// Experimental.
type ContainerOverrides struct {
	// The command to send to the container that overrides the default command from the Docker image or the job definition.
	// Experimental.
	Command *[]*string `json:"command"`
	// The environment variables to send to the container.
	//
	// You can add new environment variables, which are added to the container
	// at launch, or you can override the existing environment variables from
	// the Docker image or the job definition.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// The number of physical GPUs to reserve for the container.
	//
	// The number of GPUs reserved for all containers in a job
	// should not exceed the number of available GPUs on the compute
	// resource that the job is launched on.
	// Experimental.
	GpuCount *float64 `json:"gpuCount"`
	// The instance type to use for a multi-node parallel job.
	//
	// This parameter is not valid for single-node container jobs.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The number of MiB of memory reserved for the job.
	//
	// This value overrides the value set in the job definition.
	// Experimental.
	Memory *float64 `json:"memory"`
	// The number of vCPUs to reserve for the container.
	//
	// This value overrides the value set in the job definition.
	// Experimental.
	Vcpus *float64 `json:"vcpus"`
}

// Location of the channel data.
// Experimental.
type DataSource struct {
	// S3 location of the data source that is associated with a channel.
	// Experimental.
	S3DataSource *S3DataSource `json:"s3DataSource"`
}

// Creates `IDockerImage` instances.
// Experimental.
type DockerImage interface {
	Bind(task ISageMakerTask) *DockerImageConfig
}

// The jsii proxy struct for DockerImage
type jsiiProxy_DockerImage struct {
	_ byte // padding
}

// Experimental.
func NewDockerImage_Override(d DockerImage) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DockerImage",
		nil, // no parameters
		d,
	)
}

// Reference a Docker image that is provided as an Asset in the current app.
// Experimental.
func DockerImage_FromAsset(scope constructs.Construct, id *string, props *awsecrassets.DockerImageAssetProps) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DockerImage",
		"fromAsset",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

// Reference a Docker image stored in an ECR repository.
// Experimental.
func DockerImage_FromEcrRepository(repository awsecr.IRepository, tag *string) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DockerImage",
		"fromEcrRepository",
		[]interface{}{repository, tag},
		&returns,
	)

	return returns
}

// Reference a Docker image which URI is obtained from the task's input.
// Experimental.
func DockerImage_FromJsonExpression(expression *string, allowAnyEcrImagePull *bool) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DockerImage",
		"fromJsonExpression",
		[]interface{}{expression, allowAnyEcrImagePull},
		&returns,
	)

	return returns
}

// Reference a Docker image by it's URI.
//
// When referencing ECR images, prefer using `inEcr`.
// Experimental.
func DockerImage_FromRegistry(imageUri *string) DockerImage {
	_init_.Initialize()

	var returns DockerImage

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DockerImage",
		"fromRegistry",
		[]interface{}{imageUri},
		&returns,
	)

	return returns
}

// Called when the image is used by a SageMaker task.
// Experimental.
func (d *jsiiProxy_DockerImage) Bind(task ISageMakerTask) *DockerImageConfig {
	var returns *DockerImageConfig

	_jsii_.Invoke(
		d,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Configuration for a using Docker image.
// Experimental.
type DockerImageConfig struct {
	// The fully qualified URI of the Docker image.
	// Experimental.
	ImageUri *string `json:"imageUri"`
}

// Represents the data for an attribute.
//
// Each attribute value is described as a name-value pair.
// The name is the data type, and the value is the data itself.
// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_AttributeValue.html
//
// Experimental.
type DynamoAttributeValue interface {
	AttributeValue() interface{}
	ToObject() interface{}
}

// The jsii proxy struct for DynamoAttributeValue
type jsiiProxy_DynamoAttributeValue struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamoAttributeValue) AttributeValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeValue",
		&returns,
	)
	return returns
}


// Sets an attribute of type Boolean from state input through Json path.
//
// For example:  "BOOL": true
// Experimental.
func DynamoAttributeValue_BooleanFromJsonPath(value *string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"booleanFromJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Binary.
//
// For example:  "B": "dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk"
// Experimental.
func DynamoAttributeValue_FromBinary(value *string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBinary",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Binary Set.
//
// For example:  "BS": ["U3Vubnk=", "UmFpbnk=", "U25vd3k="]
// Experimental.
func DynamoAttributeValue_FromBinarySet(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBinarySet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Boolean.
//
// For example:  "BOOL": true
// Experimental.
func DynamoAttributeValue_FromBoolean(value *bool) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type List.
//
// For example:  "L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"N", "3.14159"}]
// Experimental.
func DynamoAttributeValue_FromList(value *[]DynamoAttributeValue) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromList",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Map.
//
// For example:  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}
// Experimental.
func DynamoAttributeValue_FromMap(value *map[string]DynamoAttributeValue) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromMap",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Null.
//
// For example:  "NULL": true
// Experimental.
func DynamoAttributeValue_FromNull(value *bool) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNull",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets a literal number.
//
// For example: 1234
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
// Experimental.
func DynamoAttributeValue_FromNumber(value *float64) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNumber",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number Set.
//
// For example:  "NS": ["42.2", "-19", "7.5", "3.14"]
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
// Experimental.
func DynamoAttributeValue_FromNumberSet(value *[]*float64) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromNumberSet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type String.
//
// For example:  "S": "Hello"
// Strings may be literal values or as JsonPath.
//
// TODO: EXAMPLE
//
// Experimental.
func DynamoAttributeValue_FromString(value *string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type String Set.
//
// For example:  "SS": ["Giraffe", "Hippo" ,"Zebra"]
// Experimental.
func DynamoAttributeValue_FromStringSet(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"fromStringSet",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Map.
//
// For example:  "M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}
// Experimental.
func DynamoAttributeValue_MapFromJsonPath(value *string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"mapFromJsonPath",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number.
//
// For example:  "N": "123.45"
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
//
// Numbers may be expressed as literal strings or as JsonPath
// Experimental.
func DynamoAttributeValue_NumberFromString(value *string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"numberFromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Sets an attribute of type Number Set.
//
// For example:  "NS": ["42.2", "-19", "7.5", "3.14"]
// Numbers are sent across the network to DynamoDB as strings,
// to maximize compatibility across languages and libraries.
// However, DynamoDB treats them as number type attributes for mathematical operations.
//
// Numbers may be expressed as literal strings or as JsonPath
// Experimental.
func DynamoAttributeValue_NumberSetFromStrings(value *[]*string) DynamoAttributeValue {
	_init_.Initialize()

	var returns DynamoAttributeValue

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoAttributeValue",
		"numberSetFromStrings",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Returns the DynamoDB attribute value.
// Experimental.
func (d *jsiiProxy_DynamoAttributeValue) ToObject() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toObject",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Determines the level of detail about provisioned throughput consumption that is returned.
// Experimental.
type DynamoConsumedCapacity string

const (
	DynamoConsumedCapacity_INDEXES DynamoConsumedCapacity = "INDEXES"
	DynamoConsumedCapacity_TOTAL DynamoConsumedCapacity = "TOTAL"
	DynamoConsumedCapacity_NONE DynamoConsumedCapacity = "NONE"
)

// A StepFunctions task to call DynamoDeleteItem.
// Experimental.
type DynamoDeleteItem interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for DynamoDeleteItem
type jsiiProxy_DynamoDeleteItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_DynamoDeleteItem) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoDeleteItem) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoDeleteItem(scope constructs.Construct, id *string, props *DynamoDeleteItemProps) DynamoDeleteItem {
	_init_.Initialize()

	j := jsiiProxy_DynamoDeleteItem{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoDeleteItem_Override(d DynamoDeleteItem, scope constructs.Construct, id *string, props *DynamoDeleteItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoDeleteItem) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoDeleteItem) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func DynamoDeleteItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func DynamoDeleteItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
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
func DynamoDeleteItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoDeleteItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func DynamoDeleteItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.DynamoDeleteItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (d *jsiiProxy_DynamoDeleteItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for DynamoDeleteItem Task.
// Experimental.
type DynamoDeleteItemProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Primary key of the item to retrieve.
	//
	// For the primary key, you must provide all of the attributes.
	// For example, with a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key and the sort key.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-Key
	//
	// Experimental.
	Key *map[string]DynamoAttributeValue `json:"key"`
	// The name of the table containing the requested item.
	// Experimental.
	Table awsdynamodb.ITable `json:"table"`
	// A condition that must be satisfied in order for a conditional DeleteItem to succeed.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ConditionExpression
	//
	// Experimental.
	ConditionExpression *string `json:"conditionExpression"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ExpressionAttributeNames
	//
	// Experimental.
	ExpressionAttributeNames *map[string]*string `json:"expressionAttributeNames"`
	// One or more values that can be substituted in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ExpressionAttributeValues
	//
	// Experimental.
	ExpressionAttributeValues *map[string]DynamoAttributeValue `json:"expressionAttributeValues"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ReturnConsumedCapacity
	//
	// Experimental.
	ReturnConsumedCapacity DynamoConsumedCapacity `json:"returnConsumedCapacity"`
	// Determines whether item collection metrics are returned.
	//
	// If set to SIZE, the response includes statistics about item collections, if any,
	// that were modified during the operation are returned in the response.
	// If set to NONE (the default), no statistics are returned.
	// Experimental.
	ReturnItemCollectionMetrics DynamoItemCollectionMetrics `json:"returnItemCollectionMetrics"`
	// Use ReturnValues if you want to get the item attributes as they appeared before they were deleted.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_DeleteItem.html#DDB-DeleteItem-request-ReturnValues
	//
	// Experimental.
	ReturnValues DynamoReturnValues `json:"returnValues"`
}

// A StepFunctions task to call DynamoGetItem.
// Experimental.
type DynamoGetItem interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for DynamoGetItem
type jsiiProxy_DynamoGetItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
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

func (j *jsiiProxy_DynamoGetItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_DynamoGetItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
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


// Experimental.
func NewDynamoGetItem(scope constructs.Construct, id *string, props *DynamoGetItemProps) DynamoGetItem {
	_init_.Initialize()

	j := jsiiProxy_DynamoGetItem{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoGetItem_Override(d DynamoGetItem, scope constructs.Construct, id *string, props *DynamoGetItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoGetItem) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoGetItem) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func DynamoGetItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func DynamoGetItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
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
func DynamoGetItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoGetItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func DynamoGetItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.DynamoGetItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
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

// Render the choices in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
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

// Render map iterator in ASL JSON format.
// Experimental.
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

// Render the default next state in ASL JSON format.
// Experimental.
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

// Render ResultSelector in ASL JSON format.
// Experimental.
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

// Render error recovery options in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (d *jsiiProxy_DynamoGetItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for DynamoGetItem Task.
// Experimental.
type DynamoGetItemProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Primary key of the item to retrieve.
	//
	// For the primary key, you must provide all of the attributes.
	// For example, with a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key and the sort key.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-Key
	//
	// Experimental.
	Key *map[string]DynamoAttributeValue `json:"key"`
	// The name of the table containing the requested item.
	// Experimental.
	Table awsdynamodb.ITable `json:"table"`
	// Determines the read consistency model: If set to true, then the operation uses strongly consistent reads;
	//
	// otherwise, the operation uses eventually consistent reads.
	// Experimental.
	ConsistentRead *bool `json:"consistentRead"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ExpressionAttributeNames
	//
	// Experimental.
	ExpressionAttributeNames *map[string]*string `json:"expressionAttributeNames"`
	// An array of DynamoProjectionExpression that identifies one or more attributes to retrieve from the table.
	//
	// These attributes can include scalars, sets, or elements of a JSON document.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ProjectionExpression
	//
	// Experimental.
	ProjectionExpression *[]DynamoProjectionExpression `json:"projectionExpression"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-ReturnConsumedCapacity
	//
	// Experimental.
	ReturnConsumedCapacity DynamoConsumedCapacity `json:"returnConsumedCapacity"`
}

// Determines whether item collection metrics are returned.
// Experimental.
type DynamoItemCollectionMetrics string

const (
	DynamoItemCollectionMetrics_SIZE DynamoItemCollectionMetrics = "SIZE"
	DynamoItemCollectionMetrics_NONE DynamoItemCollectionMetrics = "NONE"
)

// Class to generate projection expression.
// Experimental.
type DynamoProjectionExpression interface {
	AtIndex(index *float64) DynamoProjectionExpression
	ToString() *string
	WithAttribute(attr *string) DynamoProjectionExpression
}

// The jsii proxy struct for DynamoProjectionExpression
type jsiiProxy_DynamoProjectionExpression struct {
	_ byte // padding
}

// Experimental.
func NewDynamoProjectionExpression() DynamoProjectionExpression {
	_init_.Initialize()

	j := jsiiProxy_DynamoProjectionExpression{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoProjectionExpression",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoProjectionExpression_Override(d DynamoProjectionExpression) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoProjectionExpression",
		nil, // no parameters
		d,
	)
}

// Adds the array literal access for passed index.
// Experimental.
func (d *jsiiProxy_DynamoProjectionExpression) AtIndex(index *float64) DynamoProjectionExpression {
	var returns DynamoProjectionExpression

	_jsii_.Invoke(
		d,
		"atIndex",
		[]interface{}{index},
		&returns,
	)

	return returns
}

// converts and return the string expression.
// Experimental.
func (d *jsiiProxy_DynamoProjectionExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Adds the passed attribute to the chain.
// Experimental.
func (d *jsiiProxy_DynamoProjectionExpression) WithAttribute(attr *string) DynamoProjectionExpression {
	var returns DynamoProjectionExpression

	_jsii_.Invoke(
		d,
		"withAttribute",
		[]interface{}{attr},
		&returns,
	)

	return returns
}

// A StepFunctions task to call DynamoPutItem.
// Experimental.
type DynamoPutItem interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for DynamoPutItem
type jsiiProxy_DynamoPutItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_DynamoPutItem) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoPutItem) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoPutItem(scope constructs.Construct, id *string, props *DynamoPutItemProps) DynamoPutItem {
	_init_.Initialize()

	j := jsiiProxy_DynamoPutItem{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoPutItem_Override(d DynamoPutItem, scope constructs.Construct, id *string, props *DynamoPutItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoPutItem) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoPutItem) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func DynamoPutItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func DynamoPutItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
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
func DynamoPutItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoPutItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func DynamoPutItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.DynamoPutItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (d *jsiiProxy_DynamoPutItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for DynamoPutItem Task.
// Experimental.
type DynamoPutItemProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// A map of attribute name/value pairs, one for each attribute.
	//
	// Only the primary key attributes are required;
	// you can optionally provide other attribute name-value pairs for the item.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-Item
	//
	// Experimental.
	Item *map[string]DynamoAttributeValue `json:"item"`
	// The name of the table where the item should be written .
	// Experimental.
	Table awsdynamodb.ITable `json:"table"`
	// A condition that must be satisfied in order for a conditional PutItem operation to succeed.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ConditionExpression
	//
	// Experimental.
	ConditionExpression *string `json:"conditionExpression"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ExpressionAttributeNames
	//
	// Experimental.
	ExpressionAttributeNames *map[string]*string `json:"expressionAttributeNames"`
	// One or more values that can be substituted in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ExpressionAttributeValues
	//
	// Experimental.
	ExpressionAttributeValues *map[string]DynamoAttributeValue `json:"expressionAttributeValues"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ReturnConsumedCapacity
	//
	// Experimental.
	ReturnConsumedCapacity DynamoConsumedCapacity `json:"returnConsumedCapacity"`
	// The item collection metrics to returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LSI.html#LSI.ItemCollections
	//
	// Experimental.
	ReturnItemCollectionMetrics DynamoItemCollectionMetrics `json:"returnItemCollectionMetrics"`
	// Use ReturnValues if you want to get the item attributes as they appeared before they were updated with the PutItem request.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_PutItem.html#DDB-PutItem-request-ReturnValues
	//
	// Experimental.
	ReturnValues DynamoReturnValues `json:"returnValues"`
}

// Use ReturnValues if you want to get the item attributes as they appear before or after they are changed.
// Experimental.
type DynamoReturnValues string

const (
	DynamoReturnValues_NONE DynamoReturnValues = "NONE"
	DynamoReturnValues_ALL_OLD DynamoReturnValues = "ALL_OLD"
	DynamoReturnValues_UPDATED_OLD DynamoReturnValues = "UPDATED_OLD"
	DynamoReturnValues_ALL_NEW DynamoReturnValues = "ALL_NEW"
	DynamoReturnValues_UPDATED_NEW DynamoReturnValues = "UPDATED_NEW"
)

// A StepFunctions task to call DynamoUpdateItem.
// Experimental.
type DynamoUpdateItem interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for DynamoUpdateItem
type jsiiProxy_DynamoUpdateItem struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_DynamoUpdateItem) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoUpdateItem) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoUpdateItem(scope constructs.Construct, id *string, props *DynamoUpdateItemProps) DynamoUpdateItem {
	_init_.Initialize()

	j := jsiiProxy_DynamoUpdateItem{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoUpdateItem_Override(d DynamoUpdateItem, scope constructs.Construct, id *string, props *DynamoUpdateItemProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DynamoUpdateItem) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DynamoUpdateItem) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func DynamoUpdateItem_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func DynamoUpdateItem_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
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
func DynamoUpdateItem_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func DynamoUpdateItem_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func DynamoUpdateItem_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.DynamoUpdateItem",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		d,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) OnPrepare() {
	_jsii_.InvokeVoid(
		d,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) Prepare() {
	_jsii_.InvokeVoid(
		d,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		d,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (d *jsiiProxy_DynamoUpdateItem) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for DynamoUpdateItem Task.
// Experimental.
type DynamoUpdateItemProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Primary key of the item to retrieve.
	//
	// For the primary key, you must provide all of the attributes.
	// For example, with a simple primary key, you only need to provide a value for the partition key.
	// For a composite primary key, you must provide values for both the partition key and the sort key.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_GetItem.html#DDB-GetItem-request-Key
	//
	// Experimental.
	Key *map[string]DynamoAttributeValue `json:"key"`
	// The name of the table containing the requested item.
	// Experimental.
	Table awsdynamodb.ITable `json:"table"`
	// A condition that must be satisfied in order for a conditional DeleteItem to succeed.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-ConditionExpression
	//
	// Experimental.
	ConditionExpression *string `json:"conditionExpression"`
	// One or more substitution tokens for attribute names in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-ExpressionAttributeNames
	//
	// Experimental.
	ExpressionAttributeNames *map[string]*string `json:"expressionAttributeNames"`
	// One or more values that can be substituted in an expression.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-ExpressionAttributeValues
	//
	// Experimental.
	ExpressionAttributeValues *map[string]DynamoAttributeValue `json:"expressionAttributeValues"`
	// Determines the level of detail about provisioned throughput consumption that is returned in the response.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-ReturnConsumedCapacity
	//
	// Experimental.
	ReturnConsumedCapacity DynamoConsumedCapacity `json:"returnConsumedCapacity"`
	// Determines whether item collection metrics are returned.
	//
	// If set to SIZE, the response includes statistics about item collections, if any,
	// that were modified during the operation are returned in the response.
	// If set to NONE (the default), no statistics are returned.
	// Experimental.
	ReturnItemCollectionMetrics DynamoItemCollectionMetrics `json:"returnItemCollectionMetrics"`
	// Use ReturnValues if you want to get the item attributes as they appeared before they were deleted.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-ReturnValues
	//
	// Experimental.
	ReturnValues DynamoReturnValues `json:"returnValues"`
	// An expression that defines one or more attributes to be updated, the action to be performed on them, and new values for them.
	// See: https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_UpdateItem.html#DDB-UpdateItem-request-UpdateExpression
	//
	// Experimental.
	UpdateExpression *string `json:"updateExpression"`
}

// Configuration for running an ECS task on EC2.
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-ec2
//
// Experimental.
type EcsEc2LaunchTarget interface {
	IEcsLaunchTarget
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsEc2LaunchTarget
type jsiiProxy_EcsEc2LaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

// Experimental.
func NewEcsEc2LaunchTarget(options *EcsEc2LaunchTargetOptions) EcsEc2LaunchTarget {
	_init_.Initialize()

	j := jsiiProxy_EcsEc2LaunchTarget{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsEc2LaunchTarget_Override(e EcsEc2LaunchTarget, options *EcsEc2LaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsEc2LaunchTarget",
		[]interface{}{options},
		e,
	)
}

// Called when the EC2 launch type is configured on RunTask.
// Experimental.
func (e *jsiiProxy_EcsEc2LaunchTarget) Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_task, launchTargetOptions},
		&returns,
	)

	return returns
}

// Options to run an ECS task on EC2 in StepFunctions and ECS.
// Experimental.
type EcsEc2LaunchTargetOptions struct {
	// Placement constraints.
	// Experimental.
	PlacementConstraints *[]awsecs.PlacementConstraint `json:"placementConstraints"`
	// Placement strategies.
	// Experimental.
	PlacementStrategies *[]awsecs.PlacementStrategy `json:"placementStrategies"`
}

// Configuration for running an ECS task on Fargate.
// See: https://docs.aws.amazon.com/AmazonECS/latest/userguide/launch_types.html#launch-type-fargate
//
// Experimental.
type EcsFargateLaunchTarget interface {
	IEcsLaunchTarget
	Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy struct for EcsFargateLaunchTarget
type jsiiProxy_EcsFargateLaunchTarget struct {
	jsiiProxy_IEcsLaunchTarget
}

// Experimental.
func NewEcsFargateLaunchTarget(options *EcsFargateLaunchTargetOptions) EcsFargateLaunchTarget {
	_init_.Initialize()

	j := jsiiProxy_EcsFargateLaunchTarget{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsFargateLaunchTarget_Override(e EcsFargateLaunchTarget, options *EcsFargateLaunchTargetOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsFargateLaunchTarget",
		[]interface{}{options},
		e,
	)
}

// Called when the Fargate launch type configured on RunTask.
// Experimental.
func (e *jsiiProxy_EcsFargateLaunchTarget) Bind(_task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_task, launchTargetOptions},
		&returns,
	)

	return returns
}

// Properties to define an ECS service.
// Experimental.
type EcsFargateLaunchTargetOptions struct {
	// Refers to a specific runtime environment for Fargate task infrastructure.
	//
	// Fargate platform version is a combination of the kernel and container runtime versions.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/platform_versions.html
	//
	// Experimental.
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
}

// Configuration options for the ECS launch type.
// Experimental.
type EcsLaunchTargetConfig struct {
	// Additional parameters to pass to the base task.
	// Experimental.
	Parameters *map[string]interface{} `json:"parameters"`
}

// Run a Task on ECS or Fargate.
// Experimental.
type EcsRunTask interface {
	awsstepfunctions.TaskStateBase
	awsec2.IConnectable
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	Connections() awsec2.Connections
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EcsRunTask
type jsiiProxy_EcsRunTask struct {
	internal.Type__awsstepfunctionsTaskStateBase
	internal.Type__awsec2IConnectable
}

func (j *jsiiProxy_EcsRunTask) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcsRunTask) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEcsRunTask(scope constructs.Construct, id *string, props *EcsRunTaskProps) EcsRunTask {
	_init_.Initialize()

	j := jsiiProxy_EcsRunTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEcsRunTask_Override(e EcsRunTask, scope constructs.Construct, id *string, props *EcsRunTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EcsRunTask) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EcsRunTask) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EcsRunTask_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EcsRunTask_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
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
func EcsRunTask_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EcsRunTask_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EcsRunTask_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EcsRunTask",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EcsRunTask) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EcsRunTask) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EcsRunTask) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EcsRunTask) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EcsRunTask) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EcsRunTask) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EcsRunTask) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EcsRunTask) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EcsRunTask) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EcsRunTask) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EcsRunTask) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EcsRunTask) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EcsRunTask) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EcsRunTask) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// A StepFunctions Task to run a Task on ECS or Fargate.
// Deprecated: No replacement
type EcsRunTaskBase interface {
	awsec2.IConnectable
	awsstepfunctions.IStepFunctionsTask
	Connections() awsec2.Connections
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for EcsRunTaskBase
type jsiiProxy_EcsRunTaskBase struct {
	internal.Type__awsec2IConnectable
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

func (j *jsiiProxy_EcsRunTaskBase) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: No replacement
func NewEcsRunTaskBase(props *EcsRunTaskBaseProps) EcsRunTaskBase {
	_init_.Initialize()

	j := jsiiProxy_EcsRunTaskBase{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTaskBase",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: No replacement
func NewEcsRunTaskBase_Override(e EcsRunTaskBase, props *EcsRunTaskBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EcsRunTaskBase",
		[]interface{}{props},
		e,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: No replacement
func (e *jsiiProxy_EcsRunTaskBase) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Deprecated: No replacement
func (e *jsiiProxy_EcsRunTaskBase) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

// Construction properties for the BaseRunTaskProps.
// Deprecated: No replacement
type EcsRunTaskBaseProps struct {
	// The topic to run the task on.
	// Deprecated: No replacement
	Cluster awsecs.ICluster `json:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// Deprecated: No replacement
	TaskDefinition awsecs.TaskDefinition `json:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Deprecated: No replacement
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Deprecated: No replacement
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// Additional parameters to pass to the base task.
	// Deprecated: No replacement
	Parameters *map[string]interface{} `json:"parameters"`
}

// Properties for ECS Tasks.
// Experimental.
type EcsRunTaskProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ECS cluster to run the task on.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
	// An Amazon ECS launch type determines the type of infrastructure on which your tasks and services are hosted.
	// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html
	//
	// Experimental.
	LaunchTarget IEcsLaunchTarget `json:"launchTarget"`
	// [disable-awslint:ref-via-interface] Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// Experimental.
	TaskDefinition awsecs.TaskDefinition `json:"taskDefinition"`
	// Assign public IP addresses to each task.
	// Experimental.
	AssignPublicIp *bool `json:"assignPublicIp"`
	// Container setting overrides.
	//
	// Specify the container to use and the overrides to apply.
	// Experimental.
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides"`
	// Existing security groups to use for the tasks.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `json:"securityGroups"`
	// Subnets to place the task's ENIs.
	// Experimental.
	Subnets *awsec2.SubnetSelection `json:"subnets"`
}

// Call a EKS endpoint as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-eks.html
//
// Experimental.
type EksCall interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EksCall
type jsiiProxy_EksCall struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EksCall) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksCall) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEksCall(scope constructs.Construct, id *string, props *EksCallProps) EksCall {
	_init_.Initialize()

	j := jsiiProxy_EksCall{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEksCall_Override(e EksCall, scope constructs.Construct, id *string, props *EksCallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EksCall) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EksCall) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EksCall_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EksCall_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksCall",
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
func EksCall_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EksCall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EksCall_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EksCall",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EksCall) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EksCall) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EksCall) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EksCall) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EksCall) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EksCall) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EksCall) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EksCall) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EksCall) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EksCall) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EksCall) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EksCall) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EksCall) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EksCall) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EksCall) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EksCall) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EksCall) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EksCall) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EksCall) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EksCall) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EksCall) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EksCall) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EksCall) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EksCall) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EksCall) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EksCall) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EksCall) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EksCall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EksCall) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EksCall) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for calling a EKS endpoint with EksCall.
// Experimental.
type EksCallProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The EKS cluster.
	// Experimental.
	Cluster awseks.ICluster `json:"cluster"`
	// HTTP method ("GET", "POST", "PUT", ...) part of HTTP request.
	// Experimental.
	HttpMethod HttpMethods `json:"httpMethod"`
	// HTTP path of the Kubernetes REST API operation For example: /api/v1/namespaces/default/pods.
	// Experimental.
	HttpPath *string `json:"httpPath"`
	// Query Parameters part of HTTP request.
	// Experimental.
	QueryParameters *map[string]*[]*string `json:"queryParameters"`
	// Request body part of HTTP request.
	// Experimental.
	RequestBody awsstepfunctions.TaskInput `json:"requestBody"`
}

// A Step Functions Task to add a Step to an EMR Cluster.
//
// The StepConfiguration is defined as Parameters in the state machine definition.
//
// OUTPUT: the StepId
// Experimental.
type EmrAddStep interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrAddStep
type jsiiProxy_EmrAddStep struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrAddStep) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrAddStep) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrAddStep(scope constructs.Construct, id *string, props *EmrAddStepProps) EmrAddStep {
	_init_.Initialize()

	j := jsiiProxy_EmrAddStep{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrAddStep_Override(e EmrAddStep, scope constructs.Construct, id *string, props *EmrAddStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrAddStep) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrAddStep) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrAddStep_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrAddStep_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
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
func EmrAddStep_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrAddStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrAddStep_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrAddStep",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrAddStep) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrAddStep) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrAddStep) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrAddStep) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrAddStep) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrAddStep) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrAddStep) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrAddStep) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrAddStep) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrAddStep) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrAddStep) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrAddStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrAddStep) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrAddStep) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EmrAddStep.
// Experimental.
type EmrAddStepProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to add the Step to.
	// Experimental.
	ClusterId *string `json:"clusterId"`
	// A path to a JAR file run during the step.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Experimental.
	Jar *string `json:"jar"`
	// The name of the Step.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
	//
	// Experimental.
	Name *string `json:"name"`
	// The action to take when the cluster step fails.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_StepConfig.html
	//
	// Experimental.
	ActionOnFailure ActionOnFailure `json:"actionOnFailure"`
	// A list of command line arguments passed to the JAR file's main function when executed.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Experimental.
	Args *[]*string `json:"args"`
	// The name of the main class in the specified Java file.
	//
	// If not specified, the JAR file should specify a Main-Class in its manifest file.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Experimental.
	MainClass *string `json:"mainClass"`
	// A list of Java properties that are set when the step runs.
	//
	// You can use these properties to pass key value pairs to your main function.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_HadoopJarStepConfig.html
	//
	// Experimental.
	Properties *map[string]*string `json:"properties"`
}

// A Step Functions Task to to cancel a Step on an EMR Cluster.
// Experimental.
type EmrCancelStep interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrCancelStep
type jsiiProxy_EmrCancelStep struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrCancelStep) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCancelStep) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrCancelStep(scope constructs.Construct, id *string, props *EmrCancelStepProps) EmrCancelStep {
	_init_.Initialize()

	j := jsiiProxy_EmrCancelStep{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrCancelStep_Override(e EmrCancelStep, scope constructs.Construct, id *string, props *EmrCancelStepProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrCancelStep) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrCancelStep) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrCancelStep_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrCancelStep_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
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
func EmrCancelStep_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrCancelStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrCancelStep_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrCancelStep",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrCancelStep) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EmrCancelStep.
// Experimental.
type EmrCancelStepProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `json:"clusterId"`
	// The StepId to cancel.
	// Experimental.
	StepId *string `json:"stepId"`
}

// A Step Functions Task to create an EMR Cluster.
//
// The ClusterConfiguration is defined as Parameters in the state machine definition.
//
// OUTPUT: the ClusterId.
// Experimental.
type EmrCreateCluster interface {
	awsstepfunctions.TaskStateBase
	AutoScalingRole() awsiam.IRole
	Branches() *[]awsstepfunctions.StateGraph
	ClusterRole() awsiam.IRole
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	ServiceRole() awsiam.IRole
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrCreateCluster
type jsiiProxy_EmrCreateCluster struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrCreateCluster) AutoScalingRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"autoScalingRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) ClusterRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"clusterRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) ServiceRole() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"serviceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrCreateCluster) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrCreateCluster(scope constructs.Construct, id *string, props *EmrCreateClusterProps) EmrCreateCluster {
	_init_.Initialize()

	j := jsiiProxy_EmrCreateCluster{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrCreateCluster_Override(e EmrCreateCluster, scope constructs.Construct, id *string, props *EmrCreateClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrCreateCluster) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrCreateCluster) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrCreateCluster_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrCreateCluster_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
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
func EmrCreateCluster_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrCreateCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrCreateCluster_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrCreateCluster",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrCreateCluster) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for the EMR Cluster Applications.
//
// Applies to Amazon EMR releases 4.0 and later. A case-insensitive list of applications for Amazon EMR to install and configure when launching
// the cluster.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_Application.html
//
// Experimental.
type EmrCreateCluster_ApplicationConfigProperty struct {
	// The name of the application.
	// Experimental.
	Name *string `json:"name"`
	// This option is for advanced users only.
	//
	// This is meta information about third-party applications that third-party vendors use
	// for testing purposes.
	// Experimental.
	AdditionalInfo *map[string]*string `json:"additionalInfo"`
	// Arguments for Amazon EMR to pass to the application.
	// Experimental.
	Args *[]*string `json:"args"`
	// The version of the application.
	// Experimental.
	Version *string `json:"version"`
}

// An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_AutoScalingPolicy.html
//
// Experimental.
type EmrCreateCluster_AutoScalingPolicyProperty struct {
	// The upper and lower EC2 instance limits for an automatic scaling policy.
	//
	// Automatic scaling activity will not cause an instance
	// group to grow above or below these limits.
	// Experimental.
	Constraints *EmrCreateCluster_ScalingConstraintsProperty `json:"constraints"`
	// The scale-in and scale-out rules that comprise the automatic scaling policy.
	// Experimental.
	Rules *[]*EmrCreateCluster_ScalingRuleProperty `json:"rules"`
}

// Auto-termination policy for the EMR cluster.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_AutoTerminationPolicy.html
//
// Experimental.
type EmrCreateCluster_AutoTerminationPolicyProperty struct {
	// Specifies the amount of idle time after which the cluster automatically terminates.
	// Experimental.
	IdleTimeout awscdk.Duration `json:"idleTimeout"`
}

// Configuration of a bootstrap action.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_BootstrapActionConfig.html
//
// Experimental.
type EmrCreateCluster_BootstrapActionConfigProperty struct {
	// The name of the bootstrap action.
	// Experimental.
	Name *string `json:"name"`
	// The script run by the bootstrap action.
	// Experimental.
	ScriptBootstrapAction *EmrCreateCluster_ScriptBootstrapActionConfigProperty `json:"scriptBootstrapAction"`
}

// CloudWatch Alarm Comparison Operators.
// Experimental.
type EmrCreateCluster_CloudWatchAlarmComparisonOperator string

const (
	EmrCreateCluster_CloudWatchAlarmComparisonOperator_GREATER_THAN_OR_EQUAL EmrCreateCluster_CloudWatchAlarmComparisonOperator = "GREATER_THAN_OR_EQUAL"
	EmrCreateCluster_CloudWatchAlarmComparisonOperator_GREATER_THAN EmrCreateCluster_CloudWatchAlarmComparisonOperator = "GREATER_THAN"
	EmrCreateCluster_CloudWatchAlarmComparisonOperator_LESS_THAN EmrCreateCluster_CloudWatchAlarmComparisonOperator = "LESS_THAN"
	EmrCreateCluster_CloudWatchAlarmComparisonOperator_LESS_THAN_OR_EQUAL EmrCreateCluster_CloudWatchAlarmComparisonOperator = "LESS_THAN_OR_EQUAL"
)

// The definition of a CloudWatch metric alarm, which determines when an automatic scaling activity is triggered.
//
// When the defined alarm conditions
// are satisfied, scaling activity begins.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_CloudWatchAlarmDefinition.html
//
// Experimental.
type EmrCreateCluster_CloudWatchAlarmDefinitionProperty struct {
	// Determines how the metric specified by MetricName is compared to the value specified by Threshold.
	// Experimental.
	ComparisonOperator EmrCreateCluster_CloudWatchAlarmComparisonOperator `json:"comparisonOperator"`
	// The name of the CloudWatch metric that is watched to determine an alarm condition.
	// Experimental.
	MetricName *string `json:"metricName"`
	// The period, in seconds, over which the statistic is applied.
	//
	// EMR CloudWatch metrics are emitted every five minutes (300 seconds), so if
	// an EMR CloudWatch metric is specified, specify 300.
	// Experimental.
	Period awscdk.Duration `json:"period"`
	// A CloudWatch metric dimension.
	// Experimental.
	Dimensions *[]*EmrCreateCluster_MetricDimensionProperty `json:"dimensions"`
	// The number of periods, in five-minute increments, during which the alarm condition must exist before the alarm triggers automatic scaling activity.
	// Experimental.
	EvaluationPeriods *float64 `json:"evaluationPeriods"`
	// The namespace for the CloudWatch metric.
	// Experimental.
	Namespace *string `json:"namespace"`
	// The statistic to apply to the metric associated with the alarm.
	// Experimental.
	Statistic EmrCreateCluster_CloudWatchAlarmStatistic `json:"statistic"`
	// The value against which the specified statistic is compared.
	// Experimental.
	Threshold *float64 `json:"threshold"`
	// The unit of measure associated with the CloudWatch metric being watched.
	//
	// The value specified for Unit must correspond to the units
	// specified in the CloudWatch metric.
	// Experimental.
	Unit EmrCreateCluster_CloudWatchAlarmUnit `json:"unit"`
}

// CloudWatch Alarm Statistics.
// Experimental.
type EmrCreateCluster_CloudWatchAlarmStatistic string

const (
	EmrCreateCluster_CloudWatchAlarmStatistic_SAMPLE_COUNT EmrCreateCluster_CloudWatchAlarmStatistic = "SAMPLE_COUNT"
	EmrCreateCluster_CloudWatchAlarmStatistic_AVERAGE EmrCreateCluster_CloudWatchAlarmStatistic = "AVERAGE"
	EmrCreateCluster_CloudWatchAlarmStatistic_SUM EmrCreateCluster_CloudWatchAlarmStatistic = "SUM"
	EmrCreateCluster_CloudWatchAlarmStatistic_MINIMUM EmrCreateCluster_CloudWatchAlarmStatistic = "MINIMUM"
	EmrCreateCluster_CloudWatchAlarmStatistic_MAXIMUM EmrCreateCluster_CloudWatchAlarmStatistic = "MAXIMUM"
)

// CloudWatch Alarm Units.
// Experimental.
type EmrCreateCluster_CloudWatchAlarmUnit string

const (
	EmrCreateCluster_CloudWatchAlarmUnit_NONE EmrCreateCluster_CloudWatchAlarmUnit = "NONE"
	EmrCreateCluster_CloudWatchAlarmUnit_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "SECONDS"
	EmrCreateCluster_CloudWatchAlarmUnit_MICRO_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "MICRO_SECONDS"
	EmrCreateCluster_CloudWatchAlarmUnit_MILLI_SECONDS EmrCreateCluster_CloudWatchAlarmUnit = "MILLI_SECONDS"
	EmrCreateCluster_CloudWatchAlarmUnit_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "BYTES"
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BYTES"
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BYTES"
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BYTES"
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BYTES"
	EmrCreateCluster_CloudWatchAlarmUnit_BITS EmrCreateCluster_CloudWatchAlarmUnit = "BITS"
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BITS"
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BITS"
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BITS"
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BITS"
	EmrCreateCluster_CloudWatchAlarmUnit_PERCENT EmrCreateCluster_CloudWatchAlarmUnit = "PERCENT"
	EmrCreateCluster_CloudWatchAlarmUnit_COUNT EmrCreateCluster_CloudWatchAlarmUnit = "COUNT"
	EmrCreateCluster_CloudWatchAlarmUnit_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "BYTES_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BYTES_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BYTES_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BYTES_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BYTES_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BYTES_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "BITS_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_KILO_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "KILO_BITS_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_MEGA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "MEGA_BITS_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_GIGA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "GIGA_BITS_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_TERA_BITS_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "TERA_BITS_PER_SECOND"
	EmrCreateCluster_CloudWatchAlarmUnit_COUNT_PER_SECOND EmrCreateCluster_CloudWatchAlarmUnit = "COUNT_PER_SECOND"
)

// An optional configuration specification to be used when provisioning cluster instances, which can include configurations for applications and software bundled with Amazon EMR.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_Configuration.html
//
// Experimental.
type EmrCreateCluster_ConfigurationProperty struct {
	// The classification within a configuration.
	// Experimental.
	Classification *string `json:"classification"`
	// A list of additional configurations to apply within a configuration object.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `json:"configurations"`
	// A set of properties specified within a configuration classification.
	// Experimental.
	Properties *map[string]*string `json:"properties"`
}

// Configuration of requested EBS block device associated with the instance group with count of volumes that will be associated to every instance.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_EbsBlockDeviceConfig.html
//
// Experimental.
type EmrCreateCluster_EbsBlockDeviceConfigProperty struct {
	// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
	// Experimental.
	VolumeSpecification *EmrCreateCluster_VolumeSpecificationProperty `json:"volumeSpecification"`
	// Number of EBS volumes with a specific volume configuration that will be associated with every instance in the instance group.
	// Experimental.
	VolumesPerInstance *float64 `json:"volumesPerInstance"`
}

// EBS Volume Types.
// Experimental.
type EmrCreateCluster_EbsBlockDeviceVolumeType string

const (
	EmrCreateCluster_EbsBlockDeviceVolumeType_GP2 EmrCreateCluster_EbsBlockDeviceVolumeType = "GP2"
	EmrCreateCluster_EbsBlockDeviceVolumeType_IO1 EmrCreateCluster_EbsBlockDeviceVolumeType = "IO1"
	EmrCreateCluster_EbsBlockDeviceVolumeType_STANDARD EmrCreateCluster_EbsBlockDeviceVolumeType = "STANDARD"
)

// The Amazon EBS configuration of a cluster instance.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_EbsConfiguration.html
//
// Experimental.
type EmrCreateCluster_EbsConfigurationProperty struct {
	// An array of Amazon EBS volume specifications attached to a cluster instance.
	// Experimental.
	EbsBlockDeviceConfigs *[]*EmrCreateCluster_EbsBlockDeviceConfigProperty `json:"ebsBlockDeviceConfigs"`
	// Indicates whether an Amazon EBS volume is EBS-optimized.
	// Experimental.
	EbsOptimized *bool `json:"ebsOptimized"`
}

// The Cluster ScaleDownBehavior specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html#EMR-RunJobFlow-request-ScaleDownBehavior
//
// Experimental.
type EmrCreateCluster_EmrClusterScaleDownBehavior string

const (
	EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_INSTANCE_HOUR EmrCreateCluster_EmrClusterScaleDownBehavior = "TERMINATE_AT_INSTANCE_HOUR"
	EmrCreateCluster_EmrClusterScaleDownBehavior_TERMINATE_AT_TASK_COMPLETION EmrCreateCluster_EmrClusterScaleDownBehavior = "TERMINATE_AT_TASK_COMPLETION"
)

// The configuration that defines an instance fleet.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetConfig.html
//
// Experimental.
type EmrCreateCluster_InstanceFleetConfigProperty struct {
	// The node type that the instance fleet hosts.
	//
	// Valid values are MASTER,CORE,and TASK.
	// Experimental.
	InstanceFleetType EmrCreateCluster_InstanceRoleType `json:"instanceFleetType"`
	// The instance type configurations that define the EC2 instances in the instance fleet.
	// Experimental.
	InstanceTypeConfigs *[]*EmrCreateCluster_InstanceTypeConfigProperty `json:"instanceTypeConfigs"`
	// The launch specification for the instance fleet.
	// Experimental.
	LaunchSpecifications *EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty `json:"launchSpecifications"`
	// The friendly name of the instance fleet.
	// Experimental.
	Name *string `json:"name"`
	// The target capacity of On-Demand units for the instance fleet, which determines how many On-Demand instances to provision.
	// Experimental.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet, which determines how many Spot instances to provision.
	// Experimental.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetProvisioningSpecifications.html
//
// Experimental.
type EmrCreateCluster_InstanceFleetProvisioningSpecificationsProperty struct {
	// The launch specification for Spot instances in the fleet, which determines the defined duration and provisioning timeout behavior.
	// Experimental.
	SpotSpecification *EmrCreateCluster_SpotProvisioningSpecificationProperty `json:"spotSpecification"`
}

// Configuration defining a new instance group.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroupConfig.html
//
// Experimental.
type EmrCreateCluster_InstanceGroupConfigProperty struct {
	// Target number of instances for the instance group.
	// Experimental.
	InstanceCount *float64 `json:"instanceCount"`
	// The role of the instance group in the cluster.
	// Experimental.
	InstanceRole EmrCreateCluster_InstanceRoleType `json:"instanceRole"`
	// The EC2 instance type for all instances in the instance group.
	// Experimental.
	InstanceType *string `json:"instanceType"`
	// An automatic scaling policy for a core instance group or task instance group in an Amazon EMR cluster.
	// Experimental.
	AutoScalingPolicy *EmrCreateCluster_AutoScalingPolicyProperty `json:"autoScalingPolicy"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	// Experimental.
	BidPrice *string `json:"bidPrice"`
	// The list of configurations supplied for an EMR cluster instance group.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `json:"configurations"`
	// EBS configurations that will be attached to each EC2 instance in the instance group.
	// Experimental.
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `json:"ebsConfiguration"`
	// Market type of the EC2 instances used to create a cluster node.
	// Experimental.
	Market EmrCreateCluster_InstanceMarket `json:"market"`
	// Friendly name given to the instance group.
	// Experimental.
	Name *string `json:"name"`
}

// EC2 Instance Market.
// Experimental.
type EmrCreateCluster_InstanceMarket string

const (
	EmrCreateCluster_InstanceMarket_ON_DEMAND EmrCreateCluster_InstanceMarket = "ON_DEMAND"
	EmrCreateCluster_InstanceMarket_SPOT EmrCreateCluster_InstanceMarket = "SPOT"
)

// Instance Role Types.
// Experimental.
type EmrCreateCluster_InstanceRoleType string

const (
	EmrCreateCluster_InstanceRoleType_MASTER EmrCreateCluster_InstanceRoleType = "MASTER"
	EmrCreateCluster_InstanceRoleType_CORE EmrCreateCluster_InstanceRoleType = "CORE"
	EmrCreateCluster_InstanceRoleType_TASK EmrCreateCluster_InstanceRoleType = "TASK"
)

// An instance type configuration for each instance type in an instance fleet, which determines the EC2 instances Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceTypeConfig.html
//
// Experimental.
type EmrCreateCluster_InstanceTypeConfigProperty struct {
	// An EC2 instance type.
	// Experimental.
	InstanceType *string `json:"instanceType"`
	// The bid price for each EC2 Spot instance type as defined by InstanceType.
	//
	// Expressed in USD.
	// Experimental.
	BidPrice *string `json:"bidPrice"`
	// The bid price, as a percentage of On-Demand price.
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice"`
	// A configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `json:"configurations"`
	// The configuration of Amazon Elastic Block Storage (EBS) attached to each instance as defined by InstanceType.
	// Experimental.
	EbsConfiguration *EmrCreateCluster_EbsConfigurationProperty `json:"ebsConfiguration"`
	// The number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in the InstanceFleetConfig.
	// Experimental.
	WeightedCapacity *float64 `json:"weightedCapacity"`
}

// A specification of the number and type of Amazon EC2 instances.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_JobFlowInstancesConfig.html
//
// Experimental.
type EmrCreateCluster_InstancesConfigProperty struct {
	// A list of additional Amazon EC2 security group IDs for the master node.
	// Experimental.
	AdditionalMasterSecurityGroups *[]*string `json:"additionalMasterSecurityGroups"`
	// A list of additional Amazon EC2 security group IDs for the core and task nodes.
	// Experimental.
	AdditionalSlaveSecurityGroups *[]*string `json:"additionalSlaveSecurityGroups"`
	// The name of the EC2 key pair that can be used to ssh to the master node as the user called "hadoop.".
	// Experimental.
	Ec2KeyName *string `json:"ec2KeyName"`
	// Applies to clusters that use the uniform instance group configuration.
	//
	// To launch the cluster in Amazon Virtual Private Cloud (Amazon VPC),
	// set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch.
	// Experimental.
	Ec2SubnetId *string `json:"ec2SubnetId"`
	// Applies to clusters that use the instance fleet configuration.
	//
	// When multiple EC2 subnet IDs are specified, Amazon EMR evaluates them and
	// launches instances in the optimal subnet.
	// Experimental.
	Ec2SubnetIds *[]*string `json:"ec2SubnetIds"`
	// The identifier of the Amazon EC2 security group for the master node.
	// Experimental.
	EmrManagedMasterSecurityGroup *string `json:"emrManagedMasterSecurityGroup"`
	// The identifier of the Amazon EC2 security group for the core and task nodes.
	// Experimental.
	EmrManagedSlaveSecurityGroup *string `json:"emrManagedSlaveSecurityGroup"`
	// Applies only to Amazon EMR release versions earlier than 4.0. The Hadoop version for the cluster.
	// Experimental.
	HadoopVersion *string `json:"hadoopVersion"`
	// The number of EC2 instances in the cluster.
	// Experimental.
	InstanceCount *float64 `json:"instanceCount"`
	// Describes the EC2 instances and instance configurations for clusters that use the instance fleet configuration.
	//
	// The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.
	// Experimental.
	InstanceFleets *[]*EmrCreateCluster_InstanceFleetConfigProperty `json:"instanceFleets"`
	// Configuration for the instance groups in a cluster.
	// Experimental.
	InstanceGroups *[]*EmrCreateCluster_InstanceGroupConfigProperty `json:"instanceGroups"`
	// The EC2 instance type of the master node.
	// Experimental.
	MasterInstanceType *string `json:"masterInstanceType"`
	// The Availability Zone in which the cluster runs.
	// Experimental.
	Placement *EmrCreateCluster_PlacementTypeProperty `json:"placement"`
	// The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.
	// Experimental.
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup"`
	// The EC2 instance type of the core and task nodes.
	// Experimental.
	SlaveInstanceType *string `json:"slaveInstanceType"`
	// Specifies whether to lock the cluster to prevent the Amazon EC2 instances from being terminated by API call, user intervention, or in the event of a job-flow error.
	// Experimental.
	TerminationProtected *bool `json:"terminationProtected"`
}

// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_KerberosAttributes.html
//
// Experimental.
type EmrCreateCluster_KerberosAttributesProperty struct {
	// The name of the Kerberos realm to which all nodes in a cluster belong.
	//
	// For example, EC2.INTERNAL.
	// Experimental.
	Realm *string `json:"realm"`
	// The Active Directory password for ADDomainJoinUser.
	// Experimental.
	AdDomainJoinPassword *string `json:"adDomainJoinPassword"`
	// Required only when establishing a cross-realm trust with an Active Directory domain.
	//
	// A user with sufficient privileges to join
	// resources to the domain.
	// Experimental.
	AdDomainJoinUser *string `json:"adDomainJoinUser"`
	// Required only when establishing a cross-realm trust with a KDC in a different realm.
	//
	// The cross-realm principal password, which
	// must be identical across realms.
	// Experimental.
	CrossRealmTrustPrincipalPassword *string `json:"crossRealmTrustPrincipalPassword"`
	// The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster.
	// Experimental.
	KdcAdminPassword *string `json:"kdcAdminPassword"`
}

// A CloudWatch dimension, which is specified using a Key (known as a Name in CloudWatch), Value pair.
//
// By default, Amazon EMR uses
// one dimension whose Key is JobFlowID and Value is a variable representing the cluster ID, which is ${emr.clusterId}. This enables
// the rule to bootstrap when the cluster ID becomes available
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_MetricDimension.html
//
// Experimental.
type EmrCreateCluster_MetricDimensionProperty struct {
	// The dimension name.
	// Experimental.
	Key *string `json:"key"`
	// The dimension value.
	// Experimental.
	Value *string `json:"value"`
}

// The Amazon EC2 Availability Zone configuration of the cluster (job flow).
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_PlacementType.html
//
// Experimental.
type EmrCreateCluster_PlacementTypeProperty struct {
	// The Amazon EC2 Availability Zone for the cluster.
	//
	// AvailabilityZone is used for uniform instance groups, while AvailabilityZones
	// (plural) is used for instance fleets.
	// Experimental.
	AvailabilityZone *string `json:"availabilityZone"`
	// When multiple Availability Zones are specified, Amazon EMR evaluates them and launches instances in the optimal Availability Zone.
	//
	// AvailabilityZones is used for instance fleets, while AvailabilityZone (singular) is used for uniform instance groups.
	// Experimental.
	AvailabilityZones *[]*string `json:"availabilityZones"`
}

// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
//
// And an automatic scaling configuration, which describes how the policy adds or removes instances, the cooldown period,
// and the number of EC2 instances that will be added each time the CloudWatch metric alarm condition is satisfied.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingAction.html
//
// Experimental.
type EmrCreateCluster_ScalingActionProperty struct {
	// The type of adjustment the automatic scaling activity makes when triggered, and the periodicity of the adjustment.
	// Experimental.
	SimpleScalingPolicyConfiguration *EmrCreateCluster_SimpleScalingPolicyConfigurationProperty `json:"simpleScalingPolicyConfiguration"`
	// Not available for instance groups.
	//
	// Instance groups use the market type specified for the group.
	// Experimental.
	Market EmrCreateCluster_InstanceMarket `json:"market"`
}

// AutoScaling Adjustment Type.
// Experimental.
type EmrCreateCluster_ScalingAdjustmentType string

const (
	EmrCreateCluster_ScalingAdjustmentType_CHANGE_IN_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "CHANGE_IN_CAPACITY"
	EmrCreateCluster_ScalingAdjustmentType_PERCENT_CHANGE_IN_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "PERCENT_CHANGE_IN_CAPACITY"
	EmrCreateCluster_ScalingAdjustmentType_EXACT_CAPACITY EmrCreateCluster_ScalingAdjustmentType = "EXACT_CAPACITY"
)

// The upper and lower EC2 instance limits for an automatic scaling policy.
//
// Automatic scaling activities triggered by automatic scaling
// rules will not cause an instance group to grow above or below these limits.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingConstraints.html
//
// Experimental.
type EmrCreateCluster_ScalingConstraintsProperty struct {
	// The upper boundary of EC2 instances in an instance group beyond which scaling activities are not allowed to grow.
	//
	// Scale-out
	// activities will not add instances beyond this boundary.
	// Experimental.
	MaxCapacity *float64 `json:"maxCapacity"`
	// The lower boundary of EC2 instances in an instance group below which scaling activities are not allowed to shrink.
	//
	// Scale-in
	// activities will not terminate instances below this boundary.
	// Experimental.
	MinCapacity *float64 `json:"minCapacity"`
}

// A scale-in or scale-out rule that defines scaling activity, including the CloudWatch metric alarm that triggers activity, how EC2 instances are added or removed, and the periodicity of adjustments.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingRule.html
//
// Experimental.
type EmrCreateCluster_ScalingRuleProperty struct {
	// The conditions that trigger an automatic scaling activity.
	// Experimental.
	Action *EmrCreateCluster_ScalingActionProperty `json:"action"`
	// The name used to identify an automatic scaling rule.
	//
	// Rule names must be unique within a scaling policy.
	// Experimental.
	Name *string `json:"name"`
	// The CloudWatch alarm definition that determines when automatic scaling activity is triggered.
	// Experimental.
	Trigger *EmrCreateCluster_ScalingTriggerProperty `json:"trigger"`
	// A friendly, more verbose description of the automatic scaling rule.
	// Experimental.
	Description *string `json:"description"`
}

// The conditions that trigger an automatic scaling activity and the definition of a CloudWatch metric alarm.
//
// When the defined alarm conditions are met along with other trigger parameters, scaling activity begins.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScalingTrigger.html
//
// Experimental.
type EmrCreateCluster_ScalingTriggerProperty struct {
	// The definition of a CloudWatch metric alarm.
	//
	// When the defined alarm conditions are met along with other trigger parameters,
	// scaling activity begins.
	// Experimental.
	CloudWatchAlarmDefinition *EmrCreateCluster_CloudWatchAlarmDefinitionProperty `json:"cloudWatchAlarmDefinition"`
}

// Configuration of the script to run during a bootstrap action.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ScriptBootstrapActionConfig.html
//
// Experimental.
type EmrCreateCluster_ScriptBootstrapActionConfigProperty struct {
	// Location of the script to run during a bootstrap action.
	//
	// Can be either a location in Amazon S3 or on a local file system.
	// Experimental.
	Path *string `json:"path"`
	// A list of command line arguments to pass to the bootstrap action script.
	// Experimental.
	Args *[]*string `json:"args"`
}

// An automatic scaling configuration, which describes how the policy adds or removes instances, the cooldown period, and the number of EC2 instances that will be added each time the CloudWatch metric alarm condition is satisfied.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SimpleScalingPolicyConfiguration.html
//
// Experimental.
type EmrCreateCluster_SimpleScalingPolicyConfigurationProperty struct {
	// The amount by which to scale in or scale out, based on the specified AdjustmentType.
	//
	// A positive value adds to the instance group's
	// EC2 instance count while a negative number removes instances. If AdjustmentType is set to EXACT_CAPACITY, the number should only be
	// a positive integer.
	// Experimental.
	ScalingAdjustment *float64 `json:"scalingAdjustment"`
	// The way in which EC2 instances are added (if ScalingAdjustment is a positive number) or terminated (if ScalingAdjustment is a negative number) each time the scaling activity is triggered.
	// Experimental.
	AdjustmentType EmrCreateCluster_ScalingAdjustmentType `json:"adjustmentType"`
	// The amount of time, in seconds, after a scaling activity completes before any further trigger-related scaling activities can start.
	// Experimental.
	CoolDown *float64 `json:"coolDown"`
}

// Spot Allocation Strategies.
//
// Specifies the strategy to use in launching Spot Instance fleets. For example, "capacity-optimized" launches instances from Spot Instance pools with optimal capacity for the number of instances that are launching.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
// Experimental.
type EmrCreateCluster_SpotAllocationStrategy string

const (
	EmrCreateCluster_SpotAllocationStrategy_CAPACITY_OPTIMIZED EmrCreateCluster_SpotAllocationStrategy = "CAPACITY_OPTIMIZED"
)

// The launch specification for Spot instances in the instance fleet, which determines the defined duration and provisioning timeout behavior.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_SpotProvisioningSpecification.html
//
// Experimental.
type EmrCreateCluster_SpotProvisioningSpecificationProperty struct {
	// The action to take when TargetSpotCapacity has not been fulfilled when the TimeoutDurationMinutes has expired.
	// Experimental.
	TimeoutAction EmrCreateCluster_SpotTimeoutAction `json:"timeoutAction"`
	// The spot provisioning timeout period in minutes.
	// Experimental.
	TimeoutDurationMinutes *float64 `json:"timeoutDurationMinutes"`
	// Specifies the strategy to use in launching Spot Instance fleets.
	// Experimental.
	AllocationStrategy EmrCreateCluster_SpotAllocationStrategy `json:"allocationStrategy"`
	// The defined duration for Spot instances (also known as Spot blocks) in minutes.
	// Experimental.
	BlockDurationMinutes *float64 `json:"blockDurationMinutes"`
}

// Spot Timeout Actions.
// Experimental.
type EmrCreateCluster_SpotTimeoutAction string

const (
	EmrCreateCluster_SpotTimeoutAction_SWITCH_TO_ON_DEMAND EmrCreateCluster_SpotTimeoutAction = "SWITCH_TO_ON_DEMAND"
	EmrCreateCluster_SpotTimeoutAction_TERMINATE_CLUSTER EmrCreateCluster_SpotTimeoutAction = "TERMINATE_CLUSTER"
)

// EBS volume specifications such as volume type, IOPS, and size (GiB) that will be requested for the EBS volume attached to an EC2 instance in the cluster.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_VolumeSpecification.html
//
// Experimental.
type EmrCreateCluster_VolumeSpecificationProperty struct {
	// The volume size.
	//
	// If the volume type is EBS-optimized, the minimum value is 10GiB.
	// Maximum size is 1TiB
	// Experimental.
	VolumeSize awscdk.Size `json:"volumeSize"`
	// The volume type.
	//
	// Volume types supported are gp2, io1, standard.
	// Experimental.
	VolumeType EmrCreateCluster_EbsBlockDeviceVolumeType `json:"volumeType"`
	// The number of I/O operations per second (IOPS) that the volume supports.
	// Experimental.
	Iops *float64 `json:"iops"`
}

// Properties for EmrCreateCluster.
//
// See the RunJobFlow API for complete documentation on input parameters
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_RunJobFlow.html
//
// Experimental.
type EmrCreateClusterProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// A specification of the number and type of Amazon EC2 instances.
	// Experimental.
	Instances *EmrCreateCluster_InstancesConfigProperty `json:"instances"`
	// The Name of the Cluster.
	// Experimental.
	Name *string `json:"name"`
	// A JSON string for selecting additional features.
	// Experimental.
	AdditionalInfo *string `json:"additionalInfo"`
	// A case-insensitive list of applications for Amazon EMR to install and configure when launching the cluster.
	// Experimental.
	Applications *[]*EmrCreateCluster_ApplicationConfigProperty `json:"applications"`
	// An IAM role for automatic scaling policies.
	// Experimental.
	AutoScalingRole awsiam.IRole `json:"autoScalingRole"`
	// An auto-termination policy for an Amazon EMR cluster.
	//
	// An auto-termination policy defines the amount of
	// idle time in seconds after which a cluster automatically terminates. The value must be between
	// 60 seconds and 7 days.
	// Experimental.
	AutoTerminationPolicy *EmrCreateCluster_AutoTerminationPolicyProperty `json:"autoTerminationPolicy"`
	// A list of bootstrap actions to run before Hadoop starts on the cluster nodes.
	// Experimental.
	BootstrapActions *[]*EmrCreateCluster_BootstrapActionConfigProperty `json:"bootstrapActions"`
	// Also called instance profile and EC2 role.
	//
	// An IAM role for an EMR cluster. The EC2 instances of the cluster assume this role.
	//
	// This attribute has been renamed from jobFlowRole to clusterRole to align with other ERM/StepFunction integration parameters.
	// Experimental.
	ClusterRole awsiam.IRole `json:"clusterRole"`
	// The list of configurations supplied for the EMR cluster you are creating.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `json:"configurations"`
	// The ID of a custom Amazon EBS-backed Linux AMI.
	// Experimental.
	CustomAmiId *string `json:"customAmiId"`
	// The size of the EBS root device volume of the Linux AMI that is used for each EC2 instance.
	// Experimental.
	EbsRootVolumeSize awscdk.Size `json:"ebsRootVolumeSize"`
	// Attributes for Kerberos configuration when Kerberos authentication is enabled using a security configuration.
	// Experimental.
	KerberosAttributes *EmrCreateCluster_KerberosAttributesProperty `json:"kerberosAttributes"`
	// The location in Amazon S3 to write the log files of the job flow.
	// Experimental.
	LogUri *string `json:"logUri"`
	// The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster.
	// Experimental.
	ReleaseLabel *string `json:"releaseLabel"`
	// Specifies the way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.
	// Experimental.
	ScaleDownBehavior EmrCreateCluster_EmrClusterScaleDownBehavior `json:"scaleDownBehavior"`
	// The name of a security configuration to apply to the cluster.
	// Experimental.
	SecurityConfiguration *string `json:"securityConfiguration"`
	// The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.
	// Experimental.
	ServiceRole awsiam.IRole `json:"serviceRole"`
	// Specifies the step concurrency level to allow multiple steps to run in parallel.
	//
	// Requires EMR release label 5.28.0 or above.
	// Must be in range [1, 256].
	// Experimental.
	StepConcurrencyLevel *float64 `json:"stepConcurrencyLevel"`
	// A list of tags to associate with a cluster and propagate to Amazon EC2 instances.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// A value of true indicates that all IAM users in the AWS account can perform cluster actions if they have the proper IAM policy permissions.
	// Experimental.
	VisibleToAllUsers *bool `json:"visibleToAllUsers"`
}

// A Step Functions Task to to modify an InstanceFleet on an EMR Cluster.
// Experimental.
type EmrModifyInstanceFleetByName interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrModifyInstanceFleetByName
type jsiiProxy_EmrModifyInstanceFleetByName struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrModifyInstanceFleetByName(scope constructs.Construct, id *string, props *EmrModifyInstanceFleetByNameProps) EmrModifyInstanceFleetByName {
	_init_.Initialize()

	j := jsiiProxy_EmrModifyInstanceFleetByName{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrModifyInstanceFleetByName_Override(e EmrModifyInstanceFleetByName, scope constructs.Construct, id *string, props *EmrModifyInstanceFleetByNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrModifyInstanceFleetByName) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrModifyInstanceFleetByName_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrModifyInstanceFleetByName_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
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
func EmrModifyInstanceFleetByName_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrModifyInstanceFleetByName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrModifyInstanceFleetByName_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceFleetByName",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceFleetByName) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EmrModifyInstanceFleetByName.
// Experimental.
type EmrModifyInstanceFleetByNameProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `json:"clusterId"`
	// The InstanceFleetName to update.
	// Experimental.
	InstanceFleetName *string `json:"instanceFleetName"`
	// The target capacity of On-Demand units for the instance fleet.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetModifyConfig.html
	//
	// Experimental.
	TargetOnDemandCapacity *float64 `json:"targetOnDemandCapacity"`
	// The target capacity of Spot units for the instance fleet.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceFleetModifyConfig.html
	//
	// Experimental.
	TargetSpotCapacity *float64 `json:"targetSpotCapacity"`
}

// A Step Functions Task to to modify an InstanceGroup on an EMR Cluster.
// Experimental.
type EmrModifyInstanceGroupByName interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrModifyInstanceGroupByName
type jsiiProxy_EmrModifyInstanceGroupByName struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrModifyInstanceGroupByName(scope constructs.Construct, id *string, props *EmrModifyInstanceGroupByNameProps) EmrModifyInstanceGroupByName {
	_init_.Initialize()

	j := jsiiProxy_EmrModifyInstanceGroupByName{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrModifyInstanceGroupByName_Override(e EmrModifyInstanceGroupByName, scope constructs.Construct, id *string, props *EmrModifyInstanceGroupByNameProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrModifyInstanceGroupByName) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrModifyInstanceGroupByName_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrModifyInstanceGroupByName_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
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
func EmrModifyInstanceGroupByName_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrModifyInstanceGroupByName_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrModifyInstanceGroupByName_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrModifyInstanceGroupByName",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrModifyInstanceGroupByName) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Modify the size or configurations of an instance group.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroupModifyConfig.html
//
// Experimental.
type EmrModifyInstanceGroupByName_InstanceGroupModifyConfigProperty struct {
	// A list of new or modified configurations to apply for an instance group.
	// Experimental.
	Configurations *[]*EmrCreateCluster_ConfigurationProperty `json:"configurations"`
	// The EC2 InstanceIds to terminate.
	//
	// After you terminate the instances, the instance group will not return to its original requested size.
	// Experimental.
	EC2InstanceIdsToTerminate *[]*string `json:"eC2InstanceIdsToTerminate"`
	// Target size for the instance group.
	// Experimental.
	InstanceCount *float64 `json:"instanceCount"`
	// Policy for customizing shrink operations.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ShrinkPolicy.html
	//
	// Experimental.
	ShrinkPolicy *EmrModifyInstanceGroupByName_ShrinkPolicyProperty `json:"shrinkPolicy"`
}

// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceResizePolicy.html
//
// Experimental.
type EmrModifyInstanceGroupByName_InstanceResizePolicyProperty struct {
	// Specific list of instances to be protected when shrinking an instance group.
	// Experimental.
	InstancesToProtect *[]*string `json:"instancesToProtect"`
	// Specific list of instances to be terminated when shrinking an instance group.
	// Experimental.
	InstancesToTerminate *[]*string `json:"instancesToTerminate"`
	// Decommissioning timeout override for the specific list of instances to be terminated.
	// Experimental.
	InstanceTerminationTimeout awscdk.Duration `json:"instanceTerminationTimeout"`
}

// Policy for customizing shrink operations.
//
// Allows configuration of decommissioning timeout and targeted instance shrinking.
// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ShrinkPolicy.html
//
// Experimental.
type EmrModifyInstanceGroupByName_ShrinkPolicyProperty struct {
	// The desired timeout for decommissioning an instance.
	//
	// Overrides the default YARN decommissioning timeout.
	// Experimental.
	DecommissionTimeout awscdk.Duration `json:"decommissionTimeout"`
	// Custom policy for requesting termination protection or termination of specific instances when shrinking an instance group.
	// Experimental.
	InstanceResizePolicy *EmrModifyInstanceGroupByName_InstanceResizePolicyProperty `json:"instanceResizePolicy"`
}

// Properties for EmrModifyInstanceGroupByName.
// Experimental.
type EmrModifyInstanceGroupByNameProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `json:"clusterId"`
	// The JSON that you want to provide to your ModifyInstanceGroup call as input.
	//
	// This uses the same syntax as the ModifyInstanceGroups API.
	// See: https://docs.aws.amazon.com/emr/latest/APIReference/API_ModifyInstanceGroups.html
	//
	// Experimental.
	InstanceGroup *EmrModifyInstanceGroupByName_InstanceGroupModifyConfigProperty `json:"instanceGroup"`
	// The InstanceGroupName to update.
	// Experimental.
	InstanceGroupName *string `json:"instanceGroupName"`
}

// A Step Functions Task to to set Termination Protection on an EMR Cluster.
// Experimental.
type EmrSetClusterTerminationProtection interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrSetClusterTerminationProtection
type jsiiProxy_EmrSetClusterTerminationProtection struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrSetClusterTerminationProtection(scope constructs.Construct, id *string, props *EmrSetClusterTerminationProtectionProps) EmrSetClusterTerminationProtection {
	_init_.Initialize()

	j := jsiiProxy_EmrSetClusterTerminationProtection{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrSetClusterTerminationProtection_Override(e EmrSetClusterTerminationProtection, scope constructs.Construct, id *string, props *EmrSetClusterTerminationProtectionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrSetClusterTerminationProtection) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrSetClusterTerminationProtection_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrSetClusterTerminationProtection_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
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
func EmrSetClusterTerminationProtection_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrSetClusterTerminationProtection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrSetClusterTerminationProtection_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrSetClusterTerminationProtection",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrSetClusterTerminationProtection) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EmrSetClusterTerminationProtection.
// Experimental.
type EmrSetClusterTerminationProtectionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to update.
	// Experimental.
	ClusterId *string `json:"clusterId"`
	// Termination protection indicator.
	// Experimental.
	TerminationProtected *bool `json:"terminationProtected"`
}

// A Step Functions Task to terminate an EMR Cluster.
// Experimental.
type EmrTerminateCluster interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EmrTerminateCluster
type jsiiProxy_EmrTerminateCluster struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EmrTerminateCluster) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EmrTerminateCluster) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEmrTerminateCluster(scope constructs.Construct, id *string, props *EmrTerminateClusterProps) EmrTerminateCluster {
	_init_.Initialize()

	j := jsiiProxy_EmrTerminateCluster{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEmrTerminateCluster_Override(e EmrTerminateCluster, scope constructs.Construct, id *string, props *EmrTerminateClusterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EmrTerminateCluster) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EmrTerminateCluster) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EmrTerminateCluster_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EmrTerminateCluster_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
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
func EmrTerminateCluster_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EmrTerminateCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EmrTerminateCluster_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EmrTerminateCluster",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EmrTerminateCluster) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EmrTerminateCluster.
// Experimental.
type EmrTerminateClusterProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The ClusterId to terminate.
	// Experimental.
	ClusterId *string `json:"clusterId"`
}

// Encryption Configuration of the S3 bucket.
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html
//
// Experimental.
type EncryptionConfiguration struct {
	// Type of S3 server-side encryption enabled.
	// Experimental.
	EncryptionOption EncryptionOption `json:"encryptionOption"`
	// KMS key ARN or ID.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
}

// Encryption Options of the S3 bucket.
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_EncryptionConfiguration.html#athena-Type-EncryptionConfiguration-EncryptionOption
//
// Experimental.
type EncryptionOption string

const (
	EncryptionOption_S3_MANAGED EncryptionOption = "S3_MANAGED"
	EncryptionOption_KMS EncryptionOption = "KMS"
	EncryptionOption_CLIENT_SIDE_KMS EncryptionOption = "CLIENT_SIDE_KMS"
)

// A Step Functions Task to evaluate an expression.
//
// OUTPUT: the output of this task is the evaluated expression.
// Experimental.
type EvaluateExpression interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EvaluateExpression
type jsiiProxy_EvaluateExpression struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EvaluateExpression) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EvaluateExpression) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEvaluateExpression(scope constructs.Construct, id *string, props *EvaluateExpressionProps) EvaluateExpression {
	_init_.Initialize()

	j := jsiiProxy_EvaluateExpression{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEvaluateExpression_Override(e EvaluateExpression, scope constructs.Construct, id *string, props *EvaluateExpressionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EvaluateExpression) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EvaluateExpression) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EvaluateExpression_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EvaluateExpression_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
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
func EvaluateExpression_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EvaluateExpression_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EvaluateExpression_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EvaluateExpression",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EvaluateExpression) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for EvaluateExpression.
// Experimental.
type EvaluateExpressionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The expression to evaluate.
	//
	// The expression may contain state paths.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Expression *string `json:"expression"`
	// The runtime language to use to evaluate the expression.
	// Experimental.
	Runtime awslambda.Runtime `json:"runtime"`
}

// A StepFunctions Task to send events to an EventBridge event bus.
// Experimental.
type EventBridgePutEvents interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for EventBridgePutEvents
type jsiiProxy_EventBridgePutEvents struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_EventBridgePutEvents) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgePutEvents) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewEventBridgePutEvents(scope constructs.Construct, id *string, props *EventBridgePutEventsProps) EventBridgePutEvents {
	_init_.Initialize()

	j := jsiiProxy_EventBridgePutEvents{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEventBridgePutEvents_Override(e EventBridgePutEvents, scope constructs.Construct, id *string, props *EventBridgePutEventsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_EventBridgePutEvents) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_EventBridgePutEvents) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func EventBridgePutEvents_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func EventBridgePutEvents_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
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
func EventBridgePutEvents_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func EventBridgePutEvents_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func EventBridgePutEvents_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.EventBridgePutEvents",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		e,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		e,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		e,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		e,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) OnPrepare() {
	_jsii_.InvokeVoid(
		e,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) Prepare() {
	_jsii_.InvokeVoid(
		e,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		e,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (e *jsiiProxy_EventBridgePutEvents) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		e,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// An entry to be sent to EventBridge.
// See: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEventsRequestEntry.html
//
// Experimental.
type EventBridgePutEventsEntry struct {
	// The event body.
	//
	// Can either be provided as an object or as a JSON-serialized string
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	Detail awsstepfunctions.TaskInput `json:"detail"`
	// Used along with the source field to help identify the fields and values expected in the detail field.
	//
	// For example, events by CloudTrail have detail type "AWS API Call via CloudTrail"
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	// Experimental.
	DetailType *string `json:"detailType"`
	// The service or application that caused this event to be generated.
	//
	// TODO: EXAMPLE
	//
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	// Experimental.
	Source *string `json:"source"`
	// The event bus the entry will be sent to.
	// Experimental.
	EventBus awsevents.IEventBus `json:"eventBus"`
}

// Properties for sending events with PutEvents.
// Experimental.
type EventBridgePutEventsProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The entries that will be sent (must be at least 1).
	// Experimental.
	Entries *[]*EventBridgePutEventsEntry `json:"entries"`
}

// Start a Job run as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-databrew.html
//
// Experimental.
type GlueDataBrewStartJobRun interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for GlueDataBrewStartJobRun
type jsiiProxy_GlueDataBrewStartJobRun struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewGlueDataBrewStartJobRun(scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) GlueDataBrewStartJobRun {
	_init_.Initialize()

	j := jsiiProxy_GlueDataBrewStartJobRun{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueDataBrewStartJobRun_Override(g GlueDataBrewStartJobRun, scope constructs.Construct, id *string, props *GlueDataBrewStartJobRunProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_GlueDataBrewStartJobRun) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func GlueDataBrewStartJobRun_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func GlueDataBrewStartJobRun_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
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
func GlueDataBrewStartJobRun_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func GlueDataBrewStartJobRun_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func GlueDataBrewStartJobRun_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.GlueDataBrewStartJobRun",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		g,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		g,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (g *jsiiProxy_GlueDataBrewStartJobRun) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for starting a job run with StartJobRun.
// Experimental.
type GlueDataBrewStartJobRunProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Glue DataBrew Job to run.
	// Experimental.
	Name *string `json:"name"`
}

// Starts an AWS Glue job in a Task state.
//
// OUTPUT: the output of this task is a JobRun structure, for details consult
// https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-JobRun
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html
//
// Experimental.
type GlueStartJobRun interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for GlueStartJobRun
type jsiiProxy_GlueStartJobRun struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_GlueStartJobRun) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueStartJobRun) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewGlueStartJobRun(scope constructs.Construct, id *string, props *GlueStartJobRunProps) GlueStartJobRun {
	_init_.Initialize()

	j := jsiiProxy_GlueStartJobRun{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueStartJobRun_Override(g GlueStartJobRun, scope constructs.Construct, id *string, props *GlueStartJobRunProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		[]interface{}{scope, id, props},
		g,
	)
}

func (j *jsiiProxy_GlueStartJobRun) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_GlueStartJobRun) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func GlueStartJobRun_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func GlueStartJobRun_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
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
func GlueStartJobRun_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func GlueStartJobRun_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func GlueStartJobRun_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.GlueStartJobRun",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		g,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		g,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		g,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		g,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		g,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) OnPrepare() {
	_jsii_.InvokeVoid(
		g,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) Prepare() {
	_jsii_.InvokeVoid(
		g,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		g,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (g *jsiiProxy_GlueStartJobRun) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		g,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for starting an AWS Glue job as a task.
// Experimental.
type GlueStartJobRunProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Glue job name.
	// Experimental.
	GlueJobName *string `json:"glueJobName"`
	// The job arguments specifically for this run.
	//
	// For this job run, they replace the default arguments set in the job
	// definition itself.
	// Experimental.
	Arguments awsstepfunctions.TaskInput `json:"arguments"`
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	//
	// Must be at least 1 minute.
	// Experimental.
	NotifyDelayAfter awscdk.Duration `json:"notifyDelayAfter"`
	// The name of the SecurityConfiguration structure to be used with this job run.
	//
	// This must match the Glue API
	// See: https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html#aws-glue-api-regex-oneLine
	//
	// Experimental.
	SecurityConfiguration *string `json:"securityConfiguration"`
}

// Http Methods that API Gateway supports.
// Experimental.
type HttpMethod string

const (
	HttpMethod_GET HttpMethod = "GET"
	HttpMethod_POST HttpMethod = "POST"
	HttpMethod_PUT HttpMethod = "PUT"
	HttpMethod_DELETE HttpMethod = "DELETE"
	HttpMethod_PATCH HttpMethod = "PATCH"
	HttpMethod_HEAD HttpMethod = "HEAD"
	HttpMethod_OPTIONS HttpMethod = "OPTIONS"
)

// Method type of a EKS call.
// Experimental.
type HttpMethods string

const (
	HttpMethods_GET HttpMethods = "GET"
	HttpMethods_POST HttpMethods = "POST"
	HttpMethods_PUT HttpMethods = "PUT"
	HttpMethods_DELETE HttpMethods = "DELETE"
	HttpMethods_PATCH HttpMethods = "PATCH"
	HttpMethods_HEAD HttpMethods = "HEAD"
)

// Configuration of the container used to host the model.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ContainerDefinition.html
//
// Experimental.
type IContainerDefinition interface {
	// Called when the ContainerDefinition is used by a SageMaker task.
	// Experimental.
	Bind(task ISageMakerTask) *ContainerDefinitionConfig
}

// The jsii proxy for IContainerDefinition
type jsiiProxy_IContainerDefinition struct {
	_ byte // padding
}

func (i *jsiiProxy_IContainerDefinition) Bind(task ISageMakerTask) *ContainerDefinitionConfig {
	var returns *ContainerDefinitionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// An Amazon ECS launch type determines the type of infrastructure on which your tasks and services are hosted.
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/launch_types.html
//
// Experimental.
type IEcsLaunchTarget interface {
	// called when the ECS launch target is configured on RunTask.
	// Experimental.
	Bind(task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig
}

// The jsii proxy for IEcsLaunchTarget
type jsiiProxy_IEcsLaunchTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IEcsLaunchTarget) Bind(task EcsRunTask, launchTargetOptions *LaunchTargetBindOptions) *EcsLaunchTargetConfig {
	var returns *EcsLaunchTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{task, launchTargetOptions},
		&returns,
	)

	return returns
}

// Task to train a machine learning model using Amazon SageMaker.
// Experimental.
type ISageMakerTask interface {
	awsiam.IGrantable
}

// The jsii proxy for ISageMakerTask
type jsiiProxy_ISageMakerTask struct {
	internal.Type__awsiamIGrantable
}

// Input mode that the algorithm supports.
// Experimental.
type InputMode string

const (
	InputMode_PIPE InputMode = "PIPE"
	InputMode_FILE InputMode = "FILE"
)

// Invocation type of a Lambda.
// Deprecated: use `LambdaInvocationType`
type InvocationType string

const (
	InvocationType_REQUEST_RESPONSE InvocationType = "REQUEST_RESPONSE"
	InvocationType_EVENT InvocationType = "EVENT"
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
)

// A Step Functions Task to invoke an Activity worker.
//
// An Activity can be used directly as a Resource.
// Deprecated: use `StepFunctionsInvokeActivity`
type InvokeActivity interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for InvokeActivity
type jsiiProxy_InvokeActivity struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `StepFunctionsInvokeActivity`
func NewInvokeActivity(activity awsstepfunctions.IActivity, props *InvokeActivityProps) InvokeActivity {
	_init_.Initialize()

	j := jsiiProxy_InvokeActivity{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeActivity",
		[]interface{}{activity, props},
		&j,
	)

	return &j
}

// Deprecated: use `StepFunctionsInvokeActivity`
func NewInvokeActivity_Override(i InvokeActivity, activity awsstepfunctions.IActivity, props *InvokeActivityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeActivity",
		[]interface{}{activity, props},
		i,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: use `StepFunctionsInvokeActivity`
func (i *jsiiProxy_InvokeActivity) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for FunctionTask.
// Deprecated: use `StepFunctionsInvokeActivity` and `StepFunctionsInvokeActivityProps`.
type InvokeActivityProps struct {
	// Maximum time between heart beats.
	//
	// If the time between heart beats takes longer than this, a 'Timeout' error is raised.
	// Deprecated: use `StepFunctionsInvokeActivity` and `StepFunctionsInvokeActivityProps`.
	Heartbeat awscdk.Duration `json:"heartbeat"`
}

// A Step Functions Task to invoke a Lambda function.
//
// The Lambda function Arn is defined as Resource in the state machine definition.
//
// OUTPUT: the output of this task is the return value of the Lambda Function.
// Deprecated: Use `LambdaInvoke`
type InvokeFunction interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for InvokeFunction
type jsiiProxy_InvokeFunction struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `LambdaInvoke`
func NewInvokeFunction(lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) InvokeFunction {
	_init_.Initialize()

	j := jsiiProxy_InvokeFunction{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeFunction",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Deprecated: Use `LambdaInvoke`
func NewInvokeFunction_Override(i InvokeFunction, lambdaFunction awslambda.IFunction, props *InvokeFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.InvokeFunction",
		[]interface{}{lambdaFunction, props},
		i,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: Use `LambdaInvoke`
func (i *jsiiProxy_InvokeFunction) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for InvokeFunction.
// Deprecated: use `LambdaInvoke`
type InvokeFunctionProps struct {
	// The JSON that you want to provide to your Lambda function as input.
	//
	// This parameter is named as payload to keep consistent with RunLambdaTask class.
	// Deprecated: use `LambdaInvoke`
	Payload *map[string]interface{} `json:"payload"`
}

// An object representing an AWS Batch job dependency.
// Experimental.
type JobDependency struct {
	// The job ID of the AWS Batch job associated with this dependency.
	// Experimental.
	JobId *string `json:"jobId"`
	// The type of the job dependency.
	// Experimental.
	Type *string `json:"type"`
}

// Invocation type of a Lambda.
// Experimental.
type LambdaInvocationType string

const (
	LambdaInvocationType_REQUEST_RESPONSE LambdaInvocationType = "REQUEST_RESPONSE"
	LambdaInvocationType_EVENT LambdaInvocationType = "EVENT"
	LambdaInvocationType_DRY_RUN LambdaInvocationType = "DRY_RUN"
)

// Invoke a Lambda function as a Task.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html
//
// Experimental.
type LambdaInvoke interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for LambdaInvoke
type jsiiProxy_LambdaInvoke struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_LambdaInvoke) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvoke) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaInvoke(scope constructs.Construct, id *string, props *LambdaInvokeProps) LambdaInvoke {
	_init_.Initialize()

	j := jsiiProxy_LambdaInvoke{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvoke_Override(l LambdaInvoke, scope constructs.Construct, id *string, props *LambdaInvokeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaInvoke) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_LambdaInvoke) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func LambdaInvoke_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func LambdaInvoke_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
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
func LambdaInvoke_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func LambdaInvoke_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func LambdaInvoke_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.LambdaInvoke",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		l,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		l,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		l,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		l,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		l,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		l,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		l,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		l,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		l,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		l,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		l,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (l *jsiiProxy_LambdaInvoke) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		l,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for invoking a Lambda function with LambdaInvoke.
// Experimental.
type LambdaInvokeProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Lambda function to invoke.
	// Experimental.
	LambdaFunction awslambda.IFunction `json:"lambdaFunction"`
	// Up to 3583 bytes of base64-encoded data about the invoking client to pass to the function.
	// Experimental.
	ClientContext *string `json:"clientContext"`
	// Invocation type of the Lambda function.
	// Experimental.
	InvocationType LambdaInvocationType `json:"invocationType"`
	// The JSON that will be supplied as input to the Lambda function.
	// Experimental.
	Payload awsstepfunctions.TaskInput `json:"payload"`
	// Invoke the Lambda in a way that only returns the payload response without additional metadata.
	//
	// The `payloadResponseOnly` property cannot be used if `integrationPattern`, `invocationType`,
	// `clientContext`, or `qualifier` are specified.
	// It always uses the REQUEST_RESPONSE behavior.
	// Experimental.
	PayloadResponseOnly *bool `json:"payloadResponseOnly"`
	// Version or alias to invoke a published version of the function.
	//
	// You only need to supply this if you want the version of the Lambda Function to depend
	// on data in the state machine state. If not, you can pass the appropriate Alias or Version object
	// directly as the `lambdaFunction` argument.
	// Experimental.
	Qualifier *string `json:"qualifier"`
	// Whether to retry on Lambda service exceptions.
	//
	// This handles `Lambda.ServiceException`, `Lambda.AWSLambdaException` and
	// `Lambda.SdkClientException` with an interval of 2 seconds, a back-off rate
	// of 2 and 6 maximum attempts.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/bp-lambda-serviceexception.html
	//
	// Experimental.
	RetryOnServiceExceptions *bool `json:"retryOnServiceExceptions"`
}

// Options for binding a launch target to an ECS run job task.
// Experimental.
type LaunchTargetBindOptions struct {
	// Task definition to run Docker containers in Amazon ECS.
	// Experimental.
	TaskDefinition awsecs.ITaskDefinition `json:"taskDefinition"`
	// A regional grouping of one or more container instances on which you can run tasks and services.
	// Experimental.
	Cluster awsecs.ICluster `json:"cluster"`
}

// A message attribute to add to the SNS message.
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html
//
// Experimental.
type MessageAttribute struct {
	// The value of the attribute.
	// Experimental.
	Value interface{} `json:"value"`
	// The data type for the attribute.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
	//
	// Experimental.
	DataType MessageAttributeDataType `json:"dataType"`
}

// The data type set for the SNS message attributes.
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html#SNSMessageAttributes.DataTypes
//
// Experimental.
type MessageAttributeDataType string

const (
	MessageAttributeDataType_STRING MessageAttributeDataType = "STRING"
	MessageAttributeDataType_STRING_ARRAY MessageAttributeDataType = "STRING_ARRAY"
	MessageAttributeDataType_NUMBER MessageAttributeDataType = "NUMBER"
	MessageAttributeDataType_BINARY MessageAttributeDataType = "BINARY"
)

// Specifies the metric name and regular expressions used to parse algorithm logs.
// Experimental.
type MetricDefinition struct {
	// Name of the metric.
	// Experimental.
	Name *string `json:"name"`
	// Regular expression that searches the output of a training job and gets the value of the metric.
	// Experimental.
	Regex *string `json:"regex"`
}

// Specifies how many models the container hosts.
// Experimental.
type Mode string

const (
	Mode_SINGLE_MODEL Mode = "SINGLE_MODEL"
	Mode_MULTI_MODEL Mode = "MULTI_MODEL"
)

// Configures the timeout and maximum number of retries for processing a transform job invocation.
// Experimental.
type ModelClientOptions struct {
	// The maximum number of retries when invocation requests are failing.
	// Experimental.
	InvocationsMaxRetries *float64 `json:"invocationsMaxRetries"`
	// The timeout duration for an invocation request.
	// Experimental.
	InvocationsTimeout awscdk.Duration `json:"invocationsTimeout"`
}

// Configures the S3 bucket where SageMaker will save the result of model training.
// Experimental.
type OutputDataConfig struct {
	// Identifies the S3 path where you want Amazon SageMaker to store the model artifacts.
	// Experimental.
	S3OutputLocation S3Location `json:"s3OutputLocation"`
	// Optional KMS encryption key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
}

// Identifies a model that you want to host and the resources to deploy for hosting it.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ProductionVariant.html
//
// Experimental.
type ProductionVariant struct {
	// The ML compute instance type.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// The name of the model that you want to host.
	//
	// This is the name that you specified when creating the model.
	// Experimental.
	ModelName *string `json:"modelName"`
	// The name of the production variant.
	// Experimental.
	VariantName *string `json:"variantName"`
	// The size of the Elastic Inference (EI) instance to use for the production variant.
	// Experimental.
	AcceleratorType AcceleratorType `json:"acceleratorType"`
	// Number of instances to launch initially.
	// Experimental.
	InitialInstanceCount *float64 `json:"initialInstanceCount"`
	// Determines initial traffic distribution among all of the models that you specify in the endpoint configuration.
	// Experimental.
	InitialVariantWeight *float64 `json:"initialVariantWeight"`
}

// A Step Functions Task to publish messages to SNS topic.
//
// A Function can be used directly as a Resource, but this class mirrors
// integration with other AWS services via a specific class instance.
// Deprecated: Use `SnsPublish`
type PublishToTopic interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for PublishToTopic
type jsiiProxy_PublishToTopic struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `SnsPublish`
func NewPublishToTopic(topic awssns.ITopic, props *PublishToTopicProps) PublishToTopic {
	_init_.Initialize()

	j := jsiiProxy_PublishToTopic{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.PublishToTopic",
		[]interface{}{topic, props},
		&j,
	)

	return &j
}

// Deprecated: Use `SnsPublish`
func NewPublishToTopic_Override(p PublishToTopic, topic awssns.ITopic, props *PublishToTopicProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.PublishToTopic",
		[]interface{}{topic, props},
		p,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: Use `SnsPublish`
func (p *jsiiProxy_PublishToTopic) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for PublishTask.
// Deprecated: Use `SnsPublish`
type PublishToTopicProps struct {
	// The text message to send to the topic.
	// Deprecated: Use `SnsPublish`
	Message awsstepfunctions.TaskInput `json:"message"`
	// The service integration pattern indicates different ways to call Publish to SNS.
	//
	// The valid value is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN.
	// Deprecated: Use `SnsPublish`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// If true, send a different message to every subscription type.
	//
	// If this is set to true, message must be a JSON object with a
	// "default" key and a key for every subscription type (such as "sqs",
	// "email", etc.) The values are strings representing the messages
	// being sent to every subscription type.
	// See: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html#API_Publish_RequestParameters
	//
	// Deprecated: Use `SnsPublish`
	MessagePerSubscriptionType *bool `json:"messagePerSubscriptionType"`
	// Used as the "Subject" line when the message is delivered to email endpoints.
	//
	// Also included, if present, in the standard JSON messages delivered to other endpoints.
	// Deprecated: Use `SnsPublish`
	Subject *string `json:"subject"`
}

// Database and data catalog context in which the query execution occurs.
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionContext.html
//
// Experimental.
type QueryExecutionContext struct {
	// Name of catalog used in query execution.
	// Experimental.
	CatalogName *string `json:"catalogName"`
	// Name of database used in query execution.
	// Experimental.
	DatabaseName *string `json:"databaseName"`
}

// Define the format of the input data.
// Experimental.
type RecordWrapperType string

const (
	RecordWrapperType_NONE RecordWrapperType = "NONE"
	RecordWrapperType_RECORD_IO RecordWrapperType = "RECORD_IO"
)

// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
// Experimental.
type ResourceConfig struct {
	// The number of ML compute instances to use.
	// Experimental.
	InstanceCount *float64 `json:"instanceCount"`
	// ML compute instance type.
	//
	// To provide an instance type from the task input, supply an instance type in the following way
	// where the value in the task input is an EC2 instance type prepended with "ml.":
	//
	// ```ts
	// new ec2.InstanceType(sfn.JsonPath.stringAt('$.path.to.instanceType'));
	// ```
	// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_ResourceConfig.html#sagemaker-Type-ResourceConfig-InstanceType
	//
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// Size of the ML storage volume that you want to provision.
	// Experimental.
	VolumeSize awscdk.Size `json:"volumeSize"`
	// KMS key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s) that run the training job.
	// Experimental.
	VolumeEncryptionKey awskms.IKey `json:"volumeEncryptionKey"`
}

// Location of query result along with S3 bucket configuration.
// See: https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultConfiguration.html
//
// Experimental.
type ResultConfiguration struct {
	// Encryption option used if enabled in S3.
	// Experimental.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration"`
	// S3 path of query results.
	//
	// TODO: EXAMPLE
	//
	// Experimental.
	OutputLocation *awss3.Location `json:"outputLocation"`
}

// A Step Functions Task to run AWS Batch.
// Deprecated: use `BatchSubmitJob`
type RunBatchJob interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunBatchJob
type jsiiProxy_RunBatchJob struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `BatchSubmitJob`
func NewRunBatchJob(props *RunBatchJobProps) RunBatchJob {
	_init_.Initialize()

	j := jsiiProxy_RunBatchJob{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunBatchJob",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: use `BatchSubmitJob`
func NewRunBatchJob_Override(r RunBatchJob, props *RunBatchJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunBatchJob",
		[]interface{}{props},
		r,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: use `BatchSubmitJob`
func (r *jsiiProxy_RunBatchJob) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for RunBatchJob.
// Deprecated: use `BatchSubmitJob`
type RunBatchJobProps struct {
	// The arn of the job definition used by this job.
	// Deprecated: use `BatchSubmitJob`
	JobDefinitionArn *string `json:"jobDefinitionArn"`
	// The name of the job.
	//
	// The first character must be alphanumeric, and up to 128 letters (uppercase and lowercase),
	// numbers, hyphens, and underscores are allowed.
	// Deprecated: use `BatchSubmitJob`
	JobName *string `json:"jobName"`
	// The arn of the job queue into which the job is submitted.
	// Deprecated: use `BatchSubmitJob`
	JobQueueArn *string `json:"jobQueueArn"`
	// The array size can be between 2 and 10,000.
	//
	// If you specify array properties for a job, it becomes an array job.
	// For more information, see Array Jobs in the AWS Batch User Guide.
	// Deprecated: use `BatchSubmitJob`
	ArraySize *float64 `json:"arraySize"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and 10 attempts.
	// If the value of attempts is greater than one,
	// the job is retried on failure the same number of attempts as the value.
	// Deprecated: use `BatchSubmitJob`
	Attempts *float64 `json:"attempts"`
	// A list of container overrides in JSON format that specify the name of a container in the specified job definition and the overrides it should receive.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-containerOverrides
	//
	// Deprecated: use `BatchSubmitJob`
	ContainerOverrides *ContainerOverrides `json:"containerOverrides"`
	// A list of dependencies for the job.
	//
	// A job can depend upon a maximum of 20 jobs.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-dependsOn
	//
	// Deprecated: use `BatchSubmitJob`
	DependsOn *[]*JobDependency `json:"dependsOn"`
	// The service integration pattern indicates different ways to call TerminateCluster.
	//
	// The valid value is either FIRE_AND_FORGET or SYNC.
	// Deprecated: use `BatchSubmitJob`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// The payload to be passed as parametrs to the batch job.
	// Deprecated: use `BatchSubmitJob`
	Payload *map[string]interface{} `json:"payload"`
	// The timeout configuration for this SubmitJob operation.
	//
	// The minimum value for the timeout is 60 seconds.
	// See: https://docs.aws.amazon.com/batch/latest/APIReference/API_SubmitJob.html#Batch-SubmitJob-request-timeout
	//
	// Deprecated: use `BatchSubmitJob`
	Timeout awscdk.Duration `json:"timeout"`
}

// Run an ECS/EC2 Task in a StepFunctions workflow.
// Deprecated: - replaced by `EcsRunTask`
type RunEcsEc2Task interface {
	EcsRunTaskBase
	Connections() awsec2.Connections
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for RunEcsEc2Task
type jsiiProxy_RunEcsEc2Task struct {
	jsiiProxy_EcsRunTaskBase
}

func (j *jsiiProxy_RunEcsEc2Task) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: - replaced by `EcsRunTask`
func NewRunEcsEc2Task(props *RunEcsEc2TaskProps) RunEcsEc2Task {
	_init_.Initialize()

	j := jsiiProxy_RunEcsEc2Task{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsEc2Task",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: - replaced by `EcsRunTask`
func NewRunEcsEc2Task_Override(r RunEcsEc2Task, props *RunEcsEc2TaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsEc2Task",
		[]interface{}{props},
		r,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: - replaced by `EcsRunTask`
func (r *jsiiProxy_RunEcsEc2Task) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Deprecated: - replaced by `EcsRunTask`
func (r *jsiiProxy_RunEcsEc2Task) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		r,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

// Properties to run an ECS task on EC2 in StepFunctionsan ECS.
// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
type RunEcsEc2TaskProps struct {
	// The topic to run the task on.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	Cluster awsecs.ICluster `json:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	TaskDefinition awsecs.TaskDefinition `json:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// Placement constraints.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	PlacementConstraints *[]awsecs.PlacementConstraint `json:"placementConstraints"`
	// Placement strategies.
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	PlacementStrategies *[]awsecs.PlacementStrategy `json:"placementStrategies"`
	// Existing security group to use for the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// In what subnets to place the task's ENIs.
	//
	// (Only applicable in case the TaskDefinition is configured for AwsVpc networking)
	// Deprecated: use `EcsRunTask` and `EcsRunTaskProps`
	Subnets *awsec2.SubnetSelection `json:"subnets"`
}

// Start a service on an ECS cluster.
// Deprecated: replaced by `EcsRunTask`
type RunEcsFargateTask interface {
	EcsRunTaskBase
	Connections() awsec2.Connections
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
	ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup)
}

// The jsii proxy struct for RunEcsFargateTask
type jsiiProxy_RunEcsFargateTask struct {
	jsiiProxy_EcsRunTaskBase
}

func (j *jsiiProxy_RunEcsFargateTask) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}


// Deprecated: replaced by `EcsRunTask`
func NewRunEcsFargateTask(props *RunEcsFargateTaskProps) RunEcsFargateTask {
	_init_.Initialize()

	j := jsiiProxy_RunEcsFargateTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsFargateTask",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: replaced by `EcsRunTask`
func NewRunEcsFargateTask_Override(r RunEcsFargateTask, props *RunEcsFargateTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunEcsFargateTask",
		[]interface{}{props},
		r,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: replaced by `EcsRunTask`
func (r *jsiiProxy_RunEcsFargateTask) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Deprecated: replaced by `EcsRunTask`
func (r *jsiiProxy_RunEcsFargateTask) ConfigureAwsVpcNetworking(vpc awsec2.IVpc, assignPublicIp *bool, subnetSelection *awsec2.SubnetSelection, securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		r,
		"configureAwsVpcNetworking",
		[]interface{}{vpc, assignPublicIp, subnetSelection, securityGroup},
	)
}

// Properties to define an ECS service.
// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
type RunEcsFargateTaskProps struct {
	// The topic to run the task on.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	Cluster awsecs.ICluster `json:"cluster"`
	// Task Definition used for running tasks in the service.
	//
	// Note: this must be TaskDefinition, and not ITaskDefinition,
	// as it requires properties that are not known for imported task definitions
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	TaskDefinition awsecs.TaskDefinition `json:"taskDefinition"`
	// Container setting overrides.
	//
	// Key is the name of the container to override, value is the
	// values you want to override.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	ContainerOverrides *[]*ContainerOverride `json:"containerOverrides"`
	// The service integration pattern indicates different ways to call RunTask in ECS.
	//
	// The valid value for Lambda is FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// Assign public IP addresses to each task.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	AssignPublicIp *bool `json:"assignPublicIp"`
	// Fargate platform version to run this service on.
	//
	// Unless you have specific compatibility requirements, you don't need to
	// specify this.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	PlatformVersion awsecs.FargatePlatformVersion `json:"platformVersion"`
	// Existing security group to use for the tasks.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	SecurityGroup awsec2.ISecurityGroup `json:"securityGroup"`
	// In what subnets to place the task's ENIs.
	// Deprecated: replaced by `EcsRunTask` and `EcsRunTaskProps`
	Subnets *awsec2.SubnetSelection `json:"subnets"`
}

// Invoke a Glue job as a Task.
//
// OUTPUT: the output of this task is a JobRun structure, for details consult
// https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-jobs-runs.html#aws-glue-api-jobs-runs-JobRun
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-glue.html
//
// Deprecated: use `GlueStartJobRun`
type RunGlueJobTask interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunGlueJobTask
type jsiiProxy_RunGlueJobTask struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: use `GlueStartJobRun`
func NewRunGlueJobTask(glueJobName *string, props *RunGlueJobTaskProps) RunGlueJobTask {
	_init_.Initialize()

	j := jsiiProxy_RunGlueJobTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunGlueJobTask",
		[]interface{}{glueJobName, props},
		&j,
	)

	return &j
}

// Deprecated: use `GlueStartJobRun`
func NewRunGlueJobTask_Override(r RunGlueJobTask, glueJobName *string, props *RunGlueJobTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunGlueJobTask",
		[]interface{}{glueJobName, props},
		r,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: use `GlueStartJobRun`
func (r *jsiiProxy_RunGlueJobTask) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Properties for RunGlueJobTask.
// Deprecated: use `GlueStartJobRun`
type RunGlueJobTaskProps struct {
	// The job arguments specifically for this run.
	//
	// For this job run, they replace the default arguments set in the job definition itself.
	// Deprecated: use `GlueStartJobRun`
	Arguments *map[string]*string `json:"arguments"`
	// The service integration pattern indicates different ways to start the Glue job.
	//
	// The valid value for Glue is either FIRE_AND_FORGET or SYNC.
	// Deprecated: use `GlueStartJobRun`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// After a job run starts, the number of minutes to wait before sending a job run delay notification.
	//
	// Must be at least 1 minute.
	// Deprecated: use `GlueStartJobRun`
	NotifyDelayAfter awscdk.Duration `json:"notifyDelayAfter"`
	// The name of the SecurityConfiguration structure to be used with this job run.
	//
	// This must match the Glue API
	// [single-line string pattern](https://docs.aws.amazon.com/glue/latest/dg/aws-glue-api-common.html#aws-glue-api-regex-oneLine).
	// Deprecated: use `GlueStartJobRun`
	SecurityConfiguration *string `json:"securityConfiguration"`
	// The job run timeout.
	//
	// This is the maximum time that a job run can consume resources before it is terminated and enters TIMEOUT status.
	// Must be at least 1 minute.
	// Deprecated: use `GlueStartJobRun`
	Timeout awscdk.Duration `json:"timeout"`
}

// Invoke a Lambda function as a Task.
//
// OUTPUT: the output of this task is either the return value of Lambda's
// Invoke call, or whatever the Lambda Function posted back using
// `SendTaskSuccess/SendTaskFailure` in `waitForTaskToken` mode.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-lambda.html
//
// Deprecated: Use `LambdaInvoke`
type RunLambdaTask interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for RunLambdaTask
type jsiiProxy_RunLambdaTask struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `LambdaInvoke`
func NewRunLambdaTask(lambdaFunction awslambda.IFunction, props *RunLambdaTaskProps) RunLambdaTask {
	_init_.Initialize()

	j := jsiiProxy_RunLambdaTask{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunLambdaTask",
		[]interface{}{lambdaFunction, props},
		&j,
	)

	return &j
}

// Deprecated: Use `LambdaInvoke`
func NewRunLambdaTask_Override(r RunLambdaTask, lambdaFunction awslambda.IFunction, props *RunLambdaTaskProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.RunLambdaTask",
		[]interface{}{lambdaFunction, props},
		r,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: Use `LambdaInvoke`
func (r *jsiiProxy_RunLambdaTask) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for RunLambdaTask.
// Deprecated: Use `LambdaInvoke`
type RunLambdaTaskProps struct {
	// Client context to pass to the function.
	// Deprecated: Use `LambdaInvoke`
	ClientContext *string `json:"clientContext"`
	// The service integration pattern indicates different ways to invoke Lambda function.
	//
	// The valid value for Lambda is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN,
	// it determines whether to pause the workflow until a task token is returned.
	//
	// If this is set to WAIT_FOR_TASK_TOKEN, the JsonPath.taskToken value must be included
	// somewhere in the payload and the Lambda must call
	// `SendTaskSuccess/SendTaskFailure` using that token.
	// Deprecated: Use `LambdaInvoke`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// Invocation type of the Lambda function.
	// Deprecated: Use `LambdaInvoke`
	InvocationType InvocationType `json:"invocationType"`
	// The JSON that you want to provide to your Lambda function as input.
	// Deprecated: Use `LambdaInvoke`
	Payload awsstepfunctions.TaskInput `json:"payload"`
	// Version or alias of the function to be invoked.
	// Deprecated: Use `LambdaInvoke`
	Qualifier *string `json:"qualifier"`
}

// S3 Data Distribution Type.
// Experimental.
type S3DataDistributionType string

const (
	S3DataDistributionType_FULLY_REPLICATED S3DataDistributionType = "FULLY_REPLICATED"
	S3DataDistributionType_SHARDED_BY_S3_KEY S3DataDistributionType = "SHARDED_BY_S3_KEY"
)

// S3 location of the channel data.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_S3DataSource.html
//
// Experimental.
type S3DataSource struct {
	// S3 Uri.
	// Experimental.
	S3Location S3Location `json:"s3Location"`
	// List of one or more attribute names to use that are found in a specified augmented manifest file.
	// Experimental.
	AttributeNames *[]*string `json:"attributeNames"`
	// S3 Data Distribution Type.
	// Experimental.
	S3DataDistributionType S3DataDistributionType `json:"s3DataDistributionType"`
	// S3 Data Type.
	// Experimental.
	S3DataType S3DataType `json:"s3DataType"`
}

// S3 Data Type.
// Experimental.
type S3DataType string

const (
	S3DataType_MANIFEST_FILE S3DataType = "MANIFEST_FILE"
	S3DataType_S3_PREFIX S3DataType = "S3_PREFIX"
	S3DataType_AUGMENTED_MANIFEST_FILE S3DataType = "AUGMENTED_MANIFEST_FILE"
)

// Constructs `IS3Location` objects.
// Experimental.
type S3Location interface {
	Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig
}

// The jsii proxy struct for S3Location
type jsiiProxy_S3Location struct {
	_ byte // padding
}

// Experimental.
func NewS3Location_Override(s S3Location) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		nil, // no parameters
		s,
	)
}

// An `IS3Location` built with a determined bucket and key prefix.
// Experimental.
func S3Location_FromBucket(bucket awss3.IBucket, keyPrefix *string) S3Location {
	_init_.Initialize()

	var returns S3Location

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		"fromBucket",
		[]interface{}{bucket, keyPrefix},
		&returns,
	)

	return returns
}

// An `IS3Location` determined fully by a JSON Path from the task input.
//
// Due to the dynamic nature of those locations, the IAM grants that will be set by `grantRead` and `grantWrite`
// apply to the `*` resource.
// Experimental.
func S3Location_FromJsonExpression(expression *string) S3Location {
	_init_.Initialize()

	var returns S3Location

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.S3Location",
		"fromJsonExpression",
		[]interface{}{expression},
		&returns,
	)

	return returns
}

// Called when the S3Location is bound to a StepFunctions task.
// Experimental.
func (s *jsiiProxy_S3Location) Bind(task ISageMakerTask, opts *S3LocationBindOptions) *S3LocationConfig {
	var returns *S3LocationConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{task, opts},
		&returns,
	)

	return returns
}

// Options for binding an S3 Location.
// Experimental.
type S3LocationBindOptions struct {
	// Allow reading from the S3 Location.
	// Experimental.
	ForReading *bool `json:"forReading"`
	// Allow writing to the S3 Location.
	// Experimental.
	ForWriting *bool `json:"forWriting"`
}

// Stores information about the location of an object in Amazon S3.
// Experimental.
type S3LocationConfig struct {
	// Uniquely identifies the resource in Amazon S3.
	// Experimental.
	Uri *string `json:"uri"`
}

// A Step Functions Task to create a SageMaker endpoint.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateEndpoint interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SageMakerCreateEndpoint
type jsiiProxy_SageMakerCreateEndpoint struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpoint) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerCreateEndpoint(scope constructs.Construct, id *string, props *SageMakerCreateEndpointProps) SageMakerCreateEndpoint {
	_init_.Initialize()

	j := jsiiProxy_SageMakerCreateEndpoint{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerCreateEndpoint_Override(s SageMakerCreateEndpoint, scope constructs.Construct, id *string, props *SageMakerCreateEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerCreateEndpoint) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateEndpoint) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerCreateEndpoint_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerCreateEndpoint_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
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
func SageMakerCreateEndpoint_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerCreateEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerCreateEndpoint_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpoint",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpoint) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// A Step Functions Task to create a SageMaker endpoint configuration.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateEndpointConfig interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SageMakerCreateEndpointConfig
type jsiiProxy_SageMakerCreateEndpointConfig struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerCreateEndpointConfig(scope constructs.Construct, id *string, props *SageMakerCreateEndpointConfigProps) SageMakerCreateEndpointConfig {
	_init_.Initialize()

	j := jsiiProxy_SageMakerCreateEndpointConfig{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerCreateEndpointConfig_Override(s SageMakerCreateEndpointConfig, scope constructs.Construct, id *string, props *SageMakerCreateEndpointConfigProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateEndpointConfig) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerCreateEndpointConfig_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerCreateEndpointConfig_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
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
func SageMakerCreateEndpointConfig_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerCreateEndpointConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerCreateEndpointConfig_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateEndpointConfig",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerCreateEndpointConfig) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for creating an Amazon SageMaker endpoint configuration.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateEndpointConfigProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The name of the endpoint configuration.
	// Experimental.
	EndpointConfigName *string `json:"endpointConfigName"`
	// An list of ProductionVariant objects, one for each model that you want to host at this endpoint.
	//
	// Identifies a model that you want to host and the resources to deploy for hosting it.
	// If you are deploying multiple models, tell Amazon SageMaker how to distribute traffic among the models by specifying variant weights.
	// Experimental.
	ProductionVariants *[]*ProductionVariant `json:"productionVariants"`
	// AWS Key Management Service key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance that hosts the endpoint.
	// Experimental.
	KmsKey awskms.IKey `json:"kmsKey"`
	// Tags to be applied to the endpoint configuration.
	// Experimental.
	Tags awsstepfunctions.TaskInput `json:"tags"`
}

// Properties for creating an Amazon SageMaker endpoint.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateEndpointProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The name of an endpoint configuration.
	// Experimental.
	EndpointConfigName *string `json:"endpointConfigName"`
	// The name of the endpoint.
	//
	// The name must be unique within an AWS Region in your AWS account.
	// Experimental.
	EndpointName *string `json:"endpointName"`
	// Tags to be applied to the endpoint.
	// Experimental.
	Tags awsstepfunctions.TaskInput `json:"tags"`
}

// A Step Functions Task to create a SageMaker model.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateModel interface {
	awsstepfunctions.TaskStateBase
	awsec2.IConnectable
	awsiam.IGrantable
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	Connections() awsec2.Connections
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	GrantPrincipal() awsiam.IPrincipal
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	Role() awsiam.IRole
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SageMakerCreateModel
type jsiiProxy_SageMakerCreateModel struct {
	internal.Type__awsstepfunctionsTaskStateBase
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_SageMakerCreateModel) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateModel) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerCreateModel(scope constructs.Construct, id *string, props *SageMakerCreateModelProps) SageMakerCreateModel {
	_init_.Initialize()

	j := jsiiProxy_SageMakerCreateModel{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerCreateModel_Override(s SageMakerCreateModel, scope constructs.Construct, id *string, props *SageMakerCreateModelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerCreateModel) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateModel) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerCreateModel_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerCreateModel_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
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
func SageMakerCreateModel_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerCreateModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerCreateModel_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateModel",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Add the security group to all instances via the launch configuration security groups array.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		s,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerCreateModel) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for creating an Amazon SageMaker model.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerCreateModelProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The name of the new model.
	// Experimental.
	ModelName *string `json:"modelName"`
	// The definition of the primary docker image containing inference code, associated artifacts, and custom environment map that the inference code uses when the model is deployed for predictions.
	// Experimental.
	PrimaryContainer IContainerDefinition `json:"primaryContainer"`
	// Specifies the containers in the inference pipeline.
	// Experimental.
	Containers *[]IContainerDefinition `json:"containers"`
	// Isolates the model container.
	//
	// No inbound or outbound network calls can be made to or from the model container.
	// Experimental.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation"`
	// An execution role that you can pass in a CreateModel API request.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// The subnets of the VPC to which the hosted model is connected (Note this parameter is only used when VPC is provided).
	// Experimental.
	SubnetSelection *awsec2.SubnetSelection `json:"subnetSelection"`
	// Tags to be applied to the model.
	// Experimental.
	Tags awsstepfunctions.TaskInput `json:"tags"`
	// The VPC that is accessible by the hosted model.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
}

// Class representing the SageMaker Create Training Job task.
// Experimental.
type SageMakerCreateTrainingJob interface {
	awsstepfunctions.TaskStateBase
	awsec2.IConnectable
	awsiam.IGrantable
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	Connections() awsec2.Connections
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	GrantPrincipal() awsiam.IPrincipal
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	Role() awsiam.IRole
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	AddSecurityGroup(securityGroup awsec2.ISecurityGroup)
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SageMakerCreateTrainingJob
type jsiiProxy_SageMakerCreateTrainingJob struct {
	internal.Type__awsstepfunctionsTaskStateBase
	internal.Type__awsec2IConnectable
	internal.Type__awsiamIGrantable
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerCreateTrainingJob(scope constructs.Construct, id *string, props *SageMakerCreateTrainingJobProps) SageMakerCreateTrainingJob {
	_init_.Initialize()

	j := jsiiProxy_SageMakerCreateTrainingJob{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerCreateTrainingJob_Override(s SageMakerCreateTrainingJob, scope constructs.Construct, id *string, props *SageMakerCreateTrainingJobProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateTrainingJob) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerCreateTrainingJob_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerCreateTrainingJob_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
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
func SageMakerCreateTrainingJob_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerCreateTrainingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerCreateTrainingJob_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTrainingJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Add the security group to all instances via the launch configuration security groups array.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) AddSecurityGroup(securityGroup awsec2.ISecurityGroup) {
	_jsii_.InvokeVoid(
		s,
		"addSecurityGroup",
		[]interface{}{securityGroup},
	)
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTrainingJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for creating an Amazon SageMaker training job.
// Experimental.
type SageMakerCreateTrainingJobProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Identifies the training algorithm to use.
	// Experimental.
	AlgorithmSpecification *AlgorithmSpecification `json:"algorithmSpecification"`
	// Describes the various datasets (e.g. train, validation, test) and the Amazon S3 location where stored.
	// Experimental.
	InputDataConfig *[]*Channel `json:"inputDataConfig"`
	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the results of model training.
	// Experimental.
	OutputDataConfig *OutputDataConfig `json:"outputDataConfig"`
	// Training Job Name.
	// Experimental.
	TrainingJobName *string `json:"trainingJobName"`
	// Isolates the training container.
	//
	// No inbound or outbound network calls can be made to or from the training container.
	// Experimental.
	EnableNetworkIsolation *bool `json:"enableNetworkIsolation"`
	// Algorithm-specific parameters that influence the quality of the model.
	//
	// Set hyperparameters before you start the learning process.
	// For a list of hyperparameters provided by Amazon SageMaker
	// See: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
	//
	// Experimental.
	Hyperparameters *map[string]interface{} `json:"hyperparameters"`
	// Specifies the resources, ML compute instances, and ML storage volumes to deploy for model training.
	// Experimental.
	ResourceConfig *ResourceConfig `json:"resourceConfig"`
	// Role for the Training Job.
	//
	// The role must be granted all necessary permissions for the SageMaker training job to
	// be able to operate.
	//
	// See https://docs.aws.amazon.com/fr_fr/sagemaker/latest/dg/sagemaker-roles.html#sagemaker-roles-createtrainingjob-perms
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Sets a time limit for training.
	// Experimental.
	StoppingCondition *StoppingCondition `json:"stoppingCondition"`
	// Tags to be applied to the train job.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// Specifies the VPC that you want your training job to connect to.
	// Experimental.
	VpcConfig *VpcConfig `json:"vpcConfig"`
}

// Class representing the SageMaker Create Transform Job task.
// Experimental.
type SageMakerCreateTransformJob interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	Role() awsiam.IRole
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
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

func (j *jsiiProxy_SageMakerCreateTransformJob) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerCreateTransformJob) SetIteration(val awsstepfunctions.StateGraph) {
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

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerCreateTransformJob",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
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

// Render the choices in ASL JSON format.
// Experimental.
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

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
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

// Render map iterator in ASL JSON format.
// Experimental.
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

// Render the default next state in ASL JSON format.
// Experimental.
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

// Render ResultSelector in ASL JSON format.
// Experimental.
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

// Render error recovery options in ASL JSON format.
// Experimental.
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

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
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

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
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

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerCreateTransformJob) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for creating an Amazon SageMaker transform job task.
// Experimental.
type SageMakerCreateTransformJobProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Name of the model that you want to use for the transform job.
	// Experimental.
	ModelName *string `json:"modelName"`
	// Dataset to be transformed and the Amazon S3 location where it is stored.
	// Experimental.
	TransformInput *TransformInput `json:"transformInput"`
	// Transform Job Name.
	// Experimental.
	TransformJobName *string `json:"transformJobName"`
	// S3 location where you want Amazon SageMaker to save the results from the transform job.
	// Experimental.
	TransformOutput *TransformOutput `json:"transformOutput"`
	// Number of records to include in a mini-batch for an HTTP inference request.
	// Experimental.
	BatchStrategy BatchStrategy `json:"batchStrategy"`
	// Environment variables to set in the Docker container.
	// Experimental.
	Environment *map[string]*string `json:"environment"`
	// Maximum number of parallel requests that can be sent to each instance in a transform job.
	// Experimental.
	MaxConcurrentTransforms *float64 `json:"maxConcurrentTransforms"`
	// Maximum allowed size of the payload, in MB.
	// Experimental.
	MaxPayload awscdk.Size `json:"maxPayload"`
	// Configures the timeout and maximum number of retries for processing a transform job invocation.
	// Experimental.
	ModelClientOptions *ModelClientOptions `json:"modelClientOptions"`
	// Role for the Transform Job.
	// Experimental.
	Role awsiam.IRole `json:"role"`
	// Tags to be applied to the train job.
	// Experimental.
	Tags *map[string]*string `json:"tags"`
	// ML compute instances for the transform job.
	// Experimental.
	TransformResources *TransformResources `json:"transformResources"`
}

// A Step Functions Task to update a SageMaker endpoint.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerUpdateEndpoint interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SageMakerUpdateEndpoint
type jsiiProxy_SageMakerUpdateEndpoint struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSageMakerUpdateEndpoint(scope constructs.Construct, id *string, props *SageMakerUpdateEndpointProps) SageMakerUpdateEndpoint {
	_init_.Initialize()

	j := jsiiProxy_SageMakerUpdateEndpoint{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSageMakerUpdateEndpoint_Override(s SageMakerUpdateEndpoint, scope constructs.Construct, id *string, props *SageMakerUpdateEndpointProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SageMakerUpdateEndpoint) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SageMakerUpdateEndpoint_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SageMakerUpdateEndpoint_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
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
func SageMakerUpdateEndpoint_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SageMakerUpdateEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SageMakerUpdateEndpoint_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SageMakerUpdateEndpoint",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SageMakerUpdateEndpoint) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for updating Amazon SageMaker endpoint.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-sagemaker.html
//
// Experimental.
type SageMakerUpdateEndpointProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The name of the new endpoint configuration.
	// Experimental.
	EndpointConfigName *string `json:"endpointConfigName"`
	// The name of the endpoint whose configuration you want to update.
	// Experimental.
	EndpointName *string `json:"endpointName"`
}

// A StepFunctions Task to send messages to SQS queue.
//
// A Function can be used directly as a Resource, but this class mirrors
// integration with other AWS services via a specific class instance.
// Deprecated: Use `SqsSendMessage`
type SendToQueue interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for SendToQueue
type jsiiProxy_SendToQueue struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `SqsSendMessage`
func NewSendToQueue(queue awssqs.IQueue, props *SendToQueueProps) SendToQueue {
	_init_.Initialize()

	j := jsiiProxy_SendToQueue{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SendToQueue",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Deprecated: Use `SqsSendMessage`
func NewSendToQueue_Override(s SendToQueue, queue awssqs.IQueue, props *SendToQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SendToQueue",
		[]interface{}{queue, props},
		s,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: Use `SqsSendMessage`
func (s *jsiiProxy_SendToQueue) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

// Properties for SendMessageTask.
// Deprecated: Use `SqsSendMessage`
type SendToQueueProps struct {
	// The text message to send to the queue.
	// Deprecated: Use `SqsSendMessage`
	MessageBody awsstepfunctions.TaskInput `json:"messageBody"`
	// The length of time, in seconds, for which to delay a specific message.
	//
	// Valid values are 0-900 seconds.
	// Deprecated: Use `SqsSendMessage`
	Delay awscdk.Duration `json:"delay"`
	// The service integration pattern indicates different ways to call SendMessage to SQS.
	//
	// The valid value is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN.
	// Deprecated: Use `SqsSendMessage`
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// The token used for deduplication of sent messages.
	// Deprecated: Use `SqsSendMessage`
	MessageDeduplicationId *string `json:"messageDeduplicationId"`
	// The tag that specifies that a message belongs to a specific message group.
	//
	// Required for FIFO queues. FIFO ordering applies to messages in the same message
	// group.
	// Deprecated: Use `SqsSendMessage`
	MessageGroupId *string `json:"messageGroupId"`
}

// Configuration for a shuffle option for input data in a channel.
// Experimental.
type ShuffleConfig struct {
	// Determines the shuffling order.
	// Experimental.
	Seed *float64 `json:"seed"`
}

// A Step Functions Task to publish messages to SNS topic.
// Experimental.
type SnsPublish interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SnsPublish
type jsiiProxy_SnsPublish struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_SnsPublish) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsPublish) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsPublish(scope constructs.Construct, id *string, props *SnsPublishProps) SnsPublish {
	_init_.Initialize()

	j := jsiiProxy_SnsPublish{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsPublish_Override(s SnsPublish, scope constructs.Construct, id *string, props *SnsPublishProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SnsPublish) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SnsPublish) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SnsPublish_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SnsPublish_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
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
func SnsPublish_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SnsPublish_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SnsPublish_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SnsPublish",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SnsPublish) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SnsPublish) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SnsPublish) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SnsPublish) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SnsPublish) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SnsPublish) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SnsPublish) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SnsPublish) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SnsPublish) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SnsPublish) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SnsPublish) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SnsPublish) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_SnsPublish) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SnsPublish) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SnsPublish) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for publishing a message to an SNS topic.
// Experimental.
type SnsPublishProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The message you want to send.
	//
	// With the exception of SMS, messages must be UTF-8 encoded strings and
	// at most 256 KB in size.
	// For SMS, each message can contain up to 140 characters.
	// Experimental.
	Message awsstepfunctions.TaskInput `json:"message"`
	// The SNS topic that the task will publish to.
	// Experimental.
	Topic awssns.ITopic `json:"topic"`
	// Add message attributes when publishing.
	//
	// These attributes carry additional metadata about the message and may be used
	// for subscription filters.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-message-attributes.html
	//
	// Experimental.
	MessageAttributes *map[string]*MessageAttribute `json:"messageAttributes"`
	// Send different messages for each transport protocol.
	//
	// For example, you might want to send a shorter message to SMS subscribers
	// and a more verbose message to email and SQS subscribers.
	//
	// Your message must be a JSON object with a top-level JSON key of
	// "default" with a value that is a string
	// You can define other top-level keys that define the message you want to
	// send to a specific transport protocol (i.e. "sqs", "email", "http", etc)
	// See: https://docs.aws.amazon.com/sns/latest/api/API_Publish.html#API_Publish_RequestParameters
	//
	// Experimental.
	MessagePerSubscriptionType *bool `json:"messagePerSubscriptionType"`
	// Used as the "Subject" line when the message is delivered to email endpoints.
	//
	// This field will also be included, if present, in the standard JSON messages
	// delivered to other endpoints.
	// Experimental.
	Subject *string `json:"subject"`
}

// Method to use to split the transform job's data files into smaller batches.
// Experimental.
type SplitType string

const (
	SplitType_NONE SplitType = "NONE"
	SplitType_LINE SplitType = "LINE"
	SplitType_RECORD_IO SplitType = "RECORD_IO"
	SplitType_TF_RECORD SplitType = "TF_RECORD"
)

// A StepFunctions Task to send messages to SQS queue.
// Experimental.
type SqsSendMessage interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for SqsSendMessage
type jsiiProxy_SqsSendMessage struct {
	internal.Type__awsstepfunctionsTaskStateBase
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

func (j *jsiiProxy_SqsSendMessage) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

func (j *jsiiProxy_SqsSendMessage) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
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


// Experimental.
func NewSqsSendMessage(scope constructs.Construct, id *string, props *SqsSendMessageProps) SqsSendMessage {
	_init_.Initialize()

	j := jsiiProxy_SqsSendMessage{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsSendMessage_Override(s SqsSendMessage, scope constructs.Construct, id *string, props *SqsSendMessageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_SqsSendMessage) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_SqsSendMessage) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func SqsSendMessage_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func SqsSendMessage_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
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
func SqsSendMessage_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func SqsSendMessage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func SqsSendMessage_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.SqsSendMessage",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
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

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
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

// Render map iterator in ASL JSON format.
// Experimental.
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

// Render the default next state in ASL JSON format.
// Experimental.
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

// Render ResultSelector in ASL JSON format.
// Experimental.
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

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
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

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_SqsSendMessage) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for sending a message to an SQS queue.
// Experimental.
type SqsSendMessageProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The text message to send to the queue.
	// Experimental.
	MessageBody awsstepfunctions.TaskInput `json:"messageBody"`
	// The SQS queue that messages will be sent to.
	// Experimental.
	Queue awssqs.IQueue `json:"queue"`
	// The length of time, for which to delay a message.
	//
	// Messages that you send to the queue remain invisible to consumers for the duration
	// of the delay period. The maximum allowed delay is 15 minutes.
	// Experimental.
	Delay awscdk.Duration `json:"delay"`
	// The token used for deduplication of sent messages.
	//
	// Any messages sent with the same deduplication ID are accepted successfully,
	// but aren't delivered during the 5-minute deduplication interval.
	// Experimental.
	MessageDeduplicationId *string `json:"messageDeduplicationId"`
	// The tag that specifies that a message belongs to a specific message group.
	//
	// Messages that belong to the same message group are processed in a FIFO manner.
	// Messages in different message groups might be processed out of order.
	// Experimental.
	MessageGroupId *string `json:"messageGroupId"`
}

// A Step Functions Task to call StartExecution on another state machine.
//
// It supports three service integration patterns: FIRE_AND_FORGET, SYNC and WAIT_FOR_TASK_TOKEN.
// Deprecated: - use 'StepFunctionsStartExecution'
type StartExecution interface {
	awsstepfunctions.IStepFunctionsTask
	Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for StartExecution
type jsiiProxy_StartExecution struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: - use 'StepFunctionsStartExecution'
func NewStartExecution(stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) StartExecution {
	_init_.Initialize()

	j := jsiiProxy_StartExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StartExecution",
		[]interface{}{stateMachine, props},
		&j,
	)

	return &j
}

// Deprecated: - use 'StepFunctionsStartExecution'
func NewStartExecution_Override(s StartExecution, stateMachine awsstepfunctions.IStateMachine, props *StartExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StartExecution",
		[]interface{}{stateMachine, props},
		s,
	)
}

// Called when the task object is used in a workflow.
// Deprecated: - use 'StepFunctionsStartExecution'
func (s *jsiiProxy_StartExecution) Bind(task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}

// Properties for StartExecution.
// Deprecated: - use 'StepFunctionsStartExecution'
type StartExecutionProps struct {
	// The JSON input for the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'
	Input *map[string]interface{} `json:"input"`
	// The service integration pattern indicates different ways to call StartExecution to Step Functions.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `json:"integrationPattern"`
	// The name of the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Deprecated: - use 'StepFunctionsStartExecution'
	Name *string `json:"name"`
}

// A Step Functions Task to invoke an Activity worker.
//
// An Activity can be used directly as a Resource.
// Experimental.
type StepFunctionsInvokeActivity interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for StepFunctionsInvokeActivity
type jsiiProxy_StepFunctionsInvokeActivity struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionsInvokeActivity(scope constructs.Construct, id *string, props *StepFunctionsInvokeActivityProps) StepFunctionsInvokeActivity {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsInvokeActivity{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsInvokeActivity_Override(s StepFunctionsInvokeActivity, scope constructs.Construct, id *string, props *StepFunctionsInvokeActivityProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_StepFunctionsInvokeActivity) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func StepFunctionsInvokeActivity_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func StepFunctionsInvokeActivity_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
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
func StepFunctionsInvokeActivity_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func StepFunctionsInvokeActivity_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func StepFunctionsInvokeActivity_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsInvokeActivity",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_StepFunctionsInvokeActivity) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for invoking an Activity worker.
// Experimental.
type StepFunctionsInvokeActivityProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// Step Functions Activity to invoke.
	// Experimental.
	Activity awsstepfunctions.IActivity `json:"activity"`
}

// A Step Functions Task to call StartExecution on another state machine.
//
// It supports three service integration patterns: REQUEST_RESPONSE, RUN_JOB, and WAIT_FOR_TASK_TOKEN.
// Experimental.
type StepFunctionsStartExecution interface {
	awsstepfunctions.TaskStateBase
	Branches() *[]awsstepfunctions.StateGraph
	Comment() *string
	DefaultChoice() awsstepfunctions.State
	SetDefaultChoice(val awsstepfunctions.State)
	EndStates() *[]awsstepfunctions.INextable
	Id() *string
	InputPath() *string
	Iteration() awsstepfunctions.StateGraph
	SetIteration(val awsstepfunctions.StateGraph)
	Node() awscdk.ConstructNode
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	StartState() awsstepfunctions.State
	StateId() *string
	TaskMetrics() *awsstepfunctions.TaskMetricsConfig
	TaskPolicies() *[]awsiam.PolicyStatement
	AddBranch(branch awsstepfunctions.StateGraph)
	AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase
	AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State)
	AddIterator(iteration awsstepfunctions.StateGraph)
	AddPrefix(x *string)
	AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase
	BindToGraph(graph awsstepfunctions.StateGraph)
	MakeDefault(def awsstepfunctions.State)
	MakeNext(next awsstepfunctions.State)
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain
	OnPrepare()
	OnSynthesize(session constructs.ISynthesisSession)
	OnValidate() *[]*string
	Prepare()
	RenderBranches() interface{}
	RenderChoices() interface{}
	RenderInputOutput() interface{}
	RenderIterator() interface{}
	RenderNextEnd() interface{}
	RenderResultSelector() interface{}
	RenderRetryCatch() interface{}
	Synthesize(session awscdk.ISynthesisSession)
	ToStateJson() *map[string]interface{}
	ToString() *string
	Validate() *[]*string
	WhenBoundToGraph(graph awsstepfunctions.StateGraph)
}

// The jsii proxy struct for StepFunctionsStartExecution
type jsiiProxy_StepFunctionsStartExecution struct {
	internal.Type__awsstepfunctionsTaskStateBase
}

func (j *jsiiProxy_StepFunctionsStartExecution) Branches() *[]awsstepfunctions.StateGraph {
	var returns *[]awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) DefaultChoice() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) EndStates() *[]awsstepfunctions.INextable {
	var returns *[]awsstepfunctions.INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) Iteration() awsstepfunctions.StateGraph {
	var returns awsstepfunctions.StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) StartState() awsstepfunctions.State {
	var returns awsstepfunctions.State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) TaskMetrics() *awsstepfunctions.TaskMetricsConfig {
	var returns *awsstepfunctions.TaskMetricsConfig
	_jsii_.Get(
		j,
		"taskMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionsStartExecution) TaskPolicies() *[]awsiam.PolicyStatement {
	var returns *[]awsiam.PolicyStatement
	_jsii_.Get(
		j,
		"taskPolicies",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionsStartExecution(scope constructs.Construct, id *string, props *StepFunctionsStartExecutionProps) StepFunctionsStartExecution {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsStartExecution{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsStartExecution_Override(s StepFunctionsStartExecution, scope constructs.Construct, id *string, props *StepFunctionsStartExecutionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_StepFunctionsStartExecution) SetDefaultChoice(val awsstepfunctions.State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_StepFunctionsStartExecution) SetIteration(val awsstepfunctions.StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func StepFunctionsStartExecution_FilterNextables(states *[]awsstepfunctions.State) *[]awsstepfunctions.INextable {
	_init_.Initialize()

	var returns *[]awsstepfunctions.INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func StepFunctionsStartExecution_FindReachableEndStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
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
func StepFunctionsStartExecution_FindReachableStates(start awsstepfunctions.State, options *awsstepfunctions.FindStateOptions) *[]awsstepfunctions.State {
	_init_.Initialize()

	var returns *[]awsstepfunctions.State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func StepFunctionsStartExecution_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func StepFunctionsStartExecution_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions_tasks.StepFunctionsStartExecution",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

// Add a paralle branch to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

// Add a recovery handler for this state.
//
// When a particular error occurs, execution will continue at the error
// handler instead of failing the state machine execution.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddCatch(handler awsstepfunctions.IChainable, props *awsstepfunctions.CatchProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

// Add a choice branch to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next},
	)
}

// Add a map iterator to this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

// Add a prefix to the stateId of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

// Add retry configuration for this state.
//
// This controls if and how the execution will be retried if a particular
// error occurs.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) AddRetry(props *awsstepfunctions.RetryProps) awsstepfunctions.TaskStateBase {
	var returns awsstepfunctions.TaskStateBase

	_jsii_.Invoke(
		s,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Register this state as part of the given graph.
//
// Don't call this. It will be called automatically when you work
// with states normally.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

// Make the indicated state the default choice transition of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

// Make the indicated state the default transition of this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

// Return the given named metric for this Task.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity fails.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricFailed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricFailed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times the heartbeat times out for this activity.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricHeartbeatTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricHeartbeatTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the Task starts and the time it closes.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricRunTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricRunTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is scheduled.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricScheduled(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduled",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, for which the activity stays in the schedule state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricScheduleTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricScheduleTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity is started.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricStarted(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricStarted",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity succeeds.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricSucceeded(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricSucceeded",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// The interval, in milliseconds, between the time the activity is scheduled and the time it closes.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricTime(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTime",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of times this activity times out.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) MetricTimedOut(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		s,
		"metricTimedOut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Continue normal execution with the given state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) Next(next awsstepfunctions.IChainable) awsstepfunctions.Chain {
	var returns awsstepfunctions.Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Perform final modifications before synthesis.
//
// This method can be implemented by derived constructs in order to perform
// final changes before synthesis. prepare() will be called after child
// constructs have been prepared.
//
// This is an advanced framework feature. Only use this if you
// understand the implications.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

// Render parallel branches in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the choices in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render InputPath/Parameters/OutputPath in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render map iterator in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render the default next state in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render ResultSelector in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Render error recovery options in ASL JSON format.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Allows this construct to emit artifacts into the cloud assembly during synthesis.
//
// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
// as they participate in synthesizing the cloud assembly.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

// Return the Amazon States Language object for this state.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Returns a string representation of this construct.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Validate the current construct.
//
// This method can be implemented by derived constructs in order to perform
// validation logic. It is called on all constructs before synthesis.
//
// Returns: An array of validation error messages, or an empty array if the construct is valid.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Called whenever this state is bound to a graph.
//
// Can be overridden by subclasses.
// Experimental.
func (s *jsiiProxy_StepFunctionsStartExecution) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

// Properties for StartExecution.
// Experimental.
type StepFunctionsStartExecutionProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `json:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `json:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `json:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `json:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `json:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `json:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `json:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `json:"timeout"`
	// The Step Functions state machine to start the execution on.
	// Experimental.
	StateMachine awsstepfunctions.IStateMachine `json:"stateMachine"`
	// Pass the execution ID from the context object to the execution input.
	//
	// This allows the Step Functions UI to link child executions from parent executions, making it easier to trace execution flow across state machines.
	//
	// If you set this property to `true`, the `input` property must be an object (provided by `sfn.TaskInput.fromObject`) or omitted entirely.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-nested-workflows.html#nested-execution-startid
	//
	// Experimental.
	AssociateWithParent *bool `json:"associateWithParent"`
	// The JSON input for the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Experimental.
	Input awsstepfunctions.TaskInput `json:"input"`
	// The name of the execution, same as that of StartExecution.
	// See: https://docs.aws.amazon.com/step-functions/latest/apireference/API_StartExecution.html
	//
	// Experimental.
	Name *string `json:"name"`
}

// Specifies a limit to how long a model training job can run.
//
// When the job reaches the time limit, Amazon SageMaker ends the training job.
// Experimental.
type StoppingCondition struct {
	// The maximum length of time, in seconds, that the training or compilation job can run.
	// Experimental.
	MaxRuntime awscdk.Duration `json:"maxRuntime"`
}

// An environment variable to be set in the container run as a task.
// Experimental.
type TaskEnvironmentVariable struct {
	// Name for the environment variable.
	//
	// Use `JsonPath` class's static methods to specify name from a JSON path.
	// Experimental.
	Name *string `json:"name"`
	// Value of the environment variable.
	//
	// Use `JsonPath` class's static methods to specify value from a JSON path.
	// Experimental.
	Value *string `json:"value"`
}

// S3 location of the input data that the model can consume.
// Experimental.
type TransformDataSource struct {
	// S3 location of the input data.
	// Experimental.
	S3DataSource *TransformS3DataSource `json:"s3DataSource"`
}

// Dataset to be transformed and the Amazon S3 location where it is stored.
// Experimental.
type TransformInput struct {
	// S3 location of the channel data.
	// Experimental.
	TransformDataSource *TransformDataSource `json:"transformDataSource"`
	// The compression type of the transform data.
	// Experimental.
	CompressionType CompressionType `json:"compressionType"`
	// Multipurpose internet mail extension (MIME) type of the data.
	// Experimental.
	ContentType *string `json:"contentType"`
	// Method to use to split the transform job's data files into smaller batches.
	// Experimental.
	SplitType SplitType `json:"splitType"`
}

// S3 location where you want Amazon SageMaker to save the results from the transform job.
// Experimental.
type TransformOutput struct {
	// S3 path where you want Amazon SageMaker to store the results of the transform job.
	// Experimental.
	S3OutputPath *string `json:"s3OutputPath"`
	// MIME type used to specify the output data.
	// Experimental.
	Accept *string `json:"accept"`
	// Defines how to assemble the results of the transform job as a single S3 object.
	// Experimental.
	AssembleWith AssembleWith `json:"assembleWith"`
	// AWS KMS key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	// Experimental.
	EncryptionKey awskms.IKey `json:"encryptionKey"`
}

// ML compute instances for the transform job.
// Experimental.
type TransformResources struct {
	// Number of ML compute instances to use in the transform job.
	// Experimental.
	InstanceCount *float64 `json:"instanceCount"`
	// ML compute instance type for the transform job.
	// Experimental.
	InstanceType awsec2.InstanceType `json:"instanceType"`
	// AWS KMS key that Amazon SageMaker uses to encrypt data on the storage volume attached to the ML compute instance(s).
	// Experimental.
	VolumeEncryptionKey awskms.IKey `json:"volumeEncryptionKey"`
}

// Location of the channel data.
// Experimental.
type TransformS3DataSource struct {
	// Identifies either a key name prefix or a manifest.
	// Experimental.
	S3Uri *string `json:"s3Uri"`
	// S3 Data Type.
	// Experimental.
	S3DataType S3DataType `json:"s3DataType"`
}

// Specifies the VPC that you want your Amazon SageMaker training job to connect to.
// Experimental.
type VpcConfig struct {
	// VPC.
	// Experimental.
	Vpc awsec2.IVpc `json:"vpc"`
	// VPC subnets.
	// Experimental.
	Subnets *awsec2.SubnetSelection `json:"subnets"`
}

