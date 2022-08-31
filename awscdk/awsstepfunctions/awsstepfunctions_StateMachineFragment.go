package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Base class for reusable state machine fragments.
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
type StateMachineFragment interface {
	awscdk.Construct
	IChainable
	// The states to chain onto if this fragment is used.
	// Experimental.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	// Experimental.
	Id() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The start state of this state machine fragment.
	// Experimental.
	StartState() State
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
	// Prefix the IDs of all states in this state machine fragment.
	//
	// Use this to avoid multiple copies of the state machine all having the
	// same state IDs.
	// Experimental.
	PrefixStates(prefix *string) StateMachineFragment
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
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Wrap all states in this state machine fragment up into a single state.
	//
	// This can be used to add retry or error handling onto this state
	// machine fragment.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
	// Experimental.
	ToSingleState(options *SingleStateOptions) Parallel
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
}

// The jsii proxy struct for StateMachineFragment
type jsiiProxy_StateMachineFragment struct {
	internal.Type__awscdkConstruct
	jsiiProxy_IChainable
}

func (j *jsiiProxy_StateMachineFragment) EndStates() *[]INextable {
	var returns *[]INextable
	_jsii_.Get(
		j,
		"endStates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineFragment) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineFragment) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StateMachineFragment) StartState() State {
	var returns State
	_jsii_.Get(
		j,
		"startState",
		&returns,
	)
	return returns
}


// Experimental.
func NewStateMachineFragment_Override(s StateMachineFragment, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions.StateMachineFragment",
		[]interface{}{scope, id},
		s,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func StateMachineFragment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStateMachineFragment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions.StateMachineFragment",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) Next(next IChainable) Chain {
	if err := s.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns Chain

	_jsii_.Invoke(
		s,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) OnPrepare() {
	_jsii_.InvokeVoid(
		s,
		"onPrepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StateMachineFragment) OnSynthesize(session constructs.ISynthesisSession) {
	if err := s.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StateMachineFragment) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) PrefixStates(prefix *string) StateMachineFragment {
	var returns StateMachineFragment

	_jsii_.Invoke(
		s,
		"prefixStates",
		[]interface{}{prefix},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) Prepare() {
	_jsii_.InvokeVoid(
		s,
		"prepare",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StateMachineFragment) Synthesize(session awscdk.ISynthesisSession) {
	if err := s.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"synthesize",
		[]interface{}{session},
	)
}

func (s *jsiiProxy_StateMachineFragment) ToSingleState(options *SingleStateOptions) Parallel {
	if err := s.validateToSingleStateParameters(options); err != nil {
		panic(err)
	}
	var returns Parallel

	_jsii_.Invoke(
		s,
		"toSingleState",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StateMachineFragment) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

