package awsec2alpha


// Additional props needed for secondary Address.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	SecondaryAddressBlocks: []iIpAddresses{
//   		vpc_v2.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonProvidedIp"),
//   		}),
//   	},
//   })
//
//   vpc_v2.NewSubnetV2(this, jsii.String("subnetA"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("us-east-1a"),
//   	Ipv4CidrBlock: vpc_v2.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	Ipv6CidrBlock: vpc_v2.NewIpCidr(jsii.String("2a05:d02c:25:4000::/60")),
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   })
//
// Experimental.
type SecondaryAddressProps struct {
	// Required to set Secondary cidr block resource name in order to generate unique logical id for the resource.
	// Experimental.
	CidrBlockName *string `field:"required" json:"cidrBlockName" yaml:"cidrBlockName"`
}

