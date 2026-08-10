package interfacesawssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationProvider.
// Experimental.
type IApplicationProviderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ApplicationProvider resource.
	// Experimental.
	ApplicationProviderRef() *ApplicationProviderReference
}

// The jsii proxy for IApplicationProviderRef
type jsiiProxy_IApplicationProviderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IApplicationProviderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IApplicationProviderRef) ApplicationProviderRef() *ApplicationProviderReference {
	var returns *ApplicationProviderReference
	_jsii_.Get(
		j,
		"applicationProviderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationProviderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationProviderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

