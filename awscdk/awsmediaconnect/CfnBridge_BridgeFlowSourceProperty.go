package awsmediaconnect


// The source of the bridge.
//
// A flow source originates in MediaConnect as an existing cloud flow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bridgeFlowSourceProperty := &BridgeFlowSourceProperty{
//   	FlowArn: jsii.String("flowArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	FlowVpcInterfaceAttachment: &VpcInterfaceAttachmentProperty{
//   		VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html
//
type CfnBridge_BridgeFlowSourceProperty struct {
	// The ARN of the cloud flow used as a source of this bridge.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-flowarn
	//
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The name of the flow source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the VPC interface attachment to use for this source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-bridgeflowsource.html#cfn-mediaconnect-bridge-bridgeflowsource-flowvpcinterfaceattachment
	//
	FlowVpcInterfaceAttachment interface{} `field:"optional" json:"flowVpcInterfaceAttachment" yaml:"flowVpcInterfaceAttachment"`
}

