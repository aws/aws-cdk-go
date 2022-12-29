package awsdlm


// Specifies an AMI deprecation rule for cross-Region AMI copies created by a cross-Region copy rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyDeprecateRuleProperty := &crossRegionCopyDeprecateRuleProperty{
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyDeprecateRuleProperty struct {
	// The period after which to deprecate the cross-Region AMI copies.
	//
	// The period must be less than or equal to the cross-Region AMI copy retention period, and it can't be greater than 10 years. This is equivalent to 120 months, 520 weeks, or 3650 days.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The unit of time in which to measure the *Interval* .
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

