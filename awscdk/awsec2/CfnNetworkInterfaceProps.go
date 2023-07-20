package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnNetworkInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfaceProps := &CfnNetworkInterfaceProps{
//   	SubnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EnablePrimaryIpv6: jsii.Boolean(false),
//   	GroupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	InterfaceType: jsii.String("interfaceType"),
//   	Ipv6AddressCount: jsii.Number(123),
//   	Ipv6Addresses: []interface{}{
//   		&InstanceIpv6AddressProperty{
//   			Ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	PrivateIpAddress: jsii.String("privateIpAddress"),
//   	PrivateIpAddresses: []interface{}{
//   		&PrivateIpAddressSpecificationProperty{
//   			Primary: jsii.Boolean(false),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	SecondaryPrivateIpAddressCount: jsii.Number(123),
//   	SourceDestCheck: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html
//
type CfnNetworkInterfaceProps struct {
	// The ID of the subnet to associate with the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-subnetid
	//
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// A description for the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If you have instances or ENIs that rely on the IPv6 address not changing, to avoid disrupting traffic to instances or ENIs, you can enable a primary IPv6 address.
	//
	// Enable this option to automatically assign an IPv6 associated with the ENI attached to your instance to be the primary IPv6 address. When you enable an IPv6 address to be a primary IPv6, you cannot disable it. Traffic will be routed to the primary IPv6 address until the instance is terminated or the ENI is detached. If you have multiple IPv6 addresses associated with an ENI and you enable a primary IPv6 address, the first IPv6 address associated with the ENI becomes the primary IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-enableprimaryipv6
	//
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// The security group IDs associated with this network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-groupset
	//
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// The type of network interface.
	//
	// The default is `interface` . The supported values are `efa` and `trunk` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-interfacetype
	//
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the `Ipv6Addresses` property and don't specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-ipv6addresscount
	//
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface.
	//
	// If you're specifying a number of IPv6 addresses, use the `Ipv6AddressCount` property and don't specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-ipv6addresses
	//
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address.
	//
	// If you want to specify multiple private IP address, use the `PrivateIpAddresses` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-privateipaddress
	//
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Assigns private IP addresses to the network interface.
	//
	// You can specify a primary private IP address by setting the value of the `Primary` property to `true` in the `PrivateIpAddressSpecification` property. If you want EC2 to automatically assign private IP addresses, use the `SecondaryPrivateIpAddressCount` property and do not specify this property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-privateipaddresses
	//
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	//
	// When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using `privateIpAddresses` .
	//
	// You can't specify a count of private IPv4 addresses if you've specified one of the following: specific private IPv4 addresses, specific IPv4 prefixes, or a count of IPv4 prefixes.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-secondaryprivateipaddresscount
	//
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.
	//
	// If the value is `true` , source/destination checks are enabled; otherwise, they are disabled. The default value is `true` . You must disable source/destination checks if the instance runs services such as network address translation, routing, or firewalls.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-sourcedestcheck
	//
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// An arbitrary set of tags (key-value pairs) for this network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkinterface.html#cfn-ec2-networkinterface-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

