package awsautoscalingplans

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalingplans/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPlan.
// Experimental.
type IScalingPlanRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScalingPlanRef
type jsiiProxy_IScalingPlanRef struct {
	internal.Type__constructsIConstruct
}

