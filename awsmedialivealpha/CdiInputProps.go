package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Properties for a CDI (uncompressed) input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var role Role
//   var securityGroupRef ISecurityGroupRef
//   var subnetRef ISubnetRef
//
//   cdiInputProps := &CdiInputProps{
//   	Subnets: []ISubnetRef{
//   		subnetRef,
//   	},
//
//   	// the properties below are optional
//   	Role: role,
//   	SecurityGroups: []ISecurityGroupRef{
//   		securityGroupRef,
//   	},
//   }
//
// Experimental.
type CdiInputProps struct {
	// Two VPC subnets, in two different availability zones, for the CDI input network interfaces.
	// Experimental.
	Subnets *[]interfacesawsec2.ISubnetRef `field:"required" json:"subnets" yaml:"subnets"`
	// The IAM role MediaLive assumes to create network interfaces in the VPC.
	//
	// When omitted, a role is created and granted the required permissions.
	// When you provide a role, no permissions are added — you own all the permissions it needs.
	// Default: - a role is created with the medialive.amazonaws.com service principal
	//
	// [disable-awslint:prefer-ref-interface].
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Security groups to attach to the CDI input network interfaces.
	// Default: - VPC default security group.
	//
	// Experimental.
	SecurityGroups *[]interfacesawsec2.ISecurityGroupRef `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

