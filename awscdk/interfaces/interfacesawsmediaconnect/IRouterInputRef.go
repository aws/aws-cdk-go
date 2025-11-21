package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouterInput.
// Experimental.
type IRouterInputRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RouterInput resource.
	// Experimental.
	RouterInputRef() *RouterInputReference
}

// The jsii proxy for IRouterInputRef
type jsiiProxy_IRouterInputRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRouterInputRef) RouterInputRef() *RouterInputReference {
	var returns *RouterInputReference
	_jsii_.Get(
		j,
		"routerInputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInputRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterInputRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

