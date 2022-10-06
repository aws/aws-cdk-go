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
//   obdInterfaceProperty := &obdInterfaceProperty{
//   	name: jsii.String("name"),
//   	requestMessageId: requestMessageId,
//
//   	// the properties below are optional
//   	dtcRequestIntervalSeconds: dtcRequestIntervalSeconds,
//   	hasTransmissionEcu: hasTransmissionEcu,
//   	obdStandard: jsii.String("obdStandard"),
//   	pidRequestIntervalSeconds: pidRequestIntervalSeconds,
//   	useExtendedIds: useExtendedIds,
//   }
//
type CfnDecoderManifest_ObdInterfaceProperty struct {
	// `CfnDecoderManifest.ObdInterfaceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDecoderManifest.ObdInterfaceProperty.RequestMessageId`.
	RequestMessageId interface{} `field:"required" json:"requestMessageId" yaml:"requestMessageId"`
	// `CfnDecoderManifest.ObdInterfaceProperty.DtcRequestIntervalSeconds`.
	DtcRequestIntervalSeconds interface{} `field:"optional" json:"dtcRequestIntervalSeconds" yaml:"dtcRequestIntervalSeconds"`
	// `CfnDecoderManifest.ObdInterfaceProperty.HasTransmissionEcu`.
	HasTransmissionEcu interface{} `field:"optional" json:"hasTransmissionEcu" yaml:"hasTransmissionEcu"`
	// `CfnDecoderManifest.ObdInterfaceProperty.ObdStandard`.
	ObdStandard *string `field:"optional" json:"obdStandard" yaml:"obdStandard"`
	// `CfnDecoderManifest.ObdInterfaceProperty.PidRequestIntervalSeconds`.
	PidRequestIntervalSeconds interface{} `field:"optional" json:"pidRequestIntervalSeconds" yaml:"pidRequestIntervalSeconds"`
	// `CfnDecoderManifest.ObdInterfaceProperty.UseExtendedIds`.
	UseExtendedIds interface{} `field:"optional" json:"useExtendedIds" yaml:"useExtendedIds"`
}

