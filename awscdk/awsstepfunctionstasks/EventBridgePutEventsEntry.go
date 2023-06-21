package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// An entry to be sent to EventBridge.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var eventBus eventBus
//   var taskInput taskInput
//
//   eventBridgePutEventsEntry := &EventBridgePutEventsEntry{
//   	Detail: taskInput,
//   	DetailType: jsii.String("detailType"),
//   	Source: jsii.String("source"),
//
//   	// the properties below are optional
//   	EventBus: eventBus,
//   }
//
// See: https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_PutEventsRequestEntry.html
//
type EventBridgePutEventsEntry struct {
	// The event body.
	//
	// Can either be provided as an object or as a JSON-serialized string.
	//
	// Example:
	//   sfn.TaskInput_FromText(jsii.String("{\"instance-id\": \"i-1234567890abcdef0\", \"state\": \"terminated\"}"))
	//   sfn.TaskInput_FromObject(map[string]interface{}{
	//   	"Message": jsii.String("Hello from Step Functions"),
	//   })
	//   sfn.TaskInput_FromJsonPathAt(jsii.String("$.EventDetail"))
	//
	Detail awsstepfunctions.TaskInput `field:"required" json:"detail" yaml:"detail"`
	// Used along with the source field to help identify the fields and values expected in the detail field.
	//
	// For example, events by CloudTrail have detail type "AWS API Call via CloudTrail".
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// The service or application that caused this event to be generated.
	//
	// Example value: `com.example.service`
	// See: https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-events.html
	//
	Source *string `field:"required" json:"source" yaml:"source"`
	// The event bus the entry will be sent to.
	EventBus awsevents.IEventBus `field:"optional" json:"eventBus" yaml:"eventBus"`
}

