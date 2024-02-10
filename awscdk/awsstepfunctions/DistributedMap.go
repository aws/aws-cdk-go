package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Distributed Mode Map state in the state machine.
//
// A `Map` state can be used to run a set of steps for each element of an input array.
// A Map state will execute the same steps for multiple entries of an array in the state input.
//
// While the Parallel state executes multiple branches of steps using the same input, a Map state
// will execute the same steps for multiple entries of an array in the state input.
//
// A `Map` state in `Distributed` mode will execute a child workflow for each iteration of the Map state.
// This serves to increase concurrency and allows for larger workloads to be run in a single state machine.
//
// Example:
//   distributedMap := sfn.NewDistributedMap(this, jsii.String("Distributed Map State"), &DistributedMapProps{
//   	MaxConcurrency: jsii.Number(1),
//   	ItemsPath: sfn.JsonPath_StringAt(jsii.String("$.inputForMap")),
//   })
//   distributedMap.ItemProcessor(sfn.NewPass(this, jsii.String("Pass State")))
//
// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-asl-use-map-state-distributed.html
//
type DistributedMap interface {
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
	AddCatch(handler IChainable, props *CatchProps) DistributedMap
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
	AddRetry(props *RetryProps) DistributedMap
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	BindToGraph(graph StateGraph)
	// Define item processor in a Distributed Map.
	//
	// A Distributed Map must have a non-empty item processor.
	ItemProcessor(processor IChainable, config *ProcessorConfig) DistributedMap
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

// The jsii proxy struct for DistributedMap
type jsiiProxy_DistributedMap struct {
	jsiiProxy_MapBase
	jsiiProxy_INextable
}

func (j *jsiiProxy_DistributedMap) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ItemSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"itemSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ItemsPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"itemsPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DistributedMap) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewDistributedMap(scope constructs.Construct, id *string, props *DistributedMapProps) DistributedMap {
	_init_.Initialize()

	if err := validateNewDistributedMapParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DistributedMap{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewDistributedMap_Override(d DistributedMap, scope constructs.Construct, id *string, props *DistributedMapProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		[]interface{}{scope, id, props},
		d,
	)
}

func (j *jsiiProxy_DistributedMap)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_DistributedMap)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_DistributedMap)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_DistributedMap)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_DistributedMap)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func DistributedMap_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateDistributedMap_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func DistributedMap_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateDistributedMap_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func DistributedMap_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateDistributedMap_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
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
func DistributedMap_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedMap_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a DistributedMap.
func DistributedMap_IsDistributedMap(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDistributedMap_IsDistributedMapParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		"isDistributedMap",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func DistributedMap_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateDistributedMap_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.DistributedMap",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (d *jsiiProxy_DistributedMap) AddBranch(branch StateGraph) {
	if err := d.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addBranch",
		[]interface{}{branch},
	)
}

func (d *jsiiProxy_DistributedMap) AddCatch(handler IChainable, props *CatchProps) DistributedMap {
	if err := d.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns DistributedMap

	_jsii_.Invoke(
		d,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := d.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (d *jsiiProxy_DistributedMap) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := d.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (d *jsiiProxy_DistributedMap) AddIterator(iteration StateGraph) {
	if err := d.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (d *jsiiProxy_DistributedMap) AddPrefix(x *string) {
	if err := d.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addPrefix",
		[]interface{}{x},
	)
}

func (d *jsiiProxy_DistributedMap) AddRetry(props *RetryProps) DistributedMap {
	if err := d.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns DistributedMap

	_jsii_.Invoke(
		d,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) BindToGraph(graph StateGraph) {
	if err := d.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (d *jsiiProxy_DistributedMap) ItemProcessor(processor IChainable, config *ProcessorConfig) DistributedMap {
	if err := d.validateItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	var returns DistributedMap

	_jsii_.Invoke(
		d,
		"itemProcessor",
		[]interface{}{processor, config},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) MakeDefault(def State) {
	if err := d.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeDefault",
		[]interface{}{def},
	)
}

func (d *jsiiProxy_DistributedMap) MakeNext(next State) {
	if err := d.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"makeNext",
		[]interface{}{next},
	)
}

func (d *jsiiProxy_DistributedMap) Next(next IChainable) Chain {
	if err := d.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		d,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DistributedMap) WhenBoundToGraph(graph StateGraph) {
	if err := d.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

