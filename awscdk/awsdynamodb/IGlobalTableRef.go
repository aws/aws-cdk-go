package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalTable.
// Experimental.
type IGlobalTableRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGlobalTableRef
type jsiiProxy_IGlobalTableRef struct {
	internal.Type__constructsIConstruct
}

