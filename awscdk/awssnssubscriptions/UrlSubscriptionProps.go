package awssnssubscriptions

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Options for URL subscriptions.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.AddSubscription(
//   subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/"), &UrlSubscriptionProps{
//   	DeliveryPolicy: &DeliveryPolicy{
//   		HealthyRetryPolicy: &HealthyRetryPolicy{
//   			MinDelayTarget: awscdk.Duration_Seconds(jsii.Number(5)),
//   			MaxDelayTarget: awscdk.Duration_*Seconds(jsii.Number(10)),
//   			NumRetries: jsii.Number(6),
//   			BackoffFunction: sns.BackoffFunction_EXPONENTIAL,
//   		},
//   		ThrottlePolicy: &ThrottlePolicy{
//   			MaxReceivesPerSecond: jsii.Number(10),
//   		},
//   		RequestPolicy: &RequestPolicy{
//   			HeaderContentType: jsii.String("application/json"),
//   		},
//   	},
//   }))
//
type UrlSubscriptionProps struct {
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
	// The delivery policy.
	// Default: - if the initial delivery of the message fails, three retries with a delay between failed attempts set at 20 seconds.
	//
	DeliveryPolicy *awssns.DeliveryPolicy `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The subscription's protocol.
	// Default: - Protocol is derived from url.
	//
	Protocol awssns.SubscriptionProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The message to the queue is the same as it was sent to the topic.
	//
	// If false, the message will be wrapped in an SNS envelope.
	// Default: false.
	//
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
}

