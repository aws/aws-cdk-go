package previewawsiotfleetwisemixins


// Information about signal decoder using the Controller Area Network (CAN) protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canSignalDecoderProperty := &CanSignalDecoderProperty{
//   	CanSignal: &CanSignalProperty{
//   		Factor: jsii.String("factor"),
//   		IsBigEndian: jsii.String("isBigEndian"),
//   		IsSigned: jsii.String("isSigned"),
//   		Length: jsii.String("length"),
//   		MessageId: jsii.String("messageId"),
//   		Name: jsii.String("name"),
//   		Offset: jsii.String("offset"),
//   		SignalValueType: jsii.String("signalValueType"),
//   		StartBit: jsii.String("startBit"),
//   	},
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	InterfaceId: jsii.String("interfaceId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.html
//
type CfnDecoderManifestPropsMixin_CanSignalDecoderProperty struct {
	// Information about a single controller area network (CAN) signal and the messages it receives and transmits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.html#cfn-iotfleetwise-decodermanifest-cansignaldecoder-cansignal
	//
	CanSignal interface{} `field:"optional" json:"canSignal" yaml:"canSignal"`
	// The fully qualified name of a signal decoder as defined in a vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.html#cfn-iotfleetwise-decodermanifest-cansignaldecoder-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"optional" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The ID of a network interface that specifies what network protocol a vehicle follows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.html#cfn-iotfleetwise-decodermanifest-cansignaldecoder-interfaceid
	//
	InterfaceId *string `field:"optional" json:"interfaceId" yaml:"interfaceId"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignaldecoder.html#cfn-iotfleetwise-decodermanifest-cansignaldecoder-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

