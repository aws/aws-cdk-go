package awsiotfleetwise


// A single controller area network (CAN) device interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canInterfaceProperty := &CanInterfaceProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ProtocolName: jsii.String("protocolName"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html
//
type CfnDecoderManifest_CanInterfaceProperty struct {
	// The unique name of the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// (Optional) The name of the communication protocol for the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolname
	//
	ProtocolName *string `field:"optional" json:"protocolName" yaml:"protocolName"`
	// (Optional) The version of the communication protocol for the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolversion
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

