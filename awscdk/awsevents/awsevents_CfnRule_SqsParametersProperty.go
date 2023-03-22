package awsevents


// This structure includes the custom parameter to be used when the target is an SQS FIFO queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sqsParametersProperty := &sqsParametersProperty{
//   	messageGroupId: jsii.String("messageGroupId"),
//   }
//
type CfnRule_SqsParametersProperty struct {
	// The FIFO message group ID to use as the target.
	MessageGroupId *string `field:"required" json:"messageGroupId" yaml:"messageGroupId"`
}

