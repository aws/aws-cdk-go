package awsiotfleetwise


// A network interface that specifies the On-board diagnostic (OBD) II network protocol.
//
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
	// The name of the interface.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the message requesting vehicle data.
	RequestMessageId interface{} `field:"required" json:"requestMessageId" yaml:"requestMessageId"`
	// The maximum number message requests per diagnostic trouble code per second.
	DtcRequestIntervalSeconds interface{} `field:"optional" json:"dtcRequestIntervalSeconds" yaml:"dtcRequestIntervalSeconds"`
	// Whether the vehicle has a transmission control module (TCM).
	HasTransmissionEcu interface{} `field:"optional" json:"hasTransmissionEcu" yaml:"hasTransmissionEcu"`
	// The standard OBD II PID.
	ObdStandard *string `field:"optional" json:"obdStandard" yaml:"obdStandard"`
	// The maximum number message requests per second.
	PidRequestIntervalSeconds interface{} `field:"optional" json:"pidRequestIntervalSeconds" yaml:"pidRequestIntervalSeconds"`
	// Whether to use extended IDs in the message.
	UseExtendedIds interface{} `field:"optional" json:"useExtendedIds" yaml:"useExtendedIds"`
}

