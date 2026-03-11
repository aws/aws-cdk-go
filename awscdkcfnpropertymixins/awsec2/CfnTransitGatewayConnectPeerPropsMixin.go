package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a transit gateway Connect peer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayConnectPeerPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayConnectPeerPropsMixin(&CfnTransitGatewayConnectPeerMixinProps{
//   	ConnectPeerConfiguration: &TransitGatewayConnectPeerConfigurationProperty{
//   		BgpConfigurations: []interface{}{
//   			&TransitGatewayAttachmentBgpConfigurationProperty{
//   				BgpStatus: jsii.String("bgpStatus"),
//   				PeerAddress: jsii.String("peerAddress"),
//   				PeerAsn: jsii.Number(123),
//   				TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   				TransitGatewayAsn: jsii.Number(123),
//   			},
//   		},
//   		InsideCidrBlocks: []*string{
//   			jsii.String("insideCidrBlocks"),
//   		},
//   		PeerAddress: jsii.String("peerAddress"),
//   		Protocol: jsii.String("protocol"),
//   		TransitGatewayAddress: jsii.String("transitGatewayAddress"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html
//
type CfnTransitGatewayConnectPeerPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayConnectPeerMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayConnectPeerPropsMixin
type jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) Props() *CfnTransitGatewayConnectPeerMixinProps {
	var returns *CfnTransitGatewayConnectPeerMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayConnectPeer`.
func NewCfnTransitGatewayConnectPeerPropsMixin(props *CfnTransitGatewayConnectPeerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayConnectPeerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayConnectPeerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayConnectPeerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayConnectPeer`.
func NewCfnTransitGatewayConnectPeerPropsMixin_Override(c CfnTransitGatewayConnectPeerPropsMixin, props *CfnTransitGatewayConnectPeerMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayConnectPeerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayConnectPeerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayConnectPeerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayConnectPeerPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayConnectPeerPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayConnectPeerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

