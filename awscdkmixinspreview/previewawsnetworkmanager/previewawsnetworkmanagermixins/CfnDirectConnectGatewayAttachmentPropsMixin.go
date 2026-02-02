package previewawsnetworkmanagermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsnetworkmanager/previewawsnetworkmanagermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an AWS Direct Connect gateway attachment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDirectConnectGatewayAttachmentPropsMixin := awscdkmixinspreview.Mixins.NewCfnDirectConnectGatewayAttachmentPropsMixin(&CfnDirectConnectGatewayAttachmentMixinProps{
//   	CoreNetworkId: jsii.String("coreNetworkId"),
//   	DirectConnectGatewayArn: jsii.String("directConnectGatewayArn"),
//   	EdgeLocations: []*string{
//   		jsii.String("edgeLocations"),
//   	},
//   	ProposedNetworkFunctionGroupChange: &ProposedNetworkFunctionGroupChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		NetworkFunctionGroupName: jsii.String("networkFunctionGroupName"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	ProposedSegmentChange: &ProposedSegmentChangeProperty{
//   		AttachmentPolicyRuleNumber: jsii.Number(123),
//   		SegmentName: jsii.String("segmentName"),
//   		Tags: []CfnTag{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   	},
//   	RoutingPolicyLabel: jsii.String("routingPolicyLabel"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-networkmanager-directconnectgatewayattachment.html
//
type CfnDirectConnectGatewayAttachmentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDirectConnectGatewayAttachmentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDirectConnectGatewayAttachmentPropsMixin
type jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin) Props() *CfnDirectConnectGatewayAttachmentMixinProps {
	var returns *CfnDirectConnectGatewayAttachmentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::NetworkManager::DirectConnectGatewayAttachment`.
func NewCfnDirectConnectGatewayAttachmentPropsMixin(props *CfnDirectConnectGatewayAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDirectConnectGatewayAttachmentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDirectConnectGatewayAttachmentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::NetworkManager::DirectConnectGatewayAttachment`.
func NewCfnDirectConnectGatewayAttachmentPropsMixin_Override(c CfnDirectConnectGatewayAttachmentPropsMixin, props *CfnDirectConnectGatewayAttachmentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDirectConnectGatewayAttachmentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDirectConnectGatewayAttachmentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDirectConnectGatewayAttachmentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_networkmanager.mixins.CfnDirectConnectGatewayAttachmentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDirectConnectGatewayAttachmentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

