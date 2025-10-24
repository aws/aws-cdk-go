package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Consolidated return parameters to pass to VPC construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   var cfnResource CfnResource
//   var ipamPool IIpamPool
//
//   vpcCidrOptions := &VpcCidrOptions{
//   	AmazonProvided: jsii.Boolean(false),
//   	CidrBlockName: jsii.String("cidrBlockName"),
//   	Dependencies: []CfnResource{
//   		cfnResource,
//   	},
//   	Ipv4CidrBlock: jsii.String("ipv4CidrBlock"),
//   	Ipv4IpamPool: ipamPool,
//   	Ipv4IpamProvisionedCidrs: []*string{
//   		jsii.String("ipv4IpamProvisionedCidrs"),
//   	},
//   	Ipv4NetmaskLength: jsii.Number(123),
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	Ipv6IpamPool: ipamPool,
//   	Ipv6NetmaskLength: jsii.Number(123),
//   	Ipv6PoolId: jsii.String("ipv6PoolId"),
//   }
//
// Experimental.
type VpcCidrOptions struct {
	// Use amazon provided IP range.
	// Default: false.
	//
	// Experimental.
	AmazonProvided *bool `field:"optional" json:"amazonProvided" yaml:"amazonProvided"`
	// Required to set Secondary cidr block resource name in order to generate unique logical id for the resource.
	// Default: - no name for primary addresses.
	//
	// Experimental.
	CidrBlockName *string `field:"optional" json:"cidrBlockName" yaml:"cidrBlockName"`
	// Dependency to associate Ipv6 CIDR block.
	// Default: - No dependency.
	//
	// Experimental.
	Dependencies *[]awscdk.CfnResource `field:"optional" json:"dependencies" yaml:"dependencies"`
	// IPv4 CIDR Block.
	// Default: '10.0.0.0/16'
	//
	// Experimental.
	Ipv4CidrBlock *string `field:"optional" json:"ipv4CidrBlock" yaml:"ipv4CidrBlock"`
	// Ipv4 IPAM Pool.
	// Default: - Only required when using IPAM Ipv4.
	//
	// Experimental.
	Ipv4IpamPool IIpamPool `field:"optional" json:"ipv4IpamPool" yaml:"ipv4IpamPool"`
	// IPv4 CIDR provisioned under pool Required to check for overlapping CIDRs after provisioning is complete under IPAM pool.
	// Default: - no IPAM IPv4 CIDR range is provisioned using IPAM.
	//
	// Experimental.
	Ipv4IpamProvisionedCidrs *[]*string `field:"optional" json:"ipv4IpamProvisionedCidrs" yaml:"ipv4IpamProvisionedCidrs"`
	// CIDR Mask for Vpc.
	// Default: - Only required when using IPAM Ipv4.
	//
	// Experimental.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// IPv6 CIDR block from the BOYIP IPv6 address pool.
	// Default: - None.
	//
	// Experimental.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Ipv6 IPAM pool id for VPC range, can only be defined under public scope.
	// Default: - no pool id.
	//
	// Experimental.
	Ipv6IpamPool IIpamPool `field:"optional" json:"ipv6IpamPool" yaml:"ipv6IpamPool"`
	// CIDR Mask for Vpc.
	// Default: - Only required when using AWS Ipam.
	//
	// Experimental.
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// ID of the BYOIP IPv6 address pool from which to allocate the IPv6 CIDR block.
	// Default: - None.
	//
	// Experimental.
	Ipv6PoolId *string `field:"optional" json:"ipv6PoolId" yaml:"ipv6PoolId"`
}

