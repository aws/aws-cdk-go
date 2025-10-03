package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Workflow.
// Experimental.
type IWorkflowRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWorkflowRef
type jsiiProxy_IWorkflowRef struct {
	internal.Type__constructsIConstruct
}

