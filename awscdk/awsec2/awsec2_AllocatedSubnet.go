package awsec2


// Cidr Allocated Subnet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   allocatedSubnet := &allocatedSubnet{
//   	cidr: jsii.String("cidr"),
//   }
//
type AllocatedSubnet struct {
	// Cidr Allocations for a Subnet.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
}

