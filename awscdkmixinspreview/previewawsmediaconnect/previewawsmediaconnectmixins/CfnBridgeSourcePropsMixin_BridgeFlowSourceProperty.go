package previewawsmediaconnectmixins


// The source of the bridge.
//
// A flow source originates in MediaConnect as an existing cloud flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   bridgeFlowSourceProperty := &BridgeFlowSourceProperty{
//   	FlowArn: jsii.String("flowArn"),
//   	FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgesource-bridgeflowsource.html
//
type CfnBridgeSourcePropsMixin_BridgeFlowSourceProperty struct {
	// The ARN of the cloud flow used as a source of this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgesource-bridgeflowsource.html#cfn-mediaconnect-bridgesource-bridgeflowsource-flowarn
	//
	FlowArn *string `field:"optional" json:"flowArn" yaml:"flowArn"`
	// The name of the VPC interface attachment to use for this source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridgesource-bridgeflowsource.html#cfn-mediaconnect-bridgesource-bridgeflowsource-flowvpcinterfaceattachment
	//
	FlowVpcInterfaceAttachment interface{} `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

