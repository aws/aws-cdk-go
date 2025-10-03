package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalingPolicy.
// Experimental.
type IScalingPolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScalingPolicyRef
type jsiiProxy_IScalingPolicyRef struct {
	internal.Type__constructsIConstruct
}

