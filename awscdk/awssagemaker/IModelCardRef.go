package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelCard.
// Experimental.
type IModelCardRef interface {
	constructs.IConstruct
}

// The jsii proxy for IModelCardRef
type jsiiProxy_IModelCardRef struct {
	internal.Type__constructsIConstruct
}

