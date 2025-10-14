package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPolicy.
// Experimental.
type IScalingPolicyRef interface {
	constructs.IConstruct
	// A reference to a ScalingPolicy resource.
	// Experimental.
	ScalingPolicyRef() *ScalingPolicyReference
}

// The jsii proxy for IScalingPolicyRef
type jsiiProxy_IScalingPolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScalingPolicyRef) ScalingPolicyRef() *ScalingPolicyReference {
	var returns *ScalingPolicyReference
	_jsii_.Get(
		j,
		"scalingPolicyRef",
		&returns,
	)
	return returns
}

