package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Task.
// Experimental.
type ITaskRef interface {
	constructs.IConstruct
	// A reference to a Task resource.
	// Experimental.
	TaskRef() *TaskReference
}

// The jsii proxy for ITaskRef
type jsiiProxy_ITaskRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITaskRef) TaskRef() *TaskReference {
	var returns *TaskReference
	_jsii_.Get(
		j,
		"taskRef",
		&returns,
	)
	return returns
}

