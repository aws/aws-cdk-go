package interfacesawseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Addon.
// Experimental.
type IAddonRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Addon resource.
	// Experimental.
	AddonRef() *AddonReference
}

// The jsii proxy for IAddonRef
type jsiiProxy_IAddonRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAddonRef) AddonRef() *AddonReference {
	var returns *AddonReference
	_jsii_.Get(
		j,
		"addonRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddonRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAddonRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

