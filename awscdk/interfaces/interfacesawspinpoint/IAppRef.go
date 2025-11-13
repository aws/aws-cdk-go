package interfacesawspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a App.
// Experimental.
type IAppRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a App resource.
	// Experimental.
	AppRef() *AppReference
}

// The jsii proxy for IAppRef
type jsiiProxy_IAppRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAppRef) AppRef() *AppReference {
	var returns *AppReference
	_jsii_.Get(
		j,
		"appRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

