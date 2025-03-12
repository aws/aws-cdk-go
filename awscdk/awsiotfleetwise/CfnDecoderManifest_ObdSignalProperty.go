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
//   	IsSigned: jsii.String("isSigned"),
//   	SignalValueType: jsii.String("signalValueType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html
//
type CfnDecoderManifest_ObdSignalProperty struct {
	// The length of a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bytelength
	//
	ByteLength *string `field:"required" json:"byteLength" yaml:"byteLength"`
	// The offset used to calculate the signal value.
	//
	// Combined with scaling, the calculation is `value = raw_value * scaling + offset` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-offset
	//
	Offset *string `field:"required" json:"offset" yaml:"offset"`
	// The diagnostic code used to request data from a vehicle for this signal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-pid
	//
	Pid *string `field:"required" json:"pid" yaml:"pid"`
	// The length of the requested data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-pidresponselength
	//
	PidResponseLength *string `field:"required" json:"pidResponseLength" yaml:"pidResponseLength"`
	// A multiplier used to decode the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-scaling
	//
	Scaling *string `field:"required" json:"scaling" yaml:"scaling"`
	// The mode of operation (diagnostic service) in a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-servicemode
	//
	ServiceMode *string `field:"required" json:"serviceMode" yaml:"serviceMode"`
	// Indicates the beginning of the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-startbyte
	//
	StartByte *string `field:"required" json:"startByte" yaml:"startByte"`
	// The number of bits to mask in a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bitmasklength
	//
	BitMaskLength *string `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// The number of positions to shift bits in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bitrightshift
	//
	BitRightShift *string `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-issigned
	//
	IsSigned interface{} `field:"optional" json:"isSigned" yaml:"isSigned"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-signalvaluetype
	//
	SignalValueType *string `field:"optional" json:"signalValueType" yaml:"signalValueType"`
}

