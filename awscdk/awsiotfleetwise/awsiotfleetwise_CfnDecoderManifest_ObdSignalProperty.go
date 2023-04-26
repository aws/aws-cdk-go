package awsiotfleetwise


// Information about signal messages using the on-board diagnostics (OBD) II protocol in a vehicle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdSignalProperty := &ObdSignalProperty{
//   	ByteLength: jsii.String("byteLength"),
//   	Offset: jsii.String("offset"),
//   	Pid: jsii.String("pid"),
//   	PidResponseLength: jsii.String("pidResponseLength"),
//   	Scaling: jsii.String("scaling"),
//   	ServiceMode: jsii.String("serviceMode"),
//   	StartByte: jsii.String("startByte"),
//
//   	// the properties below are optional
//   	BitMaskLength: jsii.String("bitMaskLength"),
//   	BitRightShift: jsii.String("bitRightShift"),
//   }
//
type CfnDecoderManifest_ObdSignalProperty struct {
	// The length of a message.
	ByteLength *string `field:"required" json:"byteLength" yaml:"byteLength"`
	// Indicates where data appears in the message.
	Offset *string `field:"required" json:"offset" yaml:"offset"`
	// The diagnostic code used to request data from a vehicle for this signal.
	Pid *string `field:"required" json:"pid" yaml:"pid"`
	// The length of the requested data.
	PidResponseLength *string `field:"required" json:"pidResponseLength" yaml:"pidResponseLength"`
	// A multiplier used to decode the message.
	Scaling *string `field:"required" json:"scaling" yaml:"scaling"`
	// The mode of operation (diagnostic service) in a message.
	ServiceMode *string `field:"required" json:"serviceMode" yaml:"serviceMode"`
	// Indicates the beginning of the message.
	StartByte *string `field:"required" json:"startByte" yaml:"startByte"`
	// The number of bits to mask in a message.
	BitMaskLength *string `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// The number of positions to shift bits in the message.
	BitRightShift *string `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
}

