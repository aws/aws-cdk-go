package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
)

// Define a Wait state in the state machine.
//
// A Wait state can be used to delay execution of the state machine for a while.
//
// Example:
//   convertToSeconds := tasks.NewEvaluateExpression(this, jsii.String("Convert to seconds"), &EvaluateExpressionProps{
//   	Expression: jsii.String("$.waitMilliseconds / 1000"),
//   	ResultPath: jsii.String("$.waitSeconds"),
//   })
//
//   createMessage := tasks.NewEvaluateExpression(this, jsii.String("Create message"), &EvaluateExpressionProps{
//   	// Note: this is a string inside a string.
//   	Expression: jsii.String("`Now waiting ${$.waitSeconds} seconds...`"),
//   	Runtime: lambda.Runtime_NODEJS_14_X(),
//   	ResultPath: jsii.String("$.message"),
//   })
//
//   publishMessage := tasks.NewSnsPublish(this, jsii.String("Publish message"), &SnsPublishProps{
//   	Topic: sns.NewTopic(this, jsii.String("cool-topic")),
//   	Message: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
//   	ResultPath: jsii.String("$.sns"),
//   })
//
//   wait := sfn.NewWait(this, jsii.String("Wait"), &WaitProps{
//   	Time: sfn.WaitTime_SecondsPath(jsii.String("$.waitSeconds")),
//   })
//
//   sfn.NewStateMachine(this, jsii.String("StateMachine"), &StateMachineProps{
//   	Definition: convertToSeconds.Next(createMessage).Next(publishMessage).*Next(wait),
//   })
//
// Experimental.
type Wait interface {
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

// The jsii proxy struct for Wait
type jsiiProxy_Wait struct {
	jsiiProxy_State
	jsiiProxy_INextable
}

func (j *jsiiProxy_Wait) Branches() *[]StateGraph {
	var returns *[]StateGraph
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) Comment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) DefaultChoice() State {
	var returns State
	_jsii_.Get(
		j,
		"defaultChoice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) InputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) Iteration() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"iteration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) OutputPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) Parameters() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) ResultPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resultPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) ResultSelector() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"resultSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) StateId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateId",
		&returns,
	)
	return returns
}


// Experimental.
func NewWait(scope constructs.Construct, id *string, props *WaitProps) Wait {
	_init_.Initialize()

	if err := validateNewWaitParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wait{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Wait",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWait_Override(w Wait, scope constructs.Construct, id *string, props *WaitProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.Wait",
		[]interface{}{scope, id, props},
		w,
	)
}

func (j *jsiiProxy_Wait)SetDefaultChoice(val State) {
	_jsii_.Set(
		j,
		"defaultChoice",
		val,
	)
}

func (j *jsiiProxy_Wait)SetIteration(val StateGraph) {
	_jsii_.Set(
		j,
		"iteration",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
// Experimental.
func Wait_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateWait_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Wait",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
// Experimental.
func Wait_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateWait_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Wait",
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
func Wait_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateWait_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Wait",
		"findReachableStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func Wait_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWait_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.Wait",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
// Experimental.
func Wait_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateWait_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"monocdk.aws_stepfunctions.Wait",
		"prefixStates",
		[]interface{}{root, prefix},
	)
}

func (w *jsiiProxy_Wait) AddBranch(branch StateGraph) {
	if err := w.validateAddBranchParameters(branch); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addBranch",
		[]interface{}{branch},
	)
}

func (w *jsiiProxy_Wait) AddChoice(condition Condition, next State) {
	if err := w.validateAddChoiceParameters(condition, next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addChoice",
		[]interface{}{condition, next},
	)
}

func (w *jsiiProxy_Wait) AddIterator(iteration StateGraph) {
	if err := w.validateAddIteratorParameters(iteration); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addIterator",
		[]interface{}{iteration},
	)
}

func (w *jsiiProxy_Wait) AddPrefix(x *string) {
	if err := w.validateAddPrefixParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addPrefix",
		[]interface{}{x},
	)
}

func (w *jsiiProxy_Wait) BindToGraph(graph StateGraph) {
	if err := w.validateBindToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"bindToGraph",
		[]interface{}{graph},
	)
}

func (w *jsiiProxy_Wait) MakeDefault(def State) {
	if err := w.validateMakeDefaultParameters(def); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"makeDefault",
		[]interface{}{def},
	)
}

func (w *jsiiProxy_Wait) MakeNext(next State) {
	if err := w.validateMakeNextParameters(next); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"makeNext",
		[]interface{}{next},
	)
}

func (w *jsiiProxy_Wait) Next(next IChainable) Chain {
	if err := w.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		w,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) OnPrepare() {
	_jsii_.InvokeVoid(
		w,
		"onPrepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wait) OnSynthesize(session constructs.ISynthesisSession) {
	if err := w.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_Wait) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) Prepare() {
	_jsii_.InvokeVoid(
		w,
		"prepare",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wait) RenderBranches() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderBranches",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderChoices() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderChoices",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderInputOutput() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderInputOutput",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderIterator() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderIterator",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderNextEnd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderNextEnd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderResultSelector() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderResultSelector",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) RenderRetryCatch() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderRetryCatch",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) Synthesize(session awscdk.ISynthesisSession) {
	if err := w.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"synthesize",
		[]interface{}{session},
	)
}

func (w *jsiiProxy_Wait) ToStateJson() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"toStateJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wait) WhenBoundToGraph(graph StateGraph) {
	if err := w.validateWhenBoundToGraphParameters(graph); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"whenBoundToGraph",
		[]interface{}{graph},
	)
}

