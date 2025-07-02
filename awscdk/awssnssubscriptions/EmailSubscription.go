package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
)

// Use an email address as a subscription target.
//
// Email subscriptions require confirmation.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("Topic"))
//   emailAddress := awscdk.NewCfnParameter(this, jsii.String("email-param"))
//
//   myTopic.AddSubscription(subscriptions.NewEmailSubscription(emailAddress.valueAsString))
//
type EmailSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for an email address to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for EmailSubscription
type jsiiProxy_EmailSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewEmailSubscription(emailAddress *string, props *EmailSubscriptionProps) EmailSubscription {
	_init_.Initialize()

	if err := validateNewEmailSubscriptionParameters(emailAddress, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EmailSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		&j,
	)

	return &j
}

func NewEmailSubscription_Override(e EmailSubscription, emailAddress *string, props *EmailSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.EmailSubscription",
		[]interface{}{emailAddress, props},
		e,
	)
}

func (e *jsiiProxy_EmailSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	if err := e.validateBindParameters(_topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		e,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

