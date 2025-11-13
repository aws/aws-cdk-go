package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TransitGatewayConnectPeer.
// Experimental.
type ITransitGatewayConnectPeerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TransitGatewayConnectPeer resource.
	// Experimental.
	TransitGatewayConnectPeerRef() *TransitGatewayConnectPeerReference
}

// The jsii proxy for ITransitGatewayConnectPeerRef
type jsiiProxy_ITransitGatewayConnectPeerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ITransitGatewayConnectPeerRef) TransitGatewayConnectPeerRef() *TransitGatewayConnectPeerReference {
	var returns *TransitGatewayConnectPeerReference
	_jsii_.Get(
		j,
		"transitGatewayConnectPeerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayConnectPeerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITransitGatewayConnectPeerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

