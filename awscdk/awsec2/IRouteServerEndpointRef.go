package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerEndpoint.
// Experimental.
type IRouteServerEndpointRef interface {
	constructs.IConstruct
	// A reference to a RouteServerEndpoint resource.
	// Experimental.
	RouteServerEndpointRef() *RouteServerEndpointReference
}

// The jsii proxy for IRouteServerEndpointRef
type jsiiProxy_IRouteServerEndpointRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteServerEndpointRef) RouteServerEndpointRef() *RouteServerEndpointReference {
	var returns *RouteServerEndpointReference
	_jsii_.Get(
		j,
		"routeServerEndpointRef",
		&returns,
	)
	return returns
}

