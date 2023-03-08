package awsec2


// Properties for defining a `CfnIPAMAllocation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnIPAMAllocationProps := &cfnIPAMAllocationProps{
//   	ipamPoolId: jsii.String("ipamPoolId"),
//
//   	// the properties below are optional
//   	cidr: jsii.String("cidr"),
//   	description: jsii.String("description"),
//   	netmaskLength: jsii.Number(123),
//   }
//
type CfnIPAMAllocationProps struct {
	// The ID of the IPAM pool from which you would like to allocate a CIDR.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// The CIDR you would like to allocate from the IPAM pool. Note the following:.
	//
	// - If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.
	// - If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.
	//
	// Possible values: Any available IPv4 or IPv6 CIDR.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// A description for the allocation.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The netmask length of the CIDR you would like to allocate from the IPAM pool. Note the following:.
	//
	// - If there is no DefaultNetmaskLength allocation rule set on the pool, you must specify either the NetmaskLength or the CIDR.
	// - If the DefaultNetmaskLength allocation rule is set on the pool, you can specify either the NetmaskLength or the CIDR and the DefaultNetmaskLength allocation rule will be ignored.
	//
	// Possible netmask lengths for IPv4 addresses are 0 - 32. Possible netmask lengths for IPv6 addresses are 0 - 128.
	NetmaskLength *float64 `field:"optional" json:"netmaskLength" yaml:"netmaskLength"`
}

