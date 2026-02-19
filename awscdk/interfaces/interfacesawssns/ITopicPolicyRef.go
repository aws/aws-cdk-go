package interfacesawssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicPolicy.
// Experimental.
type ITopicPolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TopicPolicy resource.
	// Experimental.
	TopicPolicyRef() *TopicPolicyReference
}

// The jsii proxy for ITopicPolicyRef
type jsiiProxy_ITopicPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITopicPolicyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITopicPolicyRef) TopicPolicyRef() *TopicPolicyReference {
	var returns *TopicPolicyReference
	_jsii_.Get(
		j,
		"topicPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicPolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

