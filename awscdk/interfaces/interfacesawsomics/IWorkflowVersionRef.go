package interfacesawsomics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkflowVersion.
// Experimental.
type IWorkflowVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkflowVersion resource.
	// Experimental.
	WorkflowVersionRef() *WorkflowVersionReference
}

// The jsii proxy for IWorkflowVersionRef
type jsiiProxy_IWorkflowVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkflowVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IWorkflowVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

