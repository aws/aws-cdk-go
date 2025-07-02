package awscdk

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Base class for creating a custom resource provider.
type CustomResourceProviderBase interface {
	constructs.Construct
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
	//   var myProvider customResourceProvider
	//
	//
	//   myProvider.AddToRolePolicy(map[string]*string{
	//   	"Effect": jsii.String("Allow"),
	//   	"Action": jsii.String("s3:GetObject"),
	//   	"Resource": jsii.String("*"),
	//   })
	//
	AddToRolePolicy(statement interface{})
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for CustomResourceProviderBase
type jsiiProxy_CustomResourceProviderBase struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_CustomResourceProviderBase) CodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProviderBase) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProviderBase) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomResourceProviderBase) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


func NewCustomResourceProviderBase_Override(c CustomResourceProviderBase, scope constructs.Construct, id *string, props *CustomResourceProviderBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.CustomResourceProviderBase",
		[]interface{}{scope, id, props},
		c,
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
func CustomResourceProviderBase_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomResourceProviderBase_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.CustomResourceProviderBase",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomResourceProviderBase) AddToRolePolicy(statement interface{}) {
	if err := c.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (c *jsiiProxy_CustomResourceProviderBase) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

