package previewawsiotfleetwisemixins


// Information about a single controller area network (CAN) signal and the messages it receives and transmits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canSignalProperty := &CanSignalProperty{
//   	Factor: jsii.String("factor"),
//   	IsBigEndian: jsii.String("isBigEndian"),
//   	IsSigned: jsii.String("isSigned"),
//   	Length: jsii.String("length"),
//   	MessageId: jsii.String("messageId"),
//   	Name: jsii.String("name"),
//   	Offset: jsii.String("offset"),
//   	SignalValueType: jsii.String("signalValueType"),
//   	StartBit: jsii.String("startBit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html
//
type CfnDecoderManifestPropsMixin_CanSignalProperty struct {
	// A multiplier used to decode the CAN message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-factor
	//
	Factor *string `field:"optional" json:"factor" yaml:"factor"`
	// Whether the byte ordering of a CAN message is big-endian.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-isbigendian
	//
	IsBigEndian *string `field:"optional" json:"isBigEndian" yaml:"isBigEndian"`
	// Whether the message data is specified as a signed value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-issigned
	//
	IsSigned *string `field:"optional" json:"isSigned" yaml:"isSigned"`
	// How many bytes of data are in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-length
	//
	Length *string `field:"optional" json:"length" yaml:"length"`
	// The ID of the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-messageid
	//
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
	// The name of the signal.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The offset used to calculate the signal value.
	//
	// Combined with factor, the calculation is `value = raw_value * factor + offset` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-offset
	//
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
	// The value type of the signal.
	//
	// The default value is `INTEGER` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-signalvaluetype
	//
	SignalValueType *string `field:"optional" json:"signalValueType" yaml:"signalValueType"`
	// Indicates the beginning of the CAN message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cansignal.html#cfn-iotfleetwise-decodermanifest-cansignal-startbit
	//
	StartBit *string `field:"optional" json:"startBit" yaml:"startBit"`
}

