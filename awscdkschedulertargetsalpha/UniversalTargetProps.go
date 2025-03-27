package awscdkschedulertargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// Properties for a Universal Target.
//
// Example:
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Cron(&CronOptionsWithTimezone{
//   		Minute: jsii.String("0"),
//   		Hour: jsii.String("0"),
//   	}),
//   	Target: targets.NewUniversal(&UniversalTargetProps{
//   		Service: jsii.String("rds"),
//   		Action: jsii.String("stopDBCluster"),
//   		Input: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   			"DbClusterIdentifier": jsii.String("my-db"),
//   		}),
//   	}),
//   })
//
// Deprecated.
type UniversalTargetProps struct {
	// The SQS queue to be used as deadLetterQueue.
	//
	// The events not successfully delivered are automatically retried for a specified period of time,
	// depending on the retry policy of the target.
	// If an event is not delivered before all retry attempts are exhausted, it will be sent to the dead letter queue.
	// Default: - no dead-letter queue.
	//
	// Deprecated.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// Input passed to the target.
	// Default: - no input.
	//
	// Deprecated.
	Input awscdkscheduleralpha.ScheduleTargetInput `field:"optional" json:"input" yaml:"input"`
	// The maximum age of a request that Scheduler sends to a target for processing.
	//
	// Minimum value of 60.
	// Maximum value of 86400.
	// Default: Duration.hours(24)
	//
	// Deprecated.
	MaxEventAge awscdk.Duration `field:"optional" json:"maxEventAge" yaml:"maxEventAge"`
	// The maximum number of times to retry when the target returns an error.
	//
	// Minimum value of 0.
	// Maximum value of 185.
	// Default: 185.
	//
	// Deprecated.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// An execution role is an IAM role that EventBridge Scheduler assumes in order to interact with other AWS services on your behalf.
	//
	// If none provided templates target will automatically create an IAM role with all the minimum necessary
	// permissions to interact with the templated target. If you wish you may specify your own IAM role, then the templated targets
	// will grant minimal required permissions.
	// Default: - created by target.
	//
	// Deprecated.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// The API action to call. Must be camelCase.
	//
	// You cannot use read-only API actions such as common GET operations.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets-universal.html#unsupported-api-actions
	//
	// Deprecated.
	Action *string `field:"required" json:"action" yaml:"action"`
	// The AWS service to call.
	//
	// This must be in lowercase.
	// Deprecated.
	Service *string `field:"required" json:"service" yaml:"service"`
	// The IAM policy statements needed to invoke the target. These statements are attached to the Scheduler's role.
	//
	// Note that the default may not be the correct actions as not all AWS services follows the same IAM action pattern, or there may be more actions needed to invoke the target.
	// Default: - Policy with `service:action` action only.
	//
	// Deprecated.
	PolicyStatements *[]awsiam.PolicyStatement `field:"optional" json:"policyStatements" yaml:"policyStatements"`
}

