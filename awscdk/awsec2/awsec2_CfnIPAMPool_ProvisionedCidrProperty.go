package awsec2


// The CIDR provisioned to the IPAM pool.
//
// A CIDR is a representation of an IP address and its associated network mask (or netmask) and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23` . An IPv6 CIDR example is `2001:DB8::/32` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   provisionedCidrProperty := &provisionedCidrProperty{
//   	cidr: jsii.String("cidr"),
//   }
//
type CfnIPAMPool_ProvisionedCidrProperty struct {
	// The CIDR provisioned to the IPAM pool.
	//
	// A CIDR is a representation of an IP address and its associated network mask (or netmask) and refers to a range of IP addresses. An IPv4 CIDR example is `10.24.34.0/23` . An IPv6 CIDR example is `2001:DB8::/32` .
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

