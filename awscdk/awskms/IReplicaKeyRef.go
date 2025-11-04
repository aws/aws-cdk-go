package awskms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ReplicaKey.
// Experimental.
type IReplicaKeyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ReplicaKey resource.
	// Experimental.
	ReplicaKeyRef() *ReplicaKeyReference
}

// The jsii proxy for IReplicaKeyRef
type jsiiProxy_IReplicaKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IReplicaKeyRef) ReplicaKeyRef() *ReplicaKeyReference {
	var returns *ReplicaKeyReference
	_jsii_.Get(
		j,
		"replicaKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicaKeyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IReplicaKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

