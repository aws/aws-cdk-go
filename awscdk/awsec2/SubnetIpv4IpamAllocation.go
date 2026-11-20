package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Allocation of a subnet's IPv4 CIDR block from an Amazon VPC IP Address Manager (IPAM) pool.
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
type SubnetIpv4IpamAllocation struct {
	// The IPAM pool from which the CIDR block is allocated at deploy time.
	//
	// The pool must be a resource planning pool for the VPC the subnet belongs to (its
	// `sourceResource` is the VPC), and the allocated CIDR block must lie within the CIDR of
	// that VPC. Pools in a private scope require the IPAM Advanced Tier.
	// To reference a pool that is not defined in this app, use `CfnIPAMPool.fromIpamPoolId()`.
	IpamPool interfacesawsec2.IIPAMPoolRef `field:"required" json:"ipamPool" yaml:"ipamPool"`
	// The netmask length of the CIDR block to allocate from the pool.
	//
	// Must be between 16 and 28 (inclusive), the sizes allowed for an IPv4 subnet.
	NetmaskLength *float64 `field:"required" json:"netmaskLength" yaml:"netmaskLength"`
}

