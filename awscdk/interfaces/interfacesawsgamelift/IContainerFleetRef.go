package interfacesawsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerFleet.
// Experimental.
type IContainerFleetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContainerFleet resource.
	// Experimental.
	ContainerFleetRef() *ContainerFleetReference
}

// The jsii proxy for IContainerFleetRef
type jsiiProxy_IContainerFleetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IContainerFleetRef) ContainerFleetRef() *ContainerFleetReference {
	var returns *ContainerFleetReference
	_jsii_.Get(
		j,
		"containerFleetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerFleetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerFleetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

