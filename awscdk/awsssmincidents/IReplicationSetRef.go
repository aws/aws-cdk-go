package awsssmincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSet.
// Experimental.
type IReplicationSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReplicationSet resource.
	// Experimental.
	ReplicationSetRef() *ReplicationSetReference
}

// The jsii proxy for IReplicationSetRef
type jsiiProxy_IReplicationSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IReplicationSetRef) ReplicationSetRef() *ReplicationSetReference {
	var returns *ReplicationSetReference
	_jsii_.Get(
		j,
		"replicationSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

