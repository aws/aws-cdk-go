package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotebookInstanceLifecycleConfig.
// Experimental.
type INotebookInstanceLifecycleConfigRef interface {
	constructs.IConstruct
	// A reference to a NotebookInstanceLifecycleConfig resource.
	// Experimental.
	NotebookInstanceLifecycleConfigRef() *NotebookInstanceLifecycleConfigReference
}

// The jsii proxy for INotebookInstanceLifecycleConfigRef
type jsiiProxy_INotebookInstanceLifecycleConfigRef struct {
	internal.Type__constructsIConstruct
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

