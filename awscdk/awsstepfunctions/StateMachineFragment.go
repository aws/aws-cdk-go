package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for reusable state machine fragments.
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
type StateMachineFragment interface {
	constructs.Construct
	IChainable
	// The states to chain onto if this fragment is used.
	EndStates() *[]INextable
	// Descriptive identifier for this chainable.
	Id() *string
	// The tree node.
	Node() constructs.Node
	// The start state of this state machine fragment.
	StartState() State
	// Continue normal execution with the given state.
	Next(next IChainable) Chain
	// Prefix the IDs of all states in this state machine fragment.
	//
	// Use this to avoid multiple copies of the state machine all having the
	// same state IDs.
	PrefixStates(prefix *string) StateMachineFragment
	// Wrap all states in this state machine fragment up into a single state.
	//
	// This can be used to add retry or error handling onto this state
	// machine fragment.
	//
	// Be aware that this changes the result of the inner state machine
	// to be an array with the result of the state machine in it. Adjust
	// your paths accordingly. For example, change 'outputPath' to
	// '$[0]'.
	ToSingleState(options *SingleStateOptions) Parallel
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for StateMachineFragment
type jsiiProxy_StateMachineFragment struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_StateMachineFragment) Node() constructs.Node {
	var returns constructs.Node
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


// Creates a new construct node.
func NewStateMachineFragment_Override(s StateMachineFragment, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.StateMachineFragment",
		[]interface{}{scope, id},
		s,
	)
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
func StateMachineFragment_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStateMachineFragment_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions.StateMachineFragment",
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

