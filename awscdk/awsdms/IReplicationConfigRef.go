package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationConfig.
// Experimental.
type IReplicationConfigRef interface {
	constructs.IConstruct
	// A reference to a ReplicationConfig resource.
	// Experimental.
	ReplicationConfigRef() *ReplicationConfigReference
}

// The jsii proxy for IReplicationConfigRef
type jsiiProxy_IReplicationConfigRef struct {
	internal.Type__constructsIConstruct
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

