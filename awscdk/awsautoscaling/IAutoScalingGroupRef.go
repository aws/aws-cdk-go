package awsautoscaling

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingGroup.
// Experimental.
type IAutoScalingGroupRef interface {
	constructs.IConstruct
	// A reference to a AutoScalingGroup resource.
	// Experimental.
	AutoScalingGroupRef() *AutoScalingGroupReference
}

// The jsii proxy for IAutoScalingGroupRef
type jsiiProxy_IAutoScalingGroupRef struct {
	internal.Type__constructsIConstruct
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

