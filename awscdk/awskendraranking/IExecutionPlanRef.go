package awskendraranking

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskendraranking/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExecutionPlan.
// Experimental.
type IExecutionPlanRef interface {
	constructs.IConstruct
	// A reference to a ExecutionPlan resource.
	// Experimental.
	ExecutionPlanRef() *ExecutionPlanReference
}

// The jsii proxy for IExecutionPlanRef
type jsiiProxy_IExecutionPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IExecutionPlanRef) ExecutionPlanRef() *ExecutionPlanReference {
	var returns *ExecutionPlanReference
	_jsii_.Get(
		j,
		"executionPlanRef",
		&returns,
	)
	return returns
}

