package interfacesawsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayRoute.
// Experimental.
type IGatewayRouteRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GatewayRoute resource.
	// Experimental.
	GatewayRouteRef() *GatewayRouteReference
}

// The jsii proxy for IGatewayRouteRef
type jsiiProxy_IGatewayRouteRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IGatewayRouteRef) GatewayRouteRef() *GatewayRouteReference {
	var returns *GatewayRouteReference
	_jsii_.Get(
		j,
		"gatewayRouteRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRouteRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGatewayRouteRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

