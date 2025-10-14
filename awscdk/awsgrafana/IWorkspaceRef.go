package awsgrafana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgrafana/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workspace.
// Experimental.
type IWorkspaceRef interface {
	constructs.IConstruct
	// A reference to a Workspace resource.
	// Experimental.
	WorkspaceRef() *WorkspaceReference
}

// The jsii proxy for IWorkspaceRef
type jsiiProxy_IWorkspaceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWorkspaceRef) WorkspaceRef() *WorkspaceReference {
	var returns *WorkspaceReference
	_jsii_.Get(
		j,
		"workspaceRef",
		&returns,
	)
	return returns
}

