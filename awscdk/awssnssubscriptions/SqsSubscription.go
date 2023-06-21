package awssnssubscriptions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssnssubscriptions/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Use an SQS queue as a subscription target.
//
// Example:
//   var queue queue
//
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.AddSubscription(subscriptions.NewSqsSubscription(queue))
//
type SqsSubscription interface {
	awssns.ITopicSubscription
	// Returns a configuration for an SQS queue to subscribe to an SNS topic.
	Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig
}

// The jsii proxy struct for SqsSubscription
type jsiiProxy_SqsSubscription struct {
	internal.Type__awssnsITopicSubscription
}

func NewSqsSubscription(queue awssqs.IQueue, props *SqsSubscriptionProps) SqsSubscription {
	_init_.Initialize()

	if err := validateNewSqsSubscriptionParameters(queue, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsSubscription{}

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		&j,
	)

	return &j
}

func NewSqsSubscription_Override(s SqsSubscription, queue awssqs.IQueue, props *SqsSubscriptionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_sns_subscriptions.SqsSubscription",
		[]interface{}{queue, props},
		s,
	)
}

func (s *jsiiProxy_SqsSubscription) Bind(topic awssns.ITopic) *awssns.TopicSubscriptionConfig {
	if err := s.validateBindParameters(topic); err != nil {
		panic(err)
	}
	var returns *awssns.TopicSubscriptionConfig

	_jsii_.Invoke(
		s,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

