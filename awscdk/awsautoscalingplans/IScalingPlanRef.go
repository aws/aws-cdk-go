package awsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscalingplans/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPlan.
// Experimental.
type IScalingPlanRef interface {
	constructs.IConstruct
	// A reference to a ScalingPlan resource.
	// Experimental.
	ScalingPlanRef() *ScalingPlanReference
}

// The jsii proxy for IScalingPlanRef
type jsiiProxy_IScalingPlanRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScalingPlanRef) ScalingPlanRef() *ScalingPlanReference {
	var returns *ScalingPlanReference
	_jsii_.Get(
		j,
		"scalingPlanRef",
		&returns,
	)
	return returns
}

