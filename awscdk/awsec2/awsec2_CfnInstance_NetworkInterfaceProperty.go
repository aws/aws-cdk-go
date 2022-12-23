package awsec2


// Specifies a network interface that is to be attached to an instance.
//
// You can create a network interface when launching an instance. For an example, see the [AWS::EC2::Instance examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#aws-properties-ec2-instance--examples--Automatically_assign_a_public_IP_address) .
//
// Alternatively, you can attach an existing network interface when launching an instance. For an example, see the [AWS::EC2:NetworkInterface examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-interface.html#aws-resource-ec2-network-interface--examples--Basic_network_interface) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &networkInterfaceProperty{
//   	deviceIndex: jsii.String("deviceIndex"),
//
//   	// the properties below are optional
//   	associateCarrierIpAddress: jsii.Boolean(false),
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	deleteOnTermination: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	groupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	ipv6AddressCount: jsii.Number(123),
//   	ipv6Addresses: []interface{}{
//   		&instanceIpv6AddressProperty{
//   			ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	privateIpAddresses: []interface{}{
//   		&privateIpAddressSpecificationProperty{
//   			primary: jsii.Boolean(false),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	secondaryPrivateIpAddressCount: jsii.Number(123),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnInstance_NetworkInterfaceProperty struct {
	// The position of the network interface in the attachment order.
	//
	// A primary network interface has a device index of 0.
	//
	// If you create a network interface when launching an instance, you must specify the device index.
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// `CfnInstance.NetworkInterfaceProperty.AssociateCarrierIpAddress`.
	AssociateCarrierIpAddress interface{} `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Indicates whether to assign a public IPv4 address to an instance.
	//
	// Applies only if creating a network interface when launching an instance. The network interface must be the primary network interface. If launching into a default subnet, the default value is `true` .
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	//
	// Applies only if creating a network interface when launching an instance.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The description of the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IDs of the security groups for the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// A number of IPv6 addresses to assign to the network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of the subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more IPv6 addresses to assign to the network interface.
	//
	// You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The ID of the network interface, when attaching an existing network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The private IPv4 address of the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// One or more private IPv4 addresses to assign to the network interface.
	//
	// Only one private IPv4 address can be designated as primary.
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses.
	//
	// You can't specify this option and specify more than one private IP address using the private IP addresses option.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet associated with the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

