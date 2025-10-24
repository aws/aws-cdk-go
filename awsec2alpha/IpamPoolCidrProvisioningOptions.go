package awsec2alpha


// Options to provision CIDRs to an IPAM pool.
//
// Used to create a new IpamPoolCidr.
//
// Example:
//   stack := awscdk.Newstack()
//   ipam := awsec2alpha.NewIpam(this, jsii.String("Ipam"), &IpamProps{
//   	OperatingRegions: []*string{
//   		jsii.String("us-west-1"),
//   	},
//   })
//   ipamPublicPool := ipam.PublicScope.AddPool(jsii.String("PublicPoolA"), &PoolOptions{
//   	AddressFamily: awsec2alpha.AddressFamily_IP_V6,
//   	AwsService: awsec2alpha.AwsServiceName_EC2,
//   	Locale: jsii.String("us-west-1"),
//   	PublicIpSource: awsec2alpha.IpamPoolPublicIpSource_AMAZON,
//   })
//   ipamPublicPool.ProvisionCidr(jsii.String("PublicPoolACidrA"), &IpamPoolCidrProvisioningOptions{
//   	NetmaskLength: jsii.Number(52),
//   })
//
//   ipamPrivatePool := ipam.PrivateScope.AddPool(jsii.String("PrivatePoolA"), &PoolOptions{
//   	AddressFamily: awsec2alpha.AddressFamily_IP_V4,
//   })
//   ipamPrivatePool.ProvisionCidr(jsii.String("PrivatePoolACidrA"), &IpamPoolCidrProvisioningOptions{
//   	NetmaskLength: jsii.Number(8),
//   })
//
//   awsec2alpha.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	PrimaryAddressBlock: awsec2alpha.IpAddresses_Ipv4(jsii.String("10.0.0.0/24")),
//   	SecondaryAddressBlocks: []IIpAddresses{
//   		awsec2alpha.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonIpv6"),
//   		}),
//   		awsec2alpha.IpAddresses_Ipv6Ipam(&IpamOptions{
//   			IpamPool: ipamPublicPool,
//   			NetmaskLength: jsii.Number(52),
//   			CidrBlockName: jsii.String("ipv6Ipam"),
//   		}),
//   		awsec2alpha.IpAddresses_Ipv4Ipam(&IpamOptions{
//   			IpamPool: ipamPrivatePool,
//   			NetmaskLength: jsii.Number(8),
//   			CidrBlockName: jsii.String("ipv4Ipam"),
//   		}),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampoolcidr.html
//
// Experimental.
type IpamPoolCidrProvisioningOptions struct {
	// Ipv6 CIDR block for the IPAM pool.
	// Default: - pool provisioned without netmask length, need netmask length in this case.
	//
	// Experimental.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Ipv6 Netmask length for the CIDR.
	// Default: - pool provisioned without netmask length, need cidr range in this case.
	//
	// Experimental.
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

