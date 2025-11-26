package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVPCBlockPublicAccessExclusionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCBlockPublicAccessExclusionMixinProps := &CfnVPCBlockPublicAccessExclusionMixinProps{
//   	InternetGatewayExclusionMode: jsii.String("internetGatewayExclusionMode"),
//   	SubnetId: jsii.String("subnetId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html
//
type CfnVPCBlockPublicAccessExclusionMixinProps struct {
	// The desired VPC Block Public Access mode for a specific VPC or subnet exclusion.
	//
	// - `allow-bidirectional` : Allow all internet traffic to and from the excluded VPCs and subnets.
	// - `allow-egress` : Allow outbound internet traffic from the excluded VPCs and subnets. Block inbound internet traffic to the excluded VPCs and subnets. Only applies when VPC Block Public Access is set to `block-bidirectional` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-internetgatewayexclusionmode
	//
	InternetGatewayExclusionMode *string `field:"optional" json:"internetGatewayExclusionMode" yaml:"internetGatewayExclusionMode"`
	// The ID of the subnet you want to exclude.
	//
	// Required only if you don't specify VpcId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC you want to exclude.
	//
	// Required only if you don't specify SubnetId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcblockpublicaccessexclusion.html#cfn-ec2-vpcblockpublicaccessexclusion-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

