package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Properties for creating a new subscription.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws-samples/dummy/awscdklibawskinesisfirehose"
//   var stream DeliveryStream
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
//   })
//
type SubscriptionProps struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	Protocol SubscriptionProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The filter policy.
	FilterPolicy *map[string]SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
	// The topic to subscribe to.
	Topic ITopic `field:"required" json:"topic" yaml:"topic"`
}

