package awssqs


// Properties to associate SQS queues with a policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//
//   queuePolicyProps := &queuePolicyProps{
//   	queues: []iQueue{
//   		queue,
//   	},
//   }
//
type QueuePolicyProps struct {
	// The set of queues this policy applies to.
	Queues *[]IQueue `field:"required" json:"queues" yaml:"queues"`
}

