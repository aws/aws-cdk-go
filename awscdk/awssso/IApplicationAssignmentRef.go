package awssso

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationAssignment.
// Experimental.
type IApplicationAssignmentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IApplicationAssignmentRef
type jsiiProxy_IApplicationAssignmentRef struct {
	internal.Type__constructsIConstruct
}

