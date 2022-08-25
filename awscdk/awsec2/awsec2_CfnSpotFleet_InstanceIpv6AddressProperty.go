package awsec2


// Describes an IPv6 address.
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
type CfnSpotFleet_InstanceIpv6AddressProperty struct {
	// The IPv6 address.
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

