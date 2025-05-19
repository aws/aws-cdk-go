package awsschedulertargets

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsschedulertargets/internal"
)

// Send an event to an AWS EventBridge by AWS EventBridge Scheduler.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//
//   eventBus := events.NewEventBus(this, jsii.String("EventBus"), &EventBusProps{
//   	EventBusName: jsii.String("DomainEvents"),
//   })
//
//   eventEntry := &EventBridgePutEventsEntry{
//   	EventBus: EventBus,
//   	Source: jsii.String("PetService"),
//   	Detail: awscdk.ScheduleTargetInput_FromObject(map[string]*string{
//   		"Name": jsii.String("Fluffy"),
//   	}),
//   	DetailType: jsii.String("🐶"),
//   }
//
//   awscdk.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdk.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: targets.NewEventBridgePutEvents(eventEntry),
//   })
//
type EventBridgePutEvents interface {
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

// The jsii proxy struct for EventBridgePutEvents
type jsiiProxy_EventBridgePutEvents struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awsschedulerIScheduleTarget
}

func (j *jsiiProxy_EventBridgePutEvents) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}


func NewEventBridgePutEvents(entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) EventBridgePutEvents {
	_init_.Initialize()

	if err := validateNewEventBridgePutEventsParameters(entry, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgePutEvents{}

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.EventBridgePutEvents",
		[]interface{}{entry, props},
		&j,
	)

	return &j
}

func NewEventBridgePutEvents_Override(e EventBridgePutEvents, entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_scheduler_targets.EventBridgePutEvents",
		[]interface{}{entry, props},
		e,
	)
}

func (e *jsiiProxy_EventBridgePutEvents) AddTargetActionToRole(role awsiam.IRole) {
	if err := e.validateAddTargetActionToRoleParameters(role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addTargetActionToRole",
		[]interface{}{role},
	)
}

func (e *jsiiProxy_EventBridgePutEvents) Bind(schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := e.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgePutEvents) BindBaseTargetConfig(_schedule awsscheduler.ISchedule) *awsscheduler.ScheduleTargetConfig {
	if err := e.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awsscheduler.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

