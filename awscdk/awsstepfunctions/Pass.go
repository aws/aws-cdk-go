package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Pass in the state machine.
//
// A Pass state can be used to transform the current execution's state.
//
// Example:
//   choice := sfn.NewChoice(this, jsii.String("Did it work?"))
//
//   // Add conditions with .when()
//   successState := sfn.NewPass(this, jsii.String("SuccessState"))
//   failureState := sfn.NewPass(this, jsii.String("FailureState"))
//   choice.When(sfn.Condition_StringEquals(jsii.String("$.status"), jsii.String("SUCCESS")), successState)
//   choice.When(sfn.Condition_NumberGreaterThan(jsii.String("$.attempts"), jsii.Number(5)), failureState)
//
//   // Use .otherwise() to indicate what should be done if none of the conditions match
//   tryAgainState := sfn.NewPass(this, jsii.String("TryAgainState"))
//   choice.Otherwise(tryAgainState)
//
type Pass interface {
	State
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
	// Allows the state to validate itself.
	ValidateState() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	WhenBoundToGraph(graph StateGraph)
}

// The jsii proxy struct for Pass
type jsiiProxy_Pass struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_Pass) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Pass) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewPass(scope constructs.Construct, id *string, props *PassProps) Pass {
	_init_.Initialize()

	if err := validateNewPassParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pass{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewPass_Override(p Pass, scope constructs.Construct, id *string, props *PassProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_Pass)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Pass)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_Pass)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Pass)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Pass)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Pass_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validatePass_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Pass_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validatePass_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Pass_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validatePass_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Pass",
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
func Pass_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePass_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Pass_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validatePass_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Pass",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (p *jsiiProxy_Pass) AddBranch(branch StateGraph) {
	if err := p.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addBranch",
		[]interface{}{branch},
	)
}

func (p *jsiiProxy_Pass) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := p.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (p *jsiiProxy_Pass) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := p.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (p *jsiiProxy_Pass) AddIterator(iteration StateGraph) {
	if err := p.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (p *jsiiProxy_Pass) AddPrefix(x *string) {
	if err := p.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPrefix",
		[]interface{}{x},
	)
}

func (p *jsiiProxy_Pass) BindToGraph(graph StateGraph) {
	if err := p.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (p *jsiiProxy_Pass) MakeDefault(def State) {
	if err := p.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"makeDefault",
		[]interface{}{def},
	)
}

func (p *jsiiProxy_Pass) MakeNext(next State) {
	if err := p.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"makeNext",
		[]interface{}{next},
	)
}

func (p *jsiiProxy_Pass) Next(next IChainable) Chain {
	if err := p.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		p,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) WhenBoundToGraph(graph StateGraph) {
	if err := p.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

