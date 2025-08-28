package awsiotfleetwise


// A list of information about signal decoders.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdSignalDecoderProperty := &ObdSignalDecoderProperty{
//   	FullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	InterfaceId: jsii.String("interfaceId"),
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
//   		IsSigned: jsii.String("isSigned"),
//   		SignalValueType: jsii.String("signalValueType"),
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html
//
type CfnDecoderManifest_ObdSignalDecoderProperty struct {
	// The fully qualified name of a signal decoder as defined in a vehicle model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-fullyqualifiedname
	//
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// The ID of a network interface that specifies what network protocol a vehicle follows.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// Information about signal messages using the on-board diagnostics (OBD) II protocol in a vehicle.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-obdsignal
	//
	ObdSignal interface{} `field:"required" json:"obdSignal" yaml:"obdSignal"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignaldecoder.html#cfn-iotfleetwise-decodermanifest-obdsignaldecoder-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

