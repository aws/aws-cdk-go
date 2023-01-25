package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for SendMessageTask.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var taskInput taskInput
//
//   sendToQueueProps := &sendToQueueProps{
//   	messageBody: taskInput,
//
//   	// the properties below are optional
//   	delay: duration,
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	messageGroupId: jsii.String("messageGroupId"),
//   }
//
// Deprecated: Use `SqsSendMessage`.
type SendToQueueProps struct {
	// The text message to send to the queue.
	// Deprecated: Use `SqsSendMessage`.
	MessageBody awsstepfunctions.TaskInput `field:"required" json:"messageBody" yaml:"messageBody"`
	// The length of time, in seconds, for which to delay a specific message.
	//
	// Valid values are 0-900 seconds.
	// Deprecated: Use `SqsSendMessage`.
	Delay awscdk.Duration `field:"optional" json:"delay" yaml:"delay"`
	// The service integration pattern indicates different ways to call SendMessage to SQS.
	//
	// The valid value is either FIRE_AND_FORGET or WAIT_FOR_TASK_TOKEN.
	// Deprecated: Use `SqsSendMessage`.
	IntegrationPattern awsstepfunctions.ServiceIntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// The token used for deduplication of sent messages.
	// Deprecated: Use `SqsSendMessage`.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The tag that specifies that a message belongs to a specific message group.
	//
	// Required for FIFO queues. FIFO ordering applies to messages in the same message
	// group.
	// Deprecated: Use `SqsSendMessage`.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

