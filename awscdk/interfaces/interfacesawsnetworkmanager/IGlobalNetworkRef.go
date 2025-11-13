package interfacesawsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnetworkmanager/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalNetwork.
// Experimental.
type IGlobalNetworkRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GlobalNetwork resource.
	// Experimental.
	GlobalNetworkRef() *GlobalNetworkReference
}

// The jsii proxy for IGlobalNetworkRef
type jsiiProxy_IGlobalNetworkRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGlobalNetworkRef) GlobalNetworkRef() *GlobalNetworkReference {
	var returns *GlobalNetworkReference
	_jsii_.Get(
		j,
		"globalNetworkRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalNetworkRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalNetworkRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

