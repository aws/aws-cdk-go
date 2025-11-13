package interfacesawsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConsumableResource.
// Experimental.
type IConsumableResourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConsumableResource resource.
	// Experimental.
	ConsumableResourceRef() *ConsumableResourceReference
}

// The jsii proxy for IConsumableResourceRef
type jsiiProxy_IConsumableResourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConsumableResourceRef) ConsumableResourceRef() *ConsumableResourceReference {
	var returns *ConsumableResourceReference
	_jsii_.Get(
		j,
		"consumableResourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConsumableResourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

