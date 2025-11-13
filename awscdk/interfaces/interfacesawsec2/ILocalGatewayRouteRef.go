package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LocalGatewayRoute.
// Experimental.
type ILocalGatewayRouteRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LocalGatewayRoute resource.
	// Experimental.
	LocalGatewayRouteRef() *LocalGatewayRouteReference
}

// The jsii proxy for ILocalGatewayRouteRef
type jsiiProxy_ILocalGatewayRouteRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILocalGatewayRouteRef) LocalGatewayRouteRef() *LocalGatewayRouteReference {
	var returns *LocalGatewayRouteReference
	_jsii_.Get(
		j,
		"localGatewayRouteRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILocalGatewayRouteRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

