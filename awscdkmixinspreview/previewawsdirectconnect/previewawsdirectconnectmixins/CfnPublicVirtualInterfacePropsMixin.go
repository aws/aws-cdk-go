package previewawsdirectconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsdirectconnect/previewawsdirectconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DirectConnect::PublicVirtualInterface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPublicVirtualInterfacePropsMixin := awscdkmixinspreview.Mixins.NewCfnPublicVirtualInterfacePropsMixin(&CfnPublicVirtualInterfaceMixinProps{
//   	AllocatePublicVirtualInterfaceRoleArn: jsii.String("allocatePublicVirtualInterfaceRoleArn"),
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
//   	RouteFilterPrefixes: []*string{
//   		jsii.String("routeFilterPrefixes"),
//   	},
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-publicvirtualinterface.html
//
type CfnPublicVirtualInterfacePropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnPublicVirtualInterfaceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnPublicVirtualInterfacePropsMixin
type jsiiProxy_CfnPublicVirtualInterfacePropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnPublicVirtualInterfacePropsMixin) Props() *CfnPublicVirtualInterfaceMixinProps {
	var returns *CfnPublicVirtualInterfaceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnPublicVirtualInterfacePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DirectConnect::PublicVirtualInterface`.
func NewCfnPublicVirtualInterfacePropsMixin(props *CfnPublicVirtualInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnPublicVirtualInterfacePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnPublicVirtualInterfacePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnPublicVirtualInterfacePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnPublicVirtualInterfacePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DirectConnect::PublicVirtualInterface`.
func NewCfnPublicVirtualInterfacePropsMixin_Override(c CfnPublicVirtualInterfacePropsMixin, props *CfnPublicVirtualInterfaceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnPublicVirtualInterfacePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnPublicVirtualInterfacePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnPublicVirtualInterfacePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnPublicVirtualInterfacePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnPublicVirtualInterfacePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_directconnect.mixins.CfnPublicVirtualInterfacePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnPublicVirtualInterfacePropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnPublicVirtualInterfacePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

