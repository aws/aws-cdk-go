// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A very simple StateMachine construct highly customized to the provider framework.
//
// This is so that this package does not need to depend on aws-stepfunctions module.
//
// The state machine continuously calls the isCompleteHandler, until it succeeds or times out.
// The handler is called `maxAttempts` times with an `interval` duration and a `backoffRate` rate.
//
// For example with:
// - maxAttempts = 360 (30 minutes)
// - interval = 5
// - backoffRate = 1 (no backoff)
//
// it will make the API Call every 5 seconds and fail after 360 failures.
//
// If the backoff rate is changed to 2 (for example), it will
// - make the first call
// - wait 5 seconds
// - make the second call
// - wait 15 seconds
// - etc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   waiterStateMachine := integ_tests_alpha.NewWaiterStateMachine(this, jsii.String("MyWaiterStateMachine"), &WaiterStateMachineProps{
//   	BackoffRate: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	TotalTimeout: cdk.Duration_*Minutes(jsii.Number(30)),
//   })
//
// Experimental.
type WaiterStateMachine interface {
	constructs.Construct
	// The AssertionsProvide that handles async requests.
	// Experimental.
	IsCompleteProvider() AssertionsProvider
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The IAM Role ARN of the role used by the state machine.
	// Experimental.
	RoleArn() *string
	// The ARN of the statemachine.
	// Experimental.
	StateMachineArn() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WaiterStateMachine
type jsiiProxy_WaiterStateMachine struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_WaiterStateMachine) IsCompleteProvider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"isCompleteProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaiterStateMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaiterStateMachine) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WaiterStateMachine) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewWaiterStateMachine(scope constructs.Construct, id *string, props *WaiterStateMachineProps) WaiterStateMachine {
	_init_.Initialize()

	if err := validateNewWaiterStateMachineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WaiterStateMachine{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.WaiterStateMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWaiterStateMachine_Override(w WaiterStateMachine, scope constructs.Construct, id *string, props *WaiterStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.WaiterStateMachine",
		[]interface{}{scope, id, props},
		w,
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
// Experimental.
func WaiterStateMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaiterStateMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.WaiterStateMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaiterStateMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

