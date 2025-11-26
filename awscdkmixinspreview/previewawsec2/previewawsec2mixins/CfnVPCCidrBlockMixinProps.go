package previewawsec2mixins


// Properties for CfnVPCCidrBlockPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCCidrBlockMixinProps := &CfnVPCCidrBlockMixinProps{
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(false),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	Ipv4NetmaskLength: jsii.Number(123),
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	Ipv6CidrBlockNetworkBorderGroup: jsii.String("ipv6CidrBlockNetworkBorderGroup"),
//   	Ipv6IpamPoolId: jsii.String("ipv6IpamPoolId"),
//   	Ipv6NetmaskLength: jsii.Number(123),
//   	Ipv6Pool: jsii.String("ipv6Pool"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
//
type CfnVPCCidrBlockMixinProps struct {
	// Requests an Amazon-provided IPv6 CIDR block with a /56 prefix length for the VPC.
	//
	// You cannot specify the range of IPv6 addresses or the size of the CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-amazonprovidedipv6cidrblock
	//
	AmazonProvidedIpv6CidrBlock interface{} `field:"optional" json:"amazonProvidedIpv6CidrBlock" yaml:"amazonProvidedIpv6CidrBlock"`
	// An IPv4 CIDR block to associate with the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Associate a CIDR allocated from an IPv4 IPAM pool to a VPC.
	//
	// For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv4ipampoolid
	//
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv4netmasklength
	//
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// An IPv6 CIDR block from the IPv6 address pool. You must also specify `Ipv6Pool` in the request.
	//
	// To let Amazon choose the IPv6 CIDR block for you, omit this parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv6cidrblock
	//
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The name of the location from which we advertise the IPV6 CIDR block.
	//
	// Use this parameter to limit the CIDR block to this location.
	//
	// You must set `AmazonProvidedIpv6CidrBlock` to `true` to use this parameter.
	//
	// You can have one IPv6 CIDR block association per network border group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv6cidrblocknetworkbordergroup
	//
	Ipv6CidrBlockNetworkBorderGroup *string `field:"optional" json:"ipv6CidrBlockNetworkBorderGroup" yaml:"ipv6CidrBlockNetworkBorderGroup"`
	// Associates a CIDR allocated from an IPv6 IPAM pool to a VPC.
	//
	// For more information about Amazon VPC IP Address Manager (IPAM), see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv6ipampoolid
	//
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// The netmask length of the IPv6 CIDR you would like to associate from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv6netmasklength
	//
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// The ID of an IPv6 address pool from which to allocate the IPv6 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-ipv6pool
	//
	Ipv6Pool *string `field:"optional" json:"ipv6Pool" yaml:"ipv6Pool"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

