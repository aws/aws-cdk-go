// CDK Integration Testing Constructs
package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cloudassemblyschema"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents the "actual" results to compare.
//
// Example:
//   var myCustomResource customResource
//   var stack stack
//   var app app
//
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
//   	testCases: []*stack{
//   		stack,
//   	},
//   })
//   integ.assertions.expect(jsii.String("CustomAssertion"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"foo": jsii.String("bar"),
//   }), awscdkintegtestsalpha.ActualResult.fromCustomResource(myCustomResource, jsii.String("data")))
//
// Experimental.
type ActualResult interface {
	// The actual results as a string.
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
}

// The jsii proxy struct for ActualResult
type jsiiProxy_ActualResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ActualResult) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewActualResult_Override(a ActualResult) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		nil, // no parameters
		a,
	)
}

func (j *jsiiProxy_ActualResult) SetResult(val *string) {
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

// Get the actual results from a AwsApiCall.
// Experimental.
func ActualResult_FromAwsApiCall(query IAwsApiCall, attribute *string) ActualResult {
	_init_.Initialize()

	var returns ActualResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		"fromAwsApiCall",
		[]interface{}{query, attribute},
		&returns,
	)

	return returns
}

// Get the actual results from a CustomResource.
// Experimental.
func ActualResult_FromCustomResource(customResource awscdk.CustomResource, attribute *string) ActualResult {
	_init_.Initialize()

	var returns ActualResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ActualResult",
		"fromCustomResource",
		[]interface{}{customResource, attribute},
		&returns,
	)

	return returns
}

// A request to make an assertion that the actual value matches the expected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actual interface{}
//   var expected interface{}
//
//   assertionRequest := &assertionRequest{
//   	actual: actual,
//   	expected: expected,
//
//   	// the properties below are optional
//   	failDeployment: jsii.Boolean(false),
//   }
//
// Experimental.
type AssertionRequest struct {
	// The actual value received.
	// Experimental.
	Actual interface{} `field:"required" json:"actual" yaml:"actual"`
	// The expected value to assert.
	// Experimental.
	Expected interface{} `field:"required" json:"expected" yaml:"expected"`
	// Set this to true if a failed assertion should result in a CloudFormation deployment failure.
	//
	// This is only necessary if assertions are being
	// executed outside of `integ-runner`.
	// Experimental.
	FailDeployment *bool `field:"optional" json:"failDeployment" yaml:"failDeployment"`
}

// The result of an Assertion wrapping the actual result data in another struct.
//
// Needed to access the whole message via getAtt() on the custom resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   assertionResult := &assertionResult{
//   	data: jsii.String("data"),
//
//   	// the properties below are optional
//   	failed: jsii.Boolean(false),
//   }
//
// Experimental.
type AssertionResult struct {
	// The result of an assertion.
	// Experimental.
	Data *string `field:"required" json:"data" yaml:"data"`
	// Whether or not the assertion failed.
	// Experimental.
	Failed *bool `field:"optional" json:"failed" yaml:"failed"`
}

// The result of an assertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   assertionResultData := &assertionResultData{
//   	status: integ_tests_alpha.status_PASS,
//
//   	// the properties below are optional
//   	message: jsii.String("message"),
//   }
//
// Experimental.
type AssertionResultData struct {
	// The status of the assertion, i.e. pass or fail.
	// Experimental.
	Status Status `field:"required" json:"status" yaml:"status"`
	// Any message returned with the assertion result typically this will be the diff if there is any.
	// Experimental.
	Message *string `field:"optional" json:"message" yaml:"message"`
}

// The type of assertion to perform.
// Experimental.
type AssertionType string

const (
	// Assert that two values are equal.
	// Experimental.
	AssertionType_EQUALS AssertionType = "EQUALS"
	// The keys and their values must be present in the target but the target can be a superset.
	// Experimental.
	AssertionType_OBJECT_LIKE AssertionType = "OBJECT_LIKE"
	// Matches the specified pattern with the array The set of elements must be in the same order as would be found.
	// Experimental.
	AssertionType_ARRAY_WITH AssertionType = "ARRAY_WITH"
)

// Represents an assertions provider.
//
// The creates a singletone
// Lambda Function that will create a single function per stack
// that serves as the custom resource provider for the various
// assertion providers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   assertionsProvider := integ_tests_alpha.NewAssertionsProvider(this, jsii.String("MyAssertionsProvider"))
//
// Experimental.
type AssertionsProvider interface {
	constructs.Construct
	// A reference to the provider Lambda Function execution Role ARN.
	// Experimental.
	HandlerRoleArn() awscdk.Reference
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The ARN of the lambda function which can be used as a serviceToken to a CustomResource.
	// Experimental.
	ServiceToken() *string
	// Create a policy statement from a specific api call.
	// Experimental.
	AddPolicyStatementFromSdkCall(service *string, api *string, resources *[]*string)
	// Add an IAM policy statement to the inline policy of the lambdas function's role.
	//
	// **Please note**: this is a direct IAM JSON policy blob, *not* a `iam.PolicyStatement`
	// object like you will see in the rest of the CDK.
	//
	// Example:
	//   var provider assertionsProvider
	//
	//   provider.addToRolePolicy(map[string]*string{
	//   	"Effect": jsii.String("Allow"),
	//   	"Action": jsii.String("s3:GetObject"),
	//   	"Resources": jsii.String("*"),
	//   })
	//
	// Experimental.
	AddToRolePolicy(statement interface{})
	// Encode an object so it can be passed as custom resource parameters.
	//
	// Custom resources will convert
	// all input parameters to strings so we encode non-strings here
	// so we can then decode them correctly in the provider function.
	// Experimental.
	Encode(obj interface{}) interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AssertionsProvider
type jsiiProxy_AssertionsProvider struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AssertionsProvider) HandlerRoleArn() awscdk.Reference {
	var returns awscdk.Reference
	_jsii_.Get(
		j,
		"handlerRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssertionsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AssertionsProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewAssertionsProvider(scope constructs.Construct, id *string) AssertionsProvider {
	_init_.Initialize()

	j := jsiiProxy_AssertionsProvider{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewAssertionsProvider_Override(a AssertionsProvider, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
		[]interface{}{scope, id},
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
func AssertionsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssertionsProvider) AddPolicyStatementFromSdkCall(service *string, api *string, resources *[]*string) {
	_jsii_.InvokeVoid(
		a,
		"addPolicyStatementFromSdkCall",
		[]interface{}{service, api, resources},
	)
}

func (a *jsiiProxy_AssertionsProvider) AddToRolePolicy(statement interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (a *jsiiProxy_AssertionsProvider) Encode(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"encode",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssertionsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

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
	_jsii_.InvokeVoid(
		a,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (a *jsiiProxy_AwsApiCall) Expect(expected ExpectedResult) {
	_jsii_.InvokeVoid(
		a,
		"expect",
		[]interface{}{expected},
	)
}

func (a *jsiiProxy_AwsApiCall) GetAtt(attributeName *string) awscdk.Reference {
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

// Options to perform an AWS JavaScript V2 API call.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var parameters interface{}
//
//   awsApiCallOptions := &awsApiCallOptions{
//   	api: jsii.String("api"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	parameters: parameters,
//   }
//
// Experimental.
type AwsApiCallOptions struct {
	// The api call to make, i.e. getBucketLifecycle.
	// Experimental.
	Api *string `field:"required" json:"api" yaml:"api"`
	// The AWS service, i.e. S3.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// Options for creating an SDKQuery provider.
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
type AwsApiCallProps struct {
	// The api call to make, i.e. getBucketLifecycle.
	// Experimental.
	Api *string `field:"required" json:"api" yaml:"api"`
	// The AWS service, i.e. S3.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// A AWS JavaScript SDK V2 request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var parameters interface{}
//
//   awsApiCallRequest := &awsApiCallRequest{
//   	api: jsii.String("api"),
//   	service: jsii.String("service"),
//
//   	// the properties below are optional
//   	flattenResponse: jsii.String("flattenResponse"),
//   	parameters: parameters,
//   }
//
// Experimental.
type AwsApiCallRequest struct {
	// The AWS api call to make i.e. getBucketLifecycle.
	// Experimental.
	Api *string `field:"required" json:"api" yaml:"api"`
	// The AWS service i.e. S3.
	// Experimental.
	Service *string `field:"required" json:"service" yaml:"service"`
	// Whether or not to flatten the response from the api call.
	//
	// Valid values are 'true' or 'false' as strings
	//
	// Typically when using an SdkRequest you will be passing it as the
	// `actual` value to an assertion provider so this would be set
	// to 'false' (you want the actual response).
	//
	// If you are using the SdkRequest to perform more of a query to return
	// a single value to use, then this should be set to 'true'. For example,
	// you could make a StepFunctions.startExecution api call and retreive the
	// `executionArn` from the response.
	// Experimental.
	FlattenResponse *string `field:"optional" json:"flattenResponse" yaml:"flattenResponse"`
	// Any parameters to pass to the api call.
	// Experimental.
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

// The result from a SdkQuery.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var apiCallResponse interface{}
//
//   awsApiCallResult := &awsApiCallResult{
//   	apiCallResponse: apiCallResponse,
//   }
//
// Experimental.
type AwsApiCallResult struct {
	// The full api response.
	// Experimental.
	ApiCallResponse interface{} `field:"required" json:"apiCallResponse" yaml:"apiCallResponse"`
}

// Construct that creates a CustomResource to assert that two values are equal.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actualResult actualResult
//   var expectedResult expectedResult
//
//   equalsAssertion := integ_tests_alpha.NewEqualsAssertion(this, jsii.String("MyEqualsAssertion"), &equalsAssertionProps{
//   	actual: actualResult,
//   	expected: expectedResult,
//
//   	// the properties below are optional
//   	failDeployment: jsii.Boolean(false),
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

// Options for an EqualsAssertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//
//   var actualResult actualResult
//   var expectedResult expectedResult
//
//   equalsAssertionProps := &equalsAssertionProps{
//   	actual: actualResult,
//   	expected: expectedResult,
//
//   	// the properties below are optional
//   	failDeployment: jsii.Boolean(false),
//   }
//
// Experimental.
type EqualsAssertionProps struct {
	// The actual results to compare.
	// Experimental.
	Actual ActualResult `field:"required" json:"actual" yaml:"actual"`
	// The expected result to assert.
	// Experimental.
	Expected ExpectedResult `field:"required" json:"expected" yaml:"expected"`
	// Set this to true if a failed assertion should result in a CloudFormation deployment failure.
	//
	// This is only necessary if assertions are being
	// executed outside of `integ-runner`.
	// Experimental.
	FailDeployment *bool `field:"optional" json:"failDeployment" yaml:"failDeployment"`
}

// Represents the "expected" results to compare.
//
// Example:
//   var app app
//   var integ integTest
//
//   integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("sendMessage"), map[string]*string{
//   	"QueueUrl": jsii.String("url"),
//   	"MessageBody": jsii.String("hello"),
//   })
//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
//   	"QueueUrl": jsii.String("url"),
//   })
//   message.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"Messages": []interface{}{
//   		map[string]*string{
//   			"Body": jsii.String("hello"),
//   		},
//   	},
//   }))
//
// Experimental.
type ExpectedResult interface {
	// The expected results encoded as a string.
	// Experimental.
	Result() *string
	// Experimental.
	SetResult(val *string)
}

// The jsii proxy struct for ExpectedResult
type jsiiProxy_ExpectedResult struct {
	_ byte // padding
}

func (j *jsiiProxy_ExpectedResult) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}


// Experimental.
func NewExpectedResult_Override(e ExpectedResult) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		nil, // no parameters
		e,
	)
}

func (j *jsiiProxy_ExpectedResult) SetResult(val *string) {
	_jsii_.Set(
		j,
		"result",
		val,
	)
}

// The actual results must be a list and must contain an item with the expected results.
//
// Example:
//   // actual results
//   actual := []map[string]*string{
//   	map[string]*string{
//   		"stringParam": jsii.String("hello"),
//   	},
//   	map[string]*string{
//   		"stringParam": jsii.String("world"),
//   	},
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult.arrayWith([]interface{}{
//   	map[string]*string{
//   		"stringParam": jsii.String("hello"),
//   	},
//   })
//
// Experimental.
func ExpectedResult_ArrayWith(expected *[]interface{}) ExpectedResult {
	_init_.Initialize()

	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"arrayWith",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// The actual results must match exactly.
//
// Missing data
// will result in a failure.
//
// Example:
//   // actual results
//   actual := map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult.exact(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   })
//
//   // fail
//   awscdkintegtestsalpha.ExpectedResult.exact(map[string]*string{
//   	"stringParam": jsii.String("hello"),
//   })
//
// Experimental.
func ExpectedResult_Exact(expected interface{}) ExpectedResult {
	_init_.Initialize()

	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"exact",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// The expected results must be a subset of the actual results.
//
// Example:
//   // actual results
//   actual := map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   	"numberParam": jsii.Number(3),
//   	"booleanParam": jsii.Boolean(true),
//   }
//   // pass
//   awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"stringParam": jsii.String("hello"),
//   })
//
// Experimental.
func ExpectedResult_ObjectLike(expected *map[string]interface{}) ExpectedResult {
	_init_.Initialize()

	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"objectLike",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// Actual results is a string that matches the Expected result regex.
//
// Example:
//   // actual results
//   actual := "some string value"
//
//   // pass
//   awscdkintegtestsalpha.ExpectedResult.stringLikeRegexp(jsii.String("value"))
//
// Experimental.
func ExpectedResult_StringLikeRegexp(expected *string) ExpectedResult {
	_init_.Initialize()

	var returns ExpectedResult

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.ExpectedResult",
		"stringLikeRegexp",
		[]interface{}{expected},
		&returns,
	)

	return returns
}

// Interface for creating a custom resource that will perform an API call using the AWS SDK.
// Experimental.
type IAwsApiCall interface {
	constructs.IConstruct
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall at the given path.
	//
	// For example the SQS.receiveMessage api response would look
	// like:
	//
	// If you wanted to assert the value of `Body` you could do.
	//
	// Example:
	//   var integ integTest
	//   actual := map[string][]map[string]interface{}{
	//   	"Messages": []map[string]interface{}{
	//   		map[string]interface{}{
	//   			"MessageId": jsii.String(""),
	//   			"ReceiptHandle": jsii.String(""),
	//   			"MD5OfBody": jsii.String(""),
	//   			"Body": jsii.String("hello"),
	//   			"Attributes": map[string]interface{}{
	//   			},
	//   			"MD5OfMessageAttributes": map[string]interface{}{
	//   			},
	//   			"MessageAttributes": map[string]interface{}{
	//   			},
	//   		},
	//   	},
	//   }
	//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"))
	//
	//   message.assertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult.stringLikeRegexp(jsii.String("hello")))
	//
	// Experimental.
	AssertAtPath(path *string, expected ExpectedResult)
	// Assert that the ExpectedResult is equal to the result of the AwsApiCall.
	//
	// Example:
	//   var integ integTest
	//
	//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	//   	functionName: jsii.String("my-func"),
	//   })
	//   invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }))
	//
	// Experimental.
	Expect(expected ExpectedResult)
	// Returns the value of an attribute of the custom resource of an arbitrary type.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt`. Use `Token.asXxx` to encode the returned `Reference` as a specific type or
	// use the convenience `getAttString` for string attributes.
	// Experimental.
	GetAtt(attributeName *string) awscdk.Reference
	// Returns the value of an attribute of the custom resource of type string.
	//
	// Attributes are returned from the custom resource provider through the
	// `Data` map where the key is the attribute name.
	//
	// Returns: a token for `Fn::GetAtt` encoded as a string.
	// Experimental.
	GetAttString(attributeName *string) *string
	// access the AssertionsProvider.
	//
	// This can be used to add additional IAM policies
	// the the provider role policy.
	//
	// Example:
	//   var apiCall awsApiCall
	//
	//   apiCall.provider.addToRolePolicy(map[string]interface{}{
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
	Provider() AssertionsProvider
}

// The jsii proxy for IAwsApiCall
type jsiiProxy_IAwsApiCall struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IAwsApiCall) AssertAtPath(path *string, expected ExpectedResult) {
	_jsii_.InvokeVoid(
		i,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (i *jsiiProxy_IAwsApiCall) Expect(expected ExpectedResult) {
	_jsii_.InvokeVoid(
		i,
		"expect",
		[]interface{}{expected},
	)
}

func (i *jsiiProxy_IAwsApiCall) GetAtt(attributeName *string) awscdk.Reference {
	var returns awscdk.Reference

	_jsii_.Invoke(
		i,
		"getAtt",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAwsApiCall) GetAttString(attributeName *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getAttString",
		[]interface{}{attributeName},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAwsApiCall) Provider() AssertionsProvider {
	var returns AssertionsProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

// Interface that allows for registering a list of assertions that should be performed on a construct.
//
// This is only necessary
// when writing integration tests.
// Experimental.
type IDeployAssert interface {
	// Query AWS using JavaScript SDK V2 API calls.
	//
	// This can be used to either
	// trigger an action or to return a result that can then be asserted against
	// an expected value.
	//
	// Example:
	//   var app app
	//   var integ integTest
	//
	//   integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("sendMessage"), map[string]*string{
	//   	"QueueUrl": jsii.String("url"),
	//   	"MessageBody": jsii.String("hello"),
	//   })
	//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]*string{
	//   	"QueueUrl": jsii.String("url"),
	//   })
	//   message.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Messages": []interface{}{
	//   		map[string]*string{
	//   			"Body": jsii.String("hello"),
	//   		},
	//   	},
	//   }))
	//
	// Experimental.
	AwsApiCall(service *string, api *string, parameters interface{}) IAwsApiCall
	// Assert that the ExpectedResult is equal to the ActualResult.
	//
	// Example:
	//   var integ integTest
	//   var apiCall awsApiCall
	//
	//   integ.assertions.expect(jsii.String("invoke"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("OK"),
	//   }), awscdkintegtestsalpha.ActualResult.fromAwsApiCall(apiCall, jsii.String("Body")))
	//
	// Experimental.
	Expect(id *string, expected ExpectedResult, actual ActualResult)
	// Invoke a lambda function and return the response which can be asserted.
	//
	// Example:
	//   var app app
	//   var integ integTest
	//
	//   invoke := integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
	//   	functionName: jsii.String("my-function"),
	//   })
	//   invoke.expect(awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
	//   	"Payload": jsii.String("200"),
	//   }))
	//
	// Experimental.
	InvokeFunction(props *LambdaInvokeFunctionProps) IAwsApiCall
}

// The jsii proxy for IDeployAssert
type jsiiProxy_IDeployAssert struct {
	_ byte // padding
}

func (i *jsiiProxy_IDeployAssert) AwsApiCall(service *string, api *string, parameters interface{}) IAwsApiCall {
	var returns IAwsApiCall

	_jsii_.Invoke(
		i,
		"awsApiCall",
		[]interface{}{service, api, parameters},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDeployAssert) Expect(id *string, expected ExpectedResult, actual ActualResult) {
	_jsii_.InvokeVoid(
		i,
		"expect",
		[]interface{}{id, expected, actual},
	)
}

func (i *jsiiProxy_IDeployAssert) InvokeFunction(props *LambdaInvokeFunctionProps) IAwsApiCall {
	var returns IAwsApiCall

	_jsii_.Invoke(
		i,
		"invokeFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

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

// An integration test case. Allows the definition of test properties that apply to all stacks under this case.
//
// It is recommended that you use the IntegTest construct since that will create
// a default IntegTestCase.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   integTestCase := integ_tests_alpha.NewIntegTestCase(this, jsii.String("MyIntegTestCase"), &integTestCaseProps{
//   	stacks: []*stack{
//   		stack,
//   	},
//
//   	// the properties below are optional
//   	allowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	cdkCommandOptions: &cdkCommands{
//   		deploy: &deployCommand{
//   			args: &deployOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				changeSetName: jsii.String("changeSetName"),
//   				ci: jsii.Boolean(false),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				execute: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				notificationArns: []*string{
//   					jsii.String("notificationArns"),
//   				},
//   				output: jsii.String("output"),
//   				outputsFile: jsii.String("outputsFile"),
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				requireApproval: awscdk.Cloud_assembly_schema.requireApproval_NEVER,
//   				reuseAssets: []*string{
//   					jsii.String("reuseAssets"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				rollback: jsii.Boolean(false),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				toolkitStackName: jsii.String("toolkitStackName"),
//   				trace: jsii.Boolean(false),
//   				usePreviousParameters: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   		destroy: &destroyCommand{
//   			args: &destroyOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				output: jsii.String("output"),
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				roleArn: jsii.String("roleArn"),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				trace: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   	},
//   	diffAssets: jsii.Boolean(false),
//   	hooks: &hooks{
//   		postDeploy: []*string{
//   			jsii.String("postDeploy"),
//   		},
//   		postDestroy: []*string{
//   			jsii.String("postDestroy"),
//   		},
//   		preDeploy: []*string{
//   			jsii.String("preDeploy"),
//   		},
//   		preDestroy: []*string{
//   			jsii.String("preDestroy"),
//   		},
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	stackUpdateWorkflow: jsii.Boolean(false),
//   })
//
// Experimental.
type IntegTestCase interface {
	constructs.Construct
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
	// The integration test manifest for this test case.
	//
	// Manifests are used
	// by the integration test runner.
	// Experimental.
	Manifest() *cloudassemblyschema.IntegManifest
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegTestCase
type jsiiProxy_IntegTestCase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_IntegTestCase) Assertions() IDeployAssert {
	var returns IDeployAssert
	_jsii_.Get(
		j,
		"assertions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCase) Manifest() *cloudassemblyschema.IntegManifest {
	var returns *cloudassemblyschema.IntegManifest
	_jsii_.Get(
		j,
		"manifest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTestCase(scope constructs.Construct, id *string, props *IntegTestCaseProps) IntegTestCase {
	_init_.Initialize()

	j := jsiiProxy_IntegTestCase{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCase_Override(i IntegTestCase, scope constructs.Construct, id *string, props *IntegTestCaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
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
func IntegTestCase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of an integration test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import integ_tests_alpha "github.com/aws/aws-cdk-go/awscdkintegtestsalpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var stack stack
//
//   integTestCaseProps := &integTestCaseProps{
//   	stacks: []*stack{
//   		stack,
//   	},
//
//   	// the properties below are optional
//   	allowDestroy: []*string{
//   		jsii.String("allowDestroy"),
//   	},
//   	cdkCommandOptions: &cdkCommands{
//   		deploy: &deployCommand{
//   			args: &deployOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				changeSetName: jsii.String("changeSetName"),
//   				ci: jsii.Boolean(false),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				execute: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				notificationArns: []*string{
//   					jsii.String("notificationArns"),
//   				},
//   				output: jsii.String("output"),
//   				outputsFile: jsii.String("outputsFile"),
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				requireApproval: awscdk.Cloud_assembly_schema.requireApproval_NEVER,
//   				reuseAssets: []*string{
//   					jsii.String("reuseAssets"),
//   				},
//   				roleArn: jsii.String("roleArn"),
//   				rollback: jsii.Boolean(false),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				toolkitStackName: jsii.String("toolkitStackName"),
//   				trace: jsii.Boolean(false),
//   				usePreviousParameters: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   		destroy: &destroyCommand{
//   			args: &destroyOptions{
//   				all: jsii.Boolean(false),
//   				app: jsii.String("app"),
//   				assetMetadata: jsii.Boolean(false),
//   				caBundlePath: jsii.String("caBundlePath"),
//   				color: jsii.Boolean(false),
//   				context: map[string]*string{
//   					"contextKey": jsii.String("context"),
//   				},
//   				debug: jsii.Boolean(false),
//   				ec2Creds: jsii.Boolean(false),
//   				exclusively: jsii.Boolean(false),
//   				force: jsii.Boolean(false),
//   				ignoreErrors: jsii.Boolean(false),
//   				json: jsii.Boolean(false),
//   				lookups: jsii.Boolean(false),
//   				notices: jsii.Boolean(false),
//   				output: jsii.String("output"),
//   				pathMetadata: jsii.Boolean(false),
//   				profile: jsii.String("profile"),
//   				proxy: jsii.String("proxy"),
//   				roleArn: jsii.String("roleArn"),
//   				stacks: []*string{
//   					jsii.String("stacks"),
//   				},
//   				staging: jsii.Boolean(false),
//   				strict: jsii.Boolean(false),
//   				trace: jsii.Boolean(false),
//   				verbose: jsii.Boolean(false),
//   				versionReporting: jsii.Boolean(false),
//   			},
//   			enabled: jsii.Boolean(false),
//   			expectedMessage: jsii.String("expectedMessage"),
//   			expectError: jsii.Boolean(false),
//   		},
//   	},
//   	diffAssets: jsii.Boolean(false),
//   	hooks: &hooks{
//   		postDeploy: []*string{
//   			jsii.String("postDeploy"),
//   		},
//   		postDestroy: []*string{
//   			jsii.String("postDestroy"),
//   		},
//   		preDeploy: []*string{
//   			jsii.String("preDeploy"),
//   		},
//   		preDestroy: []*string{
//   			jsii.String("preDestroy"),
//   		},
//   	},
//   	regions: []*string{
//   		jsii.String("regions"),
//   	},
//   	stackUpdateWorkflow: jsii.Boolean(false),
//   }
//
// Experimental.
type IntegTestCaseProps struct {
	// List of CloudFormation resource types in this stack that can be destroyed as part of an update without failing the test.
	//
	// This list should only include resources that for this specific
	// integration test we are sure will not cause errors or an outage if
	// destroyed. For example, maybe we know that a new resource will be created
	// first before the old resource is destroyed which prevents any outage.
	//
	// e.g. ['AWS::IAM::Role']
	// Experimental.
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Experimental.
	CdkCommandOptions *cloudassemblyschema.CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Experimental.
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Experimental.
	Hooks *cloudassemblyschema.Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Experimental.
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// Stacks to be deployed during the test.
	// Experimental.
	Stacks *[]awscdk.Stack `field:"required" json:"stacks" yaml:"stacks"`
}

// An integration test case stack. Allows the definition of test properties that should apply to this stack.
//
// This should be used if there are multiple stacks in the integration test
// and it is necessary to specify different test case option for each. Otherwise
// normal stacks should be added to IntegTest.
//
// Example:
//   var app app
//   var stackUnderTest stack
//
//   testCaseWithAssets := awscdkintegtestsalpha.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &integTestCaseStackProps{
//   	diffAssets: jsii.Boolean(true),
//   })
//
//   awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
//   	testCases: []*stack{
//   		stackUnderTest,
//   		testCaseWithAssets,
//   	},
//   })
//
// Experimental.
type IntegTestCaseStack interface {
	awscdk.Stack
	// The AWS account into which this stack will be deployed.
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.account` when the stack is defined. This can
	//     either be a concrete account (e.g. `585695031111`) or the
	//     `Aws.ACCOUNT_ID` token.
	// 3. `Aws.ACCOUNT_ID`, which represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::AccountId" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.account)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **account-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	// Experimental.
	Account() *string
	// The ID of the cloud assembly artifact for this stack.
	// Experimental.
	ArtifactId() *string
	// Make assertions on resources in this test case.
	// Experimental.
	Assertions() IDeployAssert
	// Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack.
	//
	// If the stack is environment-agnostic (either account and/or region are
	// tokens), this property will return an array with 2 tokens that will resolve
	// at deploy-time to the first two availability zones returned from CloudFormation's
	// `Fn::GetAZs` intrinsic function.
	//
	// If they are not available in the context, returns a set of dummy values and
	// reports them as missing, and let the CLI resolve them by calling EC2
	// `DescribeAvailabilityZones` on the target environment.
	//
	// To specify a different strategy for selecting availability zones override this method.
	// Experimental.
	AvailabilityZones() *[]*string
	// Indicates whether the stack requires bundling or not.
	// Experimental.
	BundlingRequired() *bool
	// Return the stacks this stack depends on.
	// Experimental.
	Dependencies() *[]awscdk.Stack
	// The environment coordinates in which this stack is deployed.
	//
	// In the form
	// `aws://account/region`. Use `stack.account` and `stack.region` to obtain
	// the specific values, no need to parse.
	//
	// You can use this value to determine if two stacks are targeting the same
	// environment.
	//
	// If either `stack.account` or `stack.region` are not concrete values (e.g.
	// `Aws.ACCOUNT_ID` or `Aws.REGION`) the special strings `unknown-account` and/or
	// `unknown-region` will be used respectively to indicate this stack is
	// region/account-agnostic.
	// Experimental.
	Environment() *string
	// Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.
	// Experimental.
	Nested() *bool
	// If this is a nested stack, returns it's parent stack.
	// Experimental.
	NestedStackParent() awscdk.Stack
	// If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.
	//
	// `undefined` for top-level (non-nested) stacks.
	// Experimental.
	NestedStackResource() awscdk.CfnResource
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns the list of notification Amazon Resource Names (ARNs) for the current stack.
	// Experimental.
	NotificationArns() *[]*string
	// The partition in which this stack is defined.
	// Experimental.
	Partition() *string
	// The AWS region into which this stack will be deployed (e.g. `us-west-2`).
	//
	// This value is resolved according to the following rules:
	//
	// 1. The value provided to `env.region` when the stack is defined. This can
	//     either be a concerete region (e.g. `us-west-2`) or the `Aws.REGION`
	//     token.
	// 3. `Aws.REGION`, which is represents the CloudFormation intrinsic reference
	//     `{ "Ref": "AWS::Region" }` encoded as a string token.
	//
	// Preferably, you should use the return value as an opaque string and not
	// attempt to parse it to implement your logic. If you do, you must first
	// check that it is a concerete value an not an unresolved token. If this
	// value is an unresolved token (`Token.isUnresolved(stack.region)` returns
	// `true`), this implies that the user wishes that this stack will synthesize
	// into a **region-agnostic template**. In this case, your code should either
	// fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
	// implement some other region-agnostic behavior.
	// Experimental.
	Region() *string
	// The ID of the stack.
	//
	// Example:
	//   // After resolving, looks like
	//   'arn:aws:cloudformation:us-west-2:123456789012:stack/teststack/51af3dc0-da77-11e4-872e-1234567db123'
	//
	// Experimental.
	StackId() *string
	// The concrete CloudFormation physical stack name.
	//
	// This is either the name defined explicitly in the `stackName` prop or
	// allocated based on the stack's location in the construct tree. Stacks that
	// are directly defined under the app use their construct `id` as their stack
	// name. Stacks that are defined deeper within the tree will use a hashed naming
	// scheme based on the construct path to ensure uniqueness.
	//
	// If you wish to obtain the deploy-time AWS::StackName intrinsic,
	// you can use `Aws.STACK_NAME` directly.
	// Experimental.
	StackName() *string
	// Synthesis method for this stack.
	// Experimental.
	Synthesizer() awscdk.IStackSynthesizer
	// Tags to be applied to the stack.
	// Experimental.
	Tags() awscdk.TagManager
	// The name of the CloudFormation template file emitted to the output directory during synthesis.
	//
	// Example value: `MyStack.template.json`
	// Experimental.
	TemplateFile() *string
	// Options for CloudFormation template (like version, transform, description).
	// Experimental.
	TemplateOptions() awscdk.ITemplateOptions
	// Whether termination protection is enabled for this stack.
	// Experimental.
	TerminationProtection() *bool
	// The Amazon domain suffix for the region in which this stack is defined.
	// Experimental.
	UrlSuffix() *string
	// Add a dependency between this stack and another stack.
	//
	// This can be used to define dependencies between any two stacks within an
	// app, and also supports nested stacks.
	// Experimental.
	AddDependency(target awscdk.Stack, reason *string)
	// Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.
	//
	// Duplicate values are removed when stack is synthesized.
	//
	// Example:
	//   declare const stack: Stack;
	//
	//   stack.addTransform('AWS::Serverless-2016-10-31')
	//
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html
	//
	// Experimental.
	AddTransform(transform *string)
	// Returns the naming scheme used to allocate logical IDs.
	//
	// By default, uses
	// the `HashedAddressingScheme` but this method can be overridden to customize
	// this behavior.
	//
	// In order to make sure logical IDs are unique and stable, we hash the resource
	// construct tree path (i.e. toplevel/secondlevel/.../myresource) and add it as
	// a suffix to the path components joined without a separator (CloudFormation
	// IDs only allow alphanumeric characters).
	//
	// The result will be:
	//
	//    <path.join('')><md5(path.join('/')>
	//      "human"      "hash"
	//
	// If the "human" part of the ID exceeds 240 characters, we simply trim it so
	// the total ID doesn't exceed CloudFormation's 255 character limit.
	//
	// We only take 8 characters from the md5 hash (0.000005 chance of collision).
	//
	// Special cases:
	//
	// - If the path only contains a single component (i.e. it's a top-level
	//    resource), we won't add the hash to it. The hash is not needed for
	//    disamiguation and also, it allows for a more straightforward migration an
	//    existing CloudFormation template to a CDK stack without logical ID changes
	//    (or renames).
	// - For aesthetic reasons, if the last components of the path are the same
	//    (i.e. `L1/L2/Pipeline/Pipeline`), they will be de-duplicated to make the
	//    resulting human portion of the ID more pleasing: `L1L2Pipeline<HASH>`
	//    instead of `L1L2PipelinePipeline<HASH>`
	// - If a component is named "Default" it will be omitted from the path. This
	//    allows refactoring higher level abstractions around constructs without affecting
	//    the IDs of already deployed resources.
	// - If a component is named "Resource" it will be omitted from the user-visible
	//    path, but included in the hash. This reduces visual noise in the human readable
	//    part of the identifier.
	// Experimental.
	AllocateLogicalId(cfnElement awscdk.CfnElement) *string
	// Create a CloudFormation Export for a value.
	//
	// Returns a string representing the corresponding `Fn.importValue()`
	// expression for this Export. You can control the name for the export by
	// passing the `name` option.
	//
	// If you don't supply a value for `name`, the value you're exporting must be
	// a Resource attribute (for example: `bucket.bucketName`) and it will be
	// given the same name as the automatic cross-stack reference that would be created
	// if you used the attribute in another Stack.
	//
	// One of the uses for this method is to *remove* the relationship between
	// two Stacks established by automatic cross-stack references. It will
	// temporarily ensure that the CloudFormation Export still exists while you
	// remove the reference from the consuming stack. After that, you can remove
	// the resource and the manual export.
	//
	// ## Example
	//
	// Here is how the process works. Let's say there are two stacks,
	// `producerStack` and `consumerStack`, and `producerStack` has a bucket
	// called `bucket`, which is referenced by `consumerStack` (perhaps because
	// an AWS Lambda Function writes into it, or something like that).
	//
	// It is not safe to remove `producerStack.bucket` because as the bucket is being
	// deleted, `consumerStack` might still be using it.
	//
	// Instead, the process takes two deployments:
	//
	// ### Deployment 1: break the relationship
	//
	// - Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
	//    stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
	//    remove the Lambda Function altogether).
	// - In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
	//    will make sure the CloudFormation Export continues to exist while the relationship
	//    between the two stacks is being broken.
	// - Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).
	//
	// ### Deployment 2: remove the bucket resource
	//
	// - You are now free to remove the `bucket` resource from `producerStack`.
	// - Don't forget to remove the `exportValue()` call as well.
	// - Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).
	// Experimental.
	ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string
	// Creates an ARN from components.
	//
	// If `partition`, `region` or `account` are not specified, the stack's
	// partition, region and account will be used.
	//
	// If any component is the empty string, an empty string will be inserted
	// into the generated ARN at the location that component corresponds to.
	//
	// The ARN will be formatted as follows:
	//
	//    arn:{partition}:{service}:{region}:{account}:{resource}{sep}{resource-name}
	//
	// The required ARN pieces that are omitted will be taken from the stack that
	// the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
	// can be 'undefined'.
	// Experimental.
	FormatArn(components *awscdk.ArnComponents) *string
	// Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.
	//
	// This method is called when a `CfnElement` is created and used to render the
	// initial logical identity of resources. Logical ID renames are applied at
	// this stage.
	//
	// This method uses the protected method `allocateLogicalId` to render the
	// logical ID for an element. To modify the naming scheme, extend the `Stack`
	// class and override this method.
	// Experimental.
	GetLogicalId(element awscdk.CfnElement) *string
	// Look up a fact value for the given fact for the region of this stack.
	//
	// Will return a definite value only if the region of the current stack is resolved.
	// If not, a lookup map will be added to the stack and the lookup will be done at
	// CDK deployment time.
	//
	// What regions will be included in the lookup map is controlled by the
	// `@aws-cdk/core:target-partitions` context value: it must be set to a list
	// of partitions, and only regions from the given partitions will be included.
	// If no such context key is set, all regions will be included.
	//
	// This function is intended to be used by construct library authors. Application
	// builders can rely on the abstractions offered by construct libraries and do
	// not have to worry about regional facts.
	//
	// If `defaultValue` is not given, it is an error if the fact is unknown for
	// the given region.
	// Experimental.
	RegionalFact(factName *string, defaultValue *string) *string
	// Rename a generated logical identities.
	//
	// To modify the naming scheme strategy, extend the `Stack` class and
	// override the `allocateLogicalId` method.
	// Experimental.
	RenameLogicalId(oldId *string, newId *string)
	// Indicate that a context key was expected.
	//
	// Contains instructions which will be emitted into the cloud assembly on how
	// the key should be supplied.
	// Experimental.
	ReportMissingContextKey(report *cloudassemblyschema.MissingContext)
	// Resolve a tokenized value in the context of the current stack.
	// Experimental.
	Resolve(obj interface{}) interface{}
	// Splits the provided ARN into its components.
	//
	// Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
	// and a Token representing a dynamic CloudFormation expression
	// (in which case the returned components will also be dynamic CloudFormation expressions,
	// encoded as Tokens).
	// Experimental.
	SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents
	// Convert an object, potentially containing tokens, to a JSON string.
	// Experimental.
	ToJsonString(obj interface{}, space *float64) *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IntegTestCaseStack
type jsiiProxy_IntegTestCaseStack struct {
	internal.Type__awscdkStack
}

func (j *jsiiProxy_IntegTestCaseStack) Account() *string {
	var returns *string
	_jsii_.Get(
		j,
		"account",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) ArtifactId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Assertions() IDeployAssert {
	var returns IDeployAssert
	_jsii_.Get(
		j,
		"assertions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) BundlingRequired() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"bundlingRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Dependencies() *[]awscdk.Stack {
	var returns *[]awscdk.Stack
	_jsii_.Get(
		j,
		"dependencies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Nested() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"nested",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NestedStackParent() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"nestedStackParent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NestedStackResource() awscdk.CfnResource {
	var returns awscdk.CfnResource
	_jsii_.Get(
		j,
		"nestedStackResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) NotificationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notificationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Partition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) StackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Synthesizer() awscdk.IStackSynthesizer {
	var returns awscdk.IStackSynthesizer
	_jsii_.Get(
		j,
		"synthesizer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) Tags() awscdk.TagManager {
	var returns awscdk.TagManager
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TemplateFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TemplateOptions() awscdk.ITemplateOptions {
	var returns awscdk.ITemplateOptions
	_jsii_.Get(
		j,
		"templateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) TerminationProtection() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"terminationProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IntegTestCaseStack) UrlSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlSuffix",
		&returns,
	)
	return returns
}


// Experimental.
func NewIntegTestCaseStack(scope constructs.Construct, id *string, props *IntegTestCaseStackProps) IntegTestCaseStack {
	_init_.Initialize()

	j := jsiiProxy_IntegTestCaseStack{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewIntegTestCaseStack_Override(i IntegTestCaseStack, scope constructs.Construct, id *string, props *IntegTestCaseStackProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
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
func IntegTestCaseStack_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns whether the construct is a IntegTestCaseStack.
// Experimental.
func IntegTestCaseStack_IsIntegTestCaseStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isIntegTestCaseStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Return whether the given object is a Stack.
//
// We do attribute detection since we can't reliably use 'instanceof'.
// Experimental.
func IntegTestCaseStack_IsStack(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"isStack",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Looks up the first stack scope in which `construct` is defined.
//
// Fails if there is no stack up the tree.
// Experimental.
func IntegTestCaseStack_Of(construct constructs.IConstruct) awscdk.Stack {
	_init_.Initialize()

	var returns awscdk.Stack

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.IntegTestCaseStack",
		"of",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) AddDependency(target awscdk.Stack, reason *string) {
	_jsii_.InvokeVoid(
		i,
		"addDependency",
		[]interface{}{target, reason},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AddTransform(transform *string) {
	_jsii_.InvokeVoid(
		i,
		"addTransform",
		[]interface{}{transform},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) AllocateLogicalId(cfnElement awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"allocateLogicalId",
		[]interface{}{cfnElement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ExportValue(exportedValue interface{}, options *awscdk.ExportValueOptions) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"exportValue",
		[]interface{}{exportedValue, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) FormatArn(components *awscdk.ArnComponents) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"formatArn",
		[]interface{}{components},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) GetLogicalId(element awscdk.CfnElement) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"getLogicalId",
		[]interface{}{element},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) RegionalFact(factName *string, defaultValue *string) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"regionalFact",
		[]interface{}{factName, defaultValue},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) RenameLogicalId(oldId *string, newId *string) {
	_jsii_.InvokeVoid(
		i,
		"renameLogicalId",
		[]interface{}{oldId, newId},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) ReportMissingContextKey(report *cloudassemblyschema.MissingContext) {
	_jsii_.InvokeVoid(
		i,
		"reportMissingContextKey",
		[]interface{}{report},
	)
}

func (i *jsiiProxy_IntegTestCaseStack) Resolve(obj interface{}) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) SplitArn(arn *string, arnFormat awscdk.ArnFormat) *awscdk.ArnComponents {
	var returns *awscdk.ArnComponents

	_jsii_.Invoke(
		i,
		"splitArn",
		[]interface{}{arn, arnFormat},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ToJsonString(obj interface{}, space *float64) *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toJsonString",
		[]interface{}{obj, space},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IntegTestCaseStack) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Properties of an integration test case stack.
//
// Example:
//   var app app
//   var stackUnderTest stack
//
//   testCaseWithAssets := awscdkintegtestsalpha.NewIntegTestCaseStack(app, jsii.String("TestCaseAssets"), &integTestCaseStackProps{
//   	diffAssets: jsii.Boolean(true),
//   })
//
//   awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
//   	testCases: []*stack{
//   		stackUnderTest,
//   		testCaseWithAssets,
//   	},
//   })
//
// Experimental.
type IntegTestCaseStackProps struct {
	// List of CloudFormation resource types in this stack that can be destroyed as part of an update without failing the test.
	//
	// This list should only include resources that for this specific
	// integration test we are sure will not cause errors or an outage if
	// destroyed. For example, maybe we know that a new resource will be created
	// first before the old resource is destroyed which prevents any outage.
	//
	// e.g. ['AWS::IAM::Role']
	// Experimental.
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Experimental.
	CdkCommandOptions *cloudassemblyschema.CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Experimental.
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Experimental.
	Hooks *cloudassemblyschema.Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Experimental.
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// Include runtime versioning information in this Stack.
	// Experimental.
	AnalyticsReporting *bool `field:"optional" json:"analyticsReporting" yaml:"analyticsReporting"`
	// A description of the stack.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The AWS environment (account/region) where this stack will be deployed.
	//
	// Set the `region`/`account` fields of `env` to either a concrete value to
	// select the indicated environment (recommended for production stacks), or to
	// the values of environment variables
	// `CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
	// depend on the AWS credentials/configuration that the CDK CLI is executed
	// under (recommended for development stacks).
	//
	// If the `Stack` is instantiated inside a `Stage`, any undefined
	// `region`/`account` fields from `env` will default to the same field on the
	// encompassing `Stage`, if configured there.
	//
	// If either `region` or `account` are not set nor inherited from `Stage`, the
	// Stack will be considered "*environment-agnostic*"". Environment-agnostic
	// stacks can be deployed to any environment but may not be able to take
	// advantage of all features of the CDK. For example, they will not be able to
	// use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
	// automatically translate Service Principals to the right format based on the
	// environment's AWS partition, and other such enhancements.
	//
	// Example:
	//   // Use a concrete account and region to deploy this stack to:
	//   // `.account` and `.region` will simply return these values.
	//   new Stack(app, 'Stack1', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     },
	//   });
	//
	//   // Use the CLI's current credentials to determine the target environment:
	//   // `.account` and `.region` will reflect the account+region the CLI
	//   // is configured to use (based on the user CLI credentials)
	//   new Stack(app, 'Stack2', {
	//     env: {
	//       account: process.env.CDK_DEFAULT_ACCOUNT,
	//       region: process.env.CDK_DEFAULT_REGION
	//     },
	//   });
	//
	//   // Define multiple stacks stage associated with an environment
	//   const myStage = new Stage(app, 'MyStage', {
	//     env: {
	//       account: '123456789012',
	//       region: 'us-east-1'
	//     }
	//   });
	//
	//   // both of these stacks will use the stage's account/region:
	//   // `.account` and `.region` will resolve to the concrete values as above
	//   new MyStack(myStage, 'Stack1');
	//   new YourStack(myStage, 'Stack2');
	//
	//   // Define an environment-agnostic stack:
	//   // `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
	//   // which will only resolve to actual values by CloudFormation during deployment.
	//   new MyStack(app, 'Stack1');
	//
	// Experimental.
	Env *awscdk.Environment `field:"optional" json:"env" yaml:"env"`
	// Name to deploy the stack with.
	// Experimental.
	StackName *string `field:"optional" json:"stackName" yaml:"stackName"`
	// Synthesis method to use while deploying this stack.
	// Experimental.
	Synthesizer awscdk.IStackSynthesizer `field:"optional" json:"synthesizer" yaml:"synthesizer"`
	// Stack tags that will be applied to all the taggable resources and the stack itself.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Whether to enable termination protection for this stack.
	// Experimental.
	TerminationProtection *bool `field:"optional" json:"terminationProtection" yaml:"terminationProtection"`
}

// Integration test properties.
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
type IntegTestProps struct {
	// List of CloudFormation resource types in this stack that can be destroyed as part of an update without failing the test.
	//
	// This list should only include resources that for this specific
	// integration test we are sure will not cause errors or an outage if
	// destroyed. For example, maybe we know that a new resource will be created
	// first before the old resource is destroyed which prevents any outage.
	//
	// e.g. ['AWS::IAM::Role']
	// Experimental.
	AllowDestroy *[]*string `field:"optional" json:"allowDestroy" yaml:"allowDestroy"`
	// Additional options to use for each CDK command.
	// Experimental.
	CdkCommandOptions *cloudassemblyschema.CdkCommands `field:"optional" json:"cdkCommandOptions" yaml:"cdkCommandOptions"`
	// Whether or not to include asset hashes in the diff Asset hashes can introduces a lot of unneccessary noise into tests, but there are some cases where asset hashes _should_ be included.
	//
	// For example
	// any tests involving custom resources or bundling.
	// Experimental.
	DiffAssets *bool `field:"optional" json:"diffAssets" yaml:"diffAssets"`
	// Additional commands to run at predefined points in the test workflow.
	//
	// e.g. { postDeploy: ['yarn', 'test'] }
	// Experimental.
	Hooks *cloudassemblyschema.Hooks `field:"optional" json:"hooks" yaml:"hooks"`
	// Limit deployment to these regions.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Run update workflow on this test case This should only be set to false to test scenarios that are not possible to test as part of the update workflow.
	// Experimental.
	StackUpdateWorkflow *bool `field:"optional" json:"stackUpdateWorkflow" yaml:"stackUpdateWorkflow"`
	// List of test cases that make up this test.
	// Experimental.
	TestCases *[]awscdk.Stack `field:"required" json:"testCases" yaml:"testCases"`
}

// The type of invocation.
//
// Default is REQUEST_RESPONE.
//
// Example:
//   var app app
//   var stack stack
//   var queue queue
//   var fn iFunction
//
//
//   integ := awscdkintegtestsalpha.NewIntegTest(app, jsii.String("Integ"), &integTestProps{
//   	testCases: []*stack{
//   		stack,
//   	},
//   })
//
//   integ.assertions.invokeFunction(&lambdaInvokeFunctionProps{
//   	functionName: fn.functionName,
//   	invocationType: awscdkintegtestsalpha.InvocationType_EVENT,
//   	payload: jSON.stringify(map[string]*string{
//   		"status": jsii.String("OK"),
//   	}),
//   })
//
//   message := integ.assertions.awsApiCall(jsii.String("SQS"), jsii.String("receiveMessage"), map[string]interface{}{
//   	"QueueUrl": queue.queueUrl,
//   	"WaitTimeSeconds": jsii.Number(20),
//   })
//
//   message.assertAtPath(jsii.String("Messages.0.Body"), awscdkintegtestsalpha.ExpectedResult.objectLike(map[string]interface{}{
//   	"requestContext": map[string]*string{
//   		"condition": jsii.String("Success"),
//   	},
//   	"requestPayload": map[string]*string{
//   		"status": jsii.String("OK"),
//   	},
//   	"responseContext": map[string]*f64{
//   		"statusCode": jsii.Number(200),
//   	},
//   	"responsePayload": jsii.String("success"),
//   }))
//
// Experimental.
type InvocationType string

const (
	// Invoke the function asynchronously.
	//
	// Send events that fail multiple times to the function's
	// dead-letter queue (if it's configured).
	// The API response only includes a status code.
	// Experimental.
	InvocationType_EVENT InvocationType = "EVENT"
	// Invoke the function synchronously.
	//
	// Keep the connection open until the function returns a response or times out.
	// The API response includes the function response and additional data.
	// Experimental.
	InvocationType_REQUEST_RESPONE InvocationType = "REQUEST_RESPONE"
	// Validate parameter values and verify that the user or role has permission to invoke the function.
	// Experimental.
	InvocationType_DRY_RUN InvocationType = "DRY_RUN"
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
	_jsii_.InvokeVoid(
		l,
		"assertAtPath",
		[]interface{}{path, expected},
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) Expect(expected ExpectedResult) {
	_jsii_.InvokeVoid(
		l,
		"expect",
		[]interface{}{expected},
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) GetAtt(attributeName *string) awscdk.Reference {
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

// Options to pass to the Lambda invokeFunction API call.
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
type LambdaInvokeFunctionProps struct {
	// The name of the function to invoke.
	// Experimental.
	FunctionName *string `field:"required" json:"functionName" yaml:"functionName"`
	// The type of invocation to use.
	// Experimental.
	InvocationType InvocationType `field:"optional" json:"invocationType" yaml:"invocationType"`
	// Whether to return the logs as part of the response.
	// Experimental.
	LogType LogType `field:"optional" json:"logType" yaml:"logType"`
	// Payload to send as part of the invoke.
	// Experimental.
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
}

// Set to Tail to include the execution log in the response.
//
// Applies to synchronously invoked functions only.
// Experimental.
type LogType string

const (
	// The log messages are not returned in the response.
	// Experimental.
	LogType_NONE LogType = "NONE"
	// The log messages are returned in the response.
	// Experimental.
	LogType_TAIL LogType = "TAIL"
)

// Partial and special matching during assertions.
// Experimental.
type Match interface {
}

// The jsii proxy struct for Match
type jsiiProxy_Match struct {
	_ byte // padding
}

// Experimental.
func NewMatch_Override(m Match) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.Match",
		nil, // no parameters
		m,
	)
}

// Matches the specified pattern with the array found in the same relative path of the target.
//
// The set of elements (or matchers) must be in the same order as would be found.
// Experimental.
func Match_ArrayWith(pattern *[]interface{}) *map[string]*[]interface{} {
	_init_.Initialize()

	var returns *map[string]*[]interface{}

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"arrayWith",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches the specified pattern to an object found in the same relative path of the target.
//
// The keys and their values (or matchers) must be present in the target but the target can be a superset.
// Experimental.
func Match_ObjectLike(pattern *map[string]interface{}) *map[string]*map[string]interface{} {
	_init_.Initialize()

	var returns *map[string]*map[string]interface{}

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"objectLike",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// Matches targets according to a regular expression.
// Experimental.
func Match_StringLikeRegexp(pattern *string) *map[string]*string {
	_init_.Initialize()

	var returns *map[string]*string

	_jsii_.StaticInvoke(
		"@aws-cdk/integ-tests-alpha.Match",
		"stringLikeRegexp",
		[]interface{}{pattern},
		&returns,
	)

	return returns
}

// The status of the assertion.
// Experimental.
type Status string

const (
	// The assertion passed.
	// Experimental.
	Status_PASS Status = "PASS"
	// The assertion failed.
	// Experimental.
	Status_FAIL Status = "FAIL"
)

