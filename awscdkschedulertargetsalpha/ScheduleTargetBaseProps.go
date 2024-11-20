package awscdkschedulertargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// Base properties for a Schedule Target.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   payload := map[string]*string{
//   	"message": jsii.String("Hello scheduler!"),
//   }
//
//   target := targets.NewSnsPublish(topic, &ScheduleTargetBaseProps{
//   	Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(payload),
//   })
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: Target,
//   })
//
// Experimental.
type ScheduleTargetBaseProps struct {
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
}

