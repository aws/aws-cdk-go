package awsec2


// Describes the IPv6 addresses to associate with the network interface.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceIpv6AddressProperty := &InstanceIpv6AddressProperty{
//   	Ipv6Address: jsii.String("ipv6Address"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html
//
type CfnNetworkInterface_InstanceIpv6AddressProperty struct {
	// An IPv6 address to associate with the network interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-instanceipv6address.html#cfn-ec2-networkinterface-instanceipv6address-ipv6address
	//
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

