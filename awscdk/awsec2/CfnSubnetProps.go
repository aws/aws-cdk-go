package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSubnet`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var privateDnsNameOptionsOnLaunch interface{}
//
//   cfnSubnetProps := &CfnSubnetProps{
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AssignIpv6AddressOnCreation: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	EnableDns64: jsii.Boolean(false),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	Ipv4NetmaskLength: jsii.Number(123),
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	Ipv6CidrBlocks: []*string{
//   		jsii.String("ipv6CidrBlocks"),
//   	},
//   	Ipv6IpamPoolId: jsii.String("ipv6IpamPoolId"),
//   	Ipv6Native: jsii.Boolean(false),
//   	Ipv6NetmaskLength: jsii.Number(123),
//   	MapPublicIpOnLaunch: jsii.Boolean(false),
//   	OutpostArn: jsii.String("outpostArn"),
//   	PrivateDnsNameOptionsOnLaunch: privateDnsNameOptionsOnLaunch,
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html
//
type CfnSubnetProps struct {
	// The ID of the VPC the subnet is in.
	//
	// If you update this property, you must also update the `CidrBlock` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-vpcid
	//
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is `false` .
	//
	// If you specify `AssignIpv6AddressOnCreation` , you must also specify an IPv6 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-assignipv6addressoncreation
	//
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The Availability Zone of the subnet.
	//
	// If you update this property, you must also update the `CidrBlock` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The IPv4 CIDR block assigned to the subnet.
	//
	// If you update this property, we create a new subnet, and then delete the existing one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
	//
	// For more information, see [DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-nat64-dns64) in the *Amazon Virtual Private Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-enabledns64
	//
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// An IPv4 IPAM pool ID for the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv4ipampoolid
	//
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// An IPv4 netmask length for the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv4netmasklength
	//
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// The IPv6 CIDR block.
	//
	// If you specify `AssignIpv6AddressOnCreation` , you must also specify an IPv6 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6cidrblock
	//
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The IPv6 network ranges for the subnet, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6cidrblocks
	//
	Ipv6CidrBlocks *[]*string `field:"optional" json:"ipv6CidrBlocks" yaml:"ipv6CidrBlocks"`
	// An IPv6 IPAM pool ID for the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6ipampoolid
	//
	Ipv6IpamPoolId *string `field:"optional" json:"ipv6IpamPoolId" yaml:"ipv6IpamPoolId"`
	// Indicates whether this is an IPv6 only subnet.
	//
	// For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *Amazon Virtual Private Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6native
	//
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// An IPv6 netmask length for the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-ipv6netmasklength
	//
	Ipv6NetmaskLength *float64 `field:"optional" json:"ipv6NetmaskLength" yaml:"ipv6NetmaskLength"`
	// Indicates whether instances launched in this subnet receive a public IPv4 address. The default value is `false` .
	//
	// AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-mappubliciponlaunch
	//
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-outpostarn
	//
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled.
	//
	// For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// Available options:
	//
	// - EnableResourceNameDnsAAAARecord (true | false)
	// - EnableResourceNameDnsARecord (true | false)
	// - HostnameType (ip-name | resource-name).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-privatednsnameoptionsonlaunch
	//
	PrivateDnsNameOptionsOnLaunch interface{} `field:"optional" json:"privateDnsNameOptionsOnLaunch" yaml:"privateDnsNameOptionsOnLaunch"`
	// Any tags assigned to the subnet.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnet.html#cfn-ec2-subnet-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

