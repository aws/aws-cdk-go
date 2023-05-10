// An experiment to bundle the entire CDK into a single module
package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v3"
)

// An AWS-Lambda backed custom resource provider, for CDK Construct Library constructs.
//
// This is a provider for `CustomResource` constructs, backed by an AWS Lambda
// Function. It only supports NodeJS runtimes.
//
// **This is not a generic custom resource provider class**. It is specifically
// intended to be used only by constructs in the AWS CDK Construct Library, and
// only exists here because of reverse dependency issues (for example, it cannot
// use `iam.PolicyStatement` objects, since the `iam` library already depends on
// the CDK `core` library and we cannot have cyclic dependencies).
//
// If you are not writing constructs for the AWS Construct Library, you should
// use the `Provider` class in the `custom-resources` module instead, which has
// a better API and supports all Lambda runtimes, not just Node.
//
// N.B.: When you are writing Custom Resource Providers, there are a number of
// lifecycle events you have to pay attention to. These are documented in the
// README of the `custom-resources` module. Be sure to give the documentation
// in that module a read, regardless of whether you end up using the Provider
// class in there or this one.
//
// Example:
//   provider := awscdk.CustomResourceProvider_GetOrCreateProvider(this, jsii.String("Custom::MyCustomResourceType"), &CustomResourceProviderProps{
//   	CodeDirectory: fmt.Sprintf("%v/my-handler", __dirname),
//   	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_14_X,
//   	PolicyStatements: []interface{}{
//   		map[string]*string{
//   			"Effect": jsii.String("Allow"),
//   			"Action": jsii.String("s3:PutObject*"),
//   			"Resource": jsii.String("*"),
//   		},
//   	},
//   })
//
// Experimental.
type CustomResourceProvider interface {
	Construct
	// The hash of the lambda code backing this provider.
	//
	// Can be used to trigger updates
	// on code changes, even when the properties of a custom resource remain unchanged.
	// Experimental.
	CodeHash() *string
	// The construct tree node associated with this construct.
	// Experimental.
	Node() ConstructNode
	// The ARN of the provider's AWS Lambda function role.
	// Experimental.
	RoleArn() *string
	// The ARN of the provider's AWS Lambda function which should be used as the `serviceToken` when defining a custom resource.
	//
	// Example:
	//   var myProvider customResourceProvider
	//
	//
	//   awscdk.NewCustomResource(this, jsii.String("MyCustomResource"), &CustomResourceProps{
	//   	ServiceToken: myProvider.ServiceToken,
	//   	Properties: map[string]interface{}{
	//   		"myPropertyOne": jsii.String("one"),
	//   		"myPropertyTwo": jsii.String("two"),
	//   	},
	//   })
	//
	// Experimental.
	ServiceToken() *string
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
	Synthesize(session ISynthesisSession)
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

// The jsii proxy struct for CustomResourceProvider
type jsiiProxy_CustomResourceProvider struct {
	jsiiProxy_Construct
}

func (j *jsiiProxy_CustomResourceProvider) CodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProvider) Node() ConstructNode {
	var returns ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProvider) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomResourceProvider(scope constructs.Construct, id *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	if err := validateNewCustomResourceProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResourceProvider{}

	_jsii_.Create(
		"monocdk.CustomResourceProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomResourceProvider_Override(c CustomResourceProvider, scope constructs.Construct, id *string, props *CustomResourceProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.CustomResourceProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns a stack-level singleton ARN (service token) for the custom resource provider.
//
// Returns: the service token of the custom resource provider, which should be
// used when defining a `CustomResource`.
// Experimental.
func CustomResourceProvider_GetOrCreate(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) *string {
	_init_.Initialize()

	if err := validateCustomResourceProvider_GetOrCreateParameters(scope, uniqueid, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"monocdk.CustomResourceProvider",
		"getOrCreate",
		[]interface{}{scope, uniqueid, props},
		&returns,
	)

	return returns
}

// Returns a stack-level singleton for the custom resource provider.
//
// Returns: the service token of the custom resource provider, which should be
// used when defining a `CustomResource`.
// Experimental.
func CustomResourceProvider_GetOrCreateProvider(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_GetOrCreateProviderParameters(scope, uniqueid, props); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"monocdk.CustomResourceProvider",
		"getOrCreateProvider",
		[]interface{}{scope, uniqueid, props},
		&returns,
	)

	return returns
}

// Return whether the given object is a Construct.
// Experimental.
func CustomResourceProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomResourceProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.CustomResourceProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) OnPrepare() {
	_jsii_.InvokeVoid(
		c,
		"onPrepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomResourceProvider) OnSynthesize(session constructs.ISynthesisSession) {
	if err := c.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomResourceProvider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) Prepare() {
	_jsii_.InvokeVoid(
		c,
		"prepare",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomResourceProvider) Synthesize(session ISynthesisSession) {
	if err := c.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"synthesize",
		[]interface{}{session},
	)
}

func (c *jsiiProxy_CustomResourceProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}

