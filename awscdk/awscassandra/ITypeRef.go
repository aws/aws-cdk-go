package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Type.
// Experimental.
type ITypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITypeRef
type jsiiProxy_ITypeRef struct {
	internal.Type__constructsIConstruct
}

