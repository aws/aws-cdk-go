package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Registers sources (network interfaces) with the specified transit gateway multicast domain.
//
// A multicast source is a network interface attached to a supported instance that sends multicast traffic. For information about supported instances, see [Multicast Considerations](https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-limits.html#multicast-limits) in *Amazon VPC Transit Gateways* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayMulticastGroupSourcePropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayMulticastGroupSourcePropsMixin(&CfnTransitGatewayMulticastGroupSourceMixinProps{
//   	GroupIpAddress: jsii.String("groupIpAddress"),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	TransitGatewayMulticastDomainId: jsii.String("transitGatewayMulticastDomainId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaymulticastgroupsource.html
//
type CfnTransitGatewayMulticastGroupSourcePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayMulticastGroupSourceMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayMulticastGroupSourcePropsMixin
type jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin) Props() *CfnTransitGatewayMulticastGroupSourceMixinProps {
	var returns *CfnTransitGatewayMulticastGroupSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayMulticastGroupSource`.
func NewCfnTransitGatewayMulticastGroupSourcePropsMixin(props *CfnTransitGatewayMulticastGroupSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayMulticastGroupSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayMulticastGroupSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayMulticastGroupSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayMulticastGroupSource`.
func NewCfnTransitGatewayMulticastGroupSourcePropsMixin_Override(c CfnTransitGatewayMulticastGroupSourcePropsMixin, props *CfnTransitGatewayMulticastGroupSourceMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayMulticastGroupSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayMulticastGroupSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayMulticastGroupSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayMulticastGroupSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayMulticastGroupSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayMulticastGroupSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayMulticastGroupSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

