package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Options to subscribing to an SNS topic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   subscriptionProps := &subscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   }
//
// Experimental.
type SubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
}

