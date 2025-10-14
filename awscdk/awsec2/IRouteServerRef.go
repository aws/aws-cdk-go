package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServer.
// Experimental.
type IRouteServerRef interface {
	constructs.IConstruct
	// A reference to a RouteServer resource.
	// Experimental.
	RouteServerRef() *RouteServerReference
}

// The jsii proxy for IRouteServerRef
type jsiiProxy_IRouteServerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteServerRef) RouteServerRef() *RouteServerReference {
	var returns *RouteServerReference
	_jsii_.Get(
		j,
		"routeServerRef",
		&returns,
	)
	return returns
}

