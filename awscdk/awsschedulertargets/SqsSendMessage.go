package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an Amazon SQS Queue as a target for AWS EventBridge Scheduler.
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
//   	Input: awscdk.ScheduleTargetInput_FromText(payload),
//   	MessageGroupId: jsii.String(MessageGroupId),
//   })
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Minutes(jsii.Number(1))),
//   	Target: Target,
//   })
//
type SqsSendMessage interface {
	ScheduleTargetBase
	awsscheduler.IScheduleTarget
	TargetArn() *string
	AddTargetActionToRole(role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
	BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig
}

// The jsii proxy struct for SqsSendMessage
type jsiiProxy_SqsSendMessage struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_SqsSendMessage) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewSqsSendMessage(queue awssqs.IQueue, props *SqsSendMessageProps) SqsSendMessage {
	_init_.Initialize()

	if err := validateNewSqsSendMessageParameters(queue, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsSendMessage{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SqsSendMessage",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsSendMessage_Override(s SqsSendMessage, queue awssqs.IQueue, props *SqsSendMessageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.SqsSendMessage",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsSendMessage) AddTargetActionToRole(role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (s *jsiiProxy_SqsSendMessage) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

