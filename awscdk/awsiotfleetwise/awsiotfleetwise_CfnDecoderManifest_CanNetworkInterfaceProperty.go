package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canNetworkInterfaceProperty := &canNetworkInterfaceProperty{
//   	canInterface: &canInterfaceProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		protocolName: jsii.String("protocolName"),
//   		protocolVersion: jsii.String("protocolVersion"),
//   	},
//   	interfaceId: jsii.String("interfaceId"),
//   	type: jsii.String("type"),
//   }
//
type CfnDecoderManifest_CanNetworkInterfaceProperty struct {
	// `CfnDecoderManifest.CanNetworkInterfaceProperty.CanInterface`.
	CanInterface interface{} `field:"required" json:"canInterface" yaml:"canInterface"`
	// `CfnDecoderManifest.CanNetworkInterfaceProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.CanNetworkInterfaceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

