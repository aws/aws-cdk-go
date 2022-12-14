package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdSignalProperty := &obdSignalProperty{
//   	byteLength: jsii.String("byteLength"),
//   	offset: jsii.String("offset"),
//   	pid: jsii.String("pid"),
//   	pidResponseLength: jsii.String("pidResponseLength"),
//   	scaling: jsii.String("scaling"),
//   	serviceMode: jsii.String("serviceMode"),
//   	startByte: jsii.String("startByte"),
//
//   	// the properties below are optional
//   	bitMaskLength: jsii.String("bitMaskLength"),
//   	bitRightShift: jsii.String("bitRightShift"),
//   }
//
type CfnDecoderManifest_ObdSignalProperty struct {
	// `CfnDecoderManifest.ObdSignalProperty.ByteLength`.
	ByteLength *string `field:"required" json:"byteLength" yaml:"byteLength"`
	// `CfnDecoderManifest.ObdSignalProperty.Offset`.
	Offset *string `field:"required" json:"offset" yaml:"offset"`
	// `CfnDecoderManifest.ObdSignalProperty.Pid`.
	Pid *string `field:"required" json:"pid" yaml:"pid"`
	// `CfnDecoderManifest.ObdSignalProperty.PidResponseLength`.
	PidResponseLength *string `field:"required" json:"pidResponseLength" yaml:"pidResponseLength"`
	// `CfnDecoderManifest.ObdSignalProperty.Scaling`.
	Scaling *string `field:"required" json:"scaling" yaml:"scaling"`
	// `CfnDecoderManifest.ObdSignalProperty.ServiceMode`.
	ServiceMode *string `field:"required" json:"serviceMode" yaml:"serviceMode"`
	// `CfnDecoderManifest.ObdSignalProperty.StartByte`.
	StartByte *string `field:"required" json:"startByte" yaml:"startByte"`
	// `CfnDecoderManifest.ObdSignalProperty.BitMaskLength`.
	BitMaskLength *string `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// `CfnDecoderManifest.ObdSignalProperty.BitRightShift`.
	BitRightShift *string `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
}

