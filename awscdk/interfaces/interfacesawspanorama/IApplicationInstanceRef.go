package interfacesawspanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawspanorama/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationInstance.
// Experimental.
type IApplicationInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationInstance resource.
	// Experimental.
	ApplicationInstanceRef() *ApplicationInstanceReference
}

// The jsii proxy for IApplicationInstanceRef
type jsiiProxy_IApplicationInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IApplicationInstanceRef) ApplicationInstanceRef() *ApplicationInstanceReference {
	var returns *ApplicationInstanceReference
	_jsii_.Get(
		j,
		"applicationInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

