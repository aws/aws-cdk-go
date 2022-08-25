package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract action for TopicRule.
// Experimental.
type IAction interface {
	// Returns the topic rule action specification.
	// Experimental.
	Bind(topicRule ITopicRule) *ActionConfig
}

// The jsii proxy for IAction
type jsiiProxy_IAction struct {
	_ byte // padding
}

func (i *jsiiProxy_IAction) Bind(topicRule ITopicRule) *ActionConfig {
	var returns *ActionConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{topicRule},
		&returns,
	)

	return returns
}

