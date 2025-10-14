package awselasticache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticache/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationGroup.
// Experimental.
type IReplicationGroupRef interface {
	constructs.IConstruct
	// A reference to a ReplicationGroup resource.
	// Experimental.
	ReplicationGroupRef() *ReplicationGroupReference
}

// The jsii proxy for IReplicationGroupRef
type jsiiProxy_IReplicationGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReplicationGroupRef) ReplicationGroupRef() *ReplicationGroupReference {
	var returns *ReplicationGroupReference
	_jsii_.Get(
		j,
		"replicationGroupRef",
		&returns,
	)
	return returns
}

