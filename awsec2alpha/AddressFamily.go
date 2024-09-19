package awsec2alpha


// Represents the address family for IP addresses in an IPAM pool.
//
// IP_V4 - Represents the IPv4 address family.
// IP_V6 - Represents the IPv6 address family.
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
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ipampool.html#cfn-ec2-ipampool-addressfamily
//
// Experimental.
type AddressFamily string

const (
	// Represents the IPv4 address family.
	//
	// Allowed under public and private pool.
	// Experimental.
	AddressFamily_IP_V4 AddressFamily = "IP_V4"
	// Represents the IPv6 address family.
	//
	// Only allowed under public pool.
	// Experimental.
	AddressFamily_IP_V6 AddressFamily = "IP_V6"
)

