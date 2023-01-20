package awsdlm


// Specifies an AMI deprecation rule for a schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deprecateRuleProperty := &deprecateRuleProperty{
//   	count: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_DeprecateRuleProperty struct {
	// If the schedule has a count-based retention rule, this parameter specifies the number of oldest AMIs to deprecate.
	//
	// The count must be less than or equal to the schedule's retention count, and it can't be greater than 1000.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// If the schedule has an age-based retention rule, this parameter specifies the period after which to deprecate AMIs created by the schedule.
	//
	// The period must be less than or equal to the schedule's retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

