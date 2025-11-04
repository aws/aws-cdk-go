package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LifecyclePolicy.
// Experimental.
type ILifecyclePolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a LifecyclePolicy resource.
	// Experimental.
	LifecyclePolicyRef() *LifecyclePolicyReference
}

// The jsii proxy for ILifecyclePolicyRef
type jsiiProxy_ILifecyclePolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ILifecyclePolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILifecyclePolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

