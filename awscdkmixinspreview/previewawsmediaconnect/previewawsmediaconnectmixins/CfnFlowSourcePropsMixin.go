package previewawsmediaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediaconnect/previewawsmediaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MediaConnect::FlowSource` resource is usedt to add additional sources to an existing flow.
//
// Adding an additional source requires Failover to be enabled. When you enable Failover, the additional source must use the same protocol as the existing source. A source is the external video content that includes configuration information (encryption and source type) and a network address. Each flow has at least one source. A standard source comes from a source other than another AWS Elemental MediaConnect flow, such as an on-premises encoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowSourcePropsMixin := awscdkmixinspreview.Mixins.NewCfnFlowSourcePropsMixin(&CfnFlowSourceMixinProps{
//   	Decryption: &EncryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		DeviceId: jsii.String("deviceId"),
//   		KeyType: jsii.String("keyType"),
//   		Region: jsii.String("region"),
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   		Url: jsii.String("url"),
//   	},
//   	Description: jsii.String("description"),
//   	EntitlementArn: jsii.String("entitlementArn"),
//   	FlowArn: jsii.String("flowArn"),
//   	GatewayBridgeSource: &GatewayBridgeSourceProperty{
//   		BridgeArn: jsii.String("bridgeArn"),
//   		VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	IngestPort: jsii.Number(123),
//   	MaxBitrate: jsii.Number(123),
//   	MaxLatency: jsii.Number(123),
//   	MinLatency: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	Protocol: jsii.String("protocol"),
//   	SenderControlPort: jsii.Number(123),
//   	SenderIpAddress: jsii.String("senderIpAddress"),
//   	SourceListenerAddress: jsii.String("sourceListenerAddress"),
//   	SourceListenerPort: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	WhitelistCidr: jsii.String("whitelistCidr"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowsource.html
//
type CfnFlowSourcePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFlowSourceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowSourcePropsMixin
type jsiiProxy_CfnFlowSourcePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFlowSourcePropsMixin) Props() *CfnFlowSourceMixinProps {
	var returns *CfnFlowSourceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowSourcePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::FlowSource`.
func NewCfnFlowSourcePropsMixin(props *CfnFlowSourceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFlowSourcePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowSourcePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowSourcePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowSourcePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::FlowSource`.
func NewCfnFlowSourcePropsMixin_Override(c CfnFlowSourcePropsMixin, props *CfnFlowSourceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowSourcePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFlowSourcePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowSourcePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowSourcePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowSourcePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowSourcePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowSourcePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnFlowSourcePropsMixin) Supports(construct constructs.IConstruct) *bool {
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

