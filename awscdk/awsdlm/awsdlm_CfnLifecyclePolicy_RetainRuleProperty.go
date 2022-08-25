package awsdlm


// Specifies the retention rule for a lifecycle policy.
//
// You can retain snapshots based on either a count or a time interval.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   retainRuleProperty := &retainRuleProperty{
//   	count: jsii.Number(123),
//   	interval: jsii.Number(123),
//   	intervalUnit: jsii.String("intervalUnit"),
//   }
//
type CfnLifecyclePolicy_RetainRuleProperty struct {
	// The number of snapshots to retain for each volume, up to a maximum of 1000.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The amount of time to retain each snapshot.
	//
	// The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The unit of time for time-based retention.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
}

