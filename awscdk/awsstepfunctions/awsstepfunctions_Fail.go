package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Fail state in the state machine.
//
// Reaching a Fail state terminates the state execution in failure.
//
// Example:
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var submitLambda function
//   var getStatusLambda function
//
//
//   submitJob := tasks.NewLambdaInvoke(this, jsii.String("Submit Job"), &lambdaInvokeProps{
//   	lambdaFunction: submitLambda,
//   	// Lambda's result is in the attribute `Payload`
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   waitX := sfn.NewWait(this, jsii.String("Wait X Seconds"), &waitProps{
//   	time: sfn.waitTime.secondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   getStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Pass just the field named "guid" into the Lambda, put the
//   	// Lambda's result in a field called "status" in the response
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   jobFailed := sfn.NewFail(this, jsii.String("Job Failed"), &failProps{
//   	cause: jsii.String("AWS Batch Job Failed"),
//   	error: jsii.String("DescribeJob returned FAILED"),
//   })
//
//   finalStatus := tasks.NewLambdaInvoke(this, jsii.String("Get Final Job Status"), &lambdaInvokeProps{
//   	lambdaFunction: getStatusLambda,
//   	// Use "guid" field as input
//   	inputPath: jsii.String("$.guid"),
//   	outputPath: jsii.String("$.Payload"),
//   })
//
//   definition := submitJob.next(waitX).next(getStatus).next(sfn.NewChoice(this, jsii.String("Job Complete?")).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("FAILED")), jobFailed).when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("SUCCEEDED")), finalStatus).otherwise(waitX))
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &stateMachineProps{
//   	definition: definition,
//   	timeout: awscdk.Duration.minutes(jsii.Number(5)),
//   })
//
type Fail interface {
	State
	Branches() *[]StateGraph
	Comment() *string
	DefaultChoice() State
	SetDefaultChoice(val State)
	// Continuable states of this Chainable.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Iteration() StateGraph
	SetIteration(val StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	// Add a paralle branch to this state.
	AddBranch(branch StateGraph)
	// Add a choice branch to this state.
	AddChoice(condition Condition, next State)
	// Add a map iterator to this state.
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph StateGraph)
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def State)
	// Make the indicated state the default transition of this state.
	MakeNext(next State)
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
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Fail
type jsiiProxy_Fail struct {
	jsiiProxy_State
}

func (j *jsiiProxy_Fail) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Fail) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


func NewFail(scope constructs.Construct, id *string, props *FailProps) Fail {
	_init_.Initialize()

	j := jsiiProxy_Fail{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewFail_Override(f Fail, scope constructs.Construct, id *string, props *FailProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		[]interface{}{scope, id, props},
		f,
	)
}

func (j *jsiiProxy_Fail) SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Fail) SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Fail_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Fail_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Fail_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Fail",
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
func Fail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Fail_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Fail",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (f *jsiiProxy_Fail) AddBranch(branch StateGraph) {
	_jsii_.InvokeVoid(
		f,
		"addBranch",
		[]interface{}{branch},
	)
}

func (f *jsiiProxy_Fail) AddChoice(condition Condition, next State) {
	_jsii_.InvokeVoid(
		f,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (f *jsiiProxy_Fail) AddIterator(iteration StateGraph) {
	_jsii_.InvokeVoid(
		f,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (f *jsiiProxy_Fail) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		f,
		"addPrefix",
		[]interface{}{x},
	)
}

func (f *jsiiProxy_Fail) BindToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		f,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (f *jsiiProxy_Fail) MakeDefault(def State) {
	_jsii_.InvokeVoid(
		f,
		"makeDefault",
		[]interface{}{def},
	)
}

func (f *jsiiProxy_Fail) MakeNext(next State) {
	_jsii_.InvokeVoid(
		f,
		"makeNext",
		[]interface{}{next},
	)
}

func (f *jsiiProxy_Fail) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_Fail) WhenBoundToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		f,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

