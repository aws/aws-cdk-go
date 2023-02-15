package awssqs


// Reference to a queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueAttributes := &queueAttributes{
//   	queueArn: jsii.String("queueArn"),
//
//   	// the properties below are optional
//   	fifo: jsii.Boolean(false),
//   	keyArn: jsii.String("keyArn"),
//   	queueName: jsii.String("queueName"),
//   	queueUrl: jsii.String("queueUrl"),
//   }
//
type QueueAttributes struct {
	// The ARN of the queue.
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
	// Whether this queue is an Amazon SQS FIFO queue. If false, this is a standard queue.
	//
	// In case of a FIFO queue which is imported from a token, this value has to be explicitly set to true.
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// KMS encryption key, if this queue is server-side encrypted by a KMS key.
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// The name of the queue.
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The URL of the queue.
	// See: https://docs.aws.amazon.com/sdk-for-net/v2/developer-guide/QueueURL.html
	//
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
}

