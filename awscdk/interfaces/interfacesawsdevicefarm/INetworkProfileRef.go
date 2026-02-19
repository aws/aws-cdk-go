package interfacesawsdevicefarm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevicefarm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkProfile.
// Experimental.
type INetworkProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkProfile resource.
	// Experimental.
	NetworkProfileRef() *NetworkProfileReference
}

// The jsii proxy for INetworkProfileRef
type jsiiProxy_INetworkProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_INetworkProfileRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INetworkProfileRef) NetworkProfileRef() *NetworkProfileReference {
	var returns *NetworkProfileReference
	_jsii_.Get(
		j,
		"networkProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

