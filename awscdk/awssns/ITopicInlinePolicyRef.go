package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicInlinePolicy.
// Experimental.
type ITopicInlinePolicyRef interface {
	constructs.IConstruct
	// A reference to a TopicInlinePolicy resource.
	// Experimental.
	TopicInlinePolicyRef() *TopicInlinePolicyReference
}

// The jsii proxy for ITopicInlinePolicyRef
type jsiiProxy_ITopicInlinePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITopicInlinePolicyRef) TopicInlinePolicyRef() *TopicInlinePolicyReference {
	var returns *TopicInlinePolicyReference
	_jsii_.Get(
		j,
		"topicInlinePolicyRef",
		&returns,
	)
	return returns
}

