package awscdkschedulertargetsalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2"
)

// An entry to be sent to EventBridge.
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
//   	Target: targets.NewEventBridgePutEvents(eventEntry),
//   })
//
// See: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEventsRequestEntry.html
//
// Experimental.
type EventBridgePutEventsEntry struct {
	// The event body.
	//
	// Can either be provided as an object or as a JSON-serialized string.
	//
	// Example:
	//   awscdkscheduleralpha.ScheduleTargetInput_FromText(jsii.String("{\"instance-id\": \"i-1234567890abcdef0\", \"state\": \"terminated\"}"))
	//   awscdkscheduleralpha.ScheduleTargetInput_FromObject(map[string]*string{
	//   	"Message": jsii.String("Hello from a friendly event :)"),
	//   })
	//
	// Experimental.
	Detail awscdkscheduleralpha.ScheduleTargetInput `field:"required" json:"detail" yaml:"detail"`
	// Used along with the source field to help identify the fields and values expected in the detail field.
	//
	// For example, events by CloudTrail have detail type "AWS API Call via CloudTrail".
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	// Experimental.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// The event bus the entry will be sent to.
	// Experimental.
	EventBus awsevents.IEventBus `field:"required" json:"eventBus" yaml:"eventBus"`
	// The service or application that caused this event to be generated.
	//
	// Example value: `com.example.service`
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	// Experimental.
	Source *string `field:"required" json:"source" yaml:"source"`
}

