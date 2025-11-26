package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies a network address translation (NAT) gateway in the specified subnet.
//
// You can create either a public NAT gateway or a private NAT gateway. The default is a public NAT gateway. If you create a public NAT gateway, you must specify an elastic IP address.
//
// With a NAT gateway, instances in a private subnet can connect to the internet, other AWS services, or an on-premises network using the IP address of the NAT gateway. For more information, see [NAT gateways](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html) in the *Amazon VPC User Guide* .
//
// If you add a default route ( `AWS::EC2::Route` resource) that points to a NAT gateway, specify the NAT gateway ID for the route's `NatGatewayId` property.
//
// > When you associate an Elastic IP address or secondary Elastic IP address with a public NAT gateway, the network border group of the Elastic IP address must match the network border group of the Availability Zone (AZ) that the public NAT gateway is in. Otherwise, the NAT gateway fails to launch. You can see the network border group for the AZ by viewing the details of the subnet. Similarly, you can view the network border group for the Elastic IP address by viewing its details. For more information, see [Allocate an Elastic IP address](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#allocate-eip) in the *Amazon VPC User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnNatGatewayPropsMixin := awscdkmixinspreview.Mixins.NewCfnNatGatewayPropsMixin(&CfnNatGatewayMixinProps{
//   	AllocationId: jsii.String("allocationId"),
//   	AvailabilityMode: jsii.String("availabilityMode"),
//   	AvailabilityZoneAddresses: []interface{}{
//   		&AvailabilityZoneAddressProperty{
//   			AllocationIds: []*string{
//   				jsii.String("allocationIds"),
//   			},
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   		},
//   	},
//   	ConnectivityType: jsii.String("connectivityType"),
//   	MaxDrainDurationSeconds: jsii.Number(123),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	SecondaryAllocationIds: []*string{
//   		jsii.String("secondaryAllocationIds"),
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SecondaryPrivateIpAddresses: []*string{
//   		jsii.String("secondaryPrivateIpAddresses"),
//   	},
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html
//
type CfnNatGatewayPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnNatGatewayMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnNatGatewayPropsMixin
type jsiiProxy_CfnNatGatewayPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnNatGatewayPropsMixin) Props() *CfnNatGatewayMixinProps {
	var returns *CfnNatGatewayMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnNatGatewayPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::NatGateway`.
func NewCfnNatGatewayPropsMixin(props *CfnNatGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) CfnNatGatewayPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnNatGatewayPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnNatGatewayPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNatGatewayPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::NatGateway`.
func NewCfnNatGatewayPropsMixin_Override(c CfnNatGatewayPropsMixin, props *CfnNatGatewayMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNatGatewayPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnNatGatewayPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnNatGatewayPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNatGatewayPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnNatGatewayPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnNatGatewayPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnNatGatewayPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnNatGatewayPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

