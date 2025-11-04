package awssynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Canary.
// Experimental.
type ICanaryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Canary resource.
	// Experimental.
	CanaryRef() *CanaryReference
}

// The jsii proxy for ICanaryRef
type jsiiProxy_ICanaryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICanaryRef) CanaryRef() *CanaryReference {
	var returns *CanaryReference
	_jsii_.Get(
		j,
		"canaryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanaryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICanaryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

