package awsec2


// Configuration for AwsIpam.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var pool cfnIPAMPool
//
//
//   ec2.NewVpc(stack, jsii.String("TheVPC"), &vpcProps{
//   	ipAddresses: ec2.ipAddresses.awsIpamAllocation(&awsIpamProps{
//   		ipv4IpamPoolId: pool.ref,
//   		ipv4NetmaskLength: jsii.Number(18),
//   		defaultSubnetIpv4NetmaskLength: jsii.Number(24),
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
	DefaultSubnetIpv4NetmaskLength *float64 `field:"optional" json:"defaultSubnetIpv4NetmaskLength" yaml:"defaultSubnetIpv4NetmaskLength"`
}

