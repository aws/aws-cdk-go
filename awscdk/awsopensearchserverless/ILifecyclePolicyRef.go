package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecyclePolicy.
// Experimental.
type ILifecyclePolicyRef interface {
	constructs.IConstruct
	// A reference to a LifecyclePolicy resource.
	// Experimental.
	LifecyclePolicyRef() *LifecyclePolicyReference
}

// The jsii proxy for ILifecyclePolicyRef
type jsiiProxy_ILifecyclePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ILifecyclePolicyRef) LifecyclePolicyRef() *LifecyclePolicyReference {
	var returns *LifecyclePolicyReference
	_jsii_.Get(
		j,
		"lifecyclePolicyRef",
		&returns,
	)
	return returns
}

