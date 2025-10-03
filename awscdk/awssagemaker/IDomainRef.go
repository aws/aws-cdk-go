package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Domain.
// Experimental.
type IDomainRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDomainRef
type jsiiProxy_IDomainRef struct {
	internal.Type__constructsIConstruct
}

