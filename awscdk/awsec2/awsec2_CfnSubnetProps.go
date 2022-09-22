package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
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
//   cfnSubnetProps := &cfnSubnetProps{
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	assignIpv6AddressOnCreation: jsii.Boolean(false),
//   	availabilityZone: jsii.String("availabilityZone"),
//   	availabilityZoneId: jsii.String("availabilityZoneId"),
//   	cidrBlock: jsii.String("cidrBlock"),
//   	enableDns64: jsii.Boolean(false),
//   	ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	ipv6Native: jsii.Boolean(false),
//   	mapPublicIpOnLaunch: jsii.Boolean(false),
//   	outpostArn: jsii.String("outpostArn"),
//   	privateDnsNameOptionsOnLaunch: privateDnsNameOptionsOnLaunch,
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnSubnetProps struct {
	// The ID of the VPC the subnet is in.
	//
	// If you update this property, you must also update the `CidrBlock` property.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address. The default value is `false` .
	//
	// If you specify `AssignIpv6AddressOnCreation` , you must also specify `Ipv6CidrBlock` .
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The Availability Zone of the subnet.
	//
	// If you update this property, you must also update the `CidrBlock` property.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the subnet.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The IPv4 CIDR block assigned to the subnet.
	//
	// If you update this property, we create a new subnet, and then delete the existing one.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet should return synthetic IPv6 addresses for IPv4-only destinations.
	//
	// For more information, see [DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-nat64-dns64) in the *Amazon Virtual Private Cloud User Guide* .
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// The IPv6 CIDR block.
	//
	// If you specify `AssignIpv6AddressOnCreation` , you must also specify `Ipv6CidrBlock` .
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Indicates whether this is an IPv6 only subnet.
	//
	// For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the *Amazon Virtual Private Cloud User Guide* .
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// Indicates whether instances launched in this subnet receive a public IPv4 address.
	//
	// The default value is `false` .
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// The Amazon Resource Name (ARN) of the Outpost.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled.
	//
	// For more information, see [Amazon EC2 instance hostname types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-naming.html) in the *Amazon Elastic Compute Cloud User Guide* .
	PrivateDnsNameOptionsOnLaunch interface{} `field:"optional" json:"privateDnsNameOptionsOnLaunch" yaml:"privateDnsNameOptionsOnLaunch"`
	// Any tags assigned to the subnet.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

