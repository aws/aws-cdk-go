package awsentityresolution

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchingWorkflow.
// Experimental.
type IMatchingWorkflowRef interface {
	constructs.IConstruct
	// A reference to a MatchingWorkflow resource.
	// Experimental.
	MatchingWorkflowRef() *MatchingWorkflowReference
}

// The jsii proxy for IMatchingWorkflowRef
type jsiiProxy_IMatchingWorkflowRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMatchingWorkflowRef) MatchingWorkflowRef() *MatchingWorkflowReference {
	var returns *MatchingWorkflowReference
	_jsii_.Get(
		j,
		"matchingWorkflowRef",
		&returns,
	)
	return returns
}

