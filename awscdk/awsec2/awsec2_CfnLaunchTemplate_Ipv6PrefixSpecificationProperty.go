package awsec2


// Specifies an IPv6 prefix for a network interface.
//
// `Ipv6PrefixSpecification` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipv6PrefixSpecificationProperty := &ipv6PrefixSpecificationProperty{
//   	ipv6Prefix: jsii.String("ipv6Prefix"),
//   }
//
type CfnLaunchTemplate_Ipv6PrefixSpecificationProperty struct {
	// The IPv6 prefix.
	Ipv6Prefix *string `field:"optional" json:"ipv6Prefix" yaml:"ipv6Prefix"`
}

