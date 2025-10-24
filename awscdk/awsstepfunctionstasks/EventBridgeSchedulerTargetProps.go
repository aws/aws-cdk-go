package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for `EventBridgeSchedulerTarget`.
//
// Example:
//   import scheduler "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//   var key Key
//   var scheduleGroup CfnScheduleGroup
//   var targetQueue Queue
//   var deadLetterQueue Queue
//
//
//   schedulerRole := iam.NewRole(this, jsii.String("SchedulerRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("scheduler.amazonaws.com")),
//   })
//   // To send the message to the queue
//   // This policy changes depending on the type of target.
//   schedulerRole.AddToPrincipalPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("sqs:SendMessage"),
//   	},
//   	Resources: []*string{
//   		targetQueue.QueueArn,
//   	},
//   }))
//
//   createScheduleTask1 := tasks.NewEventBridgeSchedulerCreateScheduleTask(this, jsii.String("createSchedule"), &EventBridgeSchedulerCreateScheduleTaskProps{
//   	ScheduleName: jsii.String("TestSchedule"),
//   	ActionAfterCompletion: tasks.ActionAfterCompletion_NONE,
//   	ClientToken: jsii.String("testToken"),
//   	Description: jsii.String("TestDescription"),
//   	StartDate: NewDate(),
//   	EndDate: NewDate(NewDate().getTime() + 1000 * 60 * 60),
//   	FlexibleTimeWindow: awscdk.Duration_Minutes(jsii.Number(5)),
//   	GroupName: scheduleGroup.ref,
//   	KmsKey: key,
//   	Schedule: tasks.Schedule_Rate(awscdk.Duration_*Minutes(jsii.Number(5))),
//   	Timezone: jsii.String("UTC"),
//   	Enabled: jsii.Boolean(true),
//   	Target: tasks.NewEventBridgeSchedulerTarget(&EventBridgeSchedulerTargetProps{
//   		Arn: targetQueue.*QueueArn,
//   		Role: schedulerRole,
//   		RetryPolicy: &RetryPolicy{
//   			MaximumRetryAttempts: jsii.Number(2),
//   			MaximumEventAge: awscdk.Duration_*Minutes(jsii.Number(5)),
//   		},
//   		DeadLetterQueue: *DeadLetterQueue,
//   	}),
//   })
//
// See: https://docs.aws.amazon.com/scheduler/latest/APIReference/API_Target.html#API_Target_Contents
//
type EventBridgeSchedulerTargetProps struct {
	// The Amazon Resource Name (ARN) of the target.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/managing-targets.html
	//
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// The IAM role that EventBridge Scheduler will use for this target when the schedule is invoked.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// Dead letter queue for failed events.
	// Default: - No dead letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The input to the target.
	// Default: - EventBridge Scheduler delivers a default notification to the target.
	//
	Input *string `field:"optional" json:"input" yaml:"input"`
	// The retry policy settings.
	// Default: - Do not retry.
	//
	RetryPolicy *RetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
}

