package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
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
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
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
type Wait interface {
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

func (j *jsiiProxy_Wait) Node() constructs.Node {
	var returns constructs.Node
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

func (j *jsiiProxy_Wait) Processor() StateGraph {
	var returns StateGraph
	_jsii_.Get(
		j,
		"processor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) ProcessorConfig() *ProcessorConfig {
	var returns *ProcessorConfig
	_jsii_.Get(
		j,
		"processorConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wait) ProcessorMode() ProcessorMode {
	var returns ProcessorMode
	_jsii_.Get(
		j,
		"processorMode",
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

func (j *jsiiProxy_Wait) StateName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateName",
		&returns,
	)
	return returns
}


func NewWait(scope constructs.Construct, id *string, props *WaitProps) Wait {
	_init_.Initialize()

	if err := validateNewWaitParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wait{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Wait",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWait_Override(w Wait, scope constructs.Construct, id *string, props *WaitProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Wait",
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

func (j *jsiiProxy_Wait)SetProcessor(val StateGraph) {
	_jsii_.Set(
		j,
		"processor",
		val,
	)
}

func (j *jsiiProxy_Wait)SetProcessorConfig(val *ProcessorConfig) {
	if err := j.validateSetProcessorConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processorConfig",
		val,
	)
}

func (j *jsiiProxy_Wait)SetProcessorMode(val ProcessorMode) {
	_jsii_.Set(
		j,
		"processorMode",
		val,
	)
}

// Return only the states that allow chaining from an array of states.
func Wait_FilterNextables(states *[]State) *[]INextable {
	_init_.Initialize()

	if err := validateWait_FilterNextablesParameters(states); err != nil {
		panic(err)
	}
	var returns *[]INextable

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Wait",
		"filterNextables",
		[]interface{}{states},
		&returns,
	)

	return returns
}

// Find the set of end states states reachable through transitions from the given start state.
func Wait_FindReachableEndStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateWait_FindReachableEndStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Wait",
		"findReachableEndStates",
		[]interface{}{start, options},
		&returns,
	)

	return returns
}

// Find the set of states reachable through transitions from the given start state.
//
// This does not retrieve states from within sub-graphs, such as states within a Parallel state's branch.
func Wait_FindReachableStates(start State, options *FindStateOptions) *[]State {
	_init_.Initialize()

	if err := validateWait_FindReachableStatesParameters(start, options); err != nil {
		panic(err)
	}
	var returns *[]State

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Wait",
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
func Wait_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWait_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.Wait",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Add a prefix to the stateId of all States found in a construct tree.
func Wait_PrefixStates(root constructs.IConstruct, prefix *string) {
	_init_.Initialize()

	if err := validateWait_PrefixStatesParameters(root, prefix); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"aws-cdk-lib.aws_stepfunctions.Wait",
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

func (w *jsiiProxy_Wait) AddChoice(condition Condition, next State, options *ChoiceTransitionOptions) {
	if err := w.validateAddChoiceParameters(condition, next, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addChoice",
		[]interface{}{condition, next, options},
	)
}

func (w *jsiiProxy_Wait) AddItemProcessor(processor StateGraph, config *ProcessorConfig) {
	if err := w.validateAddItemProcessorParameters(processor, config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addItemProcessor",
		[]interface{}{processor, config},
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

func (w *jsiiProxy_Wait) RenderItemProcessor() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"renderItemProcessor",
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

func (w *jsiiProxy_Wait) ValidateState() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"validateState",
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

