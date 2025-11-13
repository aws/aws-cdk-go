package interfacesawsamplifyuibuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsamplifyuibuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Component.
// Experimental.
type IComponentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Component resource.
	// Experimental.
	ComponentRef() *ComponentReference
}

// The jsii proxy for IComponentRef
type jsiiProxy_IComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IComponentRef) ComponentRef() *ComponentReference {
	var returns *ComponentReference
	_jsii_.Get(
		j,
		"componentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

