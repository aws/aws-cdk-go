package awsec2


// Describes an IPv6 address.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html
//
type CfnSpotFleet_InstanceIpv6AddressProperty struct {
	// The IPv6 address.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-spotfleet-instanceipv6address.html#cfn-ec2-spotfleet-instanceipv6address-ipv6address
	//
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

