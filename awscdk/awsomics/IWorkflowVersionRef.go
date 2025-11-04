package awsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkflowVersion.
// Experimental.
type IWorkflowVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a WorkflowVersion resource.
	// Experimental.
	WorkflowVersionRef() *WorkflowVersionReference
}

// The jsii proxy for IWorkflowVersionRef
type jsiiProxy_IWorkflowVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IWorkflowVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflowVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

