package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelQualityJobDefinition.
// Experimental.
type IModelQualityJobDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ModelQualityJobDefinition resource.
	// Experimental.
	ModelQualityJobDefinitionRef() *ModelQualityJobDefinitionReference
}

// The jsii proxy for IModelQualityJobDefinitionRef
type jsiiProxy_IModelQualityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelQualityJobDefinitionRef) ModelQualityJobDefinitionRef() *ModelQualityJobDefinitionReference {
	var returns *ModelQualityJobDefinitionReference
	_jsii_.Get(
		j,
		"modelQualityJobDefinitionRef",
		&returns,
	)
	return returns
}

