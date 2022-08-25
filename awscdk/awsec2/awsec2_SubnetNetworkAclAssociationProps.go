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
//   subnetNetworkAclAssociationProps := &subnetNetworkAclAssociationProps{
//   	networkAcl: networkAcl,
//   	subnet: subnet,
//
//   	// the properties below are optional
//   	subnetNetworkAclAssociationName: jsii.String("subnetNetworkAclAssociationName"),
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
	SubnetNetworkAclAssociationName *string `field:"optional" json:"subnetNetworkAclAssociationName" yaml:"subnetNetworkAclAssociationName"`
}

