package awsmediaconnect


// Properties for defining a `CfnBridgeSource`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBridgeSourceProps := &CfnBridgeSourceProps{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	FlowSource: &BridgeFlowSourceProperty{
//   		FlowArn: jsii.String("flowArn"),
//
//   		// the properties below are optional
//   		FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	NetworkSource: &BridgeNetworkSourceProperty{
//   		MulticastIp: jsii.String("multicastIp"),
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
type CfnBridgeSourceProps struct {
	// The ARN of the bridge that you want to describe.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the network source.
	//
	// This name is used to reference the source and must be unique among sources in this bridge.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Add a flow source to an existing bridge.
	FlowSource interface{} `field:"optional" json:"flowSource" yaml:"flowSource"`
	// Add a network source to an existing bridge.
	NetworkSource interface{} `field:"optional" json:"networkSource" yaml:"networkSource"`
}

