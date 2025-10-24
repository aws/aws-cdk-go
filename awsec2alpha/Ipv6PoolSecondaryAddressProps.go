package awsec2alpha


// Additional props needed for BYOIP IPv6 address props.
//
// Example:
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.1.0.0/16")),
//   	SecondaryAddressBlocks: []IIpAddresses{
//   		awsec2alpha.IpAddresses_Ipv6ByoipPool(&Ipv6PoolSecondaryAddressProps{
//   			CidrBlockName: jsii.String("MyByoipCidrBlock"),
//   			Ipv6PoolId: jsii.String("ipv6pool-ec2-someHashValue"),
//   			Ipv6CidrBlock: jsii.String("2001:db8::/32"),
//   		}),
//   	},
//   	EnableDnsHostnames: jsii.Boolean(true),
//   	EnableDnsSupport: jsii.Boolean(true),
//   })
//
// Experimental.
type Ipv6PoolSecondaryAddressProps struct {
	// Required to set Secondary cidr block resource name in order to generate unique logical id for the resource.
	// Experimental.
	CidrBlockName *string `field:"required" json:"cidrBlockName" yaml:"cidrBlockName"`
	// A valid IPv6 CIDR block from the IPv6 address pool onboarded to AWS using BYOIP.
	//
	// The most specific IPv6 address range that you can bring is /48 for CIDRs that are publicly advertisable
	// and /56 for CIDRs that are not publicly advertisable.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-byoip.html#byoip-definitions
	//
	// Experimental.
	Ipv6CidrBlock *string `field:"required" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// ID of the IPv6 address pool from which to allocate the IPv6 CIDR block.
	//
	// Note: BYOIP Pool ID is different from the IPAM Pool ID.
	// To onboard your IPv6 address range to your AWS account please refer to the below documentation.
	// See: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/byoip-onboard.html
	//
	// Experimental.
	Ipv6PoolId *string `field:"required" json:"ipv6PoolId" yaml:"ipv6PoolId"`
}

