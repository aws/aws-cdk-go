package previewawscloudfrontmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscloudfront/previewawscloudfrontmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The request to create a new origin access identity (OAI).
//
// An origin access identity is a special CloudFront user that you can associate with Amazon S3 origins, so that you can secure all or just some of your Amazon S3 content. For more information, see [Restricting Access to Amazon S3 Content by Using an Origin Access Identity](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/private-content-restricting-access-to-s3.html) in the *Amazon CloudFront Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCloudFrontOriginAccessIdentityPropsMixin := awscdkmixinspreview.Mixins.NewCfnCloudFrontOriginAccessIdentityPropsMixin(&CfnCloudFrontOriginAccessIdentityMixinProps{
//   	CloudFrontOriginAccessIdentityConfig: &CloudFrontOriginAccessIdentityConfigProperty{
//   		Comment: jsii.String("comment"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-cloudfrontoriginaccessidentity.html
//
type CfnCloudFrontOriginAccessIdentityPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCloudFrontOriginAccessIdentityMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCloudFrontOriginAccessIdentityPropsMixin
type jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin) Props() *CfnCloudFrontOriginAccessIdentityMixinProps {
	var returns *CfnCloudFrontOriginAccessIdentityMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
func NewCfnCloudFrontOriginAccessIdentityPropsMixin(props *CfnCloudFrontOriginAccessIdentityMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCloudFrontOriginAccessIdentityPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCloudFrontOriginAccessIdentityPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CloudFront::CloudFrontOriginAccessIdentity`.
func NewCfnCloudFrontOriginAccessIdentityPropsMixin_Override(c CfnCloudFrontOriginAccessIdentityPropsMixin, props *CfnCloudFrontOriginAccessIdentityMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCloudFrontOriginAccessIdentityPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCloudFrontOriginAccessIdentityPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCloudFrontOriginAccessIdentityPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_cloudfront.mixins.CfnCloudFrontOriginAccessIdentityPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCloudFrontOriginAccessIdentityPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

