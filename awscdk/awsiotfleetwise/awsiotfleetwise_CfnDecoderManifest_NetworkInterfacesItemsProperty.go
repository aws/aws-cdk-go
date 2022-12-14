package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfacesItemsProperty := &networkInterfacesItemsProperty{
//   	interfaceId: jsii.String("interfaceId"),
//   	type: jsii.String("type"),
//
//   	// the properties below are optional
//   	canInterface: &canInterfaceProperty{
//   		name: jsii.String("name"),
//
//   		// the properties below are optional
//   		protocolName: jsii.String("protocolName"),
//   		protocolVersion: jsii.String("protocolVersion"),
//   	},
//   	obdInterface: &obdInterfaceProperty{
//   		name: jsii.String("name"),
//   		requestMessageId: jsii.String("requestMessageId"),
//
//   		// the properties below are optional
//   		dtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   		hasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   		obdStandard: jsii.String("obdStandard"),
//   		pidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   		useExtendedIds: jsii.String("useExtendedIds"),
//   	},
//   }
//
type CfnDecoderManifest_NetworkInterfacesItemsProperty struct {
	// `CfnDecoderManifest.NetworkInterfacesItemsProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.NetworkInterfacesItemsProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
	// `CfnDecoderManifest.NetworkInterfacesItemsProperty.CanInterface`.
	CanInterface interface{} `field:"optional" json:"canInterface" yaml:"canInterface"`
	// `CfnDecoderManifest.NetworkInterfacesItemsProperty.ObdInterface`.
	ObdInterface interface{} `field:"optional" json:"obdInterface" yaml:"obdInterface"`
}

