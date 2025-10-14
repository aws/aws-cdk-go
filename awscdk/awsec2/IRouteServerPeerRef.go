package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerPeer.
// Experimental.
type IRouteServerPeerRef interface {
	constructs.IConstruct
	// A reference to a RouteServerPeer resource.
	// Experimental.
	RouteServerPeerRef() *RouteServerPeerReference
}

// The jsii proxy for IRouteServerPeerRef
type jsiiProxy_IRouteServerPeerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteServerPeerRef) RouteServerPeerRef() *RouteServerPeerReference {
	var returns *RouteServerPeerReference
	_jsii_.Get(
		j,
		"routeServerPeerRef",
		&returns,
	)
	return returns
}

