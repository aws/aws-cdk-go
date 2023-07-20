package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canNetworkInterfaceProperty := &CanNetworkInterfaceProperty{
//   	CanInterface: &CanInterfaceProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		ProtocolName: jsii.String("protocolName"),
//   		ProtocolVersion: jsii.String("protocolVersion"),
//   	},
//   	InterfaceId: jsii.String("interfaceId"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html
//
type CfnDecoderManifest_CanNetworkInterfaceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-caninterface
	//
	CanInterface interface{} `field:"required" json:"canInterface" yaml:"canInterface"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-interfaceid
	//
	InterfaceId *string `field:"required" json:"interfaceId" yaml:"interfaceId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-cannetworkinterface.html#cfn-iotfleetwise-decodermanifest-cannetworkinterface-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}

