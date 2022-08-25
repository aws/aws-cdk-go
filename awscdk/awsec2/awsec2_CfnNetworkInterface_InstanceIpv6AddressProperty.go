package awsec2


// Describes the IPv6 addresses to associate with the network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceIpv6AddressProperty := &instanceIpv6AddressProperty{
//   	ipv6Address: jsii.String("ipv6Address"),
//   }
//
type CfnNetworkInterface_InstanceIpv6AddressProperty struct {
	// An IPv6 address to associate with the network interface.
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

