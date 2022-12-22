package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/aws-cdk-go/awscdk/awssqs"
)

// Options for URL subscriptions.
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
//   urlSubscriptionProps := &urlSubscriptionProps{
//   	deadLetterQueue: queue,
//   	filterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	protocol: awscdk.Aws_sns.subscriptionProtocol_HTTP,
//   	rawMessageDelivery: jsii.Boolean(false),
//   }
//
// Experimental.
type UrlSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Experimental.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Experimental.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The subscription's protocol.
	// Experimental.
	Protocol awssns.SubscriptionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	// Experimental.
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
}

