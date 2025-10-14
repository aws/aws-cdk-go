package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Topic.
// Experimental.
type ITopicRef interface {
	constructs.IConstruct
	// A reference to a Topic resource.
	// Experimental.
	TopicRef() *TopicReference
}

// The jsii proxy for ITopicRef
type jsiiProxy_ITopicRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITopicRef) TopicRef() *TopicReference {
	var returns *TopicReference
	_jsii_.Get(
		j,
		"topicRef",
		&returns,
	)
	return returns
}

