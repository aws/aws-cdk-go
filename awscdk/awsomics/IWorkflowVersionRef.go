package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkflowVersion.
// Experimental.
type IWorkflowVersionRef interface {
	constructs.IConstruct
	// A reference to a WorkflowVersion resource.
	// Experimental.
	WorkflowVersionRef() *WorkflowVersionReference
}

// The jsii proxy for IWorkflowVersionRef
type jsiiProxy_IWorkflowVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IWorkflowVersionRef) WorkflowVersionRef() *WorkflowVersionReference {
	var returns *WorkflowVersionReference
	_jsii_.Get(
		j,
		"workflowVersionRef",
		&returns,
	)
	return returns
}

