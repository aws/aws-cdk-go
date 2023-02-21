package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   signalDecodersItemsProperty := &signalDecodersItemsProperty{
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	interfaceId: jsii.String("interfaceId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	canSignal: &canSignalProperty{
//   		factor: jsii.String("factor"),
//   		isBigEndian: jsii.String("isBigEndian"),
//   		isSigned: jsii.String("isSigned"),
//   		length: jsii.String("length"),
//   		messageId: jsii.String("messageId"),
//   		offset: jsii.String("offset"),
//   		startBit: jsii.String("startBit"),
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	obdSignal: &obdSignalProperty{
//   		byteLength: jsii.String("byteLength"),
//   		offset: jsii.String("offset"),
//   		pid: jsii.String("pid"),
//   		pidResponseLength: jsii.String("pidResponseLength"),
//   		scaling: jsii.String("scaling"),
//   		serviceMode: jsii.String("serviceMode"),
//   		startByte: jsii.String("startByte"),
//
//   		// the properties below are optional
//   		bitMaskLength: jsii.String("bitMaskLength"),
//   		bitRightShift: jsii.String("bitRightShift"),
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

