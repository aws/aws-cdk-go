package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSubnetGroup.
// Experimental.
type IReplicationSubnetGroupRef interface {
	constructs.IConstruct
	// A reference to a ReplicationSubnetGroup resource.
	// Experimental.
	ReplicationSubnetGroupRef() *ReplicationSubnetGroupReference
}

// The jsii proxy for IReplicationSubnetGroupRef
type jsiiProxy_IReplicationSubnetGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReplicationSubnetGroupRef) ReplicationSubnetGroupRef() *ReplicationSubnetGroupReference {
	var returns *ReplicationSubnetGroupReference
	_jsii_.Get(
		j,
		"replicationSubnetGroupRef",
		&returns,
	)
	return returns
}

