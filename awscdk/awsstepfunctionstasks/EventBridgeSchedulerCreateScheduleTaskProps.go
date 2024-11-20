package awsstepfunctionstasks

import (
	"time"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for creating an AWS EventBridge Scheduler schedule.
//
// Example:
//   import scheduler "github.com/aws/aws-cdk-go/awscdk"
//   import kms "github.com/aws/aws-cdk-go/awscdk"
//
//   var key key
//   var scheduleGroup cfnScheduleGroup
//   var targetQueue queue
//   var deadLetterQueue queue
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
type EventBridgeSchedulerCreateScheduleTaskProps struct {
	// An optional description for this state.
	// Default: - No comment.
	//
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	// Default: - None (Task is executed using the State Machine's execution role).
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Default: - None.
	//
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Default: - The entire task input (JSON path '$').
	//
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	//
	// Depending on the AWS Service, the Service Integration Pattern availability will vary.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-supported-services.html
	//
	// Default: - `IntegrationPattern.REQUEST_RESPONSE` for most tasks.
	// `IntegrationPattern.RUN_JOB` for the following exceptions:
	// `BatchSubmitJob`, `EmrAddStep`, `EmrCreateCluster`, `EmrTerminationCluster`, and `EmrContainersStartJobRun`.
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Default: - The entire JSON node determined by the state input, the task result,
	// and resultPath is passed to the next state (JSON path '$').
	//
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Default: - Replaces the entire input with the result (JSON path '$').
	//
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Default: - None.
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Optional name for this state.
	// Default: - The construct ID will be used as state name.
	//
	StateName *string `field:"optional" json:"stateName" yaml:"stateName"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	// Default: - None.
	//
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Default: - None.
	//
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The schedule that defines when the schedule will trigger.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/schedule-types.html
	//
	Schedule Schedule `field:"required" json:"schedule" yaml:"schedule"`
	// Schedule name.
	ScheduleName *string `field:"required" json:"scheduleName" yaml:"scheduleName"`
	// The schedule's target.
	Target EventBridgeSchedulerTarget `field:"required" json:"target" yaml:"target"`
	// Specifies the action that EventBridge Scheduler applies to the schedule after the schedule completes invoking the target.
	// Default: ActionAfterCompletion.NONE
	//
	ActionAfterCompletion ActionAfterCompletion `field:"optional" json:"actionAfterCompletion" yaml:"actionAfterCompletion"`
	// Unique, case-sensitive identifier to ensure the idempotency of the request.
	// Default: - Automatically generated.
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The description for the schedule.
	// Default: - No description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies whether the schedule is enabled or disabled.
	// Default: true.
	//
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The date, in UTC, before which the schedule can invoke its target.
	//
	// Depending on the schedule's recurrence expression, invocations might stop on, or before, the EndDate you specify.
	// EventBridge Scheduler ignores EndDate for one-time schedules.
	// Default: - No end date.
	//
	EndDate *time.Time `field:"optional" json:"endDate" yaml:"endDate"`
	// The maximum time window during which a schedule can be invoked.
	//
	// Minimum value is 1 minute.
	// Maximum value is 1440 minutes (1 day).
	// Default: - Flexible time window is not enabled.
	//
	FlexibleTimeWindow awscdk.Duration `field:"optional" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// The name of the schedule group to associate with this schedule.
	// Default: - The default schedule group is used.
	//
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// The customer managed KMS key that EventBridge Scheduler will use to encrypt and decrypt payload.
	// See: https://docs.aws.amazon.com/scheduler/latest/UserGuide/encryption-rest.html
	//
	// Default: - Use automatically generated KMS key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The date, in UTC, after which the schedule can begin invoking its target.
	//
	// Depending on the schedule's recurrence expression, invocations might occur on, or after, the StartDate you specify.
	// EventBridge Scheduler ignores StartDate for one-time schedules.
	// Default: - No start date.
	//
	StartDate *time.Time `field:"optional" json:"startDate" yaml:"startDate"`
	// The timezone in which the scheduling expression is evaluated.
	// Default: - UTC.
	//
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

