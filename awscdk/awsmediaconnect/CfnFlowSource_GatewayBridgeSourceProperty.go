package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayBridgeSourceProperty := &GatewayBridgeSourceProperty{
//   	BridgeArn: jsii.String("bridgeArn"),
//
//   	// the properties below are optional
//   	VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
type CfnFlowSource_GatewayBridgeSourceProperty struct {
	// `CfnFlowSource.GatewayBridgeSourceProperty.BridgeArn`.
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// `CfnFlowSource.GatewayBridgeSourceProperty.VpcInterfaceAttachment`.
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

