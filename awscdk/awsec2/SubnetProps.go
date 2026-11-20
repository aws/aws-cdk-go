package awsec2


// Specify configuration parameters for a VPC subnet.
//
// Example:
//   var vpc Vpc
//   var ipam CfnIPAM
//
//
//   pool := ec2.NewCfnIPAMPool(this, jsii.String("SubnetPool"), &CfnIPAMPoolProps{
//   	AddressFamily: jsii.String("ipv4"),
//   	IpamScopeId: ipam.attrPrivateDefaultScopeId,
//   	Locale: this.Region,
//   	SourceResource: &SourceResourceProperty{
//   		ResourceId: vpc.VpcId,
//   		ResourceOwner: this.Account,
//   		ResourceRegion: this.*Region,
//   		ResourceType: jsii.String("vpc"),
//   	},
//   	ProvisionedCidrs: []interface{}{
//   		&ProvisionedCidrProperty{
//   			Cidr: vpc.VpcCidrBlock,
//   		},
//   	},
//   })
//
//   subnet := ec2.NewSubnet(this, jsii.String("IpamSubnet"), &SubnetProps{
//   	VpcId: vpc.*VpcId,
//   	AvailabilityZone: vpc.AvailabilityZones[jsii.Number(0)],
//   	Ipv4IpamAllocation: &SubnetIpv4IpamAllocation{
//   		IpamPool: pool,
//   		NetmaskLength: jsii.Number(24),
//   	},
//   })
//
type SubnetProps struct {
	// The availability zone for the subnet.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// The VPC which this subnet is part of.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Indicates whether a network interface created in this subnet receives an IPv6 address.
	//
	// If you specify AssignIpv6AddressOnCreation, you must also specify Ipv6CidrBlock.
	// Default: false.
	//
	AssignIpv6AddressOnCreation *bool `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// The CIDR notation for this subnet.
	//
	// Exactly one of `cidrBlock` and `ipv4IpamAllocation` must be specified.
	// Default: - the CIDR block is allocated from the IPAM pool given in `ipv4IpamAllocation`.
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Allocate the IPv4 CIDR block of this subnet from an IPAM pool at deploy time.
	//
	// When set, `ipv4CidrBlock` resolves to a deploy-time attribute of the subnet instead of a
	// concrete string, so `SubnetFilter.byCidrMask()`, `SubnetFilter.byCidrRanges()` and
	// `SubnetFilter.containsIpAddresses()` cannot be used with this subnet.
	//
	// Exactly one of `cidrBlock` and `ipv4IpamAllocation` must be specified.
	// Default: - the subnet uses the concrete `cidrBlock`.
	//
	Ipv4IpamAllocation *SubnetIpv4IpamAllocation `field:"optional" json:"ipv4IpamAllocation" yaml:"ipv4IpamAllocation"`
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

