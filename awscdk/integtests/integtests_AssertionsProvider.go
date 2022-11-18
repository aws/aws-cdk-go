package integtests

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/integtests/internal"
	"github.com/aws/constructs-go/constructs/v3"
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
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assertionsProvider := awscdk.Integ_tests.NewAssertionsProvider(this, jsii.String("MyAssertionsProvider"))
//
// Experimental.
type AssertionsProvider interface {
	awscdk.Construct
	// A reference to the provider Lambda Function execution Role ARN.
	// Experimental.
	HandlerRoleArn() awscdk.Reference
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The ARN of the lambda function which can be used as a serviceToken to a CustomResource.
	// Experimental.
	ServiceToken() *string
	// Create a policy statement from a specific api call.
	// Experimental.
	AddPolicyStatementFromSdkCall(service *string, api *string, resources *[]*string)
	// Encode an object so it can be passed as custom resource parameters.
	//
	// Custom resources will convert
	// all input parameters to strings so we encode non-strings here
	// so we can then decode them correctly in the provider function.
	// Experimental.
	Encode(obj interface{}) interface{}
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

// The jsii proxy struct for AssertionsProvider
type jsiiProxy_AssertionsProvider struct {
	internal.Type__awscdkConstruct
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

func (j *jsiiProxy_AssertionsProvider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
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

	if err := validateNewAssertionsProviderParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssertionsProvider{}

	_jsii_.Create(
		"monocdk.integ_tests.AssertionsProvider",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewAssertionsProvider_Override(a AssertionsProvider, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.integ_tests.AssertionsProvider",
		[]interface{}{scope, id},
		a,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func AssertionsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssertionsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.integ_tests.AssertionsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssertionsProvider) AddPolicyStatementFromSdkCall(service *string, api *string, resources *[]*string) {
	if err := a.validateAddPolicyStatementFromSdkCallParameters(service, api); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPolicyStatementFromSdkCall",
		[]interface{}{service, api, resources},
	)
}

func (a *jsiiProxy_AssertionsProvider) Encode(obj interface{}) interface{} {
	if err := a.validateEncodeParameters(obj); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"encode",
		[]interface{}{obj},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssertionsProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		a,
		"onPrepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssertionsProvider) OnSynthesize(session constructs.ISynthesisSession) {
	if err := a.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (a *jsiiProxy_AssertionsProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AssertionsProvider) Prepare() {
	_jsii_.InvokeVoid(
		a,
		"prepare",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AssertionsProvider) Synthesize(session awscdk.ISynthesisSession) {
	if err := a.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"synthesize",
		[]interface{}{session},
	)
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

func (a *jsiiProxy_AssertionsProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

