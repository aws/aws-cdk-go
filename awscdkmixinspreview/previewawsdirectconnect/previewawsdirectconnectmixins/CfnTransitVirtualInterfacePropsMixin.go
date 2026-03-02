package previewawsdirectconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdirectconnect/previewawsdirectconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DirectConnect::TransitVirtualInterface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnTransitVirtualInterfacePropsMixin := awscdkmixinspreview.Mixins.NewCfnTransitVirtualInterfacePropsMixin(&CfnTransitVirtualInterfaceMixinProps{
//   	AllocateTransitVirtualInterfaceRoleArn: jsii.String("allocateTransitVirtualInterfaceRoleArn"),
//   	BgpPeers: []interface{}{
//   		&BgpPeerProperty{
//   			AddressFamily: jsii.String("addressFamily"),
//   			AmazonAddress: jsii.String("amazonAddress"),
//   			Asn: jsii.String("asn"),
//   			AuthKey: jsii.String("authKey"),
//   			BgpPeerId: jsii.String("bgpPeerId"),
//   			CustomerAddress: jsii.String("customerAddress"),
//   		},
//   	},
//   	ConnectionId: jsii.String("connectionId"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   	EnableSiteLink: jsii.Boolean(false),
//   	Mtu: jsii.Number(123),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VirtualInterfaceName: jsii.String("virtualInterfaceName"),
//   	Vlan: jsii.Number(123),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-transitvirtualinterface.html
//
type CfnTransitVirtualInterfacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitVirtualInterfaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitVirtualInterfacePropsMixin
type jsiiProxy_CfnTransitVirtualInterfacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitVirtualInterfacePropsMixin) Props() *CfnTransitVirtualInterfaceMixinProps {
	var returns *CfnTransitVirtualInterfaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitVirtualInterfacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DirectConnect::TransitVirtualInterface`.
func NewCfnTransitVirtualInterfacePropsMixin(props *CfnTransitVirtualInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnTransitVirtualInterfacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitVirtualInterfacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitVirtualInterfacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnTransitVirtualInterfacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DirectConnect::TransitVirtualInterface`.
func NewCfnTransitVirtualInterfacePropsMixin_Override(c CfnTransitVirtualInterfacePropsMixin, props *CfnTransitVirtualInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnTransitVirtualInterfacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitVirtualInterfacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitVirtualInterfacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnTransitVirtualInterfacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitVirtualInterfacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnTransitVirtualInterfacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitVirtualInterfacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitVirtualInterfacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

