package awssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assignment.
// Experimental.
type IAssignmentRef interface {
	constructs.IConstruct
	// A reference to a Assignment resource.
	// Experimental.
	AssignmentRef() *AssignmentReference
}

// The jsii proxy for IAssignmentRef
type jsiiProxy_IAssignmentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssignmentRef) AssignmentRef() *AssignmentReference {
	var returns *AssignmentReference
	_jsii_.Get(
		j,
		"assignmentRef",
		&returns,
	)
	return returns
}

