package awsec2


// Describes a range of ports.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portRangeProperty := &PortRangeProperty{
//   	From: jsii.Number(123),
//   	To: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html
//
type CfnNetworkAclEntry_PortRangeProperty struct {
	// The first port in the range.
	//
	// Required if you specify 6 (TCP) or 17 (UDP) for the protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-from
	//
	From *float64 `field:"optional" json:"from" yaml:"from"`
	// The last port in the range.
	//
	// Required if you specify 6 (TCP) or 17 (UDP) for the protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-networkaclentry-portrange.html#cfn-ec2-networkaclentry-portrange-to
	//
	To *float64 `field:"optional" json:"to" yaml:"to"`
}

