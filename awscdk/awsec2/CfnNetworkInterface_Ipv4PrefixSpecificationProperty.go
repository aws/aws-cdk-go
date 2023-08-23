package awsec2


// Describes an IPv4 prefix.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-ipv4prefixspecification.html
//
type CfnNetworkInterface_Ipv4PrefixSpecificationProperty struct {
	// The IPv4 prefix.
	//
	// For information, see [Assigning prefixes to Amazon EC2 network interfaces](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-prefix-eni.html) in the *Amazon Elastic Compute Cloud User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkinterface-ipv4prefixspecification.html#cfn-ec2-networkinterface-ipv4prefixspecification-ipv4prefix
	//
	Ipv4Prefix *string `field:"required" json:"ipv4Prefix" yaml:"ipv4Prefix"`
}

