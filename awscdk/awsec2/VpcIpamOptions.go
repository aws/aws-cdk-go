package awsec2


// CIDR Allocated Vpc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcIpamOptions := &VpcIpamOptions{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	Ipv4NetmaskLength: jsii.Number(123),
//   }
//
type VpcIpamOptions struct {
	// CIDR Block for Vpc.
	// Default: - Only required when Ipam has concrete allocation available for static Vpc.
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// ipv4 IPAM Pool Id.
	// Default: - Only required when using AWS Ipam.
	//
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// CIDR Mask for Vpc.
	// Default: - Only required when using AWS Ipam.
	//
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
}

