package awsdlm


// Specifies the retention rule for cross-Region snapshot copies.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   crossRegionCopyRetainRuleProperty := &crossRegionCopyRetainRuleProperty{
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_CrossRegionCopyRetainRuleProperty struct {
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `field:"required" json:"intervalUnit" yaml:"intervalUnit"`
}

