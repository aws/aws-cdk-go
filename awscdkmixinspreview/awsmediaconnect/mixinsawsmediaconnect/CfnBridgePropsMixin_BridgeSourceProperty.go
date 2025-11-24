package mixinsawsmediaconnect


// The bridge's source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bridgeSourceProperty := &BridgeSourceProperty{
//   	FlowSource: &BridgeFlowSourceProperty{
//   		FlowArn: jsii.String("flowArn"),
//   		FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   			VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   		},
//   		Name: jsii.String("name"),
//   	},
//   	NetworkSource: &BridgeNetworkSourceProperty{
//   		MulticastIp: jsii.String("multicastIp"),
//   		MulticastSourceSettings: &MulticastSourceSettingsProperty{
//   			MulticastSourceIp: jsii.String("multicastSourceIp"),
//   		},
//   		Name: jsii.String("name"),
//   		NetworkName: jsii.String("networkName"),
//   		Port: jsii.Number(123),
//   		Protocol: jsii.String("protocol"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgesource.html
//
type CfnBridgePropsMixin_BridgeSourceProperty struct {
	// The source of the bridge.
	//
	// A flow source originates in MediaConnect as an existing cloud flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgesource.html#cfn-mediaconnect-bridge-bridgesource-flowsource
	//
	FlowSource interface{} `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge.
	//
	// A network source originates at your premises.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgesource.html#cfn-mediaconnect-bridge-bridgesource-networksource
	//
	NetworkSource interface{} `field:"optional" json:"networkSource" yaml:"networkSource"`
}

