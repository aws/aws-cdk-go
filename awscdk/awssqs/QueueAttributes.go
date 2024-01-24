package awssqs


// Reference to a queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   queueAttributes := &QueueAttributes{
//   	QueueArn: jsii.String("queueArn"),
//
//   	// the properties below are optional
//   	Fifo: jsii.Boolean(false),
//   	KeyArn: jsii.String("keyArn"),
//   	QueueName: jsii.String("queueName"),
//   	QueueUrl: jsii.String("queueUrl"),
//   }
//
type QueueAttributes struct {
	// The ARN of the queue.
	QueueArn *string `field:"required" json:"queueArn" yaml:"queueArn"`
	// Whether this queue is an Amazon SQS FIFO queue. If false, this is a standard queue.
	//
	// In case of a FIFO queue which is imported from a token, this value has to be explicitly set to true.
	// Default: - if fifo is not specified, the property will be determined based on the queue name (not possible for FIFO queues imported from a token).
	//
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// KMS encryption key, if this queue is server-side encrypted by a KMS key.
	// Default: - None.
	//
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
	// The name of the queue.
	// Default: if queue name is not specified, the name will be derived from the queue ARN.
	//
	QueueName *string `field:"optional" json:"queueName" yaml:"queueName"`
	// The URL of the queue.
	// See: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html#sqs-general-identifiers
	//
	// Default: - 'https://sqs.<region-endpoint>/<account-ID>/<queue-name>'
	//
	QueueUrl *string `field:"optional" json:"queueUrl" yaml:"queueUrl"`
}

