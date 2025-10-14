package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Route.
// Experimental.
type IRouteRef interface {
	constructs.IConstruct
	// A reference to a Route resource.
	// Experimental.
	RouteRef() *RouteReference
}

// The jsii proxy for IRouteRef
type jsiiProxy_IRouteRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteRef) RouteRef() *RouteReference {
	var returns *RouteReference
	_jsii_.Get(
		j,
		"routeRef",
		&returns,
	)
	return returns
}

