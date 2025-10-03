package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NamedQuery.
// Experimental.
type INamedQueryRef interface {
	constructs.IConstruct
}

// The jsii proxy for INamedQueryRef
type jsiiProxy_INamedQueryRef struct {
	internal.Type__constructsIConstruct
}

