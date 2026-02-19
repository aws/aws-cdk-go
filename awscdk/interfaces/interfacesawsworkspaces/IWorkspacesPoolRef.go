package interfacesawsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsworkspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkspacesPool.
// Experimental.
type IWorkspacesPoolRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkspacesPool resource.
	// Experimental.
	WorkspacesPoolRef() *WorkspacesPoolReference
}

// The jsii proxy for IWorkspacesPoolRef
type jsiiProxy_IWorkspacesPoolRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkspacesPoolRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IWorkspacesPoolRef) WorkspacesPoolRef() *WorkspacesPoolReference {
	var returns *WorkspacesPoolReference
	_jsii_.Get(
		j,
		"workspacesPoolRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspacesPoolRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspacesPoolRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

