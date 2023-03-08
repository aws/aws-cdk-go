package awsec2


// Specifies the IPv6 address for the instance.
//
// `InstanceIpv6Address` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.
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
type CfnInstance_InstanceIpv6AddressProperty struct {
	// The IPv6 address.
	Ipv6Address *string `field:"required" json:"ipv6Address" yaml:"ipv6Address"`
}

