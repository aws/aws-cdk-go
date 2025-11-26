package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a new origin access control in CloudFront.
//
// After you create an origin access control, you can add it to an origin in a CloudFront distribution so that CloudFront sends authenticated (signed) requests to the origin.
//
// This makes it possible to block public access to the origin, allowing viewers (users) to access the origin's content only through CloudFront.
//
// For more information about using a CloudFront origin access control, see [Restricting access to an AWS origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-origin.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOriginAccessControlPropsMixin := awscdkmixinspreview.Mixins.NewCfnOriginAccessControlPropsMixin(&CfnOriginAccessControlMixinProps{
//   	OriginAccessControlConfig: &OriginAccessControlConfigProperty{
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		OriginAccessControlOriginType: jsii.String("originAccessControlOriginType"),
//   		SigningBehavior: jsii.String("signingBehavior"),
//   		SigningProtocol: jsii.String("signingProtocol"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-originaccesscontrol.html
//
type CfnOriginAccessControlPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnOriginAccessControlMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnOriginAccessControlPropsMixin
type jsiiProxy_CfnOriginAccessControlPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnOriginAccessControlPropsMixin) Props() *CfnOriginAccessControlMixinProps {
	var returns *CfnOriginAccessControlMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnOriginAccessControlPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::OriginAccessControl`.
func NewCfnOriginAccessControlPropsMixin(props *CfnOriginAccessControlMixinProps, options *mixins.CfnPropertyMixinOptions) CfnOriginAccessControlPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnOriginAccessControlPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnOriginAccessControlPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::OriginAccessControl`.
func NewCfnOriginAccessControlPropsMixin_Override(c CfnOriginAccessControlPropsMixin, props *CfnOriginAccessControlMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnOriginAccessControlPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnOriginAccessControlPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnOriginAccessControlPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnOriginAccessControlPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnOriginAccessControlPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnOriginAccessControlPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

