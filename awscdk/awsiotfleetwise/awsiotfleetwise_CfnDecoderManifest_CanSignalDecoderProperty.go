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
//   canSignalDecoderProperty := &canSignalDecoderProperty{
//   	canSignal: &canSignalProperty{
//   		factor: factor,
//   		isBigEndian: isBigEndian,
//   		isSigned: isSigned,
//   		length: length,
//   		messageId: messageId,
//   		offset: offset,
//   		startBit: startBit,
//
//   		// the properties below are optional
//   		name: jsii.String("name"),
//   	},
//   	fullyQualifiedName: jsii.String("fullyQualifiedName"),
//   	interfaceId: jsii.String("interfaceId"),
//   	type: jsii.String("type"),
//   }
//
type CfnDecoderManifest_CanSignalDecoderProperty struct {
	// `CfnDecoderManifest.CanSignalDecoderProperty.CanSignal`.
	CanSignal interface{} `field:"required" json:"canSignal" yaml:"canSignal"`
	// `CfnDecoderManifest.CanSignalDecoderProperty.FullyQualifiedName`.
	FullyQualifiedName *string `field:"required" json:"fullyQualifiedName" yaml:"fullyQualifiedName"`
	// `CfnDecoderManifest.CanSignalDecoderProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.CanSignalDecoderProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

