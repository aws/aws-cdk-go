package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPCBlockPublicAccessExclusion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCBlockPublicAccessExclusionProps := &CfnVPCBlockPublicAccessExclusionProps{
//   	InternetGatewayExclusionMode: jsii.String("internetGatewayExclusionMode"),
//
//   	// the properties below are optional
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html
//
type CfnVPCBlockPublicAccessExclusionProps struct {
	// The desired Block Public Access Exclusion Mode for a specific VPC/Subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-internetgatewayexclusionmode
	//
	InternetGatewayExclusionMode *string `field:"required" json:"internetGatewayExclusionMode" yaml:"internetGatewayExclusionMode"`
	// The ID of the subnet.
	//
	// Required only if you don't specify VpcId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the vpc.
	//
	// Required only if you don't specify SubnetId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

