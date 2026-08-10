package interfacesawsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkflowExecution.
// Experimental.
type IWorkflowExecutionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkflowExecution resource.
	// Experimental.
	WorkflowExecutionRef() *WorkflowExecutionReference
}

// The jsii proxy for IWorkflowExecutionRef
type jsiiProxy_IWorkflowExecutionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkflowExecutionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkflowExecutionRef) WorkflowExecutionRef() *WorkflowExecutionReference {
	var returns *WorkflowExecutionReference
	_jsii_.Get(
		j,
		"workflowExecutionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflowExecutionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkflowExecutionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

