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

// Call REST API endpoint as a Task.
//
// Be aware that the header values must be arrays. When passing the Task Token
// in the headers field `WAIT_FOR_TASK_TOKEN` integration, use
// `JsonPath.array()` to wrap the token in an array:
//
// ```ts
// import * as apigateway from '@aws-cdk/aws-apigateway';
// declare const api: apigateway.RestApi;
//
// new tasks.CallApiGatewayRestApiEndpoint(this, 'Endpoint', {
//    api,
//    stageName: 'Stage',
//    method: tasks.HttpMethod.PUT,
//    integrationPattern: sfn.IntegrationPattern.WAIT_FOR_TASK_TOKEN,
//    headers: sfn.TaskInput.fromObject({
//      TaskToken: sfn.JsonPath.array(sfn.JsonPath.taskToken),
//    }),
// });
// ```.
//
// Example:
//   import apigateway "github.com/aws/aws-cdk-go/awscdk"
//   var api restApi
//
//
//   tasks.NewCallApiGatewayRestApiEndpoint(this, jsii.String("Endpoint"), &callApiGatewayRestApiEndpointProps{
//   	api: api,
//   	stageName: jsii.String("Stage"),
//   	method: tasks.httpMethod_PUT,
//   	integrationPattern: sfn.integrationPattern_WAIT_FOR_TASK_TOKEN,
//   	headers: sfn.taskInput.fromObject(map[string]interface{}{
//   		"TaskToken": sfn.JsonPath.array(sfn.JsonPath.taskToken),
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-api-gateway.html
//
// Experimental.
type CallApiGatewayRestApiEndpoint interface {
	awsstepfunctions.TaskStateBase
	// Experimental.
	ApiEndpoint() *string
	// Experimental.
	ArnForExecuteApi() *string
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
	// Experimental.
	StageName() *string
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
	// Experimental.
	CreatePolicyStatements() *[]awsiam.PolicyStatement
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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddBranch(branch awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddChoice(condition awsstepfunctions.Condition, next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddIterator(iteration awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) BindToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MakeDefault(def awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) MakeNext(next awsstepfunctions.State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

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

func (c *jsiiProxy_CallApiGatewayRestApiEndpoint) WhenBoundToGraph(graph awsstepfunctions.StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

