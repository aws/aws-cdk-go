package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// The target that EventBridge Scheduler will invoke.
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
type EventBridgeSchedulerTarget interface {
	// The Amazon Resource Name (ARN) of the target.
	Arn() *string
	SetArn(val *string)
	// Dead letter queue for failed events.
	DeadLetterQueue() awssqs.IQueue
	SetDeadLetterQueue(val awssqs.IQueue)
	// The input to the target.
	Input() *string
	SetInput(val *string)
	// The retry policy settings.
	RetryPolicy() *RetryPolicy
	SetRetryPolicy(val *RetryPolicy)
	// The IAM role that EventBridge Scheduler will use for this target when the schedule is invoked.
	Role() awsiam.IRole
	SetRole(val awsiam.IRole)
	// return the target object for the EventBridge Scheduler.
	RenderTargetObject() interface{}
}

// The jsii proxy struct for EventBridgeSchedulerTarget
type jsiiProxy_EventBridgeSchedulerTarget struct {
	_ byte // padding
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) RetryPolicy() *RetryPolicy {
	var returns *RetryPolicy
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventBridgeSchedulerTarget) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}


func NewEventBridgeSchedulerTarget(props *EventBridgeSchedulerTargetProps) EventBridgeSchedulerTarget {
	_init_.Initialize()

	if err := validateNewEventBridgeSchedulerTargetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgeSchedulerTarget{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EventBridgeSchedulerTarget",
		[]interface{}{props},
		&j,
	)

	return &j
}

func NewEventBridgeSchedulerTarget_Override(e EventBridgeSchedulerTarget, props *EventBridgeSchedulerTargetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions_tasks.EventBridgeSchedulerTarget",
		[]interface{}{props},
		e,
	)
}

func (j *jsiiProxy_EventBridgeSchedulerTarget)SetArn(val *string) {
	if err := j.validateSetArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"arn",
		val,
	)
}

func (j *jsiiProxy_EventBridgeSchedulerTarget)SetDeadLetterQueue(val awssqs.IQueue) {
	_jsii_.Set(
		j,
		"deadLetterQueue",
		val,
	)
}

func (j *jsiiProxy_EventBridgeSchedulerTarget)SetInput(val *string) {
	_jsii_.Set(
		j,
		"input",
		val,
	)
}

func (j *jsiiProxy_EventBridgeSchedulerTarget)SetRetryPolicy(val *RetryPolicy) {
	if err := j.validateSetRetryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryPolicy",
		val,
	)
}

func (j *jsiiProxy_EventBridgeSchedulerTarget)SetRole(val awsiam.IRole) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (e *jsiiProxy_EventBridgeSchedulerTarget) RenderTargetObject() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"renderTargetObject",
		nil, // no parameters
		&returns,
	)

	return returns
}

