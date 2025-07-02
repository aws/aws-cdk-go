package awsec2


// CIDR Allocated Subnet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allocatedSubnet := &AllocatedSubnet{
//   	Cidr: jsii.String("cidr"),
//
//   	// the properties below are optional
//   	Ipv6Cidr: jsii.String("ipv6Cidr"),
//   }
//
type AllocatedSubnet struct {
	// IPv4 CIDR Allocations for a Subnet.
	//
	// Note this is specific to the IPv4 CIDR.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// IPv6 CIDR Allocations for a Subnet.
	//
	// Note this is specific to the IPv6 CIDR.
	// Default: - no IPV6 CIDR.
	//
	Ipv6Cidr *string `field:"optional" json:"ipv6Cidr" yaml:"ipv6Cidr"`
}

