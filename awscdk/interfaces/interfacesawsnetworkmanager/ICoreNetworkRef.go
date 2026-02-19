package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CoreNetwork.
// Experimental.
type ICoreNetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a CoreNetwork resource.
	// Experimental.
	CoreNetworkRef() *CoreNetworkReference
}

// The jsii proxy for ICoreNetworkRef
type jsiiProxy_ICoreNetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ICoreNetworkRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ICoreNetworkRef) CoreNetworkRef() *CoreNetworkReference {
	var returns *CoreNetworkReference
	_jsii_.Get(
		j,
		"coreNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICoreNetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICoreNetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

