package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationInstance.
// Experimental.
type IReplicationInstanceRef interface {
	constructs.IConstruct
	// A reference to a ReplicationInstance resource.
	// Experimental.
	ReplicationInstanceRef() *ReplicationInstanceReference
}

// The jsii proxy for IReplicationInstanceRef
type jsiiProxy_IReplicationInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReplicationInstanceRef) ReplicationInstanceRef() *ReplicationInstanceReference {
	var returns *ReplicationInstanceReference
	_jsii_.Get(
		j,
		"replicationInstanceRef",
		&returns,
	)
	return returns
}

