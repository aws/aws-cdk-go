package awsiotfleetwise


// (Optional) A list of information about available network interfaces.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html
//
type CfnDecoderManifest_NetworkInterfacesItemsProperty struct {
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// (Optional) Information about a network interface specified by the Controller Area Network (CAN) protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-caninterface
	//
	CanInterface interface{} `field:"optional" json:"canInterface" yaml:"canInterface"`
	// (Optional) Information about a network interface specified by the On-board diagnostic (OBD) II protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-networkinterfacesitems.html#cfn-iotfleetwise-decodermanifest-networkinterfacesitems-obdinterface
	//
	ObdInterface interface{} `field:"optional" json:"obdInterface" yaml:"obdInterface"`
}

