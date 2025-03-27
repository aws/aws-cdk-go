package awsec2


// Request for subnet IPv6 CIDRs to be allocated for a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allocateIpv6CidrRequest := &AllocateIpv6CidrRequest{
//   	AllocatedSubnets: []allocatedSubnet{
//   		&allocatedSubnet{
//   			Cidr: jsii.String("cidr"),
//
//   			// the properties below are optional
//   			Ipv6Cidr: jsii.String("ipv6Cidr"),
//   		},
//   	},
//   	Ipv6Cidrs: []*string{
//   		jsii.String("ipv6Cidrs"),
//   	},
//   }
//
type AllocateIpv6CidrRequest struct {
	// List of subnets allocated with IPv4 CIDRs.
	AllocatedSubnets *[]*AllocatedSubnet `field:"required" json:"allocatedSubnets" yaml:"allocatedSubnets"`
	// The IPv6 CIDRs to be allocated to the subnets.
	Ipv6Cidrs *[]*string `field:"required" json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
}

