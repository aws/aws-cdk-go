package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS Lambda Invoke function API call.
//
// Use this instead of the generic AwsApiCall in order to
// invoke a lambda function. This will automatically create
// the correct permissions to invoke the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   lambdaInvokeFunction := integ_tests_alpha.NewLambdaInvokeFunction(this, jsii.String("MyLambdaInvokeFunction"), &LambdaInvokeFunctionProps{
//   	FunctionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	InvocationType: integ_tests_alpha.InvocationType_EVENT,
//   	LogType: integ_tests_alpha.LogType_NONE,
//   	Payload: jsii.String("payload"),
//   })
//
// Experimental.
type LambdaInvokeFunction interface {
	AwsApiCall
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
	//   declare const apiCall: AwsApiCall;
	//   apiCall.waiterProvider?.addToRolePolicy({
	//     Effect: 'Allow',
	//     Action: ['s3:GetObject'],
	//     Resource: ['*'],
	//   });
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

// The jsii proxy struct for LambdaInvokeFunction
type jsiiProxy_LambdaInvokeFunction struct {
	jsiiProxy_AwsApiCall
}

func (j *jsiiProxy_LambdaInvokeFunction) ApiCallResource() awscdk.CustomResource {
	var returns awscdk.CustomResource
	_jsii_.Get(
		j,
		"apiCallResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeFunction) ExpectedResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeFunction) FlattenResponse() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flattenResponse",
		&returns,
	)
	return returns
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

func (j *jsiiProxy_LambdaInvokeFunction) OutputPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outputPaths",
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

func (j *jsiiProxy_LambdaInvokeFunction) StateMachineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaInvokeFunction) WaiterProvider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"waiterProvider",
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

func (j *jsiiProxy_LambdaInvokeFunction)SetExpectedResult(val *string) {
	_jsii_.Set(
		j,
		"expectedResult",
		val,
	)
}

func (j *jsiiProxy_LambdaInvokeFunction)SetFlattenResponse(val *string) {
	if err := j.validateSetFlattenResponseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"flattenResponse",
		val,
	)
}

func (j *jsiiProxy_LambdaInvokeFunction)SetOutputPaths(val *[]*string) {
	_jsii_.Set(
		j,
		"outputPaths",
		val,
	)
}

func (j *jsiiProxy_LambdaInvokeFunction)SetStateMachineArn(val *string) {
	_jsii_.Set(
		j,
		"stateMachineArn",
		val,
	)
}

func (j *jsiiProxy_LambdaInvokeFunction)SetWaiterProvider(val AssertionsProvider) {
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

func (l *jsiiProxy_LambdaInvokeFunction) AssertAtPath(path *string, expected ExpectedResult) IApiCall {
	if err := l.validateAssertAtPathParameters(path, expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		l,
		"assertAtPath",
		[]interface{}{path, expected},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeFunction) Expect(expected ExpectedResult) IApiCall {
	if err := l.validateExpectParameters(expected); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		l,
		"expect",
		[]interface{}{expected},
		&returns,
	)

	return returns
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

func (l *jsiiProxy_LambdaInvokeFunction) Next(next IApiCall) IApiCall {
	if err := l.validateNextParameters(next); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		l,
		"next",
		[]interface{}{next},
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

func (l *jsiiProxy_LambdaInvokeFunction) WaitForAssertions(options *WaiterStateMachineOptions) IApiCall {
	if err := l.validateWaitForAssertionsParameters(options); err != nil {
		panic(err)
	}
	var returns IApiCall

	_jsii_.Invoke(
		l,
		"waitForAssertions",
		[]interface{}{options},
		&returns,
	)

	return returns
}

