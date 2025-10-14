package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicRule.
// Experimental.
type ITopicRuleRef interface {
	constructs.IConstruct
	// A reference to a TopicRule resource.
	// Experimental.
	TopicRuleRef() *TopicRuleReference
}

// The jsii proxy for ITopicRuleRef
type jsiiProxy_ITopicRuleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITopicRuleRef) TopicRuleRef() *TopicRuleReference {
	var returns *TopicRuleReference
	_jsii_.Get(
		j,
		"topicRuleRef",
		&returns,
	)
	return returns
}

