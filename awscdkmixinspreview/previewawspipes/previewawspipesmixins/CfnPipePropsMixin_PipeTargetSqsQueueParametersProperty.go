package previewawspipesmixins


// The parameters for using a Amazon SQS stream as a target.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   pipeTargetSqsQueueParametersProperty := &PipeTargetSqsQueueParametersProperty{
//   	MessageDeduplicationId: jsii.String("messageDeduplicationId"),
//   	MessageGroupId: jsii.String("messageGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html
//
type CfnPipePropsMixin_PipeTargetSqsQueueParametersProperty struct {
	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html#cfn-pipes-pipe-pipetargetsqsqueueparameters-messagededuplicationid
	//
	MessageDeduplicationId *string `field:"optional" json:"messageDeduplicationId" yaml:"messageDeduplicationId"`
	// The FIFO message group ID to use as the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-pipetargetsqsqueueparameters.html#cfn-pipes-pipe-pipetargetsqsqueueparameters-messagegroupid
	//
	MessageGroupId *string `field:"optional" json:"messageGroupId" yaml:"messageGroupId"`
}

