package awsec2


// Specify configuration parameters for a VPC subnet.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetProps := &SubnetProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	CidrBlock: jsii.String("cidrBlock"),
//   	VpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	AssignIpv6AddressOnCreation: jsii.Boolean(false),
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	MapPublicIpOnLaunch: jsii.Boolean(false),
//   }
//
type SubnetProps struct {
	// The availability zone for the subnet.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The CIDR notation for this subnet.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The VPC which this subnet is part of.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address.
	//
	// If you specify AssignIpv6AddressOnCreation, you must also specify Ipv6CidrBlock.
	// Default: false.
	//
	AssignIpv6AddressOnCreation *bool `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The IPv6 CIDR block.
	//
	// If you specify AssignIpv6AddressOnCreation, you must also specify Ipv6CidrBlock.
	// Default: - no IPv6 CIDR block.
	//
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Controls if a public IP is associated to an instance at launch.
	// Default: true in Subnet.Public, false in Subnet.Private or Subnet.Isolated.
	//
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
}

