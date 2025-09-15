package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
)

// Use an sms address as a subscription target.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//
//   myTopic.AddSubscription(subscriptions.NewSmsSubscription(jsii.String("+15551231234")))
//
type SmsSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration used to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SmsSubscription
type jsiiProxy_SmsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewSmsSubscription(phoneNumber *string, props *SmsSubscriptionProps) SmsSubscription {
	_init_.Initialize()

	if err := validateNewSmsSubscriptionParameters(phoneNumber, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SmsSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		&j,
	)

	return &j
}

func NewSmsSubscription_Override(s SmsSubscription, phoneNumber *string, props *SmsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SmsSubscription",
		[]interface{}{phoneNumber, props},
		s,
	)
}

func (s *jsiiProxy_SmsSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	if err := s.validateBindParameters(_topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

