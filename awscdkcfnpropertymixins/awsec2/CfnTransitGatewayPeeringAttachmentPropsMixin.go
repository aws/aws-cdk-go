package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Requests a transit gateway peering attachment between the specified transit gateway (requester) and a peer transit gateway (accepter).
//
// The peer transit gateway can be in your account or a different AWS account .
//
// After you create the peering attachment, the owner of the accepter transit gateway must accept the attachment request.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayPeeringAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayPeeringAttachmentPropsMixin(&CfnTransitGatewayPeeringAttachmentMixinProps{
//   	PeerAccountId: jsii.String("peerAccountId"),
//   	PeerRegion: jsii.String("peerRegion"),
//   	PeerTransitGatewayId: jsii.String("peerTransitGatewayId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: mergeStrategy,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaypeeringattachment.html
//
type CfnTransitGatewayPeeringAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayPeeringAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayPeeringAttachmentPropsMixin
type jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin) Props() *CfnTransitGatewayPeeringAttachmentMixinProps {
	var returns *CfnTransitGatewayPeeringAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayPeeringAttachment`.
func NewCfnTransitGatewayPeeringAttachmentPropsMixin(props *CfnTransitGatewayPeeringAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayPeeringAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayPeeringAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayPeeringAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayPeeringAttachment`.
func NewCfnTransitGatewayPeeringAttachmentPropsMixin_Override(c CfnTransitGatewayPeeringAttachmentPropsMixin, props *CfnTransitGatewayPeeringAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayPeeringAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayPeeringAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayPeeringAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayPeeringAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayPeeringAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayPeeringAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayPeeringAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

