package awsomics

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsomics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkflowVersion.
// Experimental.
type IWorkflowVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkflowVersionRef
type jsiiProxy_IWorkflowVersionRef struct {
	internal.Type__constructsIConstruct
}

