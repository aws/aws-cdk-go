package awseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RegistryPolicy.
// Experimental.
type IRegistryPolicyRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RegistryPolicy resource.
	// Experimental.
	RegistryPolicyRef() *RegistryPolicyReference
}

// The jsii proxy for IRegistryPolicyRef
type jsiiProxy_IRegistryPolicyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRegistryPolicyRef) RegistryPolicyRef() *RegistryPolicyReference {
	var returns *RegistryPolicyReference
	_jsii_.Get(
		j,
		"registryPolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryPolicyRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryPolicyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

