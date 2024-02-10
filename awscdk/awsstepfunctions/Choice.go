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
	// If the given condition matches, continue execution with the given state.
	When(condition Condition, next IChainable, options *ChoiceTransitionOptions) Choice
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

func (j *jsiiProxy_Choice) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Choice) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
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

func (j *jsiiProxy_Choice) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewChoice(scope constructs.Construct, id *string, props *ChoiceProps) Choice {
	_init_.Initialize()

	if err := validateNewChoiceParameters(scope, id, props); err != nil {
		panic(err)
	}
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

func (j *jsiiProxy_Choice)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Choice)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_Choice)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Choice)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Choice)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Choice_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateChoice_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
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

	if err := validateChoice_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
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

	if err := validateChoice_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
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

	if err := validateChoice_IsConstructParameters(x); err != nil {
		panic(err)
	}
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

	if err := validateChoice_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Choice",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (c *jsiiProxy_Choice) AddBranch(branch StateGraph) {
	if err := c.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addBranch",
		[]interface{}{branch},
	)
}

func (c *jsiiProxy_Choice) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := c.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (c *jsiiProxy_Choice) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := c.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (c *jsiiProxy_Choice) AddIterator(iteration StateGraph) {
	if err := c.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (c *jsiiProxy_Choice) AddPrefix(x *string) {
	if err := c.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addPrefix",
		[]interface{}{x},
	)
}

func (c *jsiiProxy_Choice) Afterwards(options *AfterwardsOptions) Chain {
	if err := c.validateAfterwardsParameters(options); err != nil {
		panic(err)
	}
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
	if err := c.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (c *jsiiProxy_Choice) MakeDefault(def State) {
	if err := c.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeDefault",
		[]interface{}{def},
	)
}

func (c *jsiiProxy_Choice) MakeNext(next State) {
	if err := c.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"makeNext",
		[]interface{}{next},
	)
}

func (c *jsiiProxy_Choice) Otherwise(def IChainable) Choice {
	if err := c.validateOtherwiseParameters(def); err != nil {
		panic(err)
	}
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

func (c *jsiiProxy_Choice) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"renderItemProcessor",
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

func (c *jsiiProxy_Choice) When(condition Condition, next IChainable, options *ChoiceTransitionOptions) Choice {
	if err := c.validateWhenParameters(condition, next, options); err != nil {
		panic(err)
	}
	var returns Choice

	_jsii_.Invoke(
		c,
		"when",
		[]interface{}{condition, next, options},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_Choice) WhenBoundToGraph(graph StateGraph) {
	if err := c.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

