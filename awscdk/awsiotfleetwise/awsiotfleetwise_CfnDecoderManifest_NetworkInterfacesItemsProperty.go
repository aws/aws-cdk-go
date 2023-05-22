package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfacesItemsProperty := &NetworkInterfacesItemsProperty{
//   	InterfaceId: jsii.String("interfaceId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	CanInterface: &CanInterfaceProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		ProtocolName: jsii.String("protocolName"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   	},
//   	ObdInterface: &ObdInterfaceProperty{
//   		Name: jsii.String("name"),
//   		RequestMessageId: jsii.String("requestMessageId"),
//
//   		// the properties below are optional
//   		DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   		HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   		ObdStandard: jsii.String("obdStandard"),
//   		PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   		UseExtendedIds: jsii.String("useExtendedIds"),
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

