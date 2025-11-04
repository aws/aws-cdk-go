package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicRule.
// Experimental.
type ITopicRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a TopicRule resource.
	// Experimental.
	TopicRuleRef() *TopicRuleReference
}

// The jsii proxy for ITopicRuleRef
type jsiiProxy_ITopicRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ITopicRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

