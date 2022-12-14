// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// A collection of test cases.
//
// Each test case file should contain exactly one
// instance of this class.
//
// Example:
//   var lambdaFunction iFunction
//   var app app
//
//
//   stack := awscdk.NewStack(app, jsii.String("cdk-integ-lambda-bundling"))
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("IntegTest"), &integTestProps{
//   	testCases: []stack{
//   		stack,
//   	},
//   })
//
//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
//   	functionName: lambdaFunction.functionName,
//   })
//   invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"Payload": jsii.String("200"),
//   }))
//
// Experimental.
type IntegTest interface {
	constructs.Construct
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegTest
type jsiiProxy_IntegTest struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IntegTest) Assertions() IDeployAssert {
	var returns IDeployAssert
	_jsii_.Get(
		j,
		"assertions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTest(scope constructs.Construct, id *string, props *IntegTestProps) IntegTest {
	_init_.Initialize()

	if err := validateNewIntegTestParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IntegTest{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTest_Override(i IntegTest, scope constructs.Construct, id *string, props *IntegTestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTest",
		[]interface{}{scope, id, props},
		i,
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
func IntegTest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIntegTest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

