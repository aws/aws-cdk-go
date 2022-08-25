package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
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
//   import sfn "github.com/aws/aws-cdk-go/awscdk"
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
//   	choice := sfn.NewChoice(this, jsii.String("Choice")).when(sfn.condition.stringEquals(jsii.String("$.branch"), jsii.String("left")), sfn.NewPass(this, jsii.String("Left Branch"))).when(sfn.condition.stringEquals(jsii.String("$.branch"), jsii.String("right")), sfn.NewPass(this, jsii.String("Right Branch")))
//
//   	// ...
//
//   	this.startState = choice
//   	this.endStates = choice.afterwards().endStates
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
//   	parallel := sfn.NewParallel(this, jsii.String("All jobs")).branch(NewMyJob(this, jsii.String("Quick"), &myJobProps{
//   		jobFlavor: jsii.String("quick"),
//   	}).prefixStates()).branch(NewMyJob(this, jsii.String("Medium"), &myJobProps{
//   		jobFlavor: jsii.String("medium"),
//   	}).prefixStates()).branch(NewMyJob(this, jsii.String("Slow"), &myJobProps{
//   		jobFlavor: jsii.String("slow"),
//   	}).prefixStates())
//
//   	sfn.NewStateMachine(this, jsii.String("MyStateMachine"), &stateMachineProps{
//   		definition: parallel,
//   	})
//   	return this
//   }
//
// Experimental.
type Parallel interface {
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
	AddCatch(handler IChainable, props *CatchProps) Parallel
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
	AddRetry(props *RetryProps) Parallel
	// Overwrites State.bindToGraph. Adds branches to the Parallel state here so that any necessary prefixes are appended first.
	// Experimental.
	BindToGraph(graph StateGraph)
	// Define one or more branches to run in parallel.
	// Experimental.
	Branch(branches ...IChainable) Parallel
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

func (j *jsiiProxy_Parallel) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewParallel(scope constructs.Construct, id *string, props *ParallelProps) Parallel {
	_init_.Initialize()

	j := jsiiProxy_Parallel{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Parallel",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewParallel_Override(p Parallel, scope constructs.Construct, id *string, props *ParallelProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Parallel",
		[]interface{}{scope, id, props},
		p,
	)
}

func (j *jsiiProxy_Parallel) SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Parallel) SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func Parallel_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	var returns *[]INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Parallel",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func Parallel_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Parallel",
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
func Parallel_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Parallel",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Parallel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Parallel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func Parallel_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions.Parallel",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (p *jsiiProxy_Parallel) AddBranch(branch StateGraph) {
	_jsii_.InvokeVoid(
		p,
		"addBranch",
		[]interface{}{branch},
	)
}

func (p *jsiiProxy_Parallel) AddCatch(handler IChainable, props *CatchProps) Parallel {
	var returns Parallel

	_jsii_.Invoke(
		p,
		"addCatch",
		[]interface{}{handler, props},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) AddChoice(condition Condition, next State) {
	_jsii_.InvokeVoid(
		p,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (p *jsiiProxy_Parallel) AddIterator(iteration StateGraph) {
	_jsii_.InvokeVoid(
		p,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (p *jsiiProxy_Parallel) AddPrefix(x *string) {
	_jsii_.InvokeVoid(
		p,
		"addPrefix",
		[]interface{}{x},
	)
}

func (p *jsiiProxy_Parallel) AddRetry(props *RetryProps) Parallel {
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
	_jsii_.InvokeVoid(
		p,
		"makeDefault",
		[]interface{}{def},
	)
}

func (p *jsiiProxy_Parallel) MakeNext(next State) {
	_jsii_.InvokeVoid(
		p,
		"makeNext",
		[]interface{}{next},
	)
}

func (p *jsiiProxy_Parallel) Next(next IChainable) Chain {
	var returns Chain

	_jsii_.Invoke(
		p,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Parallel) OnSynthesize(session constructs.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Parallel) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
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

func (p *jsiiProxy_Parallel) Synthesize(session awscdk.ISynthesisSession) {
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
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

func (p *jsiiProxy_Parallel) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Parallel) WhenBoundToGraph(graph StateGraph) {
	_jsii_.InvokeVoid(
		p,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

