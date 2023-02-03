package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Properties for sending a message to an SQS queue.
//
// Example:
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   // Use a field from the execution data as message.
//   task1 := tasks.NewSqsSendMessage(this, jsii.String("Send1"), &sqsSendMessageProps{
//   	queue: queue,
//   	messageBody: sfn.taskInput.fromJsonPathAt(jsii.String("$.message")),
//   })
//
//   // Combine a field from the execution data with
//   // a literal object.
//   task2 := tasks.NewSqsSendMessage(this, jsii.String("Send2"), &sqsSendMessageProps{
//   	queue: queue,
//   	messageBody: sfn.*taskInput.fromObject(map[string]interface{}{
//   		"field1": jsii.String("somedata"),
//   		"field2": sfn.JsonPath.stringAt(jsii.String("$.field2")),
//   	}),
//   })
//
type SqsSendMessageProps struct {
	// An optional description for this state.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Credentials for an IAM Role that the State Machine assumes for executing the task.
	//
	// This enables cross-account resource invocations.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/concepts-access-cross-acct-resources.html
	//
	Credentials *awsstepfunctions.Credentials `field:"optional" json:"credentials" yaml:"credentials"`
	// Timeout for the heartbeat.
	// Deprecated: use `heartbeatTimeout`.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// Timeout for the heartbeat.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	HeartbeatTimeout awsstepfunctions.Timeout `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the task.
	//
	// [disable-awslint:duration-prop-type] is needed because all props interface in
	// aws-stepfunctions-tasks extend this interface.
	TaskTimeout awsstepfunctions.Timeout `field:"optional" json:"taskTimeout" yaml:"taskTimeout"`
	// Timeout for the task.
	// Deprecated: use `taskTimeout`.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The text message to send to the queue.
	MessageBody awsstepfunctions.TaskInput `field:"required" json:"messageBody" yaml:"messageBody"`
	// The SQS queue that messages will be sent to.
	Queue awssqs.IQueue `field:"required" json:"queue" yaml:"queue"`
	// The length of time, for which to delay a message.
	//
	// Messages that you send to the queue remain invisible to consumers for the duration
	// of the delay period. The maximum allowed delay is 15 minutes.
	Delay awscdk.Duration `field:"optional" json:"delay" yaml:"delay"`
	// The token used for deduplication of sent messages.
	//
	// Any messages sent with the same deduplication ID are accepted successfully,
	// but aren't delivered during the 5-minute deduplication interval.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The tag that specifies that a message belongs to a specific message group.
	//
	// Messages that belong to the same message group are processed in a FIFO manner.
	// Messages in different message groups might be processed out of order.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

