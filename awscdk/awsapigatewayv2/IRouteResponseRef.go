package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouteResponse.
// Experimental.
type IRouteResponseRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RouteResponse resource.
	// Experimental.
	RouteResponseRef() *RouteResponseReference
}

// The jsii proxy for IRouteResponseRef
type jsiiProxy_IRouteResponseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IRouteResponseRef) RouteResponseRef() *RouteResponseReference {
	var returns *RouteResponseReference
	_jsii_.Get(
		j,
		"routeResponseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteResponseRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouteResponseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

