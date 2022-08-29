package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Choice in the state machine.
//
// A choice state can be used to make decisions based on the execution
// state.
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
type Choice interface {
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
	// Return a Chain that contains all reachable end states from this Choice.
	//
	// Use this to combine all possible choice paths back.
	Afterwards(options *AfterwardsOptions) Chain
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph StateGraph)
	// Make the indicated state the default choice transition of this state.
	MakeDefault(def State)
	// Make the indicated state the default transition of this state.
	MakeNext(next State)
	// If none of the given conditions match, continue execution with the given state.
	//
	// If no conditions match and no otherwise() has been given, an execution
	// error will be raised.
	Otherwise(def IChainable) Choice
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
	// If the given condition matches, continue execution with the given state.
	When(condition Condition, next IChainable) Choice
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Choice
type jsiiProxy_Choice struct {
	jsiiProxy_State
}

func (j *jsiiProxy_Choice) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


func NewChoice(scope constructs.Construct, id *string, props *ChoiceProps) Choice {
	_init_.Initialize()

	j := jsiiProxy_Choice{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewChoice_Override(c Choice, scope constructs.Construct, id *string, props *ChoiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_Choice) SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Choice) SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Choice_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Choice_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Choice_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Choice",
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
func Choice_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Choice_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (c *jsiiProxy_Choice) AddBranch(branch StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

func (c *jsiiProxy_Choice) AddChoice(condition Condition, next State) {
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (c *jsiiProxy_Choice) AddIterator(iteration StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (c *jsiiProxy_Choice) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

func (c *jsiiProxy_Choice) Afterwards(options *AfterwardsOptions) Chain {
	var returns Chain

	_jsii_.Invoke(
		c,
		"afterwards",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) BindToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (c *jsiiProxy_Choice) MakeDefault(def State) {
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

func (c *jsiiProxy_Choice) MakeNext(next State) {
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

func (c *jsiiProxy_Choice) Otherwise(def IChainable) Choice {
	var returns Choice

	_jsii_.Invoke(
		c,
		"otherwise",
		[]interface{}{def},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) When(condition Condition, next IChainable) Choice {
	var returns Choice

	_jsii_.Invoke(
		c,
		"when",
		[]interface{}{condition, next},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) WhenBoundToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

