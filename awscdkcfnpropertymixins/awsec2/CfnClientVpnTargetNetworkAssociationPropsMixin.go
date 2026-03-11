package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a target network to associate with a Client VPN endpoint.
//
// A target network is a subnet in a VPC. You can associate multiple subnets from the same VPC with a Client VPN endpoint. You can associate only one subnet in each Availability Zone. We recommend that you associate at least two subnets to provide Availability Zone redundancy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnClientVpnTargetNetworkAssociationPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnClientVpnTargetNetworkAssociationPropsMixin(&CfnClientVpnTargetNetworkAssociationMixinProps{
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	SubnetId: jsii.String("subnetId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-clientvpntargetnetworkassociation.html
//
type CfnClientVpnTargetNetworkAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnClientVpnTargetNetworkAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClientVpnTargetNetworkAssociationPropsMixin
type jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin) Props() *CfnClientVpnTargetNetworkAssociationMixinProps {
	var returns *CfnClientVpnTargetNetworkAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::ClientVpnTargetNetworkAssociation`.
func NewCfnClientVpnTargetNetworkAssociationPropsMixin(props *CfnClientVpnTargetNetworkAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnClientVpnTargetNetworkAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClientVpnTargetNetworkAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnTargetNetworkAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::ClientVpnTargetNetworkAssociation`.
func NewCfnClientVpnTargetNetworkAssociationPropsMixin_Override(c CfnClientVpnTargetNetworkAssociationPropsMixin, props *CfnClientVpnTargetNetworkAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnTargetNetworkAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnClientVpnTargetNetworkAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClientVpnTargetNetworkAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnTargetNetworkAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClientVpnTargetNetworkAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnClientVpnTargetNetworkAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnClientVpnTargetNetworkAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

