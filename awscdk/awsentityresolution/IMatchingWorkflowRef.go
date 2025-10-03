package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchingWorkflow.
// Experimental.
type IMatchingWorkflowRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMatchingWorkflowRef
type jsiiProxy_IMatchingWorkflowRef struct {
	internal.Type__constructsIConstruct
}

