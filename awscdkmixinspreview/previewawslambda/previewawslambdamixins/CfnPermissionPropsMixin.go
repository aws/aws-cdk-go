package previewawslambdamixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawslambda/previewawslambdamixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::Lambda::Permission` resource grants an AWS service or another account permission to use a function.
//
// You can apply the policy at the function level, or specify a qualifier to restrict access to a single version or alias. If you use a qualifier, the invoker must use the full Amazon Resource Name (ARN) of that version or alias to invoke the function.
//
// To grant permission to another account, specify the account ID as the `Principal` . To grant permission to an organization defined in AWS Organizations , specify the organization ID as the `PrincipalOrgID` . For AWS services, the principal is a domain-style identifier defined by the service, like `s3.amazonaws.com` or `sns.amazonaws.com` . For AWS services, you can also specify the ARN of the associated resource as the `SourceArn` . If you grant permission to a service principal without specifying the source, other accounts could potentially configure resources in their account to invoke your Lambda function.
//
// If your function has a function URL, you can specify the `FunctionUrlAuthType` parameter. This adds a condition to your permission that only applies when your function URL's `AuthType` matches the specified `FunctionUrlAuthType` . For more information about the `AuthType` parameter, see [Control access to Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html) .
//
// This resource adds a statement to a resource-based permission policy for the function. For more information about function policies, see [Lambda Function Policies](https://docs.aws.amazon.com/lambda/latest/dg/access-control-resource-based.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPermissionPropsMixin := awscdkmixinspreview.Mixins.NewCfnPermissionPropsMixin(&CfnPermissionMixinProps{
//   	Action: jsii.String("action"),
//   	EventSourceToken: jsii.String("eventSourceToken"),
//   	FunctionName: jsii.String("functionName"),
//   	FunctionUrlAuthType: jsii.String("functionUrlAuthType"),
//   	InvokedViaFunctionUrl: jsii.Boolean(false),
//   	Principal: jsii.String("principal"),
//   	PrincipalOrgId: jsii.String("principalOrgId"),
//   	SourceAccount: jsii.String("sourceAccount"),
//   	SourceArn: jsii.String("sourceArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-permission.html
//
type CfnPermissionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPermissionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPermissionPropsMixin
type jsiiProxy_CfnPermissionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPermissionPropsMixin) Props() *CfnPermissionMixinProps {
	var returns *CfnPermissionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPermissionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Lambda::Permission`.
func NewCfnPermissionPropsMixin(props *CfnPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPermissionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPermissionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPermissionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Lambda::Permission`.
func NewCfnPermissionPropsMixin_Override(c CfnPermissionPropsMixin, props *CfnPermissionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPermissionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPermissionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPermissionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_lambda.mixins.CfnPermissionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPermissionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnPermissionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

