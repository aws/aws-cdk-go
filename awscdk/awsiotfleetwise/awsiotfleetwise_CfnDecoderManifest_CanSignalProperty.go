package awsiotfleetwise


// Information about a single controller area network (CAN) signal and the messages it receives and transmits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canSignalProperty := &canSignalProperty{
//   	factor: jsii.String("factor"),
//   	isBigEndian: jsii.String("isBigEndian"),
//   	isSigned: jsii.String("isSigned"),
//   	length: jsii.String("length"),
//   	messageId: jsii.String("messageId"),
//   	offset: jsii.String("offset"),
//   	startBit: jsii.String("startBit"),
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnDecoderManifest_CanSignalProperty struct {
	// A multiplier used to decode the CAN message.
	Factor *string `field:"required" json:"factor" yaml:"factor"`
	// Whether the byte ordering of a CAN message is big-endian.
	IsBigEndian *string `field:"required" json:"isBigEndian" yaml:"isBigEndian"`
	// Whether the message data is specified as a signed value.
	IsSigned *string `field:"required" json:"isSigned" yaml:"isSigned"`
	// How many bytes of data are in the message.
	Length *string `field:"required" json:"length" yaml:"length"`
	// The ID of the message.
	MessageId *string `field:"required" json:"messageId" yaml:"messageId"`
	// Indicates where data appears in the CAN message.
	Offset *string `field:"required" json:"offset" yaml:"offset"`
	// Indicates the beginning of the CAN message.
	StartBit *string `field:"required" json:"startBit" yaml:"startBit"`
	// The name of the signal.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

