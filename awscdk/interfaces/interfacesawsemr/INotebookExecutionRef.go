package interfacesawsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsemr/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotebookExecution.
// Experimental.
type INotebookExecutionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NotebookExecution resource.
	// Experimental.
	NotebookExecutionRef() *NotebookExecutionReference
}

// The jsii proxy for INotebookExecutionRef
type jsiiProxy_INotebookExecutionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INotebookExecutionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INotebookExecutionRef) NotebookExecutionRef() *NotebookExecutionReference {
	var returns *NotebookExecutionReference
	_jsii_.Get(
		j,
		"notebookExecutionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotebookExecutionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotebookExecutionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

