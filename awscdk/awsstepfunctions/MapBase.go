package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Map state in the state machine.
//
// A `Map` state can be used to run a set of steps for each element of an input array.
// A Map state will execute the same steps for multiple entries of an array in the state input.
//
// While the Parallel state executes multiple branches of steps using the same input, a Map state
// will execute the same steps for multiple entries of an array in the state input.
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
//
type MapBase interface {
	State
	INextable
	Arguments() *map[string]interface{}
	Assign() *map[string]interface{}
	Branches() *[]StateGraph
	Comment() *string
	DefaultChoice() State
	SetDefaultChoice(val State)
	// Continuable states of this Chainable.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	Id() *string
	InputPath() *string
	Items() ProvideItems
	ItemSelector() *map[string]interface{}
	ItemsPath() *string
	Iteration() StateGraph
	SetIteration(val StateGraph)
	JsonataItemSelector() *string
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Outputs() *map[string]interface{}
	Parameters() *map[string]interface{}
	Processor() StateGraph
	SetProcessor(val StateGraph)
	ProcessorConfig() *ProcessorConfig
	SetProcessorConfig(val *ProcessorConfig)
	ProcessorMode() ProcessorMode
	SetProcessorMode(val ProcessorMode)
	QueryLanguage() QueryLanguage
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	StateName() *string
	// Add a parallel branch to this state.
	AddBranch(branch StateGraph)
	// Add a choice branch to this state.
	AddChoice(condition Condition, next State, options *ChoiceTransitionOptions)
	// Add a item processor to this state.
	AddItemProcessor(processor StateGraph, config *ProcessorConfig)
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
	// Render the assign in ASL JSON format.
	RenderAssign(topLevelQueryLanguage QueryLanguage) interface{}
	// Render parallel branches in ASL JSON format.
	RenderBranches() interface{}
	// Render the choices in ASL JSON format.
	RenderChoices(topLevelQueryLanguage QueryLanguage) interface{}
	// Render InputPath/Parameters/OutputPath/Arguments/Output in ASL JSON format.
	RenderInputOutput() interface{}
	// Render ItemProcessor in ASL JSON format.
	RenderItemProcessor() interface{}
	// Render map iterator in ASL JSON format.
	RenderIterator() interface{}
	// Render the default next state in ASL JSON format.
	RenderNextEnd() interface{}
	// Render QueryLanguage in ASL JSON format if needed.
	RenderQueryLanguage(topLevelQueryLanguage QueryLanguage) interface{}
	// Render ResultSelector in ASL JSON format.
	RenderResultSelector() interface{}
	// Render error recovery options in ASL JSON format.
	RenderRetryCatch(topLevelQueryLanguage QueryLanguage) interface{}
	// Return the Amazon States Language object for this state.
	ToStateJson(topLevelQueryLanguage QueryLanguage) *map[string]interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Validate this state.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for MapBase
type jsiiProxy_MapBase struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_MapBase) Arguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"arguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Assign() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"assign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Items() ProvideItems {
	var returns ProvideItems
	_jsii_.Get(
		j,
		"items",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ItemSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"itemSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ItemsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) JsonataItemSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jsonataItemSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Outputs() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"outputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) QueryLanguage() QueryLanguage {
	var returns QueryLanguage
	_jsii_.Get(
		j,
		"queryLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MapBase) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewMapBase_Override(m MapBase, scope constructs.Construct, id *string, props *MapBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
		[]interface{}{scope, id, props},
		m,
	)
}

func (j *jsiiProxy_MapBase)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_MapBase)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_MapBase)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_MapBase)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_MapBase)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func MapBase_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateMapBase_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func MapBase_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateMapBase_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func MapBase_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateMapBase_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
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
func MapBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMapBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func MapBase_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateMapBase_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.MapBase",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (m *jsiiProxy_MapBase) AddBranch(branch StateGraph) {
	if err := m.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addBranch",
		[]interface{}{branch},
	)
}

func (m *jsiiProxy_MapBase) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := m.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (m *jsiiProxy_MapBase) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := m.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (m *jsiiProxy_MapBase) AddIterator(iteration StateGraph) {
	if err := m.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (m *jsiiProxy_MapBase) AddPrefix(x *string) {
	if err := m.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addPrefix",
		[]interface{}{x},
	)
}

func (m *jsiiProxy_MapBase) BindToGraph(graph StateGraph) {
	if err := m.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (m *jsiiProxy_MapBase) MakeDefault(def State) {
	if err := m.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeDefault",
		[]interface{}{def},
	)
}

func (m *jsiiProxy_MapBase) MakeNext(next State) {
	if err := m.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeNext",
		[]interface{}{next},
	)
}

func (m *jsiiProxy_MapBase) Next(next IChainable) Chain {
	if err := m.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		m,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderAssign(topLevelQueryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderAssign",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderChoices(topLevelQueryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderChoices",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderQueryLanguage(topLevelQueryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderQueryLanguage",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) RenderRetryCatch(topLevelQueryLanguage QueryLanguage) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderRetryCatch",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) ToStateJson(topLevelQueryLanguage QueryLanguage) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"toStateJson",
		[]interface{}{topLevelQueryLanguage},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MapBase) WhenBoundToGraph(graph StateGraph) {
	if err := m.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

