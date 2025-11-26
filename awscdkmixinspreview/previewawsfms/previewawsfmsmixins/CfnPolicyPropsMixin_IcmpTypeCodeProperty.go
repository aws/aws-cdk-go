package previewawsfmsmixins


// ICMP protocol: The ICMP type and code.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   icmpTypeCodeProperty := &IcmpTypeCodeProperty{
//   	Code: jsii.Number(123),
//   	Type: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html
//
type CfnPolicyPropsMixin_IcmpTypeCodeProperty struct {
	// ICMP code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html#cfn-fms-policy-icmptypecode-code
	//
	Code *float64 `field:"optional" json:"code" yaml:"code"`
	// ICMP type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-icmptypecode.html#cfn-fms-policy-icmptypecode-type
	//
	Type *float64 `field:"optional" json:"type" yaml:"type"`
}

