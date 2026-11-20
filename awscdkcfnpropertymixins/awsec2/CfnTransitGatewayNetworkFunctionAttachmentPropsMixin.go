package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2"
	"github.com/aws/aws-cdk-go/awscdkcfnpropertymixins/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The AWS::EC2::TransitGatewayNetworkFunctionAttachment type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var mergeStrategy IMergeStrategy
//
//   cfnTransitGatewayNetworkFunctionAttachmentPropsMixin := awscdkcfnpropertymixins.Aws_ec2.NewCfnTransitGatewayNetworkFunctionAttachmentPropsMixin(&CfnTransitGatewayNetworkFunctionAttachmentMixinProps{
//   	AvailabilityZoneIds: []*string{
//   		jsii.String("availabilityZoneIds"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	ServiceName: jsii.String("serviceName"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewaynetworkfunctionattachment.html
//
type CfnTransitGatewayNetworkFunctionAttachmentPropsMixin interface {
	awscdk.Mixin
	constructs.IMixin
	Props() *CfnTransitGatewayNetworkFunctionAttachmentMixinProps
	Strategy() awscdk.IMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnTransitGatewayNetworkFunctionAttachmentPropsMixin
type jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin struct {
	internal.Type__awscdkMixin
	internal.Type__constructsIMixin
}

func (j *jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin) Props() *CfnTransitGatewayNetworkFunctionAttachmentMixinProps {
	var returns *CfnTransitGatewayNetworkFunctionAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin) Strategy() awscdk.IMergeStrategy {
	var returns awscdk.IMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::EC2::TransitGatewayNetworkFunctionAttachment`.
func NewCfnTransitGatewayNetworkFunctionAttachmentPropsMixin(props *CfnTransitGatewayNetworkFunctionAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) CfnTransitGatewayNetworkFunctionAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnTransitGatewayNetworkFunctionAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayNetworkFunctionAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::EC2::TransitGatewayNetworkFunctionAttachment`.
func NewCfnTransitGatewayNetworkFunctionAttachmentPropsMixin_Override(c CfnTransitGatewayNetworkFunctionAttachmentPropsMixin, props *CfnTransitGatewayNetworkFunctionAttachmentMixinProps, options *awscdkcfnpropertymixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayNetworkFunctionAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
func CfnTransitGatewayNetworkFunctionAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnTransitGatewayNetworkFunctionAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayNetworkFunctionAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnTransitGatewayNetworkFunctionAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/cfn-property-mixins.aws_ec2.CfnTransitGatewayNetworkFunctionAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnTransitGatewayNetworkFunctionAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

