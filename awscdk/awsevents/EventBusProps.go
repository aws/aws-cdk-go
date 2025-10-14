package awsevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

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
	// Dead-letter queue for the event bus.
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-rule-event-delivery.html#eb-rule-dlq
	//
	// Default: - no dead-letter queue.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The event bus description.
	//
	// The description can be up to 512 characters long.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-description
	//
	// Default: - no description.
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the event bus you are creating Note: If 'eventSourceName' is passed in, you cannot set this.
	// Default: - automatically generated name.
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The partner event source to associate with this event bus resource Note: If 'eventBusName' is passed in, you cannot set this.
	// Default: - no partner event source.
	//
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// The customer managed key that encrypt events on this event bus.
	// Default: - Use an AWS managed key.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The Logging Configuration of the Èvent Bus.
	// Default: - no logging.
	//
	LogConfig *LogConfig `field:"optional" json:"logConfig" yaml:"logConfig"`
}

