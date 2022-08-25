package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a Map state in the state machine.
//
// A `Map` state can be used to run a set of steps for each element of an input array.
// A Map state will execute the same steps for multiple entries of an array in the state input.
//
// While the Parallel state executes multiple branches of steps using the same input, a Map state
// will execute the same steps for multiple entries of an array in the state input.
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &mapProps{
//   	maxConcurrency: jsii.Number(1),
//   	itemsPath: sfn.jsonPath.stringAt(jsii.String("$.inputForMap")),
//   })
//   map.iterator(sfn.NewPass(this, jsii.String("Pass State")))
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
//
// Experimental.
type Map interface {
	State
	INextable
	// Experimental.
	Branches() *[]StateGraph
	// Experimental.
	Comment() *string
	// Experimental.
	DefaultChoice() State
	// Experimental.
	SetDefaultChoice(val State)
	// Continuable states of this Chainable.
	// Experimental.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	// Experimental.
	Id() *string
	// Experimental.
	InputPath() *string
	// Experimental.
	Iteration() StateGraph
	// Experimental.
	SetIteration(val StateGraph)
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
	// First state of this Chainable.
	// Experimental.
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	// Experimental.
	StateId() *string
	// Add a paralle branch to this state.
	// Experimental.
	AddBranch(branch StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	// Experimental.
	AddCatch(handler IChainable, props *CatchProps) Map
	// Add a choice branch to this state.
	// Experimental.
	AddChoice(condition Condition, next State)
	// Add a map iterator to this state.
	// Experimental.
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	// Experimental.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	// Experimental.
	AddRetry(props *RetryProps) Map
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	// Experimental.
	BindToGraph(graph StateGraph)
	// Define iterator state machine in Map.
	// Experimental.
	Iterator(iterator IChainable) Map
	// Make the indicated state the default choice transition of this state.
	// Experimental.
	MakeDefault(def State)
	// Make the indicated state the default transition of this state.
	// Experimental.
	MakeNext(next State)
	// Continue normal execution with the given state.
	// Experimental.
	Next(next IChainable) Chain
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
	// Validate this state.
	// Experimental.
	Validate() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	// Experimental.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Map
type jsiiProxy_Map struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_Map) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


// Experimental.
func NewMap(scope constructs.Construct, id *string, props *MapProps) Map {
	_init_.Initialize()

	j := jsiiProxy_Map{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Map",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMap_Override(m Map, scope constructs.Construct, id *string, props *MapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Map",
		[]interface{}{scope, id, props},
		m,
	)
}

func (j *jsiiProxy_Map) SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Map) SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func Map_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	var returns *[]INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Map",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func Map_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Map",
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
func Map_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Map",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Map_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Map",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func Map_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions.Map",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (m *jsiiProxy_Map) AddBranch(branch StateGraph) {
	_jsii_.InvokeVoid(
		m,
		"addBranch",
		[]interface{}{branch},
	)
}

func (m *jsiiProxy_Map) AddCatch(handler IChainable, props *CatchProps) Map {
	var returns Map

	_jsii_.Invoke(
		m,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) AddChoice(condition Condition, next State) {
	_jsii_.InvokeVoid(
		m,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (m *jsiiProxy_Map) AddIterator(iteration StateGraph) {
	_jsii_.InvokeVoid(
		m,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (m *jsiiProxy_Map) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		m,
		"addPrefix",
		[]interface{}{x},
	)
}

func (m *jsiiProxy_Map) AddRetry(props *RetryProps) Map {
	var returns Map

	_jsii_.Invoke(
		m,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) BindToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		m,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (m *jsiiProxy_Map) Iterator(iterator IChainable) Map {
	var returns Map

	_jsii_.Invoke(
		m,
		"iterator",
		[]interface{}{iterator},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) MakeDefault(def State) {
	_jsii_.InvokeVoid(
		m,
		"makeDefault",
		[]interface{}{def},
	)
}

func (m *jsiiProxy_Map) MakeNext(next State) {
	_jsii_.InvokeVoid(
		m,
		"makeNext",
		[]interface{}{next},
	)
}

func (m *jsiiProxy_Map) Next(next IChainable) Chain {
	var returns Chain

	_jsii_.Invoke(
		m,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) OnPrepare() {
	_jsii_.InvokeVoid(
		m,
		"onPrepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Map) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_Map) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) Prepare() {
	_jsii_.InvokeVoid(
		m,
		"prepare",
		nil, // no parameters
	)
}

func (m *jsiiProxy_Map) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		m,
		"synthesize",
		[]interface{}{session},
	)
}

func (m *jsiiProxy_Map) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) WhenBoundToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		m,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

