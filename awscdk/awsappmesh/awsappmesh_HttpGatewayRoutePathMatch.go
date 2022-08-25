package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Defines HTTP gateway route matching based on the URL path of the request.
//
// Example:
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-http-2"), &gatewayRouteBaseProps{
//   	routeSpec: appmesh.gatewayRouteSpec.http(&httpGatewayRouteSpecOptions{
//   		routeTarget: virtualService,
//   		match: &httpGatewayRouteMatch{
//   			// This rewrites the path from '/test' to '/rewrittenPath'.
//   			path: appmesh.httpGatewayRoutePathMatch.exactly(jsii.String("/test"), jsii.String("/rewrittenPath")),
//   		},
//   	}),
//   })
//
// Experimental.
type HttpGatewayRoutePathMatch interface {
	// Returns the gateway route path match configuration.
	// Experimental.
	Bind(scope awscdk.Construct) *HttpGatewayRoutePathMatchConfig
}

// The jsii proxy struct for HttpGatewayRoutePathMatch
type jsiiProxy_HttpGatewayRoutePathMatch struct {
	_ byte // padding
}

// Experimental.
func NewHttpGatewayRoutePathMatch_Override(h HttpGatewayRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
// Experimental.
func HttpGatewayRoutePathMatch_Exactly(path *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"exactly",
		[]interface{}{path, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
// Experimental.
func HttpGatewayRoutePathMatch_Regex(regex *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"regex",
		[]interface{}{regex, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
// Experimental.
func HttpGatewayRoutePathMatch_StartsWith(prefix *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.HttpGatewayRoutePathMatch",
		"startsWith",
		[]interface{}{prefix, rewriteTo},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpGatewayRoutePathMatch) Bind(scope awscdk.Construct) *HttpGatewayRoutePathMatchConfig {
	var returns *HttpGatewayRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

