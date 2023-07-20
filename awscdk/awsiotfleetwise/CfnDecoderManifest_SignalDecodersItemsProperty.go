package awsiotfleetwise


// Information about a signal decoder.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalDecodersItemsProperty := &SignalDecodersItemsProperty{
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	InterfaceId: jsii.String("interfaceId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CanSignal: &CanSignalProperty{
//   		Factor: jsii.String("factor"),
//   		IsBigEndian: jsii.String("isBigEndian"),
//   		IsSigned: jsii.String("isSigned"),
//   		Length: jsii.String("length"),
//   		MessageId: jsii.String("messageId"),
//   		Offset: jsii.String("offset"),
//   		StartBit: jsii.String("startBit"),
//
//   		// the properties below are optional
//   		Name: jsii.String("name"),
//   	},
//   	ObdSignal: &ObdSignalProperty{
//   		ByteLength: jsii.String("byteLength"),
//   		Offset: jsii.String("offset"),
//   		Pid: jsii.String("pid"),
//   		PidResponseLength: jsii.String("pidResponseLength"),
//   		Scaling: jsii.String("scaling"),
//   		ServiceMode: jsii.String("serviceMode"),
//   		StartByte: jsii.String("startByte"),
//
//   		// the properties below are optional
//   		BitMaskLength: jsii.String("bitMaskLength"),
//   		BitRightShift: jsii.String("bitRightShift"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html
//
type CfnDecoderManifest_SignalDecodersItemsProperty struct {
	// The fully qualified name of a signal decoder as defined in a vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The ID of a network interface that specifies what network protocol a vehicle follows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// (Optional) Information about a single controller area network (CAN) signal and the messages it receives and transmits.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-cansignal
	//
	CanSignal interface{} `field:"optional" json:"canSignal" yaml:"canSignal"`
	// (Optional) Information about signal messages using the on-board diagnostics (OBD) II protocol in a vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-signaldecodersitems.html#cfn-iotfleetwise-decodermanifest-signaldecodersitems-obdsignal
	//
	ObdSignal interface{} `field:"optional" json:"obdSignal" yaml:"obdSignal"`
}

