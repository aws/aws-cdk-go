package awsec2alpha


// Options to create a new Ipam in the account.
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
// Experimental.
type IpamProps struct {
	// Name of IPAM that can be used for tagging resource.
	// Default: - If no name provided, no tags will be added to the IPAM.
	//
	// Experimental.
	IpamName *string `field:"optional" json:"ipamName" yaml:"ipamName"`
	// The operating Regions for an IPAM.
	//
	// Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs
	// For more information about operating Regions, see [Create an IPAM](https://docs.aws.amazon.com//vpc/latest/ipam/create-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipam.html#cfn-ec2-ipam-operatingregions
	//
	// Default: - Stack.region if defined in the stack
	//
	// Experimental.
	OperatingRegions *[]*string `field:"optional" json:"operatingRegions" yaml:"operatingRegions"`
}

