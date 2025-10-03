package awscassandra

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscassandra/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Table.
// Experimental.
type ITableRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITableRef
type jsiiProxy_ITableRef struct {
	internal.Type__constructsIConstruct
}

