package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotebookInstanceLifecycleConfig.
// Experimental.
type INotebookInstanceLifecycleConfigRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NotebookInstanceLifecycleConfig resource.
	// Experimental.
	NotebookInstanceLifecycleConfigRef() *NotebookInstanceLifecycleConfigReference
}

// The jsii proxy for INotebookInstanceLifecycleConfigRef
type jsiiProxy_INotebookInstanceLifecycleConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INotebookInstanceLifecycleConfigRef) NotebookInstanceLifecycleConfigRef() *NotebookInstanceLifecycleConfigReference {
	var returns *NotebookInstanceLifecycleConfigReference
	_jsii_.Get(
		j,
		"notebookInstanceLifecycleConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotebookInstanceLifecycleConfigRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotebookInstanceLifecycleConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

