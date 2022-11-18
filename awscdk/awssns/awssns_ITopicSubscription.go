package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Topic subscription.
type ITopicSubscription interface {
	// Returns a configuration used to subscribe to an SNS topic.
	Bind(topic ITopic) *TopicSubscriptionConfig
}

// The jsii proxy for ITopicSubscription
type jsiiProxy_ITopicSubscription struct {
	_ byte // padding
}

func (i *jsiiProxy_ITopicSubscription) Bind(topic ITopic) *TopicSubscriptionConfig {
	if err := i.validateBindParameters(topic); err != nil {
		panic(err)
	}
	var returns *TopicSubscriptionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{topic},
		&returns,
	)

	return returns
}

