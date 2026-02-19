package interfacesawssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicInlinePolicy.
// Experimental.
type ITopicInlinePolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TopicInlinePolicy resource.
	// Experimental.
	TopicInlinePolicyRef() *TopicInlinePolicyReference
}

// The jsii proxy for ITopicInlinePolicyRef
type jsiiProxy_ITopicInlinePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITopicInlinePolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_ITopicInlinePolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicInlinePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

