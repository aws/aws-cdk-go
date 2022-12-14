package awsiotfleetwise


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
	// `CfnDecoderManifest.CanSignalProperty.Factor`.
	Factor *string `field:"required" json:"factor" yaml:"factor"`
	// `CfnDecoderManifest.CanSignalProperty.IsBigEndian`.
	IsBigEndian *string `field:"required" json:"isBigEndian" yaml:"isBigEndian"`
	// `CfnDecoderManifest.CanSignalProperty.IsSigned`.
	IsSigned *string `field:"required" json:"isSigned" yaml:"isSigned"`
	// `CfnDecoderManifest.CanSignalProperty.Length`.
	Length *string `field:"required" json:"length" yaml:"length"`
	// `CfnDecoderManifest.CanSignalProperty.MessageId`.
	MessageId *string `field:"required" json:"messageId" yaml:"messageId"`
	// `CfnDecoderManifest.CanSignalProperty.Offset`.
	Offset *string `field:"required" json:"offset" yaml:"offset"`
	// `CfnDecoderManifest.CanSignalProperty.StartBit`.
	StartBit *string `field:"required" json:"startBit" yaml:"startBit"`
	// `CfnDecoderManifest.CanSignalProperty.Name`.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

