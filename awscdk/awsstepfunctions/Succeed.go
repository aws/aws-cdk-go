package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Succeed state in the state machine.
//
// Reaching a Succeed state terminates the state execution in success.
//
// Example:
//   success := sfn.NewSucceed(this, jsii.String("We did it!"))
//
type Succeed interface {
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

// The jsii proxy struct for Succeed
type jsiiProxy_Succeed struct {
	jsiiProxy_State
}

func (j *jsiiProxy_Succeed) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Succeed) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewSucceed(scope constructs.Construct, id *string, props *SucceedProps) Succeed {
	_init_.Initialize()

	if err := validateNewSucceedParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Succeed{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSucceed_Override(s Succeed, scope constructs.Construct, id *string, props *SucceedProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		[]interface{}{scope, id, props},
		s,
	)
}

func (j *jsiiProxy_Succeed)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Succeed)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_Succeed)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Succeed)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Succeed)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Succeed_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateSucceed_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Succeed_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateSucceed_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Succeed_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateSucceed_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
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
func Succeed_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSucceed_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Succeed_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateSucceed_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Succeed",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (s *jsiiProxy_Succeed) AddBranch(branch StateGraph) {
	if err := s.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addBranch",
		[]interface{}{branch},
	)
}

func (s *jsiiProxy_Succeed) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := s.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (s *jsiiProxy_Succeed) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := s.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (s *jsiiProxy_Succeed) AddIterator(iteration StateGraph) {
	if err := s.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (s *jsiiProxy_Succeed) AddPrefix(x *string) {
	if err := s.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addPrefix",
		[]interface{}{x},
	)
}

func (s *jsiiProxy_Succeed) BindToGraph(graph StateGraph) {
	if err := s.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (s *jsiiProxy_Succeed) MakeDefault(def State) {
	if err := s.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeDefault",
		[]interface{}{def},
	)
}

func (s *jsiiProxy_Succeed) MakeNext(next State) {
	if err := s.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"makeNext",
		[]interface{}{next},
	)
}

func (s *jsiiProxy_Succeed) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Succeed) WhenBoundToGraph(graph StateGraph) {
	if err := s.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

