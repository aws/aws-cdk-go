package previewawsshieldmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsshield/previewawsshieldmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Provides permissions for the AWS Shield Advanced Shield response team (SRT) to access your account and your resource protections, to help you mitigate potential distributed denial of service (DDoS) attacks.
//
// *Configure `AWS::Shield::DRTAccess` for one account*
//
// To configure this resource through CloudFormation , you must be subscribed to AWS Shield Advanced . You can subscribe through the [Shield Advanced console](https://docs.aws.amazon.com/wafv2/shieldv2#/) and through the APIs. For more information, see [Subscribe to AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/enable-ddos-prem.html) .
//
// See example templates for Shield Advanced in CloudFormation at [aws-samples/aws-shield-advanced-examples](https://docs.aws.amazon.com/https://github.com/aws-samples/aws-shield-advanced-examples) .
//
// *Configure Shield Advanced using AWS CloudFormation and AWS Firewall Manager*
//
// You might be able to use Firewall Manager with AWS CloudFormation to configure Shield Advanced across multiple accounts and protected resources. To do this, your accounts must be part of an organization in AWS Organizations . You can use Firewall Manager to configure Shield Advanced protections for any resource types except for Amazon Route 53 or AWS Global Accelerator .
//
// For an example of this, see the one-click configuration guidance published by the AWS technical community at [One-click deployment of Shield Advanced](https://docs.aws.amazon.com/https://youtu.be/LCA3FwMk_QE) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDRTAccessPropsMixin := awscdkmixinspreview.Mixins.NewCfnDRTAccessPropsMixin(&CfnDRTAccessMixinProps{
//   	LogBucketList: []*string{
//   		jsii.String("logBucketList"),
//   	},
//   	RoleArn: jsii.String("roleArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-drtaccess.html
//
type CfnDRTAccessPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDRTAccessMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDRTAccessPropsMixin
type jsiiProxy_CfnDRTAccessPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDRTAccessPropsMixin) Props() *CfnDRTAccessMixinProps {
	var returns *CfnDRTAccessMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDRTAccessPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Shield::DRTAccess`.
func NewCfnDRTAccessPropsMixin(props *CfnDRTAccessMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDRTAccessPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDRTAccessPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDRTAccessPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Shield::DRTAccess`.
func NewCfnDRTAccessPropsMixin_Override(c CfnDRTAccessPropsMixin, props *CfnDRTAccessMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDRTAccessPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDRTAccessPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDRTAccessPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnDRTAccessPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDRTAccessPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDRTAccessPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

