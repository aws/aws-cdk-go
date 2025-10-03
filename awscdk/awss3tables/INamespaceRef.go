package awss3tables

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3tables/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Namespace.
// Experimental.
type INamespaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for INamespaceRef
type jsiiProxy_INamespaceRef struct {
	internal.Type__constructsIConstruct
}

