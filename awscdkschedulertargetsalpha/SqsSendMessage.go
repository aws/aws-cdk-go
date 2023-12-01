package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
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
type SqsSendMessage interface {
	ScheduleTargetBase
	awscdkscheduleralpha.IScheduleTarget
	// Experimental.
	TargetArn() *string
	// Experimental.
	AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole)
	// Create a return a Schedule Target Configuration for the given schedule.
	//
	// Returns: a Schedule Target Configuration.
	// Experimental.
	Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
	// Experimental.
	BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig
}

// The jsii proxy struct for SqsSendMessage
type jsiiProxy_SqsSendMessage struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
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


// Experimental.
func NewSqsSendMessage(queue awssqs.IQueue, props *SqsSendMessageProps) SqsSendMessage {
	_init_.Initialize()

	if err := validateNewSqsSendMessageParameters(queue, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsSendMessage{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SqsSendMessage",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsSendMessage_Override(s SqsSendMessage, queue awssqs.IQueue, props *SqsSendMessageProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.SqsSendMessage",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsSendMessage) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := s.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (s *jsiiProxy_SqsSendMessage) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsSendMessage) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := s.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		s,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

