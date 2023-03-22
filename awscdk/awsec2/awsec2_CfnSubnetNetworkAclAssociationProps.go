package awsec2


// Properties for defining a `CfnSubnetNetworkAclAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetNetworkAclAssociationProps := &cfnSubnetNetworkAclAssociationProps{
//   	networkAclId: jsii.String("networkAclId"),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnSubnetNetworkAclAssociationProps struct {
	// The ID of the network ACL.
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

