package awsec2alpha


// The IP address source for pools in the public scope.
//
// Only used for provisioning IP address CIDRs to pools in the public scope.
//
// Example:
//   stack := awscdk.Newstack()
//   ipam := awsec2alpha.NewIpam(this, jsii.String("Ipam"), &IpamProps{
//   	OperatingRegion: []*string{
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
//   	SecondaryAddressBlocks: []iIpAddresses{
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-publicipsource
//
// Experimental.
type IpamPoolPublicIpSource string

const (
	// BYOIP Ipv6 to be registered under IPAM.
	// Experimental.
	IpamPoolPublicIpSource_BYOIP IpamPoolPublicIpSource = "BYOIP"
	// Amazon Provided Ipv6 range.
	// Experimental.
	IpamPoolPublicIpSource_AMAZON IpamPoolPublicIpSource = "AMAZON"
)

