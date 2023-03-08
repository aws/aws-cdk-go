package awsiotanalytics


// The schedule for when to trigger an update.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleProperty := &scheduleProperty{
//   	scheduleExpression: jsii.String("scheduleExpression"),
//   }
//
type CfnDataset_ScheduleProperty struct {
	// The expression that defines when to trigger an update.
	//
	// For more information, see [Schedule Expressions for Rules](https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html) in the Amazon CloudWatch documentation.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
}

