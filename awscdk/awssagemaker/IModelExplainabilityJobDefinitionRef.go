package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelExplainabilityJobDefinition.
// Experimental.
type IModelExplainabilityJobDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ModelExplainabilityJobDefinition resource.
	// Experimental.
	ModelExplainabilityJobDefinitionRef() *ModelExplainabilityJobDefinitionReference
}

// The jsii proxy for IModelExplainabilityJobDefinitionRef
type jsiiProxy_IModelExplainabilityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
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

