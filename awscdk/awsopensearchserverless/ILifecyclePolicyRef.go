package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecyclePolicy.
// Experimental.
type ILifecyclePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILifecyclePolicyRef
type jsiiProxy_ILifecyclePolicyRef struct {
	internal.Type__constructsIConstruct
}

