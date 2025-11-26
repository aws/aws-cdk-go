package previewawsnetworkmanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkmanager/previewawsnetworkmanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Specifies an association between a customer gateway, a device, and optionally, a link.
//
// If you specify a link, it must be associated with the specified device. The customer gateway must be connected to a VPN attachment on a transit gateway that's registered in your global network.
//
// You cannot associate a customer gateway with more than one device and link.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCustomerGatewayAssociationPropsMixin := awscdkmixinspreview.Mixins.NewCfnCustomerGatewayAssociationPropsMixin(&CfnCustomerGatewayAssociationMixinProps{
//   	CustomerGatewayArn: jsii.String("customerGatewayArn"),
//   	DeviceId: jsii.String("deviceId"),
//   	GlobalNetworkId: jsii.String("globalNetworkId"),
//   	LinkId: jsii.String("linkId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-customergatewayassociation.html
//
type CfnCustomerGatewayAssociationPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCustomerGatewayAssociationMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCustomerGatewayAssociationPropsMixin
type jsiiProxy_CfnCustomerGatewayAssociationPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCustomerGatewayAssociationPropsMixin) Props() *CfnCustomerGatewayAssociationMixinProps {
	var returns *CfnCustomerGatewayAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCustomerGatewayAssociationPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkManager::CustomerGatewayAssociation`.
func NewCfnCustomerGatewayAssociationPropsMixin(props *CfnCustomerGatewayAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCustomerGatewayAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCustomerGatewayAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCustomerGatewayAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkManager::CustomerGatewayAssociation`.
func NewCfnCustomerGatewayAssociationPropsMixin_Override(c CfnCustomerGatewayAssociationPropsMixin, props *CfnCustomerGatewayAssociationMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCustomerGatewayAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCustomerGatewayAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCustomerGatewayAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnCustomerGatewayAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCustomerGatewayAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnCustomerGatewayAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

