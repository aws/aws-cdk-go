package previewawsmediaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediaconnect/previewawsmediaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MediaConnect::Bridge` resource defines a connection between your data center’s gateway instances and the cloud.
//
// For each bridge, you specify the type of bridge, transport protocol to use, and details for any outputs and failover.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBridgePropsMixin := awscdkmixinspreview.Mixins.NewCfnBridgePropsMixin(&CfnBridgeMixinProps{
//   	EgressGatewayBridge: &EgressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   	},
//   	IngressGatewayBridge: &IngressGatewayBridgeProperty{
//   		MaxBitrate: jsii.Number(123),
//   		MaxOutputs: jsii.Number(123),
//   	},
//   	Name: jsii.String("name"),
//   	Outputs: []interface{}{
//   		&BridgeOutputProperty{
//   			NetworkOutput: &BridgeNetworkOutputProperty{
//   				IpAddress: jsii.String("ipAddress"),
//   				Name: jsii.String("name"),
//   				NetworkName: jsii.String("networkName"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   				Ttl: jsii.Number(123),
//   			},
//   		},
//   	},
//   	PlacementArn: jsii.String("placementArn"),
//   	SourceFailoverConfig: &FailoverConfigProperty{
//   		FailoverMode: jsii.String("failoverMode"),
//   		SourcePriority: &SourcePriorityProperty{
//   			PrimarySource: jsii.String("primarySource"),
//   		},
//   		State: jsii.String("state"),
//   	},
//   	Sources: []interface{}{
//   		&BridgeSourceProperty{
//   			FlowSource: &BridgeFlowSourceProperty{
//   				FlowArn: jsii.String("flowArn"),
//   				FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   					VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   			NetworkSource: &BridgeNetworkSourceProperty{
//   				MulticastIp: jsii.String("multicastIp"),
//   				MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   					MulticastSourceIp: jsii.String("multicastSourceIp"),
//   				},
//   				Name: jsii.String("name"),
//   				NetworkName: jsii.String("networkName"),
//   				Port: jsii.Number(123),
//   				Protocol: jsii.String("protocol"),
//   			},
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridge.html
//
type CfnBridgePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnBridgeMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnBridgePropsMixin
type jsiiProxy_CfnBridgePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnBridgePropsMixin) Props() *CfnBridgeMixinProps {
	var returns *CfnBridgeMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnBridgePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::Bridge`.
func NewCfnBridgePropsMixin(props *CfnBridgeMixinProps, options *mixins.CfnPropertyMixinOptions) CfnBridgePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnBridgePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnBridgePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::Bridge`.
func NewCfnBridgePropsMixin_Override(c CfnBridgePropsMixin, props *CfnBridgeMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnBridgePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnBridgePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnBridgePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnBridgePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnBridgePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnBridgePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

