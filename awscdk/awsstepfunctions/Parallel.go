package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Define a Parallel state in the state machine.
//
// A Parallel state can be used to run one or more state machines at the same
// time.
//
// The Result of a Parallel state is an array of the results of its substatemachines.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/constructs-go/constructs"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   type myJobProps struct {
//   	jobFlavor *string
//   }
//
//   type myJob struct {
//   	stateMachineFragment
//   	startState state
//   	endStates []iNextable
//   }
//
//   func newMyJob(parent construct, id *string, props myJobProps) *myJob {
//   	this := &myJob{}
//   	sfn.NewStateMachineFragment_Override(this, parent, id)
//
//   	choice := sfn.NewChoice(this, jsii.String("Choice")).When(sfn.Condition_StringEquals(jsii.String("$.branch"), jsii.String("left")), sfn.NewPass(this, jsii.String("Left Branch"))).When(sfn.Condition_StringEquals(jsii.String("$.branch"), jsii.String("right")), sfn.NewPass(this, jsii.String("Right Branch")))
//
//   	// ...
//
//   	this.startState = choice
//   	this.endStates = choice.Afterwards().EndStates
//   	return this
//   }
//
//   type myStack struct {
//   	stack
//   }
//
//   func newMyStack(scope construct, id *string) *myStack {
//   	this := &myStack{}
//   	newStack_Override(this, scope, id)
//   	// Do 3 different variants of MyJob in parallel
//   	parallel := sfn.NewParallel(this, jsii.String("All jobs")).Branch(NewMyJob(this, jsii.String("Quick"), &myJobProps{
//   		jobFlavor: jsii.String("quick"),
//   	}).PrefixStates()).Branch(NewMyJob(this, jsii.String("Medium"), &myJobProps{
//   		jobFlavor: jsii.String("medium"),
//   	}).PrefixStates()).Branch(NewMyJob(this, jsii.String("Slow"), &myJobProps{
//   		jobFlavor: jsii.String("slow"),
//   	}).PrefixStates())
//
//   	sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &StateMachineProps{
//   		DefinitionBody: sfn.DefinitionBody_FromChainable(parallel),
//   	})
//   	return this
//   }
//
type Parallel interface {
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
	// Add a recovery handler for this state.
	//
	// When a particular error occurs, execution will continue at the error
	// handler instead of failing the state machine execution.
	AddCatch(handler IChainable, props *CatchProps) Parallel
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
	AddRetry(props *RetryProps) Parallel
	// Overwrites State.bindToGraph. Adds branches to the Parallel state here so that any necessary prefixes are appended first.
	BindToGraph(graph StateGraph)
	// Define one or more branches to run in parallel.
	Branch(branches ...IChainable) Parallel
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

// The jsii proxy struct for Parallel
type jsiiProxy_Parallel struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_Parallel) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Parallel) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewParallel(scope constructs.Construct, id *string, props *ParallelProps) Parallel {
	_init_.Initialize()

	if err := validateNewParallelParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Parallel{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewParallel_Override(p Parallel, scope constructs.Construct, id *string, props *ParallelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_Parallel)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Parallel)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

func (j *jsiiProxy_Parallel)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Parallel)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Parallel)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Parallel_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateParallel_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Parallel_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateParallel_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Parallel_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateParallel_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
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
func Parallel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateParallel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Parallel_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateParallel_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Parallel",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (p *jsiiProxy_Parallel) AddBranch(branch StateGraph) {
	if err := p.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addBranch",
		[]interface{}{branch},
	)
}

func (p *jsiiProxy_Parallel) AddCatch(handler IChainable, props *CatchProps) Parallel {
	if err := p.validateAddCatchParameters(handler, props); err != nil {
		panic(err)
	}
	var returns Parallel

	_jsii_.Invoke(
		p,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := p.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (p *jsiiProxy_Parallel) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := p.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addItemProcessor",
		[]interface{}{processor, config},
	)
}

func (p *jsiiProxy_Parallel) AddIterator(iteration StateGraph) {
	if err := p.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (p *jsiiProxy_Parallel) AddPrefix(x *string) {
	if err := p.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addPrefix",
		[]interface{}{x},
	)
}

func (p *jsiiProxy_Parallel) AddRetry(props *RetryProps) Parallel {
	if err := p.validateAddRetryParameters(props); err != nil {
		panic(err)
	}
	var returns Parallel

	_jsii_.Invoke(
		p,
		"addRetry",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) BindToGraph(graph StateGraph) {
	if err := p.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (p *jsiiProxy_Parallel) Branch(branches ...IChainable) Parallel {
	args := []interface{}{}
	for _, a := range branches {
		args = append(args, a)
	}

	var returns Parallel

	_jsii_.Invoke(
		p,
		"branch",
		args,
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) MakeDefault(def State) {
	if err := p.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"makeDefault",
		[]interface{}{def},
	)
}

func (p *jsiiProxy_Parallel) MakeNext(next State) {
	if err := p.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"makeNext",
		[]interface{}{next},
	)
}

func (p *jsiiProxy_Parallel) Next(next IChainable) Chain {
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

func (p *jsiiProxy_Parallel) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderItemProcessor",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validateState",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) WhenBoundToGraph(graph StateGraph) {
	if err := p.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

