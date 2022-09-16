// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct that creates a custom resource that will perform a query using the AWS SDK.
//
// Example:
//   var myAppStack stack
//
//
//   awscdkintegtestsalpha.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &awsApiCallProps{
//   	service: jsii.String("S3"),
//   	api: jsii.String("getObject"),
//   })
//
// Experimental.
type AwsApiCall interface {
	constructs.Construct
	IAwsApiCall
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

// The jsii proxy struct for AwsApiCall
type jsiiProxy_AwsApiCall struct {
	internal.Type__constructsConstruct
	jsiiProxy_IAwsApiCall
}

func (j *jsiiProxy_AwsApiCall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiCall) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsApiCall(scope constructs.Construct, id *string, props *AwsApiCallProps) AwsApiCall {
	_init_.Initialize()

	if err := validateNewAwsApiCallParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsApiCall{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AwsApiCall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApiCall_Override(a AwsApiCall, scope constructs.Construct, id *string, props *AwsApiCallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AwsApiCall",
		[]interface{}{scope, id, props},
		a,
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
func AwsApiCall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiCall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.AwsApiCall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiCall) AssertAtPath(path *string, expected ExpectedResult) {
	if err := a.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (a *jsiiProxy_AwsApiCall) Expect(expected ExpectedResult) {
	if err := a.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"expect",
		[]interface{}{expected},
	)
}

func (a *jsiiProxy_AwsApiCall) GetAtt(attributeName *string) awscdk.Reference {
	if err := a.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		a,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiCall) GetAttString(attributeName *string) *string {
	if err := a.validateGetAttStringParameters(attributeName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiCall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

