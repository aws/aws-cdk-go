package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct that creates a custom resource that will perform a query using the AWS SDK.
//
// Example:
//   var stack stack
//
//
//   awscdkintegtestsalpha.NewAwsApiCall(stack, jsii.String("MyAssertion"), &AwsApiCallProps{
//   	Service: jsii.String("SQS"),
//   	Api: jsii.String("receiveMessage"),
//   	Parameters: map[string]*string{
//   		"QueueUrl": jsii.String("url"),
//   	},
//   })
//
// Experimental.
type AwsApiCall interface {
	ApiCallBase
	// Experimental.
	ApiCallResource() awscdk.CustomResource
	// Experimental.
	ExpectedResult() *string
	// Experimental.
	SetExpectedResult(val *string)
	// Experimental.
	FlattenResponse() *string
	// Experimental.
	SetFlattenResponse(val *string)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputPaths() *[]*string
	// Experimental.
	SetOutputPaths(val *[]*string)
	// access the AssertionsProvider.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	// Experimental.
	Provider() AssertionsProvider
	// Experimental.
	StateMachineArn() *string
	// Experimental.
	SetStateMachineArn(val *string)
	// access the AssertionsProvider for the waiter state machine.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	//
	// Example:
	//   var apiCall awsApiCall
	//
	//   apiCall.WaiterProvider.AddToRolePolicy(map[string]interface{}{
	//   	"Effect": jsii.String("Allow"),
	//   	"Action": []*string{
	//   		jsii.String("s3:GetObject"),
	//   	},
	//   	"Resource": []*string{
	//   		jsii.String("*"),
	//   	},
	//   })
	//
	// Experimental.
	WaiterProvider() AssertionsProvider
	// Experimental.
	SetWaiterProvider(val AssertionsProvider)
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall at the given path.
	//
	// Providing a path will filter the output of the initial API call.
	//
	// For example the SQS.receiveMessage api response would look
	// like:
	//
	// If you wanted to assert the value of `Body` you could do.
	// Experimental.
	AssertAtPath(path *string, expected ExpectedResult) IApiCall
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall.
	// Experimental.
	Expect(expected ExpectedResult) IApiCall
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
	// Allows you to chain IApiCalls. This adds an explicit dependency betweent the two resources.
	//
	// Returns the IApiCall provided as `next`.
	// Experimental.
	Next(next IApiCall) IApiCall
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Wait for the IApiCall to return the expected response.
	//
	// If no expected response is specified then it will wait for
	// the IApiCall to return a success.
	// Experimental.
	WaitForAssertions(options *WaiterStateMachineOptions) IApiCall
}

// The jsii proxy struct for AwsApiCall
type jsiiProxy_AwsApiCall struct {
	jsiiProxy_ApiCallBase
}

func (j *jsiiProxy_AwsApiCall) ApiCallResource() awscdk.CustomResource {
	var returns awscdk.CustomResource
	_jsii_.Get(
		j,
		"apiCallResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiCall) ExpectedResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiCall) FlattenResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flattenResponse",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_AwsApiCall) OutputPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputPaths",
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

func (j *jsiiProxy_AwsApiCall) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsApiCall) WaiterProvider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"waiterProvider",
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

func (j *jsiiProxy_AwsApiCall)SetExpectedResult(val *string) {
	_jsii_.Set(
		j,
		"expectedResult",
		val,
	)
}

func (j *jsiiProxy_AwsApiCall)SetFlattenResponse(val *string) {
	if err := j.validateSetFlattenResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenResponse",
		val,
	)
}

func (j *jsiiProxy_AwsApiCall)SetOutputPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"outputPaths",
		val,
	)
}

func (j *jsiiProxy_AwsApiCall)SetStateMachineArn(val *string) {
	_jsii_.Set(
		j,
		"stateMachineArn",
		val,
	)
}

func (j *jsiiProxy_AwsApiCall)SetWaiterProvider(val AssertionsProvider) {
	_jsii_.Set(
		j,
		"waiterProvider",
		val,
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

func (a *jsiiProxy_AwsApiCall) AssertAtPath(path *string, expected ExpectedResult) IApiCall {
	if err := a.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		a,
		"assertAtPath",
		[]interface{}{path, expected},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiCall) Expect(expected ExpectedResult) IApiCall {
	if err := a.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		a,
		"expect",
		[]interface{}{expected},
		&returns,
	)

	return returns
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

func (a *jsiiProxy_AwsApiCall) Next(next IApiCall) IApiCall {
	if err := a.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		a,
		"next",
		[]interface{}{next},
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

func (a *jsiiProxy_AwsApiCall) WaitForAssertions(options *WaiterStateMachineOptions) IApiCall {
	if err := a.validateWaitForAssertionsParameters(options); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		a,
		"waitForAssertions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

