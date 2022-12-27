package awsec2


// Describes a network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceNetworkInterfaceSpecificationProperty := &instanceNetworkInterfaceSpecificationProperty{
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	deleteOnTermination: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	deviceIndex: jsii.Number(123),
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	ipv6AddressCount: jsii.Number(123),
//   	ipv6Addresses: []interface{}{
//   		&instanceIpv6AddressProperty{
//   			ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	privateIpAddresses: []interface{}{
//   		&privateIpAddressSpecificationProperty{
//   			privateIpAddress: jsii.String("privateIpAddress"),
//
//   			// the properties below are optional
//   			primary: jsii.Boolean(false),
//   		},
//   	},
//   	secondaryPrivateIpAddressCount: jsii.Number(123),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnSpotFleet_InstanceNetworkInterfaceSpecificationProperty struct {
	// Indicates whether to assign a public IPv4 address to an instance you launch in a VPC.
	//
	// The public IP address can only be assigned to a network interface for eth0, and can only be assigned to a new network interface, not an existing one. You cannot specify more than one network interface in the request. If launching into a default subnet, the default value is `true` .
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The description of the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The position of the network interface in the attachment order.
	//
	// A primary network interface has a device index of 0.
	//
	// If you specify a network interface when launching an instance, you must specify the device index.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The IDs of the security groups for the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// A number of IPv6 addresses to assign to the network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of the subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more IPv6 addresses to assign to the network interface.
	//
	// You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The ID of the network interface.
	//
	// If you are creating a Spot Fleet, omit this parameter because you can’t specify a network interface ID in a launch specification.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// One or more private IPv4 addresses to assign to the network interface.
	//
	// Only one private IPv4 address can be designated as primary. You cannot specify this option if you're launching more than one instance in a [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) request.
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses.
	//
	// You can't specify this option and specify more than one private IP address using the private IP addresses option. You cannot specify this option if you're launching more than one instance in a [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) request.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet associated with the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

