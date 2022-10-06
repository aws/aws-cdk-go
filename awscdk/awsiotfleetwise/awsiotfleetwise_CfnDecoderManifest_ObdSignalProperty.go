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
//   obdSignalProperty := &obdSignalProperty{
//   	byteLength: byteLength,
//   	offset: offset,
//   	pid: pid,
//   	pidResponseLength: pidResponseLength,
//   	scaling: scaling,
//   	serviceMode: serviceMode,
//   	startByte: startByte,
//
//   	// the properties below are optional
//   	bitMaskLength: bitMaskLength,
//   	bitRightShift: bitRightShift,
//   }
//
type CfnDecoderManifest_ObdSignalProperty struct {
	// `CfnDecoderManifest.ObdSignalProperty.ByteLength`.
	ByteLength interface{} `field:"required" json:"byteLength" yaml:"byteLength"`
	// `CfnDecoderManifest.ObdSignalProperty.Offset`.
	Offset interface{} `field:"required" json:"offset" yaml:"offset"`
	// `CfnDecoderManifest.ObdSignalProperty.Pid`.
	Pid interface{} `field:"required" json:"pid" yaml:"pid"`
	// `CfnDecoderManifest.ObdSignalProperty.PidResponseLength`.
	PidResponseLength interface{} `field:"required" json:"pidResponseLength" yaml:"pidResponseLength"`
	// `CfnDecoderManifest.ObdSignalProperty.Scaling`.
	Scaling interface{} `field:"required" json:"scaling" yaml:"scaling"`
	// `CfnDecoderManifest.ObdSignalProperty.ServiceMode`.
	ServiceMode interface{} `field:"required" json:"serviceMode" yaml:"serviceMode"`
	// `CfnDecoderManifest.ObdSignalProperty.StartByte`.
	StartByte interface{} `field:"required" json:"startByte" yaml:"startByte"`
	// `CfnDecoderManifest.ObdSignalProperty.BitMaskLength`.
	BitMaskLength interface{} `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// `CfnDecoderManifest.ObdSignalProperty.BitRightShift`.
	BitRightShift interface{} `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
}

