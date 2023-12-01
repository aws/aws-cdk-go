package awscdkschedulertargetsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
	"github.com/aws/aws-cdk-go/awscdkschedulertargetsalpha/v2/internal"
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
//   	Detail: awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
//   		"Name": jsii.String("Fluffy"),
//   	}),
//   	DetailType: jsii.String("🐶"),
//   }
//
//   awscdkscheduleralpha.NewSchedule(this, jsii.String("Schedule"), &ScheduleProps{
//   	Schedule: awscdkscheduleralpha.ScheduleExpression_Rate(awscdk.Duration_Hours(jsii.Number(1))),
//   	Target: targets.NewEventBridgePutEvents(eventEntry, &ScheduleTargetBaseProps{
//   	}),
//   })
//
// Experimental.
type EventBridgePutEvents interface {
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

// The jsii proxy struct for EventBridgePutEvents
type jsiiProxy_EventBridgePutEvents struct {
	jsiiProxy_ScheduleTargetBase
	internal.Type__awscdkscheduleralphaIScheduleTarget
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


// Experimental.
func NewEventBridgePutEvents(entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) EventBridgePutEvents {
	_init_.Initialize()

	if err := validateNewEventBridgePutEventsParameters(entry, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventBridgePutEvents{}

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.EventBridgePutEvents",
		[]interface{}{entry, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEventBridgePutEvents_Override(e EventBridgePutEvents, entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-scheduler-targets-alpha.EventBridgePutEvents",
		[]interface{}{entry, props},
		e,
	)
}

func (e *jsiiProxy_EventBridgePutEvents) AddTargetActionToRole(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) {
	if err := e.validateAddTargetActionToRoleParameters(schedule, role); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addTargetActionToRole",
		[]interface{}{schedule, role},
	)
}

func (e *jsiiProxy_EventBridgePutEvents) Bind(schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := e.validateBindParameters(schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{schedule},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventBridgePutEvents) BindBaseTargetConfig(_schedule awscdkscheduleralpha.ISchedule) *awscdkscheduleralpha.ScheduleTargetConfig {
	if err := e.validateBindBaseTargetConfigParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *awscdkscheduleralpha.ScheduleTargetConfig

	_jsii_.Invoke(
		e,
		"bindBaseTargetConfig",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

