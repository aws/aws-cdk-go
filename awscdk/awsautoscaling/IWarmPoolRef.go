package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WarmPool.
// Experimental.
type IWarmPoolRef interface {
	constructs.IConstruct
}

// The jsii proxy for IWarmPoolRef
type jsiiProxy_IWarmPoolRef struct {
	internal.Type__constructsIConstruct
}

