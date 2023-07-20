package awsscheduler


// An object that contains information about an Amazon SQS queue that EventBridge Scheduler uses as a dead-letter queue for your schedule.
//
// If specified, EventBridge Scheduler delivers failed events that could not be successfully delivered to a target to the queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterConfigProperty := &DeadLetterConfigProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-deadletterconfig.html
//
type CfnSchedule_DeadLetterConfigProperty struct {
	// The Amazon Resource Name (ARN) of the SQS queue specified as the destination for the dead-letter queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-deadletterconfig.html#cfn-scheduler-schedule-deadletterconfig-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

