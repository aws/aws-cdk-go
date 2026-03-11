package awsdirectconnect

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsdirectconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Resource Type definition for AWS::DirectConnect::DirectConnectGatewayAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnDirectConnectGatewayAssociationPropsMixin := awscdkcfnpropertymixins.Aws_directconnect.NewCfnDirectConnectGatewayAssociationPropsMixin(&CfnDirectConnectGatewayAssociationMixinProps{
//   	AcceptDirectConnectGatewayAssociationProposalRoleArn: jsii.String("acceptDirectConnectGatewayAssociationProposalRoleArn"),
//   	AllowedPrefixesToDirectConnectGateway: []*string{
//   		jsii.String("allowedPrefixesToDirectConnectGateway"),
//   	},
//   	AssociatedGatewayId: jsii.String("associatedGatewayId"),
//   	DirectConnectGatewayId: jsii.String("directConnectGatewayId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directconnect-directconnectgatewayassociation.html
//
type CfnDirectConnectGatewayAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnDirectConnectGatewayAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDirectConnectGatewayAssociationPropsMixin
type jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin) Props() *CfnDirectConnectGatewayAssociationMixinProps {
	var returns *CfnDirectConnectGatewayAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::DirectConnect::DirectConnectGatewayAssociation`.
func NewCfnDirectConnectGatewayAssociationPropsMixin(props *CfnDirectConnectGatewayAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnDirectConnectGatewayAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDirectConnectGatewayAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_directconnect.CfnDirectConnectGatewayAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::DirectConnect::DirectConnectGatewayAssociation`.
func NewCfnDirectConnectGatewayAssociationPropsMixin_Override(c CfnDirectConnectGatewayAssociationPropsMixin, props *CfnDirectConnectGatewayAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_directconnect.CfnDirectConnectGatewayAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnDirectConnectGatewayAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectConnectGatewayAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_directconnect.CfnDirectConnectGatewayAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDirectConnectGatewayAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_directconnect.CfnDirectConnectGatewayAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDirectConnectGatewayAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

