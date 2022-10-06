package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var bitMaskLength interface{}
//   var bitRightShift interface{}
//   var byteLength interface{}
//   var offset interface{}
//   var pid interface{}
//   var pidResponseLength interface{}
//   var scaling interface{}
//   var serviceMode interface{}
//   var startByte interface{}
//
//   obdSignalDecoderProperty := &obdSignalDecoderProperty{
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	interfaceId: jsii.String("interfaceId"),
//   	obdSignal: &obdSignalProperty{
//   		byteLength: byteLength,
//   		offset: offset,
//   		pid: pid,
//   		pidResponseLength: pidResponseLength,
//   		scaling: scaling,
//   		serviceMode: serviceMode,
//   		startByte: startByte,
//
//   		// the properties below are optional
//   		bitMaskLength: bitMaskLength,
//   		bitRightShift: bitRightShift,
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnDecoderManifest_ObdSignalDecoderProperty struct {
	// `CfnDecoderManifest.ObdSignalDecoderProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnDecoderManifest.ObdSignalDecoderProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.ObdSignalDecoderProperty.ObdSignal`.
	ObdSignal interface{} `field:"required" json:"obdSignal" yaml:"obdSignal"`
	// `CfnDecoderManifest.ObdSignalDecoderProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

