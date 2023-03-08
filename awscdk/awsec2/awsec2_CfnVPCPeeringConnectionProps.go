package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnVPCPeeringConnection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCPeeringConnectionProps := &cfnVPCPeeringConnectionProps{
//   	peerVpcId: jsii.String("peerVpcId"),
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	peerOwnerId: jsii.String("peerOwnerId"),
//   	peerRegion: jsii.String("peerRegion"),
//   	peerRoleArn: jsii.String("peerRoleArn"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnVPCPeeringConnectionProps struct {
	// The ID of the VPC with which you are creating the VPC peering connection.
	//
	// You must specify this parameter in the request.
	PeerVpcId *string `field:"required" json:"peerVpcId" yaml:"peerVpcId"`
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The AWS account ID of the owner of the accepter VPC.
	//
	// Default: Your AWS account ID.
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// The Region code for the accepter VPC, if the accepter VPC is located in a Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// The Amazon Resource Name (ARN) of the VPC peer role for the peering connection in another AWS account.
	//
	// This is required when you are peering a VPC in a different AWS account.
	PeerRoleArn *string `field:"optional" json:"peerRoleArn" yaml:"peerRoleArn"`
	// Any tags assigned to the resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

