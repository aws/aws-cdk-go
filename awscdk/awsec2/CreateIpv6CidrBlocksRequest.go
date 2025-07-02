package awsec2


// Request for IPv6 CIDR block to be split up.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   createIpv6CidrBlocksRequest := &CreateIpv6CidrBlocksRequest{
//   	Ipv6SelectedCidr: jsii.String("ipv6SelectedCidr"),
//   	SubnetCount: jsii.Number(123),
//
//   	// the properties below are optional
//   	SizeMask: jsii.String("sizeMask"),
//   }
//
type CreateIpv6CidrBlocksRequest struct {
	// The IPv6 CIDR block string representation.
	Ipv6SelectedCidr *string `field:"required" json:"ipv6SelectedCidr" yaml:"ipv6SelectedCidr"`
	// The number of subnets to assign CIDRs to.
	SubnetCount *float64 `field:"required" json:"subnetCount" yaml:"subnetCount"`
	// Size of the covered bits in the CIDR.
	// Default: - 128 - 64 = /64 CIDR.
	//
	SizeMask *string `field:"optional" json:"sizeMask" yaml:"sizeMask"`
}

