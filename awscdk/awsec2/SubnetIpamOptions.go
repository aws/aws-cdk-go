package awsec2


// CIDR Allocated Subnets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetIpamOptions := &SubnetIpamOptions{
//   	AllocatedSubnets: []allocatedSubnet{
//   		&allocatedSubnet{
//   			Cidr: jsii.String("cidr"),
//
//   			// the properties below are optional
//   			Ipv6Cidr: jsii.String("ipv6Cidr"),
//   		},
//   	},
//   }
//
type SubnetIpamOptions struct {
	// CIDR Allocations for Subnets.
	AllocatedSubnets *[]*AllocatedSubnet `field:"required" json:"allocatedSubnets" yaml:"allocatedSubnets"`
}

