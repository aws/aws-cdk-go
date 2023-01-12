package awspipes


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
	// `CfnPipe.PipeTargetSqsQueueParametersProperty.MessageDeduplicationId`.
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// `CfnPipe.PipeTargetSqsQueueParametersProperty.MessageGroupId`.
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

