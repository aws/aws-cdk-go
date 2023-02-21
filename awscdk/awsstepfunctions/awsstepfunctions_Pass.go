package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
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
//   choice.when(sfn.condition.stringEquals(jsii.String("$.status"), jsii.String("SUCCESS")), successState)
//   choice.when(sfn.condition.numberGreaterThan(jsii.String("$.attempts"), jsii.Number(5)), failureState)
//
//   // Use .otherwise() to indicate what should be done if none of the conditions match
//   tryAgainState := sfn.NewPass(this, jsii.String("TryAgainState"))
//   choice.otherwise(tryAgainState)
//
// Experimental.
type Pass interface {
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
	// Add a choice branch to this state.
	// Experimental.
	AddChoice(condition Condition, next State)
	// Add a map iterator to this state.
	// Experimental.
	AddIterator(iteration StateGraph)
	// Add a prefix to the stateId of this state.
	// Experimental.
	AddPrefix(x *string)
	// Register this state as part of the given graph.
	//
	// Don't call this. It will be called automatically when you work
	// with states normally.
	// Experimental.
	BindToGraph(graph StateGraph)
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
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
	// Called whenever this state is bound to a graph.
	//
	// Can be overridden by subclasses.
	// Experimental.
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

func (j *jsiiProxy_Pass) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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


// Experimental.
func NewPass(scope constructs.Construct, id *string, props *PassProps) Pass {
	_init_.Initialize()

	if err := validateNewPassParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Pass{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Pass",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewPass_Override(p Pass, scope constructs.Construct, id *string, props *PassProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Pass",
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

// Return only the states that allow chaining from an array of states.
// Experimental.
func Pass_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validatePass_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Pass",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func Pass_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validatePass_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Pass",
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
func Pass_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validatePass_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Pass",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Pass_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePass_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Pass",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func Pass_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validatePass_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions.Pass",
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

func (p *jsiiProxy_Pass) AddChoice(condition Condition, next State) {
	if err := p.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addChoice",
		[]interface{}{condition, next},
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

func (p *jsiiProxy_Pass) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Pass) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Pass) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Pass) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
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

func (p *jsiiProxy_Pass) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
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

func (p *jsiiProxy_Pass) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
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

