package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// State defined by supplying Amazon States Language (ASL) in the state machine.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   // create a table
//   table := dynamodb.NewTable(this, jsii.String("montable"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
//   finalStatus := sfn.NewPass(this, jsii.String("final step"))
//
//   // States language JSON to put an item into DynamoDB
//   // snippet generated from https://docs.aws.amazon.com/step-functions/latest/dg/tutorial-code-snippet.html#tutorial-code-snippet-1
//   stateJson := map[string]interface{}{
//   	"Type": jsii.String("Task"),
//   	"Resource": jsii.String("arn:aws:states:::dynamodb:putItem"),
//   	"Parameters": map[string]interface{}{
//   		"TableName": table.tableName,
//   		"Item": map[string]map[string]*string{
//   			"id": map[string]*string{
//   				"S": jsii.String("MyEntry"),
//   			},
//   		},
//   	},
//   	"ResultPath": nil,
//   }
//
//   // custom state which represents a task to insert data into DynamoDB
//   custom := sfn.NewCustomState(this, jsii.String("my custom task"), &CustomStateProps{
//   	StateJson: StateJson,
//   })
//
//   chain := sfn.Chain_Start(custom).Next(finalStatus)
//
//   sm := sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	Definition: chain,
//   	Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
//   })
//
//   // don't forget permissions. You need to assign them
//   table.grantWriteData(sm)
//
type CustomState interface {
	State
	IChainable
	INextable
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
	// Continue normal execution with the given state.
	Next(next IChainable) Chain
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
	// Returns the Amazon States Language object for this state.
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

// The jsii proxy struct for CustomState
type jsiiProxy_CustomState struct {
	jsiiProxy_State
	jsiiProxy_IChainable
	jsiiProxy_INextable
}

func (j *jsiiProxy_CustomState) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomState) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


func NewCustomState(scope constructs.Construct, id *string, props *CustomStateProps) CustomState {
	_init_.Initialize()

	if err := validateNewCustomStateParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomState{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomState_Override(c CustomState, scope constructs.Construct, id *string, props *CustomStateProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		[]interface{}{scope, id, props},
		c,
	)
}

func (j *jsiiProxy_CustomState)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_CustomState)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func CustomState_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateCustomState_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func CustomState_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateCustomState_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func CustomState_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateCustomState_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
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
func CustomState_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomState_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func CustomState_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateCustomState_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.CustomState",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (c *jsiiProxy_CustomState) AddBranch(branch StateGraph) {
	if err := c.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

func (c *jsiiProxy_CustomState) AddChoice(condition Condition, next State) {
	if err := c.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (c *jsiiProxy_CustomState) AddIterator(iteration StateGraph) {
	if err := c.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (c *jsiiProxy_CustomState) AddPrefix(x *string) {
	if err := c.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

func (c *jsiiProxy_CustomState) BindToGraph(graph StateGraph) {
	if err := c.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (c *jsiiProxy_CustomState) MakeDefault(def State) {
	if err := c.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

func (c *jsiiProxy_CustomState) MakeNext(next State) {
	if err := c.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

func (c *jsiiProxy_CustomState) Next(next IChainable) Chain {
	if err := c.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		c,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomState) WhenBoundToGraph(graph StateGraph) {
	if err := c.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

