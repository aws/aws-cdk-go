package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dtcRequestIntervalSeconds interface{}
//   var hasTransmissionEcu interface{}
//   var pidRequestIntervalSeconds interface{}
//   var requestMessageId interface{}
//   var useExtendedIds interface{}
//
//   obdNetworkInterfaceProperty := &obdNetworkInterfaceProperty{
//   	interfaceId: jsii.String("interfaceId"),
//   	obdInterface: &obdInterfaceProperty{
//   		name: jsii.String("name"),
//   		requestMessageId: requestMessageId,
//
//   		// the properties below are optional
//   		dtcRequestIntervalSeconds: dtcRequestIntervalSeconds,
//   		hasTransmissionEcu: hasTransmissionEcu,
//   		obdStandard: jsii.String("obdStandard"),
//   		pidRequestIntervalSeconds: pidRequestIntervalSeconds,
//   		useExtendedIds: useExtendedIds,
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnDecoderManifest_ObdNetworkInterfaceProperty struct {
	// `CfnDecoderManifest.ObdNetworkInterfaceProperty.InterfaceId`.
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// `CfnDecoderManifest.ObdNetworkInterfaceProperty.ObdInterface`.
	ObdInterface interface{} `field:"required" json:"obdInterface" yaml:"obdInterface"`
	// `CfnDecoderManifest.ObdNetworkInterfaceProperty.Type`.
	Type *string `field:"required" json:"type" yaml:"type"`
}

