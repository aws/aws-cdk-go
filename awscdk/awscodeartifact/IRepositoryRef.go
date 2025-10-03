package awscodeartifact

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeartifact/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Repository.
// Experimental.
type IRepositoryRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRepositoryRef
type jsiiProxy_IRepositoryRef struct {
	internal.Type__constructsIConstruct
}

