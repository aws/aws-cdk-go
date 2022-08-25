package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctionstasks/internal"
)

// A StepFunctions Task to send messages to SQS queue.
//
// A Function can be used directly as a Resource, but this class mirrors
// integration with other AWS services via a specific class instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var duration duration
//   var queue queue
//   var taskInput taskInput
//
//   sendToQueue := awscdk.Aws_stepfunctions_tasks.NewSendToQueue(queue, &sendToQueueProps{
//   	messageBody: taskInput,
//
//   	// the properties below are optional
//   	delay: duration,
//   	integrationPattern: awscdk.Aws_stepfunctions.serviceIntegrationPattern_FIRE_AND_FORGET,
//   	messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	messageGroupId: jsii.String("messageGroupId"),
//   })
//
// Deprecated: Use `SqsSendMessage`.
type SendToQueue interface {
	awsstepfunctions.IStepFunctionsTask
	// Called when the task object is used in a workflow.
	// Deprecated: Use `SqsSendMessage`.
	Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig
}

// The jsii proxy struct for SendToQueue
type jsiiProxy_SendToQueue struct {
	internal.Type__awsstepfunctionsIStepFunctionsTask
}

// Deprecated: Use `SqsSendMessage`.
func NewSendToQueue(queue awssqs.IQueue, props *SendToQueueProps) SendToQueue {
	_init_.Initialize()

	j := jsiiProxy_SendToQueue{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SendToQueue",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

// Deprecated: Use `SqsSendMessage`.
func NewSendToQueue_Override(s SendToQueue, queue awssqs.IQueue, props *SendToQueueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.SendToQueue",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SendToQueue) Bind(_task awsstepfunctions.Task) *awsstepfunctions.StepFunctionsTaskConfig {
	var returns *awsstepfunctions.StepFunctionsTaskConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_task},
		&returns,
	)

	return returns
}

