package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
	"github.com/aws/aws-cdk-go/awscdk/awsstepfunctions"
)

// Properties for sending a message to an SQS queue.
//
// Example:
//   queue := sqs.NewQueue(this, jsii.String("Queue"))
//
//   // Use a field from the execution data as message.
//   task1 := tasks.NewSqsSendMessage(this, jsii.String("Send1"), &SqsSendMessageProps{
//   	Queue: Queue,
//   	MessageBody: sfn.TaskInput_FromJsonPathAt(jsii.String("$.message")),
//   })
//
//   // Combine a field from the execution data with
//   // a literal object.
//   task2 := tasks.NewSqsSendMessage(this, jsii.String("Send2"), &SqsSendMessageProps{
//   	Queue: Queue,
//   	MessageBody: sfn.TaskInput_FromObject(map[string]interface{}{
//   		"field1": jsii.String("somedata"),
//   		"field2": sfn.JsonPath_stringAt(jsii.String("$.field2")),
//   	}),
//   })
//
// Experimental.
type SqsSendMessageProps struct {
	// An optional description for this state.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Timeout for the heartbeat.
	// Experimental.
	Heartbeat awscdk.Duration `field:"optional" json:"heartbeat" yaml:"heartbeat"`
	// JSONPath expression to select part of the state to be the input to this state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// input to be the empty object {}.
	// Experimental.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// AWS Step Functions integrates with services directly in the Amazon States Language.
	//
	// You can control these AWS services using service integration patterns.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/connect-to-resource.html#connect-wait-token
	//
	// Experimental.
	IntegrationPattern awsstepfunctions.IntegrationPattern `field:"optional" json:"integrationPattern" yaml:"integrationPattern"`
	// JSONPath expression to select select a portion of the state output to pass to the next state.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the effective
	// output to be the empty object {}.
	// Experimental.
	OutputPath *string `field:"optional" json:"outputPath" yaml:"outputPath"`
	// JSONPath expression to indicate where to inject the state's output.
	//
	// May also be the special value JsonPath.DISCARD, which will cause the state's
	// input to become its output.
	// Experimental.
	ResultPath *string `field:"optional" json:"resultPath" yaml:"resultPath"`
	// The JSON that will replace the state's raw result and become the effective result before ResultPath is applied.
	//
	// You can use ResultSelector to create a payload with values that are static
	// or selected from the state's raw result.
	// See: https://docs.aws.amazon.com/step-functions/latest/dg/input-output-inputpath-params.html#input-output-resultselector
	//
	// Experimental.
	ResultSelector *map[string]interface{} `field:"optional" json:"resultSelector" yaml:"resultSelector"`
	// Timeout for the state machine.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
	// The text message to send to the queue.
	// Experimental.
	MessageBody awsstepfunctions.TaskInput `field:"required" json:"messageBody" yaml:"messageBody"`
	// The SQS queue that messages will be sent to.
	// Experimental.
	Queue awssqs.IQueue `field:"required" json:"queue" yaml:"queue"`
	// The length of time, for which to delay a message.
	//
	// Messages that you send to the queue remain invisible to consumers for the duration
	// of the delay period. The maximum allowed delay is 15 minutes.
	// Experimental.
	Delay awscdk.Duration `field:"optional" json:"delay" yaml:"delay"`
	// The token used for deduplication of sent messages.
	//
	// Any messages sent with the same deduplication ID are accepted successfully,
	// but aren't delivered during the 5-minute deduplication interval.
	// Experimental.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The tag that specifies that a message belongs to a specific message group.
	//
	// Messages that belong to the same message group are processed in a FIFO manner.
	// Messages in different message groups might be processed out of order.
	// Experimental.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

