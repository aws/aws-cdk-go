package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelBiasJobDefinition.
// Experimental.
type IModelBiasJobDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ModelBiasJobDefinition resource.
	// Experimental.
	ModelBiasJobDefinitionRef() *ModelBiasJobDefinitionReference
}

// The jsii proxy for IModelBiasJobDefinitionRef
type jsiiProxy_IModelBiasJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelBiasJobDefinitionRef) ModelBiasJobDefinitionRef() *ModelBiasJobDefinitionReference {
	var returns *ModelBiasJobDefinitionReference
	_jsii_.Get(
		j,
		"modelBiasJobDefinitionRef",
		&returns,
	)
	return returns
}

