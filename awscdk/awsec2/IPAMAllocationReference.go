package awsec2


// A reference to a IPAMAllocation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   iPAMAllocationReference := map[string]*string{
//   	"cidr": jsii.String("cidr"),
//   	"ipamPoolAllocationId": jsii.String("ipamPoolAllocationId"),
//   	"ipamPoolId": jsii.String("ipamPoolId"),
//   }
//
type IPAMAllocationReference struct {
	// The Cidr of the IPAMAllocation resource.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// The IpamPoolAllocationId of the IPAMAllocation resource.
	IpamPoolAllocationId *string `field:"required" json:"ipamPoolAllocationId" yaml:"ipamPoolAllocationId"`
	// The IpamPoolId of the IPAMAllocation resource.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
}

