package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Requests a VPC peering connection between two VPCs: a requester VPC that you own and an accepter VPC with which to create the connection.
//
// The accepter VPC can belong to a different AWS account and can be in a different Region than the requester VPC.
//
// The requester VPC and accepter VPC cannot have overlapping CIDR blocks. If you create a VPC peering connection request between VPCs with overlapping CIDR blocks, the VPC peering connection has a status of `failed` .
//
// If the VPCs belong to different accounts, the acceptor account must have a role that allows the requester account to accept the VPC peering connection. For an example, see [Walkthrough: Peer with a VPC in another AWS account](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/peer-with-vpc-in-another-account.html) .
//
// If the requester and acceptor VPCs are in the same account, the peering request is accepted without a peering role.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCPeeringConnectionPropsMixin := awscdkmixinspreview.Mixins.NewCfnVPCPeeringConnectionPropsMixin(&CfnVPCPeeringConnectionMixinProps{
//   	PeerOwnerId: jsii.String("peerOwnerId"),
//   	PeerRegion: jsii.String("peerRegion"),
//   	PeerRoleArn: jsii.String("peerRoleArn"),
//   	PeerVpcId: jsii.String("peerVpcId"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
//
type CfnVPCPeeringConnectionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnVPCPeeringConnectionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnVPCPeeringConnectionPropsMixin
type jsiiProxy_CfnVPCPeeringConnectionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) Props() *CfnVPCPeeringConnectionMixinProps {
	var returns *CfnVPCPeeringConnectionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::VPCPeeringConnection`.
func NewCfnVPCPeeringConnectionPropsMixin(props *CfnVPCPeeringConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnVPCPeeringConnectionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnVPCPeeringConnectionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnVPCPeeringConnectionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCPeeringConnectionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::VPCPeeringConnection`.
func NewCfnVPCPeeringConnectionPropsMixin_Override(c CfnVPCPeeringConnectionPropsMixin, props *CfnVPCPeeringConnectionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCPeeringConnectionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnVPCPeeringConnectionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnVPCPeeringConnectionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCPeeringConnectionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnVPCPeeringConnectionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnVPCPeeringConnectionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnVPCPeeringConnectionPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

