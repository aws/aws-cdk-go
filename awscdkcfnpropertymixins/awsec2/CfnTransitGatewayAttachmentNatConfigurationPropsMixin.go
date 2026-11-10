package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::EC2::TransitGatewayAttachmentNatConfiguration.
//
// Configures static 1:1 NAT on a transit gateway attachment. Each rule translates a source CIDR to a NAT CIDR. NatGatewaySupport controls whether creating this resource activates the rules or stages them; deleting this resource disables NAT gateway support and removes all NAT rules. This resource must be managed from the account that owns the transit gateway. NAT support depends on the attachment type, and the transit gateway rejects a request to configure NAT on an attachment type that does not support it.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayAttachmentNatConfigurationPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayAttachmentNatConfigurationPropsMixin(&CfnTransitGatewayAttachmentNatConfigurationMixinProps{
//   	NatGatewayRules: []interface{}{
//   		&NatRuleProperty{
//   			MatchSourceCidr: jsii.String("matchSourceCidr"),
//   			NatCidr: jsii.String("natCidr"),
//   			RuleNumber: jsii.Number(123),
//   		},
//   	},
//   	NatGatewaySupport: jsii.String("natGatewaySupport"),
//   	NatSystemRouteSupport: jsii.String("natSystemRouteSupport"),
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayattachmentnatconfiguration.html
//
type CfnTransitGatewayAttachmentNatConfigurationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayAttachmentNatConfigurationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayAttachmentNatConfigurationPropsMixin
type jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin) Props() *CfnTransitGatewayAttachmentNatConfigurationMixinProps {
	var returns *CfnTransitGatewayAttachmentNatConfigurationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayAttachmentNatConfiguration`.
func NewCfnTransitGatewayAttachmentNatConfigurationPropsMixin(props *CfnTransitGatewayAttachmentNatConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayAttachmentNatConfigurationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayAttachmentNatConfigurationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayAttachmentNatConfigurationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayAttachmentNatConfiguration`.
func NewCfnTransitGatewayAttachmentNatConfigurationPropsMixin_Override(c CfnTransitGatewayAttachmentNatConfigurationPropsMixin, props *CfnTransitGatewayAttachmentNatConfigurationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayAttachmentNatConfigurationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayAttachmentNatConfigurationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayAttachmentNatConfigurationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayAttachmentNatConfigurationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayAttachmentNatConfigurationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayAttachmentNatConfigurationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayAttachmentNatConfigurationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

