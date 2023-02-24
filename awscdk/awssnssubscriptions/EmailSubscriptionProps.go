package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Options for email subscriptions.
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
//   emailSubscriptionProps := &EmailSubscriptionProps{
//   	DeadLetterQueue: queue,
//   	FilterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	Json: jsii.Boolean(false),
//   }
//
type EmailSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// Indicates if the full notification JSON should be sent to the email address or just the message text.
	Json *bool `field:"optional" json:"json" yaml:"json"`
}

