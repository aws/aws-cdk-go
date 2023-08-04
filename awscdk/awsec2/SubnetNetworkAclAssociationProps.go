package awsec2


// Properties to create a SubnetNetworkAclAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkAcl networkAcl
//   var subnet subnet
//
//   subnetNetworkAclAssociationProps := &SubnetNetworkAclAssociationProps{
//   	NetworkAcl: networkAcl,
//   	Subnet: subnet,
//
//   	// the properties below are optional
//   	SubnetNetworkAclAssociationName: jsii.String("subnetNetworkAclAssociationName"),
//   }
//
type SubnetNetworkAclAssociationProps struct {
	// The Network ACL this association is defined for.
	NetworkAcl INetworkAcl `field:"required" json:"networkAcl" yaml:"networkAcl"`
	// ID of the Subnet.
	Subnet ISubnet `field:"required" json:"subnet" yaml:"subnet"`
	// The name of the SubnetNetworkAclAssociation.
	//
	// It is not recommended to use an explicit name.
	// Default: If you don't specify a SubnetNetworkAclAssociationName, AWS CloudFormation generates a
	// unique physical ID and uses that ID for the group name.
	//
	SubnetNetworkAclAssociationName *string `field:"optional" json:"subnetNetworkAclAssociationName" yaml:"subnetNetworkAclAssociationName"`
}

