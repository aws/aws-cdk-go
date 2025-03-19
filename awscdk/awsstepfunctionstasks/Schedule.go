package awsstepfunctionstasks

import (
	"time"

	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Schedule for EventBridge Scheduler.
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
type Schedule interface {
	// The Schedule expression.
	ExpressionString() *string
}

// The jsii proxy struct for Schedule
type jsiiProxy_Schedule struct {
	_ byte // padding
}

func (j *jsiiProxy_Schedule) ExpressionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expressionString",
		&returns,
	)
	return returns
}


// Create a cron-based schedule from a set of cron fields.
func Schedule_Cron(options *CronOptions) Schedule {
	_init_.Initialize()

	if err := validateSchedule_CronParameters(options); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.Schedule",
		"cron",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Construct a one-time schedule from a Date.
func Schedule_OneTime(time *time.Time) Schedule {
	_init_.Initialize()

	if err := validateSchedule_OneTimeParameters(time); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.Schedule",
		"oneTime",
		[]interface{}{time},
		&returns,
	)

	return returns
}

// Construct a rate-based schedule from an interval.
//
// The minimum interval is 1 minute.
func Schedule_Rate(duration awscdk.Duration) Schedule {
	_init_.Initialize()

	if err := validateSchedule_RateParameters(duration); err != nil {
		panic(err)
	}
	var returns Schedule

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.Schedule",
		"rate",
		[]interface{}{duration},
		&returns,
	)

	return returns
}

