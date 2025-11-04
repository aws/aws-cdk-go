package awseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Registry.
// Experimental.
type IRegistryRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Registry resource.
	// Experimental.
	RegistryRef() *RegistryReference
}

// The jsii proxy for IRegistryRef
type jsiiProxy_IRegistryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRegistryRef) RegistryRef() *RegistryReference {
	var returns *RegistryReference
	_jsii_.Get(
		j,
		"registryRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRegistryRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

