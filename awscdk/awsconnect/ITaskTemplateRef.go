package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskTemplate.
// Experimental.
type ITaskTemplateRef interface {
	constructs.IConstruct
	// A reference to a TaskTemplate resource.
	// Experimental.
	TaskTemplateRef() *TaskTemplateReference
}

// The jsii proxy for ITaskTemplateRef
type jsiiProxy_ITaskTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITaskTemplateRef) TaskTemplateRef() *TaskTemplateReference {
	var returns *TaskTemplateReference
	_jsii_.Get(
		j,
		"taskTemplateRef",
		&returns,
	)
	return returns
}

