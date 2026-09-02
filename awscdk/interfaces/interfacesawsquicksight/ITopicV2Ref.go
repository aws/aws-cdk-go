package interfacesawsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TopicV2.
// Experimental.
type ITopicV2Ref interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TopicV2 resource.
	// Experimental.
	TopicV2Ref() *TopicV2Reference
}

// The jsii proxy for ITopicV2Ref
type jsiiProxy_ITopicV2Ref struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITopicV2Ref) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITopicV2Ref) TopicV2Ref() *TopicV2Reference {
	var returns *TopicV2Reference
	_jsii_.Get(
		j,
		"topicV2Ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicV2Ref) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITopicV2Ref) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

