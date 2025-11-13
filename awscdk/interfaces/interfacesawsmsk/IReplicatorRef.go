package interfacesawsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Replicator.
// Experimental.
type IReplicatorRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Replicator resource.
	// Experimental.
	ReplicatorRef() *ReplicatorReference
}

// The jsii proxy for IReplicatorRef
type jsiiProxy_IReplicatorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IReplicatorRef) ReplicatorRef() *ReplicatorReference {
	var returns *ReplicatorReference
	_jsii_.Get(
		j,
		"replicatorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicatorRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicatorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

