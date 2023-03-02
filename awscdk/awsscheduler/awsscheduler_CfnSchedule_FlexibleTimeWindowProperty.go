package awsscheduler


// Allows you to configure a time window during which EventBridge Scheduler invokes the schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flexibleTimeWindowProperty := &flexibleTimeWindowProperty{
//   	mode: jsii.String("mode"),
//
//   	// the properties below are optional
//   	maximumWindowInMinutes: jsii.Number(123),
//   }
//
type CfnSchedule_FlexibleTimeWindowProperty struct {
	// Determines whether the schedule is invoked within a flexible time window.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// The maximum time window during which a schedule can be invoked.
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
}

