package awsiotfleetwise


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
	// `CfnDecoderManifest.CanSignalProperty.Factor`.
	Factor interface{} `field:"required" json:"factor" yaml:"factor"`
	// `CfnDecoderManifest.CanSignalProperty.IsBigEndian`.
	IsBigEndian interface{} `field:"required" json:"isBigEndian" yaml:"isBigEndian"`
	// `CfnDecoderManifest.CanSignalProperty.IsSigned`.
	IsSigned interface{} `field:"required" json:"isSigned" yaml:"isSigned"`
	// `CfnDecoderManifest.CanSignalProperty.Length`.
	Length interface{} `field:"required" json:"length" yaml:"length"`
	// `CfnDecoderManifest.CanSignalProperty.MessageId`.
	MessageId interface{} `field:"required" json:"messageId" yaml:"messageId"`
	// `CfnDecoderManifest.CanSignalProperty.Offset`.
	Offset interface{} `field:"required" json:"offset" yaml:"offset"`
	// `CfnDecoderManifest.CanSignalProperty.StartBit`.
	StartBit interface{} `field:"required" json:"startBit" yaml:"startBit"`
	// `CfnDecoderManifest.CanSignalProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

