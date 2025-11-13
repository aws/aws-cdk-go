package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ClientVpnEndpoint.
// Experimental.
type IClientVpnEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ClientVpnEndpoint resource.
	// Experimental.
	ClientVpnEndpointRef() *ClientVpnEndpointReference
}

// The jsii proxy for IClientVpnEndpointRef
type jsiiProxy_IClientVpnEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IClientVpnEndpointRef) ClientVpnEndpointRef() *ClientVpnEndpointReference {
	var returns *ClientVpnEndpointReference
	_jsii_.Get(
		j,
		"clientVpnEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IClientVpnEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

