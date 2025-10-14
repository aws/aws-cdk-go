package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationTask.
// Experimental.
type IReplicationTaskRef interface {
	constructs.IConstruct
	// A reference to a ReplicationTask resource.
	// Experimental.
	ReplicationTaskRef() *ReplicationTaskReference
}

// The jsii proxy for IReplicationTaskRef
type jsiiProxy_IReplicationTaskRef struct {
	internal.Type__constructsIConstruct
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

