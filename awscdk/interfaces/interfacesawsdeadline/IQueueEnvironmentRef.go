package interfacesawsdeadline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdeadline/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueEnvironment.
// Experimental.
type IQueueEnvironmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a QueueEnvironment resource.
	// Experimental.
	QueueEnvironmentRef() *QueueEnvironmentReference
}

// The jsii proxy for IQueueEnvironmentRef
type jsiiProxy_IQueueEnvironmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IQueueEnvironmentRef) QueueEnvironmentRef() *QueueEnvironmentReference {
	var returns *QueueEnvironmentReference
	_jsii_.Get(
		j,
		"queueEnvironmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueEnvironmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQueueEnvironmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

