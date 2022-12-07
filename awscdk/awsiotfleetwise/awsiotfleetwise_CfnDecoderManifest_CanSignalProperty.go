package awsiotfleetwise


// Information about a single controller area network (CAN) signal and the messages it receives and transmits.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var factor interface{}
//   var isBigEndian interface{}
//   var isSigned interface{}
//   var length interface{}
//   var messageId interface{}
//   var offset interface{}
//   var startBit interface{}
//
//   canSignalProperty := &canSignalProperty{
//   	factor: factor,
//   	isBigEndian: isBigEndian,
//   	isSigned: isSigned,
//   	length: length,
//   	messageId: messageId,
//   	offset: offset,
//   	startBit: startBit,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnDecoderManifest_CanSignalProperty struct {
	// A multiplier used to decode the CAN message.
	Factor interface{} `field:"required" json:"factor" yaml:"factor"`
	// Whether the byte ordering of a CAN message is big-endian.
	IsBigEndian interface{} `field:"required" json:"isBigEndian" yaml:"isBigEndian"`
	// Whether the message data is specified as a signed value.
	IsSigned interface{} `field:"required" json:"isSigned" yaml:"isSigned"`
	// How many bytes of data are in the message.
	Length interface{} `field:"required" json:"length" yaml:"length"`
	// The ID of the message.
	MessageId interface{} `field:"required" json:"messageId" yaml:"messageId"`
	// Indicates where data appears in the CAN message.
	Offset interface{} `field:"required" json:"offset" yaml:"offset"`
	// Indicates the beginning of the CAN message.
	StartBit interface{} `field:"required" json:"startBit" yaml:"startBit"`
	// The name of the signal.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

