package awsevents


// Properties to define an event bus.
//
// Example:
//   import events "github.com/aws/aws-cdk-go/awscdk"
//
//
//   myEventBus := events.NewEventBus(this, jsii.String("EventBus"), &EventBusProps{
//   	EventBusName: jsii.String("MyEventBus1"),
//   })
//
//   tasks.NewEventBridgePutEvents(this, jsii.String("Send an event to EventBridge"), &EventBridgePutEventsProps{
//   	Entries: []eventBridgePutEventsEntry{
//   		&eventBridgePutEventsEntry{
//   			Detail: sfn.TaskInput_FromObject(map[string]interface{}{
//   				"Message": jsii.String("Hello from Step Functions!"),
//   			}),
//   			EventBus: myEventBus,
//   			DetailType: jsii.String("MessageFromStepFunctions"),
//   			Source: jsii.String("step.functions"),
//   		},
//   	},
//   })
//
type EventBusProps struct {
	// The name of the event bus you are creating Note: If 'eventSourceName' is passed in, you cannot set this.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The partner event source to associate with this event bus resource Note: If 'eventBusName' is passed in, you cannot set this.
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
}

