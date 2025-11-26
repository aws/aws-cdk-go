package previewawsec2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnVPCPeeringConnectionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCPeeringConnectionMixinProps := &CfnVPCPeeringConnectionMixinProps{
//   	PeerOwnerId: jsii.String("peerOwnerId"),
//   	PeerRegion: jsii.String("peerRegion"),
//   	PeerRoleArn: jsii.String("peerRoleArn"),
//   	PeerVpcId: jsii.String("peerVpcId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html
//
type CfnVPCPeeringConnectionMixinProps struct {
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
	// The ID of the VPC with which you are creating the VPC peering connection.
	//
	// You must specify this parameter in the request.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-peervpcid
	//
	PeerVpcId *string `field:"optional" json:"peerVpcId" yaml:"peerVpcId"`
	// Any tags assigned to the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcpeeringconnection.html#cfn-ec2-vpcpeeringconnection-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

