package awsapplicationautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScalableTarget.
// Experimental.
type IScalableTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScalableTargetRef
type jsiiProxy_IScalableTargetRef struct {
	internal.Type__constructsIConstruct
}

