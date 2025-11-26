package previewawsmediaconnectmixins


// Properties for CfnBridgeSourcePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnBridgeSourceMixinProps := &CfnBridgeSourceMixinProps{
//   	BridgeArn: jsii.String("bridgeArn"),
//   	FlowSource: &BridgeFlowSourceProperty{
//   		FlowArn: jsii.String("flowArn"),
//   		FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	NetworkSource: &BridgeNetworkSourceProperty{
//   		MulticastIp: jsii.String("multicastIp"),
//   		MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   			MulticastSourceIp: jsii.String("multicastSourceIp"),
//   		},
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html
//
type CfnBridgeSourceMixinProps struct {
	// The ARN of the bridge feeding this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-bridgearn
	//
	BridgeArn *string `field:"optional" json:"bridgeArn" yaml:"bridgeArn"`
	// The source of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-flowsource
	//
	FlowSource interface{} `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The name of the flow source.
	//
	// This name is used to reference the source and must be unique among sources in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The source of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-networksource
	//
	NetworkSource interface{} `field:"optional" json:"networkSource" yaml:"networkSource"`
}

