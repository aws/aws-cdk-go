package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Mesh.
// Experimental.
type IMeshRef interface {
	constructs.IConstruct
	// A reference to a Mesh resource.
	// Experimental.
	MeshRef() *MeshReference
}

// The jsii proxy for IMeshRef
type jsiiProxy_IMeshRef struct {
	internal.Type__constructsIConstruct
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

