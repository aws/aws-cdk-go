package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CapacityProvider.
// Experimental.
type ICapacityProviderRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CapacityProvider resource.
	// Experimental.
	CapacityProviderRef() *CapacityProviderReference
}

// The jsii proxy for ICapacityProviderRef
type jsiiProxy_ICapacityProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICapacityProviderRef) CapacityProviderRef() *CapacityProviderReference {
	var returns *CapacityProviderReference
	_jsii_.Get(
		j,
		"capacityProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProviderRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapacityProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

