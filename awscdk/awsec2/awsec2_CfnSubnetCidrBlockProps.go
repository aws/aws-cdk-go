package awsec2


// Properties for defining a `CfnSubnetCidrBlock`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSubnetCidrBlockProps := &cfnSubnetCidrBlockProps{
//   	ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnSubnetCidrBlockProps struct {
	// The IPv6 network range for the subnet, in CIDR notation. The subnet size must use a /64 prefix length.
	//
	// This parameter is required for an IPv6 only subnet.
	Ipv6CidrBlock *string `field:"required" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The ID of the subnet.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}

