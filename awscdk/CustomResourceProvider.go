package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// An AWS-Lambda backed custom resource provider, for CDK Construct Library constructs.
//
// This is a provider for `CustomResource` constructs, backed by an AWS Lambda
// Function. It only supports NodeJS runtimes.
//
// > **Application builders do not need to use this provider type**. This is not
// > a generic custom resource provider class. It is specifically
// > intended to be used only by constructs in the AWS CDK Construct Library, and
// > only exists here because of reverse dependency issues (for example, it cannot
// > use `iam.PolicyStatement` objects, since the `iam` library already depends on
// > the CDK `core` library and we cannot have cyclic dependencies).
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
//   	Runtime: awscdk.CustomResourceProviderRuntime_NODEJS_18_X,
//   })
//   provider.AddToRolePolicy(map[string]*string{
//   	"Effect": jsii.String("Allow"),
//   	"Action": jsii.String("s3:GetObject"),
//   	"Resource": jsii.String("*"),
//   })
//
type CustomResourceProvider interface {
	CustomResourceProviderBase
	// The hash of the lambda code backing this provider.
	//
	// Can be used to trigger updates
	// on code changes, even when the properties of a custom resource remain unchanged.
	CodeHash() *string
	// The tree node.
	Node() constructs.Node
	// The ARN of the provider's AWS Lambda function role.
	RoleArn() *string
	// The ARN of the provider's AWS Lambda function which should be used as the `serviceToken` when defining a custom resource.
	ServiceToken() *string
	// Add an IAM policy statement to the inline policy of the provider's lambda function's role.
	//
	// **Please note**: this is a direct IAM JSON policy blob, *not* a `iam.PolicyStatement`
	// object like you will see in the rest of the CDK.
	//
	// Example:
	//   declare const myProvider: CustomResourceProvider;
	//
	//   myProvider.addToRolePolicy({
	//     Effect: 'Allow',
	//     Action: 's3:GetObject',
	//     Resource: '*',
	//   });
	//
	AddToRolePolicy(statement interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomResourceProvider
type jsiiProxy_CustomResourceProvider struct {
	jsiiProxy_CustomResourceProviderBase
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

func (j *jsiiProxy_CustomResourceProvider) Node() constructs.Node {
	var returns constructs.Node
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


func NewCustomResourceProvider(scope constructs.Construct, id *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	if err := validateNewCustomResourceProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomResourceProvider{}

	_jsii_.Create(
		"aws-cdk-lib.CustomResourceProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewCustomResourceProvider_Override(c CustomResourceProvider, scope constructs.Construct, id *string, props *CustomResourceProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CustomResourceProvider",
		[]interface{}{scope, id, props},
		c,
	)
}

// Returns a stack-level singleton ARN (service token) for the custom resource provider.
//
// Returns: the service token of the custom resource provider, which should be
// used when defining a `CustomResource`.
func CustomResourceProvider_GetOrCreate(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) *string {
	_init_.Initialize()

	if err := validateCustomResourceProvider_GetOrCreateParameters(scope, uniqueid, props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
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
func CustomResourceProvider_GetOrCreateProvider(scope constructs.Construct, uniqueid *string, props *CustomResourceProviderProps) CustomResourceProvider {
	_init_.Initialize()

	if err := validateCustomResourceProvider_GetOrCreateProviderParameters(scope, uniqueid, props); err != nil {
		panic(err)
	}
	var returns CustomResourceProvider

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
		"getOrCreateProvider",
		[]interface{}{scope, uniqueid, props},
		&returns,
	)

	return returns
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
func CustomResourceProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomResourceProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProvider) AddToRolePolicy(statement interface{}) {
	if err := c.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToRolePolicy",
		[]interface{}{statement},
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

