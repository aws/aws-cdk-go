package interfacesawssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelExplainabilityJobDefinition.
// Experimental.
type IModelExplainabilityJobDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ModelExplainabilityJobDefinition resource.
	// Experimental.
	ModelExplainabilityJobDefinitionRef() *ModelExplainabilityJobDefinitionReference
}

// The jsii proxy for IModelExplainabilityJobDefinitionRef
type jsiiProxy_IModelExplainabilityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IModelExplainabilityJobDefinitionRef) ModelExplainabilityJobDefinitionRef() *ModelExplainabilityJobDefinitionReference {
	var returns *ModelExplainabilityJobDefinitionReference
	_jsii_.Get(
		j,
		"modelExplainabilityJobDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelExplainabilityJobDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IModelExplainabilityJobDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

