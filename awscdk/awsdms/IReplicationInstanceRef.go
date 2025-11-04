package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicationInstance.
// Experimental.
type IReplicationInstanceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReplicationInstance resource.
	// Experimental.
	ReplicationInstanceRef() *ReplicationInstanceReference
}

// The jsii proxy for IReplicationInstanceRef
type jsiiProxy_IReplicationInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IReplicationInstanceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicationInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

