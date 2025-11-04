package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workflow.
// Experimental.
type IWorkflowRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Workflow resource.
	// Experimental.
	WorkflowRef() *WorkflowReference
}

// The jsii proxy for IWorkflowRef
type jsiiProxy_IWorkflowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IWorkflowRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

