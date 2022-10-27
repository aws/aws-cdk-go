package awsec2


// Cidr Allocated Subnets.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetIpamOptions := &subnetIpamOptions{
//   	allocatedSubnets: []allocatedSubnet{
//   		&allocatedSubnet{
//   			cidr: jsii.String("cidr"),
//   		},
//   	},
//   }
//
type SubnetIpamOptions struct {
	// Cidr Allocations for Subnets.
	AllocatedSubnets *[]*AllocatedSubnet `field:"required" json:"allocatedSubnets" yaml:"allocatedSubnets"`
}

