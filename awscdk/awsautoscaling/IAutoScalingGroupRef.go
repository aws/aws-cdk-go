package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingGroup.
// Experimental.
type IAutoScalingGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AutoScalingGroup resource.
	// Experimental.
	AutoScalingGroupRef() *AutoScalingGroupReference
}

// The jsii proxy for IAutoScalingGroupRef
type jsiiProxy_IAutoScalingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAutoScalingGroupRef) AutoScalingGroupRef() *AutoScalingGroupReference {
	var returns *AutoScalingGroupReference
	_jsii_.Get(
		j,
		"autoScalingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAutoScalingGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

