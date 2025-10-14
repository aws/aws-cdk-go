package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerGroupDefinition.
// Experimental.
type IContainerGroupDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ContainerGroupDefinition resource.
	// Experimental.
	ContainerGroupDefinitionRef() *ContainerGroupDefinitionReference
}

// The jsii proxy for IContainerGroupDefinitionRef
type jsiiProxy_IContainerGroupDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContainerGroupDefinitionRef) ContainerGroupDefinitionRef() *ContainerGroupDefinitionReference {
	var returns *ContainerGroupDefinitionReference
	_jsii_.Get(
		j,
		"containerGroupDefinitionRef",
		&returns,
	)
	return returns
}

