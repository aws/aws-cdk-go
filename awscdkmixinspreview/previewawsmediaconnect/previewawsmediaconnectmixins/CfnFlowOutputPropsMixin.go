package previewawsmediaconnectmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsmediaconnect/previewawsmediaconnectmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::MediaConnect::FlowOutput` resource defines the destination address, protocol, and port that AWS Elemental MediaConnect sends the ingested video to.
//
// Each flow can have up to 50 outputs. An output can have the same protocol or a different protocol from the source. The following protocols are supported: RIST, RTP, RTP-FEC, SRT-listener, SRT-caller, Zixi pull, and Zixi push. CDI and ST 2110 JPEG XS protocols are not currently supported by AWS CloudFormation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var automatic interface{}
//
//   cfnFlowOutputPropsMixin := awscdkmixinspreview.Mixins.NewCfnFlowOutputPropsMixin(&CfnFlowOutputMixinProps{
//   	CidrAllowList: []*string{
//   		jsii.String("cidrAllowList"),
//   	},
//   	Description: jsii.String("description"),
//   	Destination: jsii.String("destination"),
//   	Encryption: &EncryptionProperty{
//   		Algorithm: jsii.String("algorithm"),
//   		KeyType: jsii.String("keyType"),
//   		RoleArn: jsii.String("roleArn"),
//   		SecretArn: jsii.String("secretArn"),
//   	},
//   	FlowArn: jsii.String("flowArn"),
//   	MaxLatency: jsii.Number(123),
//   	MediaStreamOutputConfigurations: []interface{}{
//   		&MediaStreamOutputConfigurationProperty{
//   			DestinationConfigurations: []interface{}{
//   				&DestinationConfigurationProperty{
//   					DestinationIp: jsii.String("destinationIp"),
//   					DestinationPort: jsii.Number(123),
//   					Interface: &InterfaceProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   			EncodingName: jsii.String("encodingName"),
//   			EncodingParameters: &EncodingParametersProperty{
//   				CompressionFactor: jsii.Number(123),
//   				EncoderProfile: jsii.String("encoderProfile"),
//   			},
//   			MediaStreamName: jsii.String("mediaStreamName"),
//   		},
//   	},
//   	MinLatency: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	NdiProgramName: jsii.String("ndiProgramName"),
//   	NdiSpeedHqQuality: jsii.Number(123),
//   	OutputStatus: jsii.String("outputStatus"),
//   	Port: jsii.Number(123),
//   	Protocol: jsii.String("protocol"),
//   	RemoteId: jsii.String("remoteId"),
//   	RouterIntegrationState: jsii.String("routerIntegrationState"),
//   	RouterIntegrationTransitEncryption: &FlowTransitEncryptionProperty{
//   		EncryptionKeyConfiguration: &FlowTransitEncryptionKeyConfigurationProperty{
//   			Automatic: automatic,
//   			SecretsManager: &SecretsManagerEncryptionKeyConfigurationProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SecretArn: jsii.String("secretArn"),
//   			},
//   		},
//   		EncryptionKeyType: jsii.String("encryptionKeyType"),
//   	},
//   	SmoothingLatency: jsii.Number(123),
//   	StreamId: jsii.String("streamId"),
//   	VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-flowoutput.html
//
type CfnFlowOutputPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnFlowOutputMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnFlowOutputPropsMixin
type jsiiProxy_CfnFlowOutputPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnFlowOutputPropsMixin) Props() *CfnFlowOutputMixinProps {
	var returns *CfnFlowOutputMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnFlowOutputPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::MediaConnect::FlowOutput`.
func NewCfnFlowOutputPropsMixin(props *CfnFlowOutputMixinProps, options *mixins.CfnPropertyMixinOptions) CfnFlowOutputPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnFlowOutputPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnFlowOutputPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowOutputPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::MediaConnect::FlowOutput`.
func NewCfnFlowOutputPropsMixin_Override(c CfnFlowOutputPropsMixin, props *CfnFlowOutputMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowOutputPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnFlowOutputPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnFlowOutputPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowOutputPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnFlowOutputPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_mediaconnect.mixins.CfnFlowOutputPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnFlowOutputPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnFlowOutputPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

