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
//
// Example:
//   map := sfn.NewMap(this, jsii.String("Map State"), &MapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   	ItemSelector: map[string]interface{}{
//   		"item": sfn.JsonPath_*StringAt(jsii.String("$.Map.Item.Value")),
//   	},
//   	ResultPath: jsii.String("$.mapOutput"),
//   })
//
//   // The Map iterator can contain a IChainable, which can be an individual or multiple steps chained together.
//   // Below example is with a Choice and Pass step
//   choice := sfn.NewChoice(this, jsii.String("Choice"))
//   condition1 := sfn.Condition_StringEquals(jsii.String("$.item.status"), jsii.String("SUCCESS"))
//   step1 := sfn.NewPass(this, jsii.String("Step1"))
//   step2 := sfn.NewPass(this, jsii.String("Step2"))
//   finish := sfn.NewPass(this, jsii.String("Finish"))
//
//   definition := choice.When(condition1, step1).Otherwise(step2).Afterwards().Next(finish)
//
//   map.ItemProcessor(definition)
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/amazon-states-language-map-state.html
//
type Map interface {
	MapBase
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
	ItemSelector() *map[string]interface{}
	ItemsPath() *string
	Iteration() StateGraph
	SetIteration(val StateGraph)
	// The tree node.
	Node() constructs.Node
	OutputPath() *string
	Parameters() *map[string]interface{}
	Processor() StateGraph
	SetProcessor(val StateGraph)
	ProcessorConfig() *ProcessorConfig
	SetProcessorConfig(val *ProcessorConfig)
	ProcessorMode() ProcessorMode
	SetProcessorMode(val ProcessorMode)
	ResultPath() *string
	ResultSelector() *map[string]interface{}
	// First state of this Chainable.
	StartState() State
	// Tokenized string that evaluates to the state's ID.
	StateId() *string
	StateName() *string
	// Add a parallel branch to this state.
	AddBranch(branch StateGraph)
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler IChainable, props *CatchProps) Map
	// Add a choice branch to this state.
	AddChoice(condition Condition, next State, options *ChoiceTransitionOptions)
	// Add a item processor to this state.
	AddItemProcessor(processor StateGraph, config *ProcessorConfig)
	// Add a map iterator to this state.
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	AddPrefix(x *string)
	// Add retry configuration for this state.
	//
	// This controls if and how the execution will be retried if a particular
	// error occurs.
	AddRetry(props *RetryProps) Map
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph StateGraph)
	// Define item processor in Map.
	//
	// A Map must either have a non-empty iterator or a non-empty item processor (mutually exclusive  with `iterator`).
	ItemProcessor(processor IChainable, config *ProcessorConfig) Map
	// Define iterator state machine in Map.
	//
	// A Map must either have a non-empty iterator or a non-empty item processor (mutually exclusive  with `itemProcessor`).
	// Deprecated: - use `itemProcessor` instead.
	Iterator(iterator IChainable) Map
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
	// Render ItemProcessor in ASL JSON format.
	RenderItemProcessor() interface{}
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
	// Validate this state.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Map
type jsiiProxy_Map struct {
	jsiiProxy_MapBase
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

func (j *jsiiProxy_Map) ItemSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"itemSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) ItemsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsPath",
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

func (j *jsiiProxy_Map) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_Map) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Map) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
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

func (j *jsiiProxy_Map) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewMap(scope constructs.Construct, id *string, props *MapProps) Map {
	_init_.Initialize()

	if err := validateNewMapParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Map{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Map",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewMap_Override(m Map, scope constructs.Construct, id *string, props *MapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Map",
		[]interface{}{scope, id, props},
		m,
	)
}

func (j *jsiiProxy_Map)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Map)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_Map)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Map)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Map)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Map_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateMap_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Map",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Map_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateMap_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Map",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Map_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateMap_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Map",
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
func Map_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Map",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Map_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateMap_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Map",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (m *jsiiProxy_Map) AddBranch(branch StateGraph) {
	if err := m.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addBranch",
		[]interface{}{branch},
	)
}

func (m *jsiiProxy_Map) AddCatch(handler IChainable, props *CatchProps) Map {
	if err := m.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns Map

	_jsii_.Invoke(
		m,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := m.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (m *jsiiProxy_Map) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := m.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (m *jsiiProxy_Map) AddIterator(iteration StateGraph) {
	if err := m.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (m *jsiiProxy_Map) AddPrefix(x *string) {
	if err := m.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addPrefix",
		[]interface{}{x},
	)
}

func (m *jsiiProxy_Map) AddRetry(props *RetryProps) Map {
	if err := m.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
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
	if err := m.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (m *jsiiProxy_Map) ItemProcessor(processor IChainable, config *ProcessorConfig) Map {
	if err := m.validateItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	var returns Map

	_jsii_.Invoke(
		m,
		"itemProcessor",
		[]interface{}{processor, config},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) Iterator(iterator IChainable) Map {
	if err := m.validateIteratorParameters(iterator); err != nil {
		panic(err)
	}
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
	if err := m.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeDefault",
		[]interface{}{def},
	)
}

func (m *jsiiProxy_Map) MakeNext(next State) {
	if err := m.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"makeNext",
		[]interface{}{next},
	)
}

func (m *jsiiProxy_Map) Next(next IChainable) Chain {
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

func (m *jsiiProxy_Map) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"renderItemProcessor",
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

func (m *jsiiProxy_Map) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Map) WhenBoundToGraph(graph StateGraph) {
	if err := m.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

