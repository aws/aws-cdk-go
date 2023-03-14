package awsiotfleetwise


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
type CfnDecoderManifest_SignalDecodersItemsProperty struct {
	// `CfnDecoderManifest.SignalDecodersItemsProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnDecoderManifest.SignalDecodersItemsProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.SignalDecodersItemsProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnDecoderManifest.SignalDecodersItemsProperty.CanSignal`.
	CanSignal interface{} `field:"optional" json:"canSignal" yaml:"canSignal"`
	// `CfnDecoderManifest.SignalDecodersItemsProperty.ObdSignal`.
	ObdSignal interface{} `field:"optional" json:"obdSignal" yaml:"obdSignal"`
}

