// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Lambda Invoke function API call.
//
// Use this istead of the generic AwsApiCall in order to
// invoke a lambda function. This will automatically create
// the correct permissions to invoke the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   lambdaInvokeFunction := integ_tests_alpha.NewLambdaInvokeFunction(this, jsii.String("MyLambdaInvokeFunction"), &lambdaInvokeFunctionProps{
//   	functionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	invocationType: integ_tests_alpha.invocationType_EVENT,
//   	logType: integ_tests_alpha.logType_NONE,
//   	payload: jsii.String("payload"),
//   })
//
// Experimental.
type LambdaInvokeFunction interface {
	AwsApiCall
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// access the AssertionsProvider.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	// Experimental.
	Provider() AssertionsProvider
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall at the given path.
	//
	// For example the SQS.receiveMessage api response would look
	// like:
	//
	// If you wanted to assert the value of `Body` you could do.
	// Experimental.
	AssertAtPath(path *string, expected ExpectedResult)
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall.
	// Experimental.
	Expect(expected ExpectedResult)
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	// Experimental.
	GetAttString(attributeName *string) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LambdaInvokeFunction
type jsiiProxy_LambdaInvokeFunction struct {
	jsiiProxy_AwsApiCall
}

func (j *jsiiProxy_LambdaInvokeFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeFunction) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaInvokeFunction(scope constructs.Construct, id *string, props *LambdaInvokeFunctionProps) LambdaInvokeFunction {
	_init_.Initialize()

	if err := validateNewLambdaInvokeFunctionParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaInvokeFunction{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.LambdaInvokeFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeFunction_Override(l LambdaInvokeFunction, scope constructs.Construct, id *string, props *LambdaInvokeFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.LambdaInvokeFunction",
		[]interface{}{scope, id, props},
		l,
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
func LambdaInvokeFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaInvokeFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.LambdaInvokeFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeFunction) AssertAtPath(path *string, expected ExpectedResult) {
	if err := l.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) Expect(expected ExpectedResult) {
	if err := l.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"expect",
		[]interface{}{expected},
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) GetAtt(attributeName *string) awscdk.Reference {
	if err := l.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		l,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeFunction) GetAttString(attributeName *string) *string {
	if err := l.validateGetAttStringParameters(attributeName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

