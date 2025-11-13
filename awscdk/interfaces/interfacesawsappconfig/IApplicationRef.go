package interfacesawsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Application.
// Experimental.
type IApplicationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Application resource.
	// Experimental.
	ApplicationRef() *ApplicationReference
}

// The jsii proxy for IApplicationRef
type jsiiProxy_IApplicationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IApplicationRef) ApplicationRef() *ApplicationReference {
	var returns *ApplicationReference
	_jsii_.Get(
		j,
		"applicationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

