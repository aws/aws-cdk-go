package awskendraranking

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskendraranking/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ExecutionPlan.
// Experimental.
type IExecutionPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IExecutionPlanRef
type jsiiProxy_IExecutionPlanRef struct {
	internal.Type__constructsIConstruct
}

