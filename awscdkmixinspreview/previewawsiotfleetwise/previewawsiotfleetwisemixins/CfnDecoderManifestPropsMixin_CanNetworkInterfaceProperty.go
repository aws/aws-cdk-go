package previewawsiotfleetwisemixins


// Represents a node and its specifications in an in-vehicle communication network.
//
// All signal decoders must be associated with a network node.
//
// To return this information about all the network interfaces specified in a decoder manifest, use the [ListDecoderManifestNetworkInterfaces](https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_ListDecoderManifestNetworkInterfaces.html) in the *AWS IoT FleetWise API Reference* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canNetworkInterfaceProperty := &CanNetworkInterfaceProperty{
//   	CanInterface: &CanInterfaceProperty{
//   		Name: jsii.String("name"),
//   		ProtocolName: jsii.String("protocolName"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   	},
//   	InterfaceId: jsii.String("interfaceId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html
//
type CfnDecoderManifestPropsMixin_CanNetworkInterfaceProperty struct {
	// Information about a network interface specified by the Controller Area Network (CAN) protocol.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-caninterface
	//
	CanInterface interface{} `field:"optional" json:"canInterface" yaml:"canInterface"`
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-interfaceid
	//
	InterfaceId *string `field:"optional" json:"interfaceId" yaml:"interfaceId"`
	// The network protocol for the vehicle.
	//
	// For example, `CAN_SIGNAL` specifies a protocol that defines how data is communicated between electronic control units (ECUs). `OBD_SIGNAL` specifies a protocol that defines how self-diagnostic data is communicated between ECUs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

