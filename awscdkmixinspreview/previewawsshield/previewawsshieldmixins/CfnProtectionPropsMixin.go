package previewawsshieldmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsshield/previewawsshieldmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Enables AWS Shield Advanced for a specific AWS resource.
//
// The resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone, AWS Global Accelerator standard accelerator, Elastic IP Address, Application Load Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and Network Load Balancers by association with protected Amazon EC2 Elastic IP addresses.
//
// *Configure a single `AWS::Shield::Protection`*
//
// Use this protection to protect a single resource at a time.
//
// To configure this Shield Advanced protection through CloudFormation , you must be subscribed to Shield Advanced . You can subscribe through the [Shield Advanced console](https://docs.aws.amazon.com/wafv2/shieldv2#/) and through the APIs. For more information, see [Subscribe to AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/enable-ddos-prem.html) .
//
// See example templates for Shield Advanced in CloudFormation at [aws-samples/aws-shield-advanced-examples](https://docs.aws.amazon.com/https://github.com/aws-samples/aws-shield-advanced-examples) .
//
// *Configure Shield Advanced using AWS CloudFormation and AWS Firewall Manager*
//
// You might be able to use Firewall Manager with AWS CloudFormation to configure Shield Advanced across multiple accounts and protected resources. To do this, your accounts must be part of an organization in AWS Organizations . You can use Firewall Manager to configure Shield Advanced protections for any resource types except for Amazon Route 53 or AWS Global Accelerator .
//
// For an example of this, see the one-click configuration guidance published by the AWS technical community at [One-click deployment of Shield Advanced](https://docs.aws.amazon.com/https://youtu.be/LCA3FwMk_QE) .
//
// *Configure multiple protections through the Shield Advanced console*
//
// You can add protection to multiple resources at once through the [Shield Advanced console](https://docs.aws.amazon.com/wafv2/shieldv2#/) . For more information see [Getting Started with AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html) and [Managing resource protections in AWS Shield Advanced](https://docs.aws.amazon.com/waf/latest/developerguide/ddos-manage-protected-resources.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var block interface{}
//   var count interface{}
//
//   cfnProtectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnProtectionPropsMixin(&CfnProtectionMixinProps{
//   	ApplicationLayerAutomaticResponseConfiguration: &ApplicationLayerAutomaticResponseConfigurationProperty{
//   		Action: &ActionProperty{
//   			Block: block,
//   			Count: count,
//   		},
//   		Status: jsii.String("status"),
//   	},
//   	HealthCheckArns: []*string{
//   		jsii.String("healthCheckArns"),
//   	},
//   	Name: jsii.String("name"),
//   	ResourceArn: jsii.String("resourceArn"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-shield-protection.html
//
type CfnProtectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnProtectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnProtectionPropsMixin
type jsiiProxy_CfnProtectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnProtectionPropsMixin) Props() *CfnProtectionMixinProps {
	var returns *CfnProtectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnProtectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::Shield::Protection`.
func NewCfnProtectionPropsMixin(props *CfnProtectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnProtectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnProtectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnProtectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::Shield::Protection`.
func NewCfnProtectionPropsMixin_Override(c CfnProtectionPropsMixin, props *CfnProtectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnProtectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnProtectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnProtectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_shield.mixins.CfnProtectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnProtectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnProtectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

