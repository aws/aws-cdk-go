package awsiotfleetwise


// A network interface that specifies the On-board diagnostic (OBD) II network protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdInterfaceProperty := &ObdInterfaceProperty{
//   	Name: jsii.String("name"),
//   	RequestMessageId: jsii.String("requestMessageId"),
//
//   	// the properties below are optional
//   	DtcRequestIntervalSeconds: jsii.String("dtcRequestIntervalSeconds"),
//   	HasTransmissionEcu: jsii.String("hasTransmissionEcu"),
//   	ObdStandard: jsii.String("obdStandard"),
//   	PidRequestIntervalSeconds: jsii.String("pidRequestIntervalSeconds"),
//   	UseExtendedIds: jsii.String("useExtendedIds"),
//   }
//
type CfnDecoderManifest_ObdInterfaceProperty struct {
	// The name of the interface.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the message requesting vehicle data.
	RequestMessageId *string `field:"required" json:"requestMessageId" yaml:"requestMessageId"`
	// The maximum number message requests per diagnostic trouble code per second.
	DtcRequestIntervalSeconds *string `field:"optional" json:"dtcRequestIntervalSeconds" yaml:"dtcRequestIntervalSeconds"`
	// Whether the vehicle has a transmission control module (TCM).
	HasTransmissionEcu *string `field:"optional" json:"hasTransmissionEcu" yaml:"hasTransmissionEcu"`
	// The standard OBD II PID.
	ObdStandard *string `field:"optional" json:"obdStandard" yaml:"obdStandard"`
	// The maximum number message requests per second.
	PidRequestIntervalSeconds *string `field:"optional" json:"pidRequestIntervalSeconds" yaml:"pidRequestIntervalSeconds"`
	// Whether to use extended IDs in the message.
	UseExtendedIds *string `field:"optional" json:"useExtendedIds" yaml:"useExtendedIds"`
}

