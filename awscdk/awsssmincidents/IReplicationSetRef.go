package awsssmincidents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssmincidents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSet.
// Experimental.
type IReplicationSetRef interface {
	constructs.IConstruct
	// A reference to a ReplicationSet resource.
	// Experimental.
	ReplicationSetRef() *ReplicationSetReference
}

// The jsii proxy for IReplicationSetRef
type jsiiProxy_IReplicationSetRef struct {
	internal.Type__constructsIConstruct
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

