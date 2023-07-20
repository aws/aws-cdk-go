package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPCPeeringConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCPeeringConnectionProps := &CfnVPCPeeringConnectionProps{
//   	PeerVpcId: jsii.String("peerVpcId"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	PeerOwnerId: jsii.String("peerOwnerId"),
//   	PeerRegion: jsii.String("peerRegion"),
//   	PeerRoleArn: jsii.String("peerRoleArn"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
//
type CfnVPCPeeringConnectionProps struct {
	// The ID of the VPC with which you are creating the VPC peering connection.
	//
	// You must specify this parameter in the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peervpcid
	//
	PeerVpcId *string `field:"required" json:"peerVpcId" yaml:"peerVpcId"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The AWS account ID of the owner of the accepter VPC.
	//
	// Default: Your AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerownerid
	//
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerregion
	//
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account.
	//
	// This is required when you are peering a VPC in a different AWS account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peerrolearn
	//
	PeerRoleArn *string `field:"optional" json:"peerRoleArn" yaml:"peerRoleArn"`
	// Any tags assigned to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

