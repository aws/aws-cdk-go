package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerEndpoint.
// Experimental.
type IRouteServerEndpointRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRouteServerEndpointRef
type jsiiProxy_IRouteServerEndpointRef struct {
	internal.Type__constructsIConstruct
}

