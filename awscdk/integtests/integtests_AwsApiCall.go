package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/integtests/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Construct that creates a custom resource that will perform a query using the AWS SDK.
//
// Example:
//   var myAppStack stack
//
//
//   awscdk.NewAwsApiCall(myAppStack, jsii.String("GetObject"), &AwsApiCallProps{
//   	Service: jsii.String("S3"),
//   	Api: jsii.String("getObject"),
//   })
//
// Experimental.
type AwsApiCall interface {
	awscdk.Construct
	IAwsApiCall
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Experimental.
	Provider() AssertionsProvider
	// Experimental.
	SetProvider(val AssertionsProvider)
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
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for AwsApiCall
type jsiiProxy_AwsApiCall struct {
	internal.Type__awscdkConstruct
	jsiiProxy_IAwsApiCall
}

func (j *jsiiProxy_AwsApiCall) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.integ_tests.AwsApiCall",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsApiCall_Override(a AwsApiCall, scope constructs.Construct, id *string, props *AwsApiCallProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.AwsApiCall",
		[]interface{}{scope, id, props},
		a,
	)
}

func (j *jsiiProxy_AwsApiCall)SetProvider(val AssertionsProvider) {
	if err := j.validateSetProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func AwsApiCall_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsApiCall_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.AwsApiCall",
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

func (a *jsiiProxy_AwsApiCall) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiCall) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AwsApiCall) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsApiCall) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsApiCall) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_AwsApiCall) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

