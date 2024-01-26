package awsec2


// Describes the ICMP type and code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icmpProperty := &IcmpProperty{
//   	Code: jsii.Number(123),
//   	Type: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html
//
type CfnNetworkAclEntry_IcmpProperty struct {
	// The Internet Control Message Protocol (ICMP) code.
	//
	// You can use -1 to specify all ICMP codes for the given ICMP type. Required if you specify 1 (ICMP) for the protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-code
	//
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// The Internet Control Message Protocol (ICMP) type.
	//
	// You can use -1 to specify all ICMP types. Conditional requirement: Required if you specify 1 (ICMP) for the `CreateNetworkAclEntry` protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-icmp.html#cfn-ec2-networkaclentry-icmp-type
	//
	Type *float64 `field:"optional" json:"type" yaml:"type"`
}

