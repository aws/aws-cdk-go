package interfacesawsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerGroupDefinition.
// Experimental.
type IContainerGroupDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ContainerGroupDefinition resource.
	// Experimental.
	ContainerGroupDefinitionRef() *ContainerGroupDefinitionReference
}

// The jsii proxy for IContainerGroupDefinitionRef
type jsiiProxy_IContainerGroupDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IContainerGroupDefinitionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IContainerGroupDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerGroupDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

