package awssqs


// Dead letter queue settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//
//   deadLetterQueue := &deadLetterQueue{
//   	maxReceiveCount: jsii.Number(123),
//   	queue: queue,
//   }
//
type DeadLetterQueue struct {
	// The number of times a message can be unsuccesfully dequeued before being moved to the dead-letter queue.
	MaxReceiveCount *float64 `field:"required" json:"maxReceiveCount" yaml:"maxReceiveCount"`
	// The dead-letter queue to which Amazon SQS moves messages after the value of maxReceiveCount is exceeded.
	Queue IQueue `field:"required" json:"queue" yaml:"queue"`
}

