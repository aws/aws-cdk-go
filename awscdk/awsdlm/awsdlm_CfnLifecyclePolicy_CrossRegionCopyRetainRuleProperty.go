package awsdlm


// Specifies a retention rule for cross-Region snapshot copies created by snapshot or event-based policies, or cross-Region AMI copies created by AMI policies.
//
// After the retention period expires, the cross-Region copy is deleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyRetainRuleProperty := &CrossRegionCopyRetainRuleProperty{
//   	Interval: jsii.Number(123),
//   	IntervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty struct {
	// The amount of time to retain a cross-Region snapshot or AMI copy.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	//
	// For example, to retain a cross-Region copy for 3 months, specify `Interval=3` and `IntervalUnit=MONTHS` .
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

