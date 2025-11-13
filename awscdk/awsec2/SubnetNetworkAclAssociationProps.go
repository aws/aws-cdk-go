package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Properties to create a SubnetNetworkAclAssociation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var networkAclRef INetworkAclRef
//   var subnet Subnet
//
//   subnetNetworkAclAssociationProps := &SubnetNetworkAclAssociationProps{
//   	NetworkAcl: networkAclRef,
//   	Subnet: subnet,
//
//   	// the properties below are optional
//   	SubnetNetworkAclAssociationName: jsii.String("subnetNetworkAclAssociationName"),
//   }
//
type SubnetNetworkAclAssociationProps struct {
	// The Network ACL this association is defined for.
	NetworkAcl interfacesawsec2.INetworkAclRef `field:"required" json:"networkAcl" yaml:"networkAcl"`
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

