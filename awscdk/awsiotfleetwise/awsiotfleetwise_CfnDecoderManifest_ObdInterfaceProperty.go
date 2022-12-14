package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdInterfaceProperty := &obdInterfaceProperty{
//   	name: jsii.String("name"),
//   	requestMessageId: jsii.String("requestMessageId"),
//
//   	// the properties below are optional
//   	dtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   	hasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   	obdStandard: jsii.String("obdStandard"),
//   	pidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   	useExtendedIds: jsii.String("useExtendedIds"),
//   }
//
type CfnDecoderManifest_ObdInterfaceProperty struct {
	// `CfnDecoderManifest.ObdInterfaceProperty.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `CfnDecoderManifest.ObdInterfaceProperty.RequestMessageId`.
	RequestMessageId *string `field:"required" json:"requestMessageId" yaml:"requestMessageId"`
	// `CfnDecoderManifest.ObdInterfaceProperty.DtcRequestIntervalSeconds`.
	DtcRequestIntervalSeconds *string `field:"optional" json:"dtcRequestIntervalSeconds" yaml:"dtcRequestIntervalSeconds"`
	// `CfnDecoderManifest.ObdInterfaceProperty.HasTransmissionEcu`.
	HasTransmissionEcu *string `field:"optional" json:"hasTransmissionEcu" yaml:"hasTransmissionEcu"`
	// `CfnDecoderManifest.ObdInterfaceProperty.ObdStandard`.
	ObdStandard *string `field:"optional" json:"obdStandard" yaml:"obdStandard"`
	// `CfnDecoderManifest.ObdInterfaceProperty.PidRequestIntervalSeconds`.
	PidRequestIntervalSeconds *string `field:"optional" json:"pidRequestIntervalSeconds" yaml:"pidRequestIntervalSeconds"`
	// `CfnDecoderManifest.ObdInterfaceProperty.UseExtendedIds`.
	UseExtendedIds *string `field:"optional" json:"useExtendedIds" yaml:"useExtendedIds"`
}

