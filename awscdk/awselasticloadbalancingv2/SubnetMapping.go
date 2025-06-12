package awselasticloadbalancingv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Specifies a subnet for a load balancer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sourceNatIpv6Prefix sourceNatIpv6Prefix
//   var subnet subnet
//
//   subnetMapping := &SubnetMapping{
//   	Subnet: subnet,
//
//   	// the properties below are optional
//   	AllocationId: jsii.String("allocationId"),
//   	Ipv6Address: jsii.String("ipv6Address"),
//   	PrivateIpv4Address: jsii.String("privateIpv4Address"),
//   	SourceNatIpv6Prefix: sourceNatIpv6Prefix,
//   }
//
type SubnetMapping struct {
	// The subnet.
	Subnet awsec2.ISubnet `field:"required" json:"subnet" yaml:"subnet"`
	// The allocation ID of the Elastic IP address for an internet-facing load balancer.
	// Default: undefined - AWS default is to allocate a new IP address for internet-facing load balancers.
	//
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// The IPv6 address.
	// Default: undefined - AWS default is to allocate an IPv6 address from the subnet's pool.
	//
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// The private IPv4 address for an internal load balancer.
	// Default: undefined - AWS default is to allocate a private IPv4 address from the subnet's pool.
	//
	PrivateIpv4Address *string `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
	// The IPv6 prefix to use for source NAT for a dual-stack network load balancer with UDP listeners.
	//
	// Specify an IPv6 prefix (/80 netmask) from the subnet CIDR block
	// or `SourceNatIpv6Prefix.autoAssigned()` to use an IPv6 prefix selected at random from the subnet CIDR block.
	// Default: undefined - AWS default is `SourceNatIpv6Prefix.autoAssigned()` for IPv6 load balancers
	//
	SourceNatIpv6Prefix SourceNatIpv6Prefix `field:"optional" json:"sourceNatIpv6Prefix" yaml:"sourceNatIpv6Prefix"`
}

