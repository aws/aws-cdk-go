package triggers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/triggers/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Triggers an AWS Lambda function during deployment.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   func := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Handler: jsii.String("index.handler"),
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Code: lambda.Code_FromInline(jsii.String("foo")),
//   })
//
//   triggers.NewTrigger(this, jsii.String("MyTrigger"), &TriggerProps{
//   	Handler: func,
//   	Timeout: awscdk.Duration_Minutes(jsii.Number(10)),
//   	InvocationType: triggers.InvocationType_EVENT,
//   })
//
type Trigger interface {
	constructs.Construct
	ITrigger
	// The tree node.
	Node() constructs.Node
	// Adds trigger dependencies.
	//
	// Execute this trigger only after these construct
	// scopes have been provisioned.
	ExecuteAfter(scopes ...constructs.Construct)
	// Adds this trigger as a dependency on other constructs.
	//
	// This means that this
	// trigger will get executed *before* the given construct(s).
	ExecuteBefore(scopes ...constructs.Construct)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Trigger
type jsiiProxy_Trigger struct {
	internal.Type__constructsConstruct
	jsiiProxy_ITrigger
}

func (j *jsiiProxy_Trigger) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewTrigger(scope constructs.Construct, id *string, props *TriggerProps) Trigger {
	_init_.Initialize()

	if err := validateNewTriggerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Trigger{}

	_jsii_.Create(
		"aws-cdk-lib.triggers.Trigger",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewTrigger_Override(t Trigger, scope constructs.Construct, id *string, props *TriggerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.triggers.Trigger",
		[]interface{}{scope, id, props},
		t,
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
func Trigger_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTrigger_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.triggers.Trigger",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_Trigger) ExecuteAfter(scopes ...constructs.Construct) {
	args := []interface{}{}
	for _, a := range scopes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"executeAfter",
		args,
	)
}

func (t *jsiiProxy_Trigger) ExecuteBefore(scopes ...constructs.Construct) {
	args := []interface{}{}
	for _, a := range scopes {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		t,
		"executeBefore",
		args,
	)
}

func (t *jsiiProxy_Trigger) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

