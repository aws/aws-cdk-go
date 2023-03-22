package awsec2


// Properties for defining a `CfnVPCCidrBlock`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCCidrBlockProps := &cfnVPCCidrBlockProps{
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	amazonProvidedIpv6CidrBlock: jsii.Boolean(false),
//   	cidrBlock: jsii.String("cidrBlock"),
//   	ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	ipv4NetmaskLength: jsii.Number(123),
//   	ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	ipv6IpamPoolId: jsii.String("ipv6IpamPoolId"),
//   	ipv6NetmaskLength: jsii.Number(123),
//   	ipv6Pool: jsii.String("ipv6Pool"),
//   }
//
type CfnVPCCidrBlockProps struct {
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.
	//
	// You cannot specify the range of IPv6 addresses, or the size of the CIDR block.
	AmazonProvidedIpv6CidrBlock interface{} `field:"optional" json:"amazonProvidedIpv6CidrBlock" yaml:"amazonProvidedIpv6CidrBlock"`
	// An IPv4 CIDR block to associate with the VPC.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Associate a CIDR allocated from an IPv4 IPAM pool to a VPC.
	//
	// For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// An IPv6 CIDR block from the IPv6 address pool. You must also specify `Ipv6Pool` in the request.
	//
	// To let Amazon choose the IPv6 CIDR block for you, omit this parameter.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Associates a CIDR allocated from an IPv6 IPAM pool to a VPC.
	//
	// For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	Ipv6Pool *string `field:"optional" json:"ipv6Pool" yaml:"ipv6Pool"`
}

