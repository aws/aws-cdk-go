package mixinsawsiotfleetwise


// A single controller area network (CAN) device interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   canInterfaceProperty := &CanInterfaceProperty{
//   	Name: jsii.String("name"),
//   	ProtocolName: jsii.String("protocolName"),
//   	ProtocolVersion: jsii.String("protocolVersion"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html
//
type CfnDecoderManifestPropsMixin_CanInterfaceProperty struct {
	// The unique name of the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the communication protocol for the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolname
	//
	ProtocolName *string `field:"optional" json:"protocolName" yaml:"protocolName"`
	// The version of the communication protocol for the interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-decodermanifest-caninterface.html#cfn-iotfleetwise-decodermanifest-caninterface-protocolversion
	//
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

