package awsmedialivealpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawselasticache"
)

// VPC output settings for the channel.
//
// When configured, all output endpoints are created within the specified VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroupRef ISecurityGroupRef
//   var subnetRef ISubnetRef
//
//   vpcOutputSettings := &VpcOutputSettings{
//   	Subnets: []ISubnetRef{
//   		subnetRef,
//   	},
//
//   	// the properties below are optional
//   	PublicAddressAllocationIds: []*string{
//   		jsii.String("publicAddressAllocationIds"),
//   	},
//   	SecurityGroups: []ISecurityGroupRef{
//   		securityGroupRef,
//   	},
//   }
//
// Experimental.
type VpcOutputSettings struct {
	// The subnets to use for the channel's output endpoints.
	//
	// For STANDARD channels, provide subnets in two different availability zones.
	// For SINGLE_PIPELINE channels, provide at least one subnet.
	// Experimental.
	Subnets *[]interfacesawsec2.ISubnetRef `field:"required" json:"subnets" yaml:"subnets"`
	// Public address allocation IDs to associate with ENIs created in the output VPC.
	//
	// Must specify one for SINGLE_PIPELINE, two for STANDARD channels.
	// Default: - no public addresses.
	//
	// Experimental.
	PublicAddressAllocationIds *[]*string `field:"optional" json:"publicAddressAllocationIds" yaml:"publicAddressAllocationIds"`
	// The security groups to attach to the output VPC network interfaces.
	// Default: - VPC default security group.
	//
	// Experimental.
	SecurityGroups *[]interfacesawselasticache.ISecurityGroupRef `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

