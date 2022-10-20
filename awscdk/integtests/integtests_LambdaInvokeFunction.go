package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/constructs-go/constructs/v3"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaInvokeFunction := awscdk.Integ_tests.NewLambdaInvokeFunction(this, jsii.String("MyLambdaInvokeFunction"), &lambdaInvokeFunctionProps{
//   	functionName: jsii.String("functionName"),
//
//   	// the properties below are optional
//   	invocationType: awscdk.*Integ_tests.invocationType_EVENT,
//   	logType: awscdk.*Integ_tests.logType_NONE,
//   	payload: jsii.String("payload"),
//   })
//
// Experimental.
type LambdaInvokeFunction interface {
	AwsApiCall
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

// The jsii proxy struct for LambdaInvokeFunction
type jsiiProxy_LambdaInvokeFunction struct {
	jsiiProxy_AwsApiCall
}

func (j *jsiiProxy_LambdaInvokeFunction) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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
		"monocdk.integ_tests.LambdaInvokeFunction",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaInvokeFunction_Override(l LambdaInvokeFunction, scope constructs.Construct, id *string, props *LambdaInvokeFunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.LambdaInvokeFunction",
		[]interface{}{scope, id, props},
		l,
	)
}

func (j *jsiiProxy_LambdaInvokeFunction)SetProvider(val AssertionsProvider) {
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
func LambdaInvokeFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaInvokeFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.LambdaInvokeFunction",
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

func (l *jsiiProxy_LambdaInvokeFunction) OnPrepare() {
	_jsii_.InvokeVoid(
		l,
		"onPrepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) OnSynthesize(session constructs.ISynthesisSession) {
	if err := l.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaInvokeFunction) Prepare() {
	_jsii_.InvokeVoid(
		l,
		"prepare",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaInvokeFunction) Synthesize(session awscdk.ISynthesisSession) {
	if err := l.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"synthesize",
		[]interface{}{session},
	)
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

func (l *jsiiProxy_LambdaInvokeFunction) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

