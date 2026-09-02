package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Use the ``AWS::Lambda::ResourcePolicy`` resource to attach a resource-based policy to a LAM resource.
//
// A resource-based policy applies to a single LAM resource, for example, a function, function version, or function alias. To learn more about using resource-based policies with LAM, see [Working with resource-based policies in](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) in the *Developer Guide*.
//  You can use resource-based policies to grant permissions to other AWS services, AWS accounts and organizations, and IAM users and roles to access your LAM resource. You can also deny access to specific entities, and use the full range of IAM global condition keys to further restrict who has access to your LAM resource. For example, you can limit access to calls originating from a specified IP address or VPC.
//  A resource-based policy is a JSON document containing a number of statements. Each statement defines the entities you want to grant permission to, the API actions you want to allow or deny, and the LAM resource you want the statement to apply to. A statement can also optionally include an array of logical conditions using the IAM global condition keys.
//  To use the ``AWS::Lambda::ResourcePolicy`` resource, make sure that you have the [resource-based policy permissions for Lambda](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html#access-control-resource-based-permissions).
//  To learn more about creating resource-based policies, see [Policies and permissions in](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html) in the *User Guide*. For more information about example policies for providing permissions to AWS services, other AWS accounts, and IAM users and roles, see [Example resource-based policies for functions](https://docs.aws.amazon.com/lambda/latest/dg/permissions-function-examples.html) in the *Developer Guide*.
//   **Avoid mixing permission resource types**
//  To grant permissions to access your function, we recommend using the ``AWS::Lambda::ResourcePolicy`` resource to set access permissions. With this resource, you have more flexibility and fine-grained control than ``AWS::Lambda::Permission``. This resource grants an AWS service or another account permission to call a particular API action on a function.
//  You can also use the ``AWS::Lambda::Permission`` resource, however using both ``AWS::Lambda::Permission`` and ``AWS::Lambda::ResourcePolicy`` to set permissions on a function can result in errors. Permissions defined in ``AWS::Lambda::Permission`` can be unintentionally overwritten, whether in a single CFN stack or across multiple stacks. Don't use both resource types to set permissions on a function.
//  To migrate existing permissions for a function from ``AWS::Lambda::Permission`` to ``AWS::Lambda::ResourcePolicy``, do the following:
//   1.  Set a ``Retain``[deletion policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) on the ``AWS::Lambda::Permission`` resources you want to migrate. This is necessary so that Lambda does not delete statements with the same statement ID when you delete these resources.
//   1.  Use the [GetResourcePolicy](https://docs.aws.amazon.com/lambda/latest/api/API_GetResourcePolicy.html)LAM API to retrieve the resource-based policy currently attached to the function.
//   1.  Use this policy to create a new ``AWS::Lambda::ResourcePolicy`` resource.
//   1.  Delete all the existing ``AWS::Lambda::Permission`` resources for the function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//   var policyDocument interface{}
//
//   cfnResourcePolicyPropsMixin := awscdkcfnpropertymixins.Aws_lambda.NewCfnResourcePolicyPropsMixin(&CfnResourcePolicyMixinProps{
//   	PolicyDocument: policyDocument,
//   	ResourceArn: jsii.String("resourceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-resourcepolicy.html
//
type CfnResourcePolicyPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnResourcePolicyMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnResourcePolicyPropsMixin
type jsiiProxy_CfnResourcePolicyPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Props() *CfnResourcePolicyMixinProps {
	var returns *CfnResourcePolicyMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnResourcePolicyPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin(props *CfnResourcePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnResourcePolicyPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnResourcePolicyPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnResourcePolicyPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::ResourcePolicy`.
func NewCfnResourcePolicyPropsMixin_Override(c CfnResourcePolicyPropsMixin, props *CfnResourcePolicyMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnResourcePolicyPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnResourcePolicyPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnResourcePolicyPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnResourcePolicyPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnResourcePolicyPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_lambda.CfnResourcePolicyPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnResourcePolicyPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

