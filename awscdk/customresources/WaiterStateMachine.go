package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/customresources/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A very simple StateMachine construct highly customized to the provider framework.
//
// We previously used `CfnResource` instead of `CfnStateMachine` to avoid depending
// on `aws-stepfunctions` module, but now it is okay.
//
// The state machine continuously calls the isCompleteHandler, until it succeeds or times out.
// The handler is called `maxAttempts` times with an `interval` duration and a `backoffRate` rate.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var function_ function
//   var logGroup logGroup
//
//   waiterStateMachine := awscdk.Custom_resources.NewWaiterStateMachine(this, jsii.String("MyWaiterStateMachine"), &WaiterStateMachineProps{
//   	BackoffRate: jsii.Number(123),
//   	Interval: cdk.Duration_Minutes(jsii.Number(30)),
//   	IsCompleteHandler: function_,
//   	MaxAttempts: jsii.Number(123),
//   	TimeoutHandler: function_,
//
//   	// the properties below are optional
//   	DisableLogging: jsii.Boolean(false),
//   	LogOptions: &LogOptions{
//   		Destination: logGroup,
//   		IncludeExecutionData: jsii.Boolean(false),
//   		Level: awscdk.Aws_stepfunctions.LogLevel_OFF,
//   	},
//   })
//
type WaiterStateMachine interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The ARN of the state machine.
	StateMachineArn() *string
	// Grant the given identity permissions on StartExecution of the state machine.
	GrantStartExecution(identity awsiam.IGrantable) awsiam.Grant
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for WaiterStateMachine
type jsiiProxy_WaiterStateMachine struct {
	internal.Type__constructsConstruct
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

func (j *jsiiProxy_WaiterStateMachine) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}


func NewWaiterStateMachine(scope constructs.Construct, id *string, props *WaiterStateMachineProps) WaiterStateMachine {
	_init_.Initialize()

	if err := validateNewWaiterStateMachineParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WaiterStateMachine{}

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.WaiterStateMachine",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewWaiterStateMachine_Override(w WaiterStateMachine, scope constructs.Construct, id *string, props *WaiterStateMachineProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.custom_resources.WaiterStateMachine",
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
func WaiterStateMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWaiterStateMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.custom_resources.WaiterStateMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WaiterStateMachine) GrantStartExecution(identity awsiam.IGrantable) awsiam.Grant {
	if err := w.validateGrantStartExecutionParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		w,
		"grantStartExecution",
		[]interface{}{identity},
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

