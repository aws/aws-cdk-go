package interfacesawsdatabrew

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Project.
// Experimental.
type IProjectRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Project resource.
	// Experimental.
	ProjectRef() *ProjectReference
}

// The jsii proxy for IProjectRef
type jsiiProxy_IProjectRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IProjectRef) ProjectRef() *ProjectReference {
	var returns *ProjectReference
	_jsii_.Get(
		j,
		"projectRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProjectRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProjectRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

