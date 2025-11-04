package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct that creates a custom resource that will perform an HTTP API Call.
//
// Example:
//   var stack Stack
//
//
//   awscdkintegtestsalpha.NewHttpApiCall(stack, jsii.String("MyAsssertion"), &HttpCallProps{
//   	Url: jsii.String("https://example-api.com/abc"),
//   })
//
// Experimental.
type HttpApiCall interface {
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

// The jsii proxy struct for HttpApiCall
type jsiiProxy_HttpApiCall struct {
	jsiiProxy_ApiCallBase
}

func (j *jsiiProxy_HttpApiCall) ApiCallResource() awscdk.CustomResource {
	var returns awscdk.CustomResource
	_jsii_.Get(
		j,
		"apiCallResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) ExpectedResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) FlattenResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flattenResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) OutputPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HttpApiCall) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}


// Experimental.
func NewHttpApiCall(scope constructs.Construct, id *string, props *HttpCallProps) HttpApiCall {
	_init_.Initialize()

	if err := validateNewHttpApiCallParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_HttpApiCall{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.HttpApiCall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewHttpApiCall_Override(h HttpApiCall, scope constructs.Construct, id *string, props *HttpCallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.HttpApiCall",
		[]interface{}{scope, id, props},
		h,
	)
}

func (j *jsiiProxy_HttpApiCall)SetExpectedResult(val *string) {
	_jsii_.Set(
		j,
		"expectedResult",
		val,
	)
}

func (j *jsiiProxy_HttpApiCall)SetFlattenResponse(val *string) {
	if err := j.validateSetFlattenResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenResponse",
		val,
	)
}

func (j *jsiiProxy_HttpApiCall)SetOutputPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"outputPaths",
		val,
	)
}

func (j *jsiiProxy_HttpApiCall)SetStateMachineArn(val *string) {
	_jsii_.Set(
		j,
		"stateMachineArn",
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
func HttpApiCall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateHttpApiCall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.HttpApiCall",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) AssertAtPath(path *string, expected ExpectedResult) IApiCall {
	if err := h.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		h,
		"assertAtPath",
		[]interface{}{path, expected},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) Expect(expected ExpectedResult) IApiCall {
	if err := h.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		h,
		"expect",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) GetAtt(attributeName *string) awscdk.Reference {
	if err := h.validateGetAttParameters(attributeName); err != nil {
		panic(err)
	}
	var returns awscdk.Reference

	_jsii_.Invoke(
		h,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) GetAttString(attributeName *string) *string {
	if err := h.validateGetAttStringParameters(attributeName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) Next(next IApiCall) IApiCall {
	if err := h.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		h,
		"next",
		[]interface{}{next},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpApiCall) WaitForAssertions(options *WaiterStateMachineOptions) IApiCall {
	if err := h.validateWaitForAssertionsParameters(options); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		h,
		"waitForAssertions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

