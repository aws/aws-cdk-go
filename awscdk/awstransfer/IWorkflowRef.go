package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workflow.
// Experimental.
type IWorkflowRef interface {
	constructs.IConstruct
	// A reference to a Workflow resource.
	// Experimental.
	WorkflowRef() *WorkflowReference
}

// The jsii proxy for IWorkflowRef
type jsiiProxy_IWorkflowRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWorkflowRef) WorkflowRef() *WorkflowReference {
	var returns *WorkflowReference
	_jsii_.Get(
		j,
		"workflowRef",
		&returns,
	)
	return returns
}

