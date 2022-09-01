package awsevents


// Properties to define an event bus.
//
// Example:
//   bus := events.NewEventBus(this, jsii.String("bus"), &eventBusProps{
//   	eventBusName: jsii.String("MyCustomEventBus"),
//   })
//
//   bus.archive(jsii.String("MyArchive"), &baseArchiveProps{
//   	archiveName: jsii.String("MyCustomEventBusArchive"),
//   	description: jsii.String("MyCustomerEventBus Archive"),
//   	eventPattern: &eventPattern{
//   		account: []*string{
//   			awscdk.*stack.of(this).account,
//   		},
//   	},
//   	retention: awscdk.Duration.days(jsii.Number(365)),
//   })
//
type EventBusProps struct {
	// The name of the event bus you are creating Note: If 'eventSourceName' is passed in, you cannot set this.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The partner event source to associate with this event bus resource Note: If 'eventBusName' is passed in, you cannot set this.
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
}

