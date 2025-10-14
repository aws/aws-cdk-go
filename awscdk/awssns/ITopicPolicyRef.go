package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicPolicy.
// Experimental.
type ITopicPolicyRef interface {
	constructs.IConstruct
	// A reference to a TopicPolicy resource.
	// Experimental.
	TopicPolicyRef() *TopicPolicyReference
}

// The jsii proxy for ITopicPolicyRef
type jsiiProxy_ITopicPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITopicPolicyRef) TopicPolicyRef() *TopicPolicyReference {
	var returns *TopicPolicyReference
	_jsii_.Get(
		j,
		"topicPolicyRef",
		&returns,
	)
	return returns
}

