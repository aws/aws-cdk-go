package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Keyspace.
// Experimental.
type IKeyspaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for IKeyspaceRef
type jsiiProxy_IKeyspaceRef struct {
	internal.Type__constructsIConstruct
}

