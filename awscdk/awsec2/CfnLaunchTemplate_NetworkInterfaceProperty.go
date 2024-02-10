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
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	AssociateCarrierIpAddress: jsii.Boolean(false),
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	ConnectionTrackingSpecification: &ConnectionTrackingSpecificationProperty{
//   		TcpEstablishedTimeout: jsii.Number(123),
//   		UdpStreamTimeout: jsii.Number(123),
//   		UdpTimeout: jsii.Number(123),
//   	},
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	DeviceIndex: jsii.Number(123),
//   	EnaSrdSpecification: &EnaSrdSpecificationProperty{
//   		EnaSrdEnabled: jsii.Boolean(false),
//   		EnaSrdUdpSpecification: &EnaSrdUdpSpecificationProperty{
//   			EnaSrdUdpEnabled: jsii.Boolean(false),
//   		},
//   	},
//   	Groups: []*string{
//   		jsii.String("groups"),
//   	},
//   	InterfaceType: jsii.String("interfaceType"),
//   	Ipv4PrefixCount: jsii.Number(123),
//   	Ipv4Prefixes: []interface{}{
//   		&Ipv4PrefixSpecificationProperty{
//   			Ipv4Prefix: jsii.String("ipv4Prefix"),
//   		},
//   	},
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&Ipv6AddProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	Ipv6PrefixCount: jsii.Number(123),
//   	Ipv6Prefixes: []interface{}{
//   		&Ipv6PrefixSpecificationProperty{
//   			Ipv6Prefix: jsii.String("ipv6Prefix"),
//   		},
//   	},
//   	NetworkCardIndex: jsii.Number(123),
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrimaryIpv6: jsii.Boolean(false),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PrivateIpAddresses: []interface{}{
//   		&PrivateIpAddProperty{
//   			Primary: jsii.Boolean(false),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html
//
type CfnLaunchTemplate_NetworkInterfaceProperty struct {
	// Associates a Carrier IP address with eth0 for a new network interface.
	//
	// Use this option when you launch an instance in a Wavelength Zone and want to associate a Carrier IP address with the network interface. For more information about Carrier IP addresses, see [Carrier IP addresses](https://docs.aws.amazon.com/wavelength/latest/developerguide/how-wavelengths-work.html#provider-owned-ip) in the *AWS Wavelength Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-associatecarrieripaddress
	//
	AssociateCarrierIpAddress interface{} `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Associates a public IPv4 address with eth0 for a new network interface.
	//
	// AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [Amazon VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-associatepublicipaddress
	//
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// A connection tracking specification for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-connectiontrackingspecification
	//
	ConnectionTrackingSpecification interface{} `field:"optional" json:"connectionTrackingSpecification" yaml:"connectionTrackingSpecification"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-deleteontermination
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// A description for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The device index for the network interface attachment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-deviceindex
	//
	DeviceIndex *float64 `field:"optional" json:"deviceIndex" yaml:"deviceIndex"`
	// The ENA Express configuration for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-enasrdspecification
	//
	EnaSrdSpecification interface{} `field:"optional" json:"enaSrdSpecification" yaml:"enaSrdSpecification"`
	// The IDs of one or more security groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-groups
	//
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// The type of network interface.
	//
	// To create an Elastic Fabric Adapter (EFA), specify `efa` . For more information, see [Elastic Fabric Adapter](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/efa.html) in the *Amazon Elastic Compute Cloud User Guide* .
	//
	// If you are not creating an EFA, specify `interface` or omit this parameter.
	//
	// Valid values: `interface` | `efa`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-interfacetype
	//
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv4 prefixes to be automatically assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv4Prefix` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv4prefixcount
	//
	Ipv4PrefixCount *float64 `field:"optional" json:"ipv4PrefixCount" yaml:"ipv4PrefixCount"`
	// One or more IPv4 prefixes to be assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv4PrefixCount` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv4prefixes
	//
	Ipv4Prefixes interface{} `field:"optional" json:"ipv4Prefixes" yaml:"ipv4Prefixes"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Amazon EC2 automatically selects the IPv6 addresses from the subnet range. You can't use this option if specifying specific IPv6 addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6addresscount
	//
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.
	//
	// You can't use this option if you're specifying a number of IPv6 addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6addresses
	//
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The number of IPv6 prefixes to be automatically assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv6Prefix` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6prefixcount
	//
	Ipv6PrefixCount *float64 `field:"optional" json:"ipv6PrefixCount" yaml:"ipv6PrefixCount"`
	// One or more IPv6 prefixes to be assigned to the network interface.
	//
	// You cannot use this option if you use the `Ipv6PrefixCount` option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-ipv6prefixes
	//
	Ipv6Prefixes interface{} `field:"optional" json:"ipv6Prefixes" yaml:"ipv6Prefixes"`
	// The index of the network card.
	//
	// Some instance types support multiple network cards. The primary network interface must be assigned to network card index 0. The default is network card index 0.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-networkcardindex
	//
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
	// The ID of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The primary IPv6 address of the network interface.
	//
	// When you enable an IPv6 GUA address to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address until the instance is terminated or the network interface is detached. For more information about primary IPv6 addresses, see [RunInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_RunInstances.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-primaryipv6
	//
	PrimaryIpv6 interface{} `field:"optional" json:"primaryIpv6" yaml:"primaryIpv6"`
	// The primary private IPv4 address of the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// One or more private IPv4 addresses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-privateipaddresses
	//
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-secondaryprivateipaddresscount
	//
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html#cfn-ec2-launchtemplate-networkinterface-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

