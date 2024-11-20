package awscdkschedulertargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// Properties for a SQS Queue Target.
//
// Example:
//   payload := "test"
//   messageGroupId := "id"
//   queue := sqs.NewQueue(this, jsii.String("MyQueue"), &QueueProps{
//   	Fifo: jsii.Boolean(true),
//   	ContentBasedDeduplication: jsii.Boolean(true),
//   })
//
//   target := targets.NewSqsSendMessage(queue, &SqsSendMessageProps{
//   	Input: awscdkscheduleralpha.ScheduleTargetInput_FromText(payload),
//   	MessageGroupId: jsii.String(MessageGroupId),
//   })
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   	Target: Target,
//   })
//
// Experimental.
type SqsSendMessageProps struct {
	// The SQS queue to be used as deadLetterQueue.
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Input passed to the target.
	// Default: - no input.
	//
	// Experimental.
	Input awscdkscheduleralpha.ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The maximum age of a request that Scheduler sends to a target for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	// Experimental.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the target returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// An execution role is an IAM role that EventBridge Scheduler assumes in order to interact with other AWS services on your behalf.
	//
	// If none provided templates target will automatically create an IAM role with all the minimum necessary
	// permissions to interact with the templated target. If you wish you may specify your own IAM role, then the templated targets
	// will grant minimal required permissions.
	// Default: - created by target.
	//
	// Experimental.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The FIFO message group ID to use as the target.
	//
	// This must be specified when the target is a FIFO queue. If you specify
	// a FIFO queue as a target, the queue must have content-based deduplication enabled.
	//
	// A length of `messageGroupId` must be between 1 and 128.
	// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-scheduler-schedule-sqsparameters.html#cfn-scheduler-schedule-sqsparameters-messagegroupid
	//
	// Default: - no message group ID.
	//
	// Experimental.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

