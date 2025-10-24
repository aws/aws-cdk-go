package awsevents


// The custom parameters for EventBridge to use for a target that is an Amazon SQS fair or FIFO queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsParametersProperty := &SqsParametersProperty{
//   	MessageGroupId: jsii.String("messageGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sqsparameters.html
//
type CfnRule_SqsParametersProperty struct {
	// The ID of the message group to use as the target.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sqsparameters.html#cfn-events-rule-sqsparameters-messagegroupid
	//
	MessageGroupId *string `field:"required" json:"messageGroupId" yaml:"messageGroupId"`
}

