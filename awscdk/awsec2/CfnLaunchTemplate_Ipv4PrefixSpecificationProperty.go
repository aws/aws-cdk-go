package awsec2


// Specifies an IPv4 prefix for a network interface.
//
// `Ipv4PrefixSpecification` is a property of [AWS::EC2::LaunchTemplate NetworkInterface](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-networkinterface.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipv4PrefixSpecificationProperty := &Ipv4PrefixSpecificationProperty{
//   	Ipv4Prefix: jsii.String("ipv4Prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-ipv4prefixspecification.html
//
type CfnLaunchTemplate_Ipv4PrefixSpecificationProperty struct {
	// The IPv4 prefix.
	//
	// For information, see [Assigning prefixes to Amazon EC2 network interfaces](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-ipv4prefixspecification.html#cfn-ec2-launchtemplate-ipv4prefixspecification-ipv4prefix
	//
	Ipv4Prefix *string `field:"optional" json:"ipv4Prefix" yaml:"ipv4Prefix"`
}

