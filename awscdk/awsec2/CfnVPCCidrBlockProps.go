package awsec2


// Properties for defining a `CfnVPCCidrBlock`.
//
// Example:
//   var vpc vpc
//
//
//   func associateSubnetWithV6Cidr(vpc *vpc, count *f64, subnet iSubnet) {
//   	cfnSubnet := *subnet.Node.defaultChild.(cfnSubnet)
//   	cfnSubnet.Ipv6CidrBlock = awscdk.Fn_Select(count, awscdk.Fn_Cidr(awscdk.Fn_Select(jsii.Number(0), vpc.VpcIpv6CidrBlocks), jsii.Number(256), (jsii.Number(128 - 64)).toString()))
//   	cfnSubnet.AssignIpv6AddressOnCreation = true
//   }
//
//   // make an ipv6 cidr
//   ipv6cidr := ec2.NewCfnVPCCidrBlock(this, jsii.String("CIDR6"), &CfnVPCCidrBlockProps{
//   	VpcId: vpc.VpcId,
//   	AmazonProvidedIpv6CidrBlock: jsii.Boolean(true),
//   })
//
//   // connect the ipv6 cidr to all vpc subnets
//   subnetcount := 0
//   subnets := vpc.PublicSubnets.concat(vpc.PrivateSubnets)
//   for _, subnet := range subnets {
//   	// Wait for the ipv6 cidr to complete
//   	subnet.Node.AddDependency(ipv6cidr)
//   	associateSubnetWithV6Cidr(vpc, subnetcount, subnet)
//   	subnetcount = subnetcount + 1
//   }
//
//   cluster := eks.NewCluster(this, jsii.String("hello-eks"), &ClusterProps{
//   	Version: eks.KubernetesVersion_V1_29(),
//   	Vpc: vpc,
//   	IpFamily: eks.IpFamily_IP_V6,
//   	VpcSubnets: []subnetSelection{
//   		&subnetSelection{
//   			Subnets: vpc.*PublicSubnets,
//   		},
//   	},
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
//
type CfnVPCCidrBlockProps struct {
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
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
}

