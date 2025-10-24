package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct that creates a CustomResource to assert that two values are equal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actualResult ActualResult
//   var expectedResult ExpectedResult
//
//   equalsAssertion := integ_tests_alpha.NewEqualsAssertion(this, jsii.String("MyEqualsAssertion"), &EqualsAssertionProps{
//   	Actual: actualResult,
//   	Expected: expectedResult,
//
//   	// the properties below are optional
//   	FailDeployment: jsii.Boolean(false),
//   })
//
// Experimental.
type EqualsAssertion interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The result of the assertion.
	// Experimental.
	Result() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EqualsAssertion
type jsiiProxy_EqualsAssertion struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_EqualsAssertion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EqualsAssertion) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewEqualsAssertion(scope constructs.Construct, id *string, props *EqualsAssertionProps) EqualsAssertion {
	_init_.Initialize()

	if err := validateNewEqualsAssertionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EqualsAssertion{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.EqualsAssertion",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEqualsAssertion_Override(e EqualsAssertion, scope constructs.Construct, id *string, props *EqualsAssertionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.EqualsAssertion",
		[]interface{}{scope, id, props},
		e,
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
func EqualsAssertion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEqualsAssertion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.EqualsAssertion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EqualsAssertion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

