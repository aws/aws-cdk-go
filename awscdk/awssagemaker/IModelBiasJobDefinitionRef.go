package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelBiasJobDefinition.
// Experimental.
type IModelBiasJobDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelBiasJobDefinitionRef
type jsiiProxy_IModelBiasJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

