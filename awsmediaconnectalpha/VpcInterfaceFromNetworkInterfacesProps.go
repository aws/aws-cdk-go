package awsmediaconnectalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for creating a VPC Interface from existing network interfaces.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediaconnect_alpha "github.com/aws/aws-cdk-go/awsmediaconnectalpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var securityGroup SecurityGroup
//   var subnet Subnet
//
//   vpcInterfaceFromNetworkInterfacesProps := &VpcInterfaceFromNetworkInterfacesProps{
//   	NetworkInterfaceIds: []*string{
//   		jsii.String("networkInterfaceIds"),
//   	},
//   	Role: role,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Subnet: subnet,
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   }
//
// Experimental.
type VpcInterfaceFromNetworkInterfacesProps struct {
	// IDs of the pre-created network interfaces to reuse.
	// Experimental.
	NetworkInterfaceIds *[]*string `field:"required" json:"networkInterfaceIds" yaml:"networkInterfaceIds"`
	// IAM role that MediaConnect assumes to access the ENIs.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Security groups applied to the existing ENIs.
	// Experimental.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Subnet where the existing ENIs live.
	//
	// Must be in the same Availability Zone as the flow.
	// Experimental.
	Subnet awsec2.ISubnet `field:"required" json:"subnet" yaml:"subnet"`
	// Unique name for this VPC interface within the flow.
	//
	// Cannot be changed after creation.
	// Experimental.
	VpcInterfaceName *string `field:"required" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

