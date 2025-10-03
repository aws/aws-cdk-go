package awssam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SimpleTable.
// Experimental.
type ISimpleTableRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISimpleTableRef
type jsiiProxy_ISimpleTableRef struct {
	internal.Type__constructsIConstruct
}

