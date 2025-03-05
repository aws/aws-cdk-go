package awsec2


// Describes the IPv6 prefix.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ipv6PrefixSpecificationProperty := &Ipv6PrefixSpecificationProperty{
//   	Ipv6Prefix: jsii.String("ipv6Prefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-ipv6prefixspecification.html
//
type CfnNetworkInterface_Ipv6PrefixSpecificationProperty struct {
	// The IPv6 prefix.
	//
	// For information, see [Assigning prefixes to Amazon EC2 network interfaces](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-ipv6prefixspecification.html#cfn-ec2-networkinterface-ipv6prefixspecification-ipv6prefix
	//
	Ipv6Prefix *string `field:"required" json:"ipv6Prefix" yaml:"ipv6Prefix"`
}

