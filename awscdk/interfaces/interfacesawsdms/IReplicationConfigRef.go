package interfacesawsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfig.
// Experimental.
type IReplicationConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReplicationConfig resource.
	// Experimental.
	ReplicationConfigRef() *ReplicationConfigReference
}

// The jsii proxy for IReplicationConfigRef
type jsiiProxy_IReplicationConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IReplicationConfigRef) ReplicationConfigRef() *ReplicationConfigReference {
	var returns *ReplicationConfigReference
	_jsii_.Get(
		j,
		"replicationConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

