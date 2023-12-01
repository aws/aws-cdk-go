package awsevents


// Properties to define an event bus.
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
type EventBusProps struct {
	// The name of the event bus you are creating Note: If 'eventSourceName' is passed in, you cannot set this.
	// Default: - automatically generated name.
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The partner event source to associate with this event bus resource Note: If 'eventBusName' is passed in, you cannot set this.
	// Default: - no partner event source.
	//
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
}

