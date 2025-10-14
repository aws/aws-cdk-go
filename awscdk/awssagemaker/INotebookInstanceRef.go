package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotebookInstance.
// Experimental.
type INotebookInstanceRef interface {
	constructs.IConstruct
	// A reference to a NotebookInstance resource.
	// Experimental.
	NotebookInstanceRef() *NotebookInstanceReference
}

// The jsii proxy for INotebookInstanceRef
type jsiiProxy_INotebookInstanceRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INotebookInstanceRef) NotebookInstanceRef() *NotebookInstanceReference {
	var returns *NotebookInstanceReference
	_jsii_.Get(
		j,
		"notebookInstanceRef",
		&returns,
	)
	return returns
}

