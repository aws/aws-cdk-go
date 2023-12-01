package awsdlm


// *[Custom AMI policies only]* Specifies an AMI deprecation rule for AMIs created by an AMI lifecycle policy.
//
// For age-based schedules, you must specify *Interval* and *IntervalUnit* . For count-based schedules, you must specify *Count* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deprecateRuleProperty := &DeprecateRuleProperty{
//   	Count: jsii.Number(123),
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-deprecaterule.html
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate.
	//
	// The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-deprecaterule.html#cfn-dlm-lifecyclepolicy-deprecaterule-count
	//
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule.
	//
	// The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-deprecaterule.html#cfn-dlm-lifecyclepolicy-deprecaterule-interval
	//
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-deprecaterule.html#cfn-dlm-lifecyclepolicy-deprecaterule-intervalunit
	//
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

