package awsec2


// Specifies the parameters for a network interface.
//
// `NetworkInterface` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-launchtemplatedata.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &networkInterfaceProperty{
//   	associateCarrierIpAddress: jsii.Boolean(false),
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	deleteOnTermination: jsii.Boolean(false),
//   	description: jsii.String("description"),
//   	deviceIndex: jsii.Number(123),
//   	groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	interfaceType: jsii.String("interfaceType"),
//   	ipv4PrefixCount: jsii.Number(123),
//   	ipv4Prefixes: []interface{}{
//   		&ipv4PrefixSpecificationProperty{
//   			ipv4Prefix: jsii.String("ipv4Prefix"),
//   		},
//   	},
//   	ipv6AddressCount: jsii.Number(123),
//   	ipv6Addresses: []interface{}{
//   		&ipv6AddProperty{
//   			ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	ipv6PrefixCount: jsii.Number(123),
//   	ipv6Prefixes: []interface{}{
//   		&ipv6PrefixSpecificationProperty{
//   			ipv6Prefix: jsii.String("ipv6Prefix"),
//   		},
//   	},
//   	networkCardIndex: jsii.Number(123),
//   	networkInterfaceId: jsii.String("networkInterfaceId"),
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	privateIpAddresses: []interface{}{
//   		&privateIpAddProperty{
//   			primary: jsii.Boolean(false),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	secondaryPrivateIpAddressCount: jsii.Number(123),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnLaunchTemplate_NetworkInterfaceProperty struct {
	// Indicates whether to associate a Carrier IP address with eth0 for a new network interface.
	//
	// Use this option when you launch an instance in a Wavelength Zone and want to associate a Carrier IP address with the network interface. For more information about Carrier IP addresses, see [Carrier IP addresses](https://docs.aws.amazon.com/wavelength/latest/developerguide/how-wavelengths-work.html#provider-owned-ip) in the *AWS Wavelength Developer Guide* .
	AssociateCarrierIpAddress interface{} `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Associates a public IPv4 address with eth0 for a new network interface.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// A description for the network interface.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The device index for the network interface attachment.
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The IDs of one or more security groups.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The type of network interface.
	//
	// To create an Elastic Fabric Adapter (EFA), specify `efa` . For more information, see [Elastic Fabric Adapter](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you are not creating an EFA, specify `interface` or omit this parameter.
	//
	// Valid values: `interface` | `efa`.
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv4 prefixes to be automatically assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv4Prefix` option.
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// One or more IPv4 prefixes to be assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv4PrefixCount` option.
	Ipv4Prefixes interface{} `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Amazon EC2 automatically selects the IPv6 addresses from the subnet range. You can't use this option if specifying specific IPv6 addresses.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.
	//
	// You can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The number of IPv6 prefixes to be automatically assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv6Prefix` option.
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// One or more IPv6 prefixes to be assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv6PrefixCount` option.
	Ipv6Prefixes interface{} `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// The index of the network card.
	//
	// Some instance types support multiple network cards. The primary network interface must be assigned to network card index 0. The default is network card index 0.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// The ID of the network interface.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The primary private IPv4 address of the network interface.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// One or more private IPv4 addresses.
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet for the network interface.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

