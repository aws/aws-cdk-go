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
//   var filterOrPolicy filterOrPolicy
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   emailSubscriptionProps := &EmailSubscriptionProps{
//   	DeadLetterQueue: queue,
//   	FilterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	FilterPolicyWithMessageBody: map[string]*filterOrPolicy{
//   		"filterPolicyWithMessageBodyKey": filterOrPolicy,
//   	},
//   	Json: jsii.Boolean(false),
//   }
//
type EmailSubscriptionProps struct {
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Default: - No dead letter queue enabled.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	// Default: - all messages are delivered.
	//
	FilterPolicy *map[string]awssns.SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The filter policy that is applied on the message body.
	//
	// To apply a filter policy to the message attributes, use `filterPolicy`. A maximum of one of `filterPolicyWithMessageBody` and `filterPolicy` may be used.
	// Default: - all messages are delivered.
	//
	FilterPolicyWithMessageBody *map[string]awssns.FilterOrPolicy `field:"optional" json:"filterPolicyWithMessageBody" yaml:"filterPolicyWithMessageBody"`
	// Indicates if the full notification JSON should be sent to the email address or just the message text.
	// Default: false (Message text).
	//
	Json *bool `field:"optional" json:"json" yaml:"json"`
}

