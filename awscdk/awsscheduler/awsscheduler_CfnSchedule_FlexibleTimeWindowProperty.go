package awsscheduler


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
	// `CfnSchedule.FlexibleTimeWindowProperty.Mode`.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// `CfnSchedule.FlexibleTimeWindowProperty.MaximumWindowInMinutes`.
	MaximumWindowInMinutes *float64 `field:"optional" json:"maximumWindowInMinutes" yaml:"maximumWindowInMinutes"`
}

