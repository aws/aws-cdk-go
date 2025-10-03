package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AutoScalingGroup.
// Experimental.
type IAutoScalingGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAutoScalingGroupRef
type jsiiProxy_IAutoScalingGroupRef struct {
	internal.Type__constructsIConstruct
}

