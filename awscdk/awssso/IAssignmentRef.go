package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assignment.
// Experimental.
type IAssignmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssignmentRef
type jsiiProxy_IAssignmentRef struct {
	internal.Type__constructsIConstruct
}

