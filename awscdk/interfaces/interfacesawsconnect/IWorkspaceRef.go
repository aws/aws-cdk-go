package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workspace.
// Experimental.
type IWorkspaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Workspace resource.
	// Experimental.
	WorkspaceRef() *WorkspaceReference
}

// The jsii proxy for IWorkspaceRef
type jsiiProxy_IWorkspaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IWorkspaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

