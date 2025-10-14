package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteServerPropagation.
// Experimental.
type IRouteServerPropagationRef interface {
	constructs.IConstruct
	// A reference to a RouteServerPropagation resource.
	// Experimental.
	RouteServerPropagationRef() *RouteServerPropagationReference
}

// The jsii proxy for IRouteServerPropagationRef
type jsiiProxy_IRouteServerPropagationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRouteServerPropagationRef) RouteServerPropagationRef() *RouteServerPropagationReference {
	var returns *RouteServerPropagationReference
	_jsii_.Get(
		j,
		"routeServerPropagationRef",
		&returns,
	)
	return returns
}

