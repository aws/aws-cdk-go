package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTransitGatewayVpcAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   cfnTransitGatewayVpcAttachmentProps := &CfnTransitGatewayVpcAttachmentProps{
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	TransitGatewayId: jsii.String("transitGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AddSubnetIds: []*string{
//   		jsii.String("addSubnetIds"),
//   	},
//   	Options: options,
//   	RemoveSubnetIds: []*string{
//   		jsii.String("removeSubnetIds"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html
//
type CfnTransitGatewayVpcAttachmentProps struct {
	// The IDs of the subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-subnetids
	//
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// The ID of the transit gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-transitgatewayid
	//
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The IDs of one or more subnets to add.
	//
	// You can specify at most one subnet per Availability Zone.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-addsubnetids
	//
	AddSubnetIds *[]*string `field:"optional" json:"addSubnetIds" yaml:"addSubnetIds"`
	// The VPC attachment options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The IDs of one or more subnets to remove.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-removesubnetids
	//
	RemoveSubnetIds *[]*string `field:"optional" json:"removeSubnetIds" yaml:"removeSubnetIds"`
	// The tags for the VPC attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-transitgatewayvpcattachment.html#cfn-ec2-transitgatewayvpcattachment-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

