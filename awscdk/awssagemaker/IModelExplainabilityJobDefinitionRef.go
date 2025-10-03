package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelExplainabilityJobDefinition.
// Experimental.
type IModelExplainabilityJobDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelExplainabilityJobDefinitionRef
type jsiiProxy_IModelExplainabilityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

