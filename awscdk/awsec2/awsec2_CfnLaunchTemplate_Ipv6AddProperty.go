package awsec2


// Specifies an IPv6 address in an Amazon EC2 launch template.
//
// `Ipv6Add` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipv6AddProperty := &ipv6AddProperty{
//   	ipv6Address: jsii.String("ipv6Address"),
//   }
//
type CfnLaunchTemplate_Ipv6AddProperty struct {
	// One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet.
	//
	// You can't use this option if you're specifying a number of IPv6 addresses.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
}

