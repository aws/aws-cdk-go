package awsfms


// TCP or UDP protocols: The range of ports the rule applies to.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-portrange.html
//
type CfnPolicy_PortRangeProperty struct {
	// The beginning port number of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-portrange.html#cfn-fms-policy-portrange-from
	//
	From *float64 `field:"required" json:"from" yaml:"from"`
	// The ending port number of the range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-portrange.html#cfn-fms-policy-portrange-to
	//
	To *float64 `field:"required" json:"to" yaml:"to"`
}

