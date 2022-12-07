package awsiotfleetwise


// Information about signal messages using the on-board diagnostics (OBD) II protocol in a vehicle.
//
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
	// The length of a message.
	ByteLength interface{} `field:"required" json:"byteLength" yaml:"byteLength"`
	// Indicates where data appears in the message.
	Offset interface{} `field:"required" json:"offset" yaml:"offset"`
	// The diagnostic code used to request data from a vehicle for this signal.
	Pid interface{} `field:"required" json:"pid" yaml:"pid"`
	// The length of the requested data.
	PidResponseLength interface{} `field:"required" json:"pidResponseLength" yaml:"pidResponseLength"`
	// A multiplier used to decode the message.
	Scaling interface{} `field:"required" json:"scaling" yaml:"scaling"`
	// The mode of operation (diagnostic service) in a message.
	ServiceMode interface{} `field:"required" json:"serviceMode" yaml:"serviceMode"`
	// Indicates the beginning of the message.
	StartByte interface{} `field:"required" json:"startByte" yaml:"startByte"`
	// The number of bits to mask in a message.
	BitMaskLength interface{} `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// The number of positions to shift bits in the message.
	BitRightShift interface{} `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
}

