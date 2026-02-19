package interfacesawseventschemas

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseventschemas/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Registry.
// Experimental.
type IRegistryRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Registry resource.
	// Experimental.
	RegistryRef() *RegistryReference
}

// The jsii proxy for IRegistryRef
type jsiiProxy_IRegistryRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRegistryRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IRegistryRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

