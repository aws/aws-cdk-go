package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayRoute.
// Experimental.
type IGatewayRouteRef interface {
	constructs.IConstruct
	// A reference to a GatewayRoute resource.
	// Experimental.
	GatewayRouteRef() *GatewayRouteReference
}

// The jsii proxy for IGatewayRouteRef
type jsiiProxy_IGatewayRouteRef struct {
	internal.Type__constructsIConstruct
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

