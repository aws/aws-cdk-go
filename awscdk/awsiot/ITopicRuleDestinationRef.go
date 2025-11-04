package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicRuleDestination.
// Experimental.
type ITopicRuleDestinationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TopicRuleDestination resource.
	// Experimental.
	TopicRuleDestinationRef() *TopicRuleDestinationReference
}

// The jsii proxy for ITopicRuleDestinationRef
type jsiiProxy_ITopicRuleDestinationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITopicRuleDestinationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicRuleDestinationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

