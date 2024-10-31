package awsec2alpha


// Attributes for VPCCidrBlock used for defining a new CIDR Block and also for importing an existing CIDR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//
//   vPCCidrBlockattributes := &VPCCidrBlockattributes{
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(false),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	CidrBlockName: jsii.String("cidrBlockName"),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	Ipv4IpamProvisionedCidrs: []*string{
//   		jsii.String("ipv4IpamProvisionedCidrs"),
//   	},
//   	Ipv4NetmaskLength: jsii.Number(123),
//   	Ipv6IpamPoolId: jsii.String("ipv6IpamPoolId"),
//   	Ipv6NetmaskLength: jsii.Number(123),
//   }
//
// Experimental.
type VPCCidrBlockattributes struct {
	// Amazon Provided Ipv6.
	// Default: false.
	//
	// Experimental.
	AmazonProvidedIpv6CidrBlock *bool `field:"optional" json:"amazonProvidedIpv6CidrBlock" yaml:"amazonProvidedIpv6CidrBlock"`
	// The secondary IPv4 CIDR Block.
	// Default: - no CIDR block provided.
	//
	// Experimental.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// The secondary IPv4 CIDR Block.
	// Default: - no CIDR block provided.
	//
	// Experimental.
	CidrBlockName *string `field:"optional" json:"cidrBlockName" yaml:"cidrBlockName"`
	// IPAM pool for IPv4 address type.
	// Default: - no IPAM pool Id provided for IPv4.
	//
	// Experimental.
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// IPv4 CIDR provisioned under pool Required to check for overlapping CIDRs after provisioning is complete under IPAM pool.
	// Default: - no IPAM IPv4 CIDR range is provisioned using IPAM.
	//
	// Experimental.
	Ipv4IpamProvisionedCidrs *[]*string `field:"optional" json:"ipv4IpamProvisionedCidrs" yaml:"ipv4IpamProvisionedCidrs"`
	// Net mask length for IPv4 address type.
	// Default: - no Net mask length configured for IPv4.
	//
	// Experimental.
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// IPAM pool for IPv6 address type.
	// Default: - no IPAM pool Id provided for IPv6.
	//
	// Experimental.
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Net mask length for IPv6 address type.
	// Default: - no Net mask length configured for IPv6.
	//
	// Experimental.
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
}

