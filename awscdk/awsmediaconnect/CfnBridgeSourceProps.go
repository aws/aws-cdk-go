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
//
//   		// the properties below are optional
//   		MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   			MulticastSourceIp: jsii.String("multicastSourceIp"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html
//
type CfnBridgeSourceProps struct {
	// The ARN of the bridge feeding this flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-bridgearn
	//
	BridgeArn *string `field:"required" json:"bridgeArn" yaml:"bridgeArn"`
	// The name of the flow source.
	//
	// This name is used to reference the source and must be unique among sources in this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The source of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-flowsource
	//
	FlowSource interface{} `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-mediaconnect-bridgesource.html#cfn-mediaconnect-bridgesource-networksource
	//
	NetworkSource interface{} `field:"optional" json:"networkSource" yaml:"networkSource"`
}

