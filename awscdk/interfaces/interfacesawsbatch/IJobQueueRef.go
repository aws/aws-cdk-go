package interfacesawsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobQueue.
// Experimental.
type IJobQueueRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a JobQueue resource.
	// Experimental.
	JobQueueRef() *JobQueueReference
}

// The jsii proxy for IJobQueueRef
type jsiiProxy_IJobQueueRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IJobQueueRef) JobQueueRef() *JobQueueReference {
	var returns *JobQueueReference
	_jsii_.Get(
		j,
		"jobQueueRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueueRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IJobQueueRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

