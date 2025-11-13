package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskDefinition.
// Experimental.
type ITaskDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TaskDefinition resource.
	// Experimental.
	TaskDefinitionRef() *TaskDefinitionReference
}

// The jsii proxy for ITaskDefinitionRef
type jsiiProxy_ITaskDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITaskDefinitionRef) TaskDefinitionRef() *TaskDefinitionReference {
	var returns *TaskDefinitionReference
	_jsii_.Get(
		j,
		"taskDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITaskDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

