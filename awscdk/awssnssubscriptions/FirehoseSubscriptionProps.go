package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for an Amazon Data Firehose subscription.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var stream deliveryStream
//
//
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   myTopic.AddSubscription(subscriptions.NewFirehoseSubscription(stream, &FirehoseSubscriptionProps{
//   	RawMessageDelivery: jsii.Boolean(true),
//   }))
//
type FirehoseSubscriptionProps struct {
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
	// Whether to remove any Amazon SNS metadata from published messages.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-large-payload-raw-message-delivery.html
	//
	// Default: false.
	//
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The role to assume to write messages to the Amazon Data Firehose delivery stream.
	// Default: - A new Role is created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}

