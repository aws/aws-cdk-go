package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskinesisfirehose"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
)

// Use an Amazon Data Firehose delivery stream as a subscription target.
//
// Example:
//   import firehose "github.com/aws/aws-cdk-go/awscdk"
//   var stream DeliveryStream
//
//
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   myTopic.AddSubscription(subscriptions.NewFirehoseSubscription(stream))
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html
//
type FirehoseSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for a Lambda function to subscribe to an SNS topic.
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for FirehoseSubscription
type jsiiProxy_FirehoseSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewFirehoseSubscription(deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseSubscriptionProps) FirehoseSubscription {
	_init_.Initialize()

	if err := validateNewFirehoseSubscriptionParameters(deliveryStream, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FirehoseSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.FirehoseSubscription",
		[]interface{}{deliveryStream, props},
		&j,
	)

	return &j
}

func NewFirehoseSubscription_Override(f FirehoseSubscription, deliveryStream awskinesisfirehose.IDeliveryStream, props *FirehoseSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.FirehoseSubscription",
		[]interface{}{deliveryStream, props},
		f,
	)
}

func (f *jsiiProxy_FirehoseSubscription) Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	if err := f.validateBindParameters(topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		f,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

