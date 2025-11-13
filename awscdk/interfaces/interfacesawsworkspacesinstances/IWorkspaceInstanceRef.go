package interfacesawsworkspacesinstances

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspacesinstances/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkspaceInstance.
// Experimental.
type IWorkspaceInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkspaceInstance resource.
	// Experimental.
	WorkspaceInstanceRef() *WorkspaceInstanceReference
}

// The jsii proxy for IWorkspaceInstanceRef
type jsiiProxy_IWorkspaceInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IWorkspaceInstanceRef) WorkspaceInstanceRef() *WorkspaceInstanceReference {
	var returns *WorkspaceInstanceReference
	_jsii_.Get(
		j,
		"workspaceInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspaceInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspaceInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

