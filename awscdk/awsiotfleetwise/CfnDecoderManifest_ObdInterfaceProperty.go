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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html
//
type CfnDecoderManifest_ObdInterfaceProperty struct {
	// The name of the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the message requesting vehicle data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-requestmessageid
	//
	RequestMessageId *string `field:"required" json:"requestMessageId" yaml:"requestMessageId"`
	// The maximum number message requests per diagnostic trouble code per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-dtcrequestintervalseconds
	//
	DtcRequestIntervalSeconds *string `field:"optional" json:"dtcRequestIntervalSeconds" yaml:"dtcRequestIntervalSeconds"`
	// Whether the vehicle has a transmission control module (TCM).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-hastransmissionecu
	//
	HasTransmissionEcu *string `field:"optional" json:"hasTransmissionEcu" yaml:"hasTransmissionEcu"`
	// The standard OBD II PID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-obdstandard
	//
	ObdStandard *string `field:"optional" json:"obdStandard" yaml:"obdStandard"`
	// The maximum number message requests per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-pidrequestintervalseconds
	//
	PidRequestIntervalSeconds *string `field:"optional" json:"pidRequestIntervalSeconds" yaml:"pidRequestIntervalSeconds"`
	// Whether to use extended IDs in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdinterface.html#cfn-iotfleetwise-decodermanifest-obdinterface-useextendedids
	//
	UseExtendedIds *string `field:"optional" json:"useExtendedIds" yaml:"useExtendedIds"`
}

