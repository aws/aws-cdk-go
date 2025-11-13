package interfacesawssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueInlinePolicy.
// Experimental.
type IQueueInlinePolicyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QueueInlinePolicy resource.
	// Experimental.
	QueueInlinePolicyRef() *QueueInlinePolicyReference
}

// The jsii proxy for IQueueInlinePolicyRef
type jsiiProxy_IQueueInlinePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IQueueInlinePolicyRef) QueueInlinePolicyRef() *QueueInlinePolicyReference {
	var returns *QueueInlinePolicyReference
	_jsii_.Get(
		j,
		"queueInlinePolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueInlinePolicyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueInlinePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

