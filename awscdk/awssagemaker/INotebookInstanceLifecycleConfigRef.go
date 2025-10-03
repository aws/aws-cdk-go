package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotebookInstanceLifecycleConfig.
// Experimental.
type INotebookInstanceLifecycleConfigRef interface {
	constructs.IConstruct
}

// The jsii proxy for INotebookInstanceLifecycleConfigRef
type jsiiProxy_INotebookInstanceLifecycleConfigRef struct {
	internal.Type__constructsIConstruct
}

