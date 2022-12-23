package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canInterfaceProperty := &canInterfaceProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	protocolName: jsii.String("protocolName"),
//   	protocolVersion: jsii.String("protocolVersion"),
//   }
//
type CfnDecoderManifest_CanInterfaceProperty struct {
	// `CfnDecoderManifest.CanInterfaceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDecoderManifest.CanInterfaceProperty.ProtocolName`.
	ProtocolName *string `field:"optional" json:"protocolName" yaml:"protocolName"`
	// `CfnDecoderManifest.CanInterfaceProperty.ProtocolVersion`.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

