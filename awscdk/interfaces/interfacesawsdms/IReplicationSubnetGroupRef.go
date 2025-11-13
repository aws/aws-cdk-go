package interfacesawsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationSubnetGroup.
// Experimental.
type IReplicationSubnetGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ReplicationSubnetGroup resource.
	// Experimental.
	ReplicationSubnetGroupRef() *ReplicationSubnetGroupReference
}

// The jsii proxy for IReplicationSubnetGroupRef
type jsiiProxy_IReplicationSubnetGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IReplicationSubnetGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationSubnetGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

