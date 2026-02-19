package interfacesawsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationTask.
// Experimental.
type IReplicationTaskRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReplicationTask resource.
	// Experimental.
	ReplicationTaskRef() *ReplicationTaskReference
}

// The jsii proxy for IReplicationTaskRef
type jsiiProxy_IReplicationTaskRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IReplicationTaskRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IReplicationTaskRef) ReplicationTaskRef() *ReplicationTaskReference {
	var returns *ReplicationTaskReference
	_jsii_.Get(
		j,
		"replicationTaskRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationTaskRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationTaskRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

