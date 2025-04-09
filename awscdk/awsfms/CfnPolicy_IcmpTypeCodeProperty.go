package awsfms


// ICMP protocol: The ICMP type and code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   icmpTypeCodeProperty := &IcmpTypeCodeProperty{
//   	Code: jsii.Number(123),
//   	Type: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html
//
type CfnPolicy_IcmpTypeCodeProperty struct {
	// ICMP code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html#cfn-fms-policy-icmptypecode-code
	//
	Code *float64 `field:"required" json:"code" yaml:"code"`
	// ICMP type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html#cfn-fms-policy-icmptypecode-type
	//
	Type *float64 `field:"required" json:"type" yaml:"type"`
}

