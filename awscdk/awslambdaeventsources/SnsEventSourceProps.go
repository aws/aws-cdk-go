package awslambdaeventsources

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties forwarded to the Lambda Subscription.
//
// Example:
//   import sns "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var topic topic
//
//   var fn function
//
//   deadLetterQueue := sqs.NewQueue(this, jsii.String("deadLetterQueue"))
//   fn.AddEventSource(awscdk.NewSnsEventSource(topic, &SnsEventSourceProps{
//   	FilterPolicy: map[string]interface{}{
//   	},
//   	DeadLetterQueue: deadLetterQueue,
//   }))
//
type SnsEventSourceProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

