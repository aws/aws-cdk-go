package mixinsawsiotfleetwise


// Information about signal messages using the on-board diagnostics (OBD) II protocol in a vehicle.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   obdSignalProperty := &ObdSignalProperty{
//   	BitMaskLength: jsii.String("bitMaskLength"),
//   	BitRightShift: jsii.String("bitRightShift"),
//   	ByteLength: jsii.String("byteLength"),
//   	IsSigned: jsii.String("isSigned"),
//   	Offset: jsii.String("offset"),
//   	Pid: jsii.String("pid"),
//   	PidResponseLength: jsii.String("pidResponseLength"),
//   	Scaling: jsii.String("scaling"),
//   	ServiceMode: jsii.String("serviceMode"),
//   	SignalValueType: jsii.String("signalValueType"),
//   	StartByte: jsii.String("startByte"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html
//
type CfnDecoderManifestPropsMixin_ObdSignalProperty struct {
	// The number of bits to mask in a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bitmasklength
	//
	BitMaskLength *string `field:"optional" json:"bitMaskLength" yaml:"bitMaskLength"`
	// The number of positions to shift bits in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bitrightshift
	//
	BitRightShift *string `field:"optional" json:"bitRightShift" yaml:"bitRightShift"`
	// The length of a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-bytelength
	//
	ByteLength *string `field:"optional" json:"byteLength" yaml:"byteLength"`
	// Determines whether the message is signed ( `true` ) or not ( `false` ).
	//
	// If it's signed, the message can represent both positive and negative numbers. The `isSigned` parameter only applies to the `INTEGER` raw signal type, and it doesn't affect the `FLOATING_POINT` raw signal type. The default value is `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-issigned
	//
	IsSigned interface{} `field:"optional" json:"isSigned" yaml:"isSigned"`
	// The offset used to calculate the signal value.
	//
	// Combined with scaling, the calculation is `value = raw_value * scaling + offset` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-offset
	//
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
	// The diagnostic code used to request data from a vehicle for this signal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-pid
	//
	Pid *string `field:"optional" json:"pid" yaml:"pid"`
	// The length of the requested data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-pidresponselength
	//
	PidResponseLength *string `field:"optional" json:"pidResponseLength" yaml:"pidResponseLength"`
	// A multiplier used to decode the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-scaling
	//
	Scaling *string `field:"optional" json:"scaling" yaml:"scaling"`
	// The mode of operation (diagnostic service) in a message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-servicemode
	//
	ServiceMode *string `field:"optional" json:"serviceMode" yaml:"serviceMode"`
	// The value type of the signal.
	//
	// The default value is `INTEGER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-signalvaluetype
	//
	SignalValueType *string `field:"optional" json:"signalValueType" yaml:"signalValueType"`
	// Indicates the beginning of the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdsignal.html#cfn-iotfleetwise-decodermanifest-obdsignal-startbyte
	//
	StartByte *string `field:"optional" json:"startByte" yaml:"startByte"`
}

