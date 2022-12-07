package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Defines HTTP gateway route matching based on the URL path of the request.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type HttpGatewayRoutePathMatch interface {
	// Returns the gateway route path match configuration.
	Bind(scope constructs.Construct) *HttpGatewayRoutePathMatchConfig
}

// The jsii proxy struct for HttpGatewayRoutePathMatch
type jsiiProxy_HttpGatewayRoutePathMatch struct {
	_ byte // padding
}

func NewHttpGatewayRoutePathMatch_Override(h HttpGatewayRoutePathMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		nil, // no parameters
		h,
	)
}

// The value of the path must match the specified value exactly.
//
// The provided `path` must start with the '/' character.
func HttpGatewayRoutePathMatch_Exactly(path *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpGatewayRoutePathMatch_ExactlyParameters(path); err != nil {
		panic(err)
	}
	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"exactly",
		[]interface{}{path, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified regex.
func HttpGatewayRoutePathMatch_Regex(regex *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpGatewayRoutePathMatch_RegexParameters(regex); err != nil {
		panic(err)
	}
	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"regex",
		[]interface{}{regex, rewriteTo},
		&returns,
	)

	return returns
}

// The value of the path must match the specified prefix.
func HttpGatewayRoutePathMatch_StartsWith(prefix *string, rewriteTo *string) HttpGatewayRoutePathMatch {
	_init_.Initialize()

	if err := validateHttpGatewayRoutePathMatch_StartsWithParameters(prefix); err != nil {
		panic(err)
	}
	var returns HttpGatewayRoutePathMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.HttpGatewayRoutePathMatch",
		"startsWith",
		[]interface{}{prefix, rewriteTo},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HttpGatewayRoutePathMatch) Bind(scope constructs.Construct) *HttpGatewayRoutePathMatchConfig {
	if err := h.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *HttpGatewayRoutePathMatchConfig

	_jsii_.Invoke(
		h,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

