package awsstepfunctionstasks


// The action that EventBridge Scheduler applies to the schedule after the schedule completes invoking the target.
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
type ActionAfterCompletion string

const (
	// Takes no action.
	ActionAfterCompletion_NONE ActionAfterCompletion = "NONE"
	// Deletes the schedule.
	ActionAfterCompletion_DELETE ActionAfterCompletion = "DELETE"
)

