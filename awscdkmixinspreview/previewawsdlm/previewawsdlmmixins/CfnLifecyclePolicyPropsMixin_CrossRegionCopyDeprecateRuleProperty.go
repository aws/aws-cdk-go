package previewawsdlmmixins


// *[Custom AMI policies only]* Specifies an AMI deprecation rule for cross-Region AMI copies created by an AMI policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   crossRegionCopyDeprecateRuleProperty := &CrossRegionCopyDeprecateRuleProperty{
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopydeprecaterule.html
//
type CfnLifecyclePolicyPropsMixin_CrossRegionCopyDeprecateRuleProperty struct {
	// The period after which to deprecate the cross-Region AMI copies.
	//
	// The period must be less than or equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopydeprecaterule.html#cfn-dlm-lifecyclepolicy-crossregioncopydeprecaterule-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	//
	// For example, to deprecate a cross-Region AMI copy after 3 months, specify `Interval=3` and `IntervalUnit=MONTHS` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-crossregioncopydeprecaterule.html#cfn-dlm-lifecyclepolicy-crossregioncopydeprecaterule-intervalunit
	//
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

