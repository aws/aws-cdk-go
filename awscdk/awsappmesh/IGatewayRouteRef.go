package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappmesh/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayRoute.
// Experimental.
type IGatewayRouteRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGatewayRouteRef
type jsiiProxy_IGatewayRouteRef struct {
	internal.Type__constructsIConstruct
}

