package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Mesh.
// Experimental.
type IMeshRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Mesh resource.
	// Experimental.
	MeshRef() *MeshReference
}

// The jsii proxy for IMeshRef
type jsiiProxy_IMeshRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IMeshRef) MeshRef() *MeshReference {
	var returns *MeshReference
	_jsii_.Get(
		j,
		"meshRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMeshRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMeshRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

