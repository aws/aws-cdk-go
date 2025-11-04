package awsapplicationautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPolicy.
// Experimental.
type IScalingPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ScalingPolicy resource.
	// Experimental.
	ScalingPolicyRef() *ScalingPolicyReference
}

// The jsii proxy for IScalingPolicyRef
type jsiiProxy_IScalingPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IScalingPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScalingPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

