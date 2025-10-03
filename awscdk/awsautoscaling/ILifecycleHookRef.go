package awsautoscaling

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecycleHook.
// Experimental.
type ILifecycleHookRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILifecycleHookRef
type jsiiProxy_ILifecycleHookRef struct {
	internal.Type__constructsIConstruct
}

