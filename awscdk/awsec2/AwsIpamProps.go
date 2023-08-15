package awsec2


// Configuration for AwsIpam.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pool cfnIPAMPool
//
//
//   ec2.NewVpc(this, jsii.String("TheVPC"), &VpcProps{
//   	IpAddresses: awscdk.IpAddresses_AwsIpamAllocation(&AwsIpamProps{
//   		Ipv4IpamPoolId: pool.ref,
//   		Ipv4NetmaskLength: jsii.Number(18),
//   		DefaultSubnetIpv4NetmaskLength: jsii.Number(24),
//   	}),
//   })
//
type AwsIpamProps struct {
	// Ipam Pool Id for ipv4 allocation.
	Ipv4IpamPoolId *string `field:"required" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// Netmask length for Vpc.
	Ipv4NetmaskLength *float64 `field:"required" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// Default length for Subnet ipv4 Network mask.
	//
	// Specify this option only if you do not specify all Subnets using SubnetConfiguration with a cidrMask.
	// Default: - Default ipv4 Subnet Mask for subnets in Vpc.
	//
	DefaultSubnetIpv4NetmaskLength *float64 `field:"optional" json:"defaultSubnetIpv4NetmaskLength" yaml:"defaultSubnetIpv4NetmaskLength"`
}

