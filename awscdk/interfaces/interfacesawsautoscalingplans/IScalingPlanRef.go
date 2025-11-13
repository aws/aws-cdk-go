package interfacesawsautoscalingplans

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsautoscalingplans/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPlan.
// Experimental.
type IScalingPlanRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ScalingPlan resource.
	// Experimental.
	ScalingPlanRef() *ScalingPlanReference
}

// The jsii proxy for IScalingPlanRef
type jsiiProxy_IScalingPlanRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IScalingPlanRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScalingPlanRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

