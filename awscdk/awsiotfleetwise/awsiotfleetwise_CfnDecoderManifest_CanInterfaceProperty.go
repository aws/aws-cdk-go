package awsiotfleetwise


// A single controller area network (CAN) device interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   canInterfaceProperty := &canInterfaceProperty{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	protocolName: jsii.String("protocolName"),
//   	protocolVersion: jsii.String("protocolVersion"),
//   }
//
type CfnDecoderManifest_CanInterfaceProperty struct {
	// The unique name of the interface.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the communication protocol for the interface.
	ProtocolName *string `field:"optional" json:"protocolName" yaml:"protocolName"`
	// The version of the communication protocol for the interface.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
}

