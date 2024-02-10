package awsec2


// Specifies a network interface that is to be attached to an instance.
//
// You can create a network interface when launching an instance. For an example, see the [AWS::EC2::Instance examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html#aws-properties-ec2-instance--examples--Automatically_assign_a_public_IP_address) .
//
// Alternatively, you can attach an existing network interface when launching an instance. For an example, see the [AWS::EC2:NetworkInterface examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#aws-resource-ec2-networkinterface--examples) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   networkInterfaceProperty := &NetworkInterfaceProperty{
//   	DeviceIndex: jsii.String("deviceIndex"),
//
//   	// the properties below are optional
//   	AssociateCarrierIpAddress: jsii.Boolean(false),
//   	AssociatePublicIpAddress: jsii.Boolean(false),
//   	DeleteOnTermination: jsii.Boolean(false),
//   	Description: jsii.String("description"),
//   	GroupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&InstanceIpv6AddressProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PrivateIpAddresses: []interface{}{
//   		&PrivateIpAddressSpecificationProperty{
//   			Primary: jsii.Boolean(false),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SubnetId: jsii.String("subnetId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html
//
type CfnInstance_NetworkInterfaceProperty struct {
	// The position of the network interface in the attachment order.
	//
	// A primary network interface has a device index of 0.
	//
	// If you create a network interface when launching an instance, you must specify the device index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-deviceindex
	//
	DeviceIndex *string `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Indicates whether to assign a carrier IP address to the network interface.
	//
	// You can only assign a carrier IP address to a network interface that is in a subnet in a Wavelength Zone. For more information about carrier IP addresses, see [Carrier IP address](https://docs.aws.amazon.com/wavelength/latest/developerguide/how-wavelengths-work.html#provider-owned-ip) in the *AWS Wavelength Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-associatecarrieripaddress
	//
	AssociateCarrierIpAddress interface{} `field:"optional" json:"associateCarrierIpAddress" yaml:"associateCarrierIpAddress"`
	// Indicates whether to assign a public IPv4 address to an instance.
	//
	// Applies only if creating a network interface when launching an instance. The network interface must be the primary network interface. If launching into a default subnet, the default value is `true` .
	//
	// AWS charges for all public IPv4 addresses, including public IPv4 addresses associated with running instances and Elastic IP addresses. For more information, see the *Public IPv4 Address* tab on the [VPC pricing page](https://docs.aws.amazon.com/vpc/pricing/) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-associatepublicipaddress
	//
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Indicates whether the network interface is deleted when the instance is terminated.
	//
	// Applies only if creating a network interface when launching an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-deleteontermination
	//
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// The description of the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IDs of the security groups for the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-groupset
	//
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// A number of IPv6 addresses to assign to the network interface.
	//
	// Amazon EC2 chooses the IPv6 addresses from the range of the subnet. You cannot specify this option and the option to assign specific IPv6 addresses in the same request. You can specify this option if you've specified a minimum number of instances to launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-ipv6addresscount
	//
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// The IPv6 addresses to assign to the network interface.
	//
	// You cannot specify this option and the option to assign a number of IPv6 addresses in the same request. You cannot specify this option if you've specified a minimum number of instances to launch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-ipv6addresses
	//
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// The ID of the network interface, when attaching an existing network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-networkinterfaceid
	//
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// The private IPv4 address of the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// One or more private IPv4 addresses to assign to the network interface.
	//
	// Only one private IPv4 address can be designated as primary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-privateipaddresses
	//
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses.
	//
	// You can't specify this option and specify more than one private IP address using the private IP addresses option.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-secondaryprivateipaddresscount
	//
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// The ID of the subnet associated with the network interface.
	//
	// Applies only if creating a network interface when launching an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-networkinterface.html#cfn-ec2-instance-networkinterface-subnetid
	//
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

