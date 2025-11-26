package previewawsmediaconnectmixins


// The source configuration for cloud flows receiving a stream from a bridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   gatewayBridgeSourceProperty := &GatewayBridgeSourceProperty{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	VpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html
//
type CfnFlowSourcePropsMixin_GatewayBridgeSourceProperty struct {
	// The ARN of the bridge feeding this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html#cfn-mediaconnect-flowsource-gatewaybridgesource-bridgearn
	//
	BridgeArn *string `field:"optional" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the VPC interface attachment to use for this bridge source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-gatewaybridgesource.html#cfn-mediaconnect-flowsource-gatewaybridgesource-vpcinterfaceattachment
	//
	VpcInterfaceAttachment interface{} `field:"optional" json:"vpcInterfaceAttachment" yaml:"vpcInterfaceAttachment"`
}

