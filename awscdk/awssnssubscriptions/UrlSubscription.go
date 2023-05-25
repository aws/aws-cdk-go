package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
)

// Use a URL as a subscription target.
//
// The message will be POSTed to the given URL.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.AddSubscription(subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/")))
//
// See: https://docs.aws.amazon.com/sns/latest/dg/sns-http-https-endpoint-as-subscriber.html
//
type UrlSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for a URL to subscribe to an SNS topic.
	Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for UrlSubscription
type jsiiProxy_UrlSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewUrlSubscription(url *string, props *UrlSubscriptionProps) UrlSubscription {
	_init_.Initialize()

	if err := validateNewUrlSubscriptionParameters(url, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_UrlSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		&j,
	)

	return &j
}

func NewUrlSubscription_Override(u UrlSubscription, url *string, props *UrlSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.UrlSubscription",
		[]interface{}{url, props},
		u,
	)
}

func (u *jsiiProxy_UrlSubscription) Bind(_topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	if err := u.validateBindParameters(_topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		u,
		"bind",
		[]interface{}{_topic},
		&returns,
	)

	return returns
}

