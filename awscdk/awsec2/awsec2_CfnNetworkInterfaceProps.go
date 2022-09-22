package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnNetworkInterface`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkInterfaceProps := &cfnNetworkInterfaceProps{
//   	subnetId: jsii.String("subnetId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	groupSet: []*string{
//   		jsii.String("groupSet"),
//   	},
//   	interfaceType: jsii.String("interfaceType"),
//   	ipv6AddressCount: jsii.Number(123),
//   	ipv6Addresses: []interface{}{
//   		&instanceIpv6AddressProperty{
//   			ipv6Address: jsii.String("ipv6Address"),
//   		},
//   	},
//   	privateIpAddress: jsii.String("privateIpAddress"),
//   	privateIpAddresses: []interface{}{
//   		&privateIpAddressSpecificationProperty{
//   			primary: jsii.Boolean(false),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   		},
//   	},
//   	secondaryPrivateIpAddressCount: jsii.Number(123),
//   	sourceDestCheck: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnNetworkInterfaceProps struct {
	// The ID of the subnet to associate with the network interface.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// A description for the network interface.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The security group IDs associated with this network interface.
	GroupSet *[]*string `field:"optional" json:"groupSet" yaml:"groupSet"`
	// The type of network interface.
	//
	// The default is `interface` . The supported values are `efa` and `trunk` .
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
	// The number of IPv6 addresses to assign to a network interface.
	//
	// Amazon EC2 automatically selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use the `Ipv6Addresses` property and don't specify this property.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet to associate with the network interface.
	//
	// If you're specifying a number of IPv6 addresses, use the `Ipv6AddressCount` property and don't specify this property.
	Ipv6Addresses interface{} `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Assigns a single private IP address to the network interface, which is used as the primary private IP address.
	//
	// If you want to specify multiple private IP address, use the `PrivateIpAddresses` property.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Assigns private IP addresses to the network interface.
	//
	// You can specify a primary private IP address by setting the value of the `Primary` property to `true` in the `PrivateIpAddressSpecification` property. If you want EC2 to automatically assign private IP addresses, use the `SecondaryPrivateIpAddressCount` property and do not specify this property.
	PrivateIpAddresses interface{} `field:"optional" json:"privateIpAddresses" yaml:"privateIpAddresses"`
	// The number of secondary private IPv4 addresses to assign to a network interface.
	//
	// When you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses within the subnet's IPv4 CIDR range. You can't specify this option and specify more than one private IP address using `privateIpAddresses` .
	//
	// The number of IP addresses you can assign to a network interface varies by instance type. For more information, see [IP Addresses Per ENI Per Instance Type](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-eni.html#AvailableIpPerENI) in the *Amazon Virtual Private Cloud User Guide* .
	SecondaryPrivateIpAddressCount *float64 `field:"optional" json:"secondaryPrivateIpAddressCount" yaml:"secondaryPrivateIpAddressCount"`
	// Enable or disable source/destination checks, which ensure that the instance is either the source or the destination of any traffic that it receives.
	//
	// If the value is `true` , source/destination checks are enabled; otherwise, they are disabled. The default value is `true` . You must disable source/destination checks if the instance runs services such as network address translation, routing, or firewalls.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// An arbitrary set of tags (key-value pairs) for this network interface.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

