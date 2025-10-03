package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CodeRepository.
// Experimental.
type ICodeRepositoryRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICodeRepositoryRef
type jsiiProxy_ICodeRepositoryRef struct {
	internal.Type__constructsIConstruct
}

