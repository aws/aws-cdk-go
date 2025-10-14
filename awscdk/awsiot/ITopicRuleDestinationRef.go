package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicRuleDestination.
// Experimental.
type ITopicRuleDestinationRef interface {
	constructs.IConstruct
	// A reference to a TopicRuleDestination resource.
	// Experimental.
	TopicRuleDestinationRef() *TopicRuleDestinationReference
}

// The jsii proxy for ITopicRuleDestinationRef
type jsiiProxy_ITopicRuleDestinationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITopicRuleDestinationRef) TopicRuleDestinationRef() *TopicRuleDestinationReference {
	var returns *TopicRuleDestinationReference
	_jsii_.Get(
		j,
		"topicRuleDestinationRef",
		&returns,
	)
	return returns
}

