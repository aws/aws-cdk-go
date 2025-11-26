package previewawss3mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawss3/previewawss3mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::S3::AccessGrantInstance` resource creates an S3 Access Grants instance, which serves as a logical grouping for access grants.
//
// You can create one S3 Access Grants instance per Region per account.
//
// - **Permissions** - You must have the `s3:CreateAccessGrantsInstance` permission to use this resource.
// - **Additional Permissions** - To associate an IAM Identity Center instance with your S3 Access Grants instance, you must also have the `sso:DescribeInstance` , `sso:CreateApplication` , `sso:PutApplicationGrant` , and `sso:PutApplicationAuthenticationMethod` permissions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAccessGrantsInstancePropsMixin := awscdkmixinspreview.Mixins.NewCfnAccessGrantsInstancePropsMixin(&CfnAccessGrantsInstanceMixinProps{
//   	IdentityCenterArn: jsii.String("identityCenterArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3-accessgrantsinstance.html
//
type CfnAccessGrantsInstancePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnAccessGrantsInstanceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnAccessGrantsInstancePropsMixin
type jsiiProxy_CfnAccessGrantsInstancePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnAccessGrantsInstancePropsMixin) Props() *CfnAccessGrantsInstanceMixinProps {
	var returns *CfnAccessGrantsInstanceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnAccessGrantsInstancePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::S3::AccessGrantsInstance`.
func NewCfnAccessGrantsInstancePropsMixin(props *CfnAccessGrantsInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnAccessGrantsInstancePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnAccessGrantsInstancePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnAccessGrantsInstancePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstancePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::S3::AccessGrantsInstance`.
func NewCfnAccessGrantsInstancePropsMixin_Override(c CfnAccessGrantsInstancePropsMixin, props *CfnAccessGrantsInstanceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstancePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnAccessGrantsInstancePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnAccessGrantsInstancePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstancePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnAccessGrantsInstancePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_s3.mixins.CfnAccessGrantsInstancePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnAccessGrantsInstancePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnAccessGrantsInstancePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

