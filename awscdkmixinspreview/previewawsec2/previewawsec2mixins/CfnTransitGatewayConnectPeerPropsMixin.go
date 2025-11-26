package previewawsec2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsec2/previewawsec2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Describes a transit gateway Connect peer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitGatewayConnectPeerPropsMixin := awscdkmixinspreview.Mixins.NewCfnTransitGatewayConnectPeerPropsMixin(&CfnTransitGatewayConnectPeerMixinProps{
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
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayconnectpeer.html
//
type CfnTransitGatewayConnectPeerPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnTransitGatewayConnectPeerMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayConnectPeerPropsMixin
type jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
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

func (j *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayConnectPeer`.
func NewCfnTransitGatewayConnectPeerPropsMixin(props *CfnTransitGatewayConnectPeerMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTransitGatewayConnectPeerPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayConnectPeerPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayConnectPeerPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayConnectPeer`.
func NewCfnTransitGatewayConnectPeerPropsMixin_Override(c CfnTransitGatewayConnectPeerPropsMixin, props *CfnTransitGatewayConnectPeerMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayConnectPeerPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnTransitGatewayConnectPeerPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayConnectPeerPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayConnectPeerPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_ec2.mixins.CfnTransitGatewayConnectPeerPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayConnectPeerPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

