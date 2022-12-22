package awspipes


// The parameters for using a Amazon SQS stream as a source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipeTargetSqsQueueParametersProperty := &pipeTargetSqsQueueParametersProperty{
//   	messageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	messageGroupId: jsii.String("messageGroupId"),
//   }
//
type CfnPipe_PipeTargetSqsQueueParametersProperty struct {
	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The FIFO message group ID to use as the target.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

