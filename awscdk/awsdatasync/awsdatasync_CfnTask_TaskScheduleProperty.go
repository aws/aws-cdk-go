package awsdatasync


// Specifies the schedule you want your task to use for repeated executions.
//
// For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskScheduleProperty := &taskScheduleProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnTask_TaskScheduleProperty struct {
	// A cron expression that specifies when AWS DataSync initiates a scheduled transfer from a source to a destination location.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

