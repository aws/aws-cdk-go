package awsec2alpha


// Properties to define a VPC peering connection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   var vpcV2 vpcV2
//
//   vPCPeeringConnectionProps := &VPCPeeringConnectionProps{
//   	AcceptorVpc: vpcV2,
//   	RequestorVpc: vpcV2,
//
//   	// the properties below are optional
//   	PeerRoleArn: jsii.String("peerRoleArn"),
//   	VpcPeeringConnectionName: jsii.String("vpcPeeringConnectionName"),
//   }
//
// Experimental.
type VPCPeeringConnectionProps struct {
	// The VPC that is accepting the peering connection.
	// Experimental.
	AcceptorVpc IVpcV2 `field:"required" json:"acceptorVpc" yaml:"acceptorVpc"`
	// The role arn created in the acceptor account.
	// Default: - no peerRoleArn needed if not cross account connection.
	//
	// Experimental.
	PeerRoleArn *string `field:"optional" json:"peerRoleArn" yaml:"peerRoleArn"`
	// The resource name of the peering connection.
	// Default: - peering connection provisioned without any name.
	//
	// Experimental.
	VpcPeeringConnectionName *string `field:"optional" json:"vpcPeeringConnectionName" yaml:"vpcPeeringConnectionName"`
	// The VPC that is requesting the peering connection.
	// Experimental.
	RequestorVpc IVpcV2 `field:"required" json:"requestorVpc" yaml:"requestorVpc"`
}

