package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnVPC`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCProps := &CfnVPCProps{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	EnableDnsHostnames: jsii.Boolean(false),
//   	EnableDnsSupport: jsii.Boolean(false),
//   	InstanceTenancy: jsii.String("instanceTenancy"),
//   	Ipv4IpamPoolId: jsii.String("ipv4IpamPoolId"),
//   	Ipv4NetmaskLength: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html
//
type CfnVPCProps struct {
	// The IPv4 network range for the VPC, in CIDR notation.
	//
	// For example, `10.0.0.0/16` . We modify the specified CIDR block to its canonical form; for example, if you specify `100.68.0.18/18` , we modify it to `100.68.0.0/18` .
	//
	// You must specify either `CidrBlock` or `Ipv4IpamPoolId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether the instances launched in the VPC get DNS hostnames.
	//
	// If enabled, instances in the VPC get DNS hostnames; otherwise, they do not. Disabled by default for nondefault VPCs. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support) .
	//
	// You can only enable DNS hostnames if you've enabled DNS support.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-enablednshostnames
	//
	EnableDnsHostnames interface{} `field:"optional" json:"enableDnsHostnames" yaml:"enableDnsHostnames"`
	// Indicates whether the DNS resolution is supported for the VPC.
	//
	// If enabled, queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is not enabled. Enabled by default. For more information, see [DNS attributes in your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-support) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-enablednssupport
	//
	EnableDnsSupport interface{} `field:"optional" json:"enableDnsSupport" yaml:"enableDnsSupport"`
	// The allowed tenancy of instances launched into the VPC.
	//
	// - `default` : An instance launched into the VPC runs on shared hardware by default, unless you explicitly specify a different tenancy during instance launch.
	// - `dedicated` : An instance launched into the VPC runs on dedicated hardware by default, unless you explicitly specify a tenancy of `host` during instance launch. You cannot specify a tenancy of `default` during instance launch.
	//
	// Updating `InstanceTenancy` requires no replacement only if you are updating its value from `dedicated` to `default` . Updating `InstanceTenancy` from `default` to `dedicated` requires replacement.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-instancetenancy
	//
	InstanceTenancy *string `field:"optional" json:"instanceTenancy" yaml:"instanceTenancy"`
	// The ID of an IPv4 IPAM pool you want to use for allocating this VPC's CIDR.
	//
	// For more information, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	//
	// You must specify either `CidrBlock` or `Ipv4IpamPoolId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-ipv4ipampoolid
	//
	Ipv4IpamPoolId *string `field:"optional" json:"ipv4IpamPoolId" yaml:"ipv4IpamPoolId"`
	// The netmask length of the IPv4 CIDR you want to allocate to this VPC from an Amazon VPC IP Address Manager (IPAM) pool.
	//
	// For more information about IPAM, see [What is IPAM?](https://docs.aws.amazon.com//vpc/latest/ipam/what-is-it-ipam.html) in the *Amazon VPC IPAM User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-ipv4netmasklength
	//
	Ipv4NetmaskLength *float64 `field:"optional" json:"ipv4NetmaskLength" yaml:"ipv4NetmaskLength"`
	// The tags for the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc.html#cfn-ec2-vpc-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

