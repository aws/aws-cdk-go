package awsec2


// Specify configuration parameters for a single subnet group in a VPC.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subnetConfiguration := &subnetConfiguration{
//   	name: jsii.String("name"),
//   	subnetType: awscdk.Aws_ec2.subnetType_PRIVATE_ISOLATED,
//
//   	// the properties below are optional
//   	cidrMask: jsii.Number(123),
//   	mapPublicIpOnLaunch: jsii.Boolean(false),
//   	reserved: jsii.Boolean(false),
//   }
//
type SubnetConfiguration struct {
	// Logical name for the subnet group.
	//
	// This name can be used when selecting VPC subnets to distinguish
	// between different subnet groups of the same type.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The type of Subnet to configure.
	//
	// The Subnet type will control the ability to route and connect to the
	// Internet.
	SubnetType SubnetType `field:"required" json:"subnetType" yaml:"subnetType"`
	// The number of leading 1 bits in the routing mask.
	//
	// The number of available IP addresses in each subnet of this group
	// will be equal to `2^(32 - cidrMask) - 2`.
	//
	// Valid values are `16--28`.
	CidrMask *float64 `field:"optional" json:"cidrMask" yaml:"cidrMask"`
	// Controls if a public IP is associated to an instance at launch.
	MapPublicIpOnLaunch *bool `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Controls if subnet IP space needs to be reserved.
	//
	// When true, the IP space for the subnet is reserved but no actual
	// resources are provisioned. This space is only dependent on the
	// number of availability zones and on `cidrMask` - all other subnet
	// properties are ignored.
	Reserved *bool `field:"optional" json:"reserved" yaml:"reserved"`
}

