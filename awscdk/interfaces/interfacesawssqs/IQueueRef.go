package interfacesawssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Queue.
// Experimental.
type IQueueRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Queue resource.
	// Experimental.
	QueueRef() *QueueReference
}

// The jsii proxy for IQueueRef
type jsiiProxy_IQueueRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IQueueRef) QueueRef() *QueueReference {
	var returns *QueueReference
	_jsii_.Get(
		j,
		"queueRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

