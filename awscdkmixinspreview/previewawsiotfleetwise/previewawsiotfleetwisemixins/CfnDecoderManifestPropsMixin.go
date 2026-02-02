package previewawsiotfleetwisemixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsiotfleetwise/previewawsiotfleetwisemixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates the decoder manifest associated with a model manifest. To create a decoder manifest, the following must be true:.
//
// - Every signal decoder has a unique name.
// - Each signal decoder is associated with a network interface.
// - Each network interface has a unique ID.
// - The signal decoders are specified in the model manifest.
//
// For more information, see [Decoder manifests](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/decoder-manifests.html) in the *AWS IoT FleetWise Developer Guide* .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDecoderManifestPropsMixin := awscdkmixinspreview.Mixins.NewCfnDecoderManifestPropsMixin(&CfnDecoderManifestMixinProps{
//   	DefaultForUnmappedSignals: jsii.String("defaultForUnmappedSignals"),
//   	Description: jsii.String("description"),
//   	ModelManifestArn: jsii.String("modelManifestArn"),
//   	Name: jsii.String("name"),
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfacesItemsProperty{
//   			CanInterface: &CanInterfaceProperty{
//   				Name: jsii.String("name"),
//   				ProtocolName: jsii.String("protocolName"),
//   				ProtocolVersion: jsii.String("protocolVersion"),
//   			},
//   			InterfaceId: jsii.String("interfaceId"),
//   			ObdInterface: &ObdInterfaceProperty{
//   				DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   				HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   				Name: jsii.String("name"),
//   				ObdStandard: jsii.String("obdStandard"),
//   				PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   				RequestMessageId: jsii.String("requestMessageId"),
//   				UseExtendedIds: jsii.String("useExtendedIds"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	SignalDecoders: []interface{}{
//   		&SignalDecodersItemsProperty{
//   			CanSignal: &CanSignalProperty{
//   				Factor: jsii.String("factor"),
//   				IsBigEndian: jsii.String("isBigEndian"),
//   				IsSigned: jsii.String("isSigned"),
//   				Length: jsii.String("length"),
//   				MessageId: jsii.String("messageId"),
//   				Name: jsii.String("name"),
//   				Offset: jsii.String("offset"),
//   				SignalValueType: jsii.String("signalValueType"),
//   				StartBit: jsii.String("startBit"),
//   			},
//   			FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   			InterfaceId: jsii.String("interfaceId"),
//   			ObdSignal: &ObdSignalProperty{
//   				BitMaskLength: jsii.String("bitMaskLength"),
//   				BitRightShift: jsii.String("bitRightShift"),
//   				ByteLength: jsii.String("byteLength"),
//   				IsSigned: jsii.String("isSigned"),
//   				Offset: jsii.String("offset"),
//   				Pid: jsii.String("pid"),
//   				PidResponseLength: jsii.String("pidResponseLength"),
//   				Scaling: jsii.String("scaling"),
//   				ServiceMode: jsii.String("serviceMode"),
//   				SignalValueType: jsii.String("signalValueType"),
//   				StartByte: jsii.String("startByte"),
//   			},
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	Status: jsii.String("status"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotfleetwise-decodermanifest.html
//
type CfnDecoderManifestPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDecoderManifestMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDecoderManifestPropsMixin
type jsiiProxy_CfnDecoderManifestPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDecoderManifestPropsMixin) Props() *CfnDecoderManifestMixinProps {
	var returns *CfnDecoderManifestMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDecoderManifestPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::IoTFleetWise::DecoderManifest`.
func NewCfnDecoderManifestPropsMixin(props *CfnDecoderManifestMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDecoderManifestPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDecoderManifestPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDecoderManifestPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::IoTFleetWise::DecoderManifest`.
func NewCfnDecoderManifestPropsMixin_Override(c CfnDecoderManifestPropsMixin, props *CfnDecoderManifestMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDecoderManifestPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDecoderManifestPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDecoderManifestPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_iotfleetwise.mixins.CfnDecoderManifestPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDecoderManifestPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnDecoderManifestPropsMixin) Supports(construct constructs.IConstruct) *bool {
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

