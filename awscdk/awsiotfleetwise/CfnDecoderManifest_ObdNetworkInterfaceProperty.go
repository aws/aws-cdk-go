package awsiotfleetwise


// Information about a network interface specified by the On-board diagnostic (OBD) II protocol.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   obdNetworkInterfaceProperty := &ObdNetworkInterfaceProperty{
//   	InterfaceId: jsii.String("interfaceId"),
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
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdnetworkinterface.html
//
type CfnDecoderManifest_ObdNetworkInterfaceProperty struct {
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdnetworkinterface.html#cfn-iotfleetwise-decodermanifest-obdnetworkinterface-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// (Optional) Information about a network interface specified by the On-board diagnostic (OBD) II protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdnetworkinterface.html#cfn-iotfleetwise-decodermanifest-obdnetworkinterface-obdinterface
	//
	ObdInterface interface{} `field:"required" json:"obdInterface" yaml:"obdInterface"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-obdnetworkinterface.html#cfn-iotfleetwise-decodermanifest-obdnetworkinterface-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

