package awssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationAssignment.
// Experimental.
type IApplicationAssignmentRef interface {
	constructs.IConstruct
	// A reference to a ApplicationAssignment resource.
	// Experimental.
	ApplicationAssignmentRef() *ApplicationAssignmentReference
}

// The jsii proxy for IApplicationAssignmentRef
type jsiiProxy_IApplicationAssignmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApplicationAssignmentRef) ApplicationAssignmentRef() *ApplicationAssignmentReference {
	var returns *ApplicationAssignmentReference
	_jsii_.Get(
		j,
		"applicationAssignmentRef",
		&returns,
	)
	return returns
}

