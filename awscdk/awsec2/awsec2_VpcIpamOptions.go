package awsec2


// Cidr Allocated Vpc.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcIpamOptions := &vpcIpamOptions{
//   	cidrBlock: jsii.String("cidrBlock"),
//   	ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	ipv4NetmaskLength: jsii.Number(123),
//   }
//
type VpcIpamOptions struct {
	// Cidr Block for Vpc.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// ipv4 IPAM Pool Id.
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// Cidr Mask for Vpc.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
}

