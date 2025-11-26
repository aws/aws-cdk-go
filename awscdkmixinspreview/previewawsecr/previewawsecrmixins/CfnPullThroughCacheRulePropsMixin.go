package previewawsecrmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecr/previewawsecrmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ECR::PullThroughCacheRule` resource creates or updates a pull through cache rule.
//
// A pull through cache rule provides a way to cache images from an upstream registry in your Amazon ECR private registry.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPullThroughCacheRulePropsMixin := awscdkmixinspreview.Mixins.NewCfnPullThroughCacheRulePropsMixin(&CfnPullThroughCacheRuleMixinProps{
//   	CredentialArn: jsii.String("credentialArn"),
//   	CustomRoleArn: jsii.String("customRoleArn"),
//   	EcrRepositoryPrefix: jsii.String("ecrRepositoryPrefix"),
//   	UpstreamRegistry: jsii.String("upstreamRegistry"),
//   	UpstreamRegistryUrl: jsii.String("upstreamRegistryUrl"),
//   	UpstreamRepositoryPrefix: jsii.String("upstreamRepositoryPrefix"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-pullthroughcacherule.html
//
type CfnPullThroughCacheRulePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnPullThroughCacheRuleMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPullThroughCacheRulePropsMixin
type jsiiProxy_CfnPullThroughCacheRulePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnPullThroughCacheRulePropsMixin) Props() *CfnPullThroughCacheRuleMixinProps {
	var returns *CfnPullThroughCacheRuleMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPullThroughCacheRulePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECR::PullThroughCacheRule`.
func NewCfnPullThroughCacheRulePropsMixin(props *CfnPullThroughCacheRuleMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPullThroughCacheRulePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPullThroughCacheRulePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPullThroughCacheRulePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullThroughCacheRulePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECR::PullThroughCacheRule`.
func NewCfnPullThroughCacheRulePropsMixin_Override(c CfnPullThroughCacheRulePropsMixin, props *CfnPullThroughCacheRuleMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullThroughCacheRulePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnPullThroughCacheRulePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPullThroughCacheRulePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullThroughCacheRulePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPullThroughCacheRulePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecr.mixins.CfnPullThroughCacheRulePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPullThroughCacheRulePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnPullThroughCacheRulePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

