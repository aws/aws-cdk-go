package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Associates the specified attachment with the specified transit gateway route table.
//
// You can associate one route table with an attachment.
//
// Before you can update the route table associated with an attachment, you must disassociate the transit gateway route table that is currently associated with the attachment. First update the stack to remove the associated transit gateway route table, and then update the stack with the ID of the new transit gateway route table to associate. In addition, the attachment must be in an `available` state; otherwise, the request will return an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayRouteTableAssociationPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayRouteTableAssociationPropsMixin(&CfnTransitGatewayRouteTableAssociationMixinProps{
//   	TransitGatewayAttachmentId: jsii.String("transitGatewayAttachmentId"),
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayroutetableassociation.html
//
type CfnTransitGatewayRouteTableAssociationPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayRouteTableAssociationMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayRouteTableAssociationPropsMixin
type jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin) Props() *CfnTransitGatewayRouteTableAssociationMixinProps {
	var returns *CfnTransitGatewayRouteTableAssociationMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayRouteTableAssociation`.
func NewCfnTransitGatewayRouteTableAssociationPropsMixin(props *CfnTransitGatewayRouteTableAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayRouteTableAssociationPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayRouteTableAssociationPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayRouteTableAssociationPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayRouteTableAssociation`.
func NewCfnTransitGatewayRouteTableAssociationPropsMixin_Override(c CfnTransitGatewayRouteTableAssociationPropsMixin, props *CfnTransitGatewayRouteTableAssociationMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayRouteTableAssociationPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayRouteTableAssociationPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayRouteTableAssociationPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayRouteTableAssociationPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayRouteTableAssociationPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayRouteTableAssociationPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayRouteTableAssociationPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

