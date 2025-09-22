package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskSet.
// Experimental.
type ITaskSetRef interface {
	constructs.IConstruct
	// A reference to a TaskSet resource.
	// Experimental.
	TaskSetRef() *TaskSetReference
}

// The jsii proxy for ITaskSetRef
type jsiiProxy_ITaskSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITaskSetRef) TaskSetRef() *TaskSetReference {
	var returns *TaskSetReference
	_jsii_.Get(
		j,
		"taskSetRef",
		&returns,
	)
	return returns
}

