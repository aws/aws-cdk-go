package interfacesawsiottwinmaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entity.
// Experimental.
type IEntityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Entity resource.
	// Experimental.
	EntityRef() *EntityReference
}

// The jsii proxy for IEntityRef
type jsiiProxy_IEntityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEntityRef) EntityRef() *EntityReference {
	var returns *EntityReference
	_jsii_.Get(
		j,
		"entityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEntityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

