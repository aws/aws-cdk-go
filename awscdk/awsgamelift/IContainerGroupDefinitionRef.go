package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerGroupDefinition.
// Experimental.
type IContainerGroupDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ContainerGroupDefinition resource.
	// Experimental.
	ContainerGroupDefinitionRef() *ContainerGroupDefinitionReference
}

// The jsii proxy for IContainerGroupDefinitionRef
type jsiiProxy_IContainerGroupDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IContainerGroupDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

