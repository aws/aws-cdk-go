package awscdkintegtestsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkintegtestsalpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
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
//   assertionsProvider := integ_tests_alpha.NewAssertionsProvider(this, jsii.String("MyAssertionsProvider"), &AssertionsProviderProps{
//   	Handler: jsii.String("handler"),
//   	Uuid: jsii.String("uuid"),
//   })
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
	//   provider.AddToRolePolicy(map[string]interface{}{
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
	AddToRolePolicy(statement interface{})
	// Encode an object so it can be passed as custom resource parameters.
	//
	// Custom resources will convert
	// all input parameters to strings so we encode non-strings here
	// so we can then decode them correctly in the provider function.
	// Experimental.
	Encode(obj interface{}) interface{}
	// Grant a principal access to invoke the assertion provider lambda function.
	// Experimental.
	GrantInvoke(principalArn *string)
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
func NewAssertionsProvider(scope constructs.Construct, id *string, props *AssertionsProviderProps) AssertionsProvider {
	_init_.Initialize()

	if err := validateNewAssertionsProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AssertionsProvider{}

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAssertionsProvider_Override(a AssertionsProvider, scope constructs.Construct, id *string, props *AssertionsProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/integ-tests-alpha.AssertionsProvider",
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
func AssertionsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAssertionsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
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
	if err := a.validateAddPolicyStatementFromSdkCallParameters(service, api); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addPolicyStatementFromSdkCall",
		[]interface{}{service, api, resources},
	)
}

func (a *jsiiProxy_AssertionsProvider) AddToRolePolicy(statement interface{}) {
	if err := a.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addToRolePolicy",
		[]interface{}{statement},
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

func (a *jsiiProxy_AssertionsProvider) GrantInvoke(principalArn *string) {
	if err := a.validateGrantInvokeParameters(principalArn); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"grantInvoke",
		[]interface{}{principalArn},
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

