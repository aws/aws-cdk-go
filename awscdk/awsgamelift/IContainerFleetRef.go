package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerFleet.
// Experimental.
type IContainerFleetRef interface {
	constructs.IConstruct
	// A reference to a ContainerFleet resource.
	// Experimental.
	ContainerFleetRef() *ContainerFleetReference
}

// The jsii proxy for IContainerFleetRef
type jsiiProxy_IContainerFleetRef struct {
	internal.Type__constructsIConstruct
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

