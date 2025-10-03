package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelQualityJobDefinition.
// Experimental.
type IModelQualityJobDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelQualityJobDefinitionRef
type jsiiProxy_IModelQualityJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

