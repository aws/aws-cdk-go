package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Used to generate host name matching methods.
//
// Example:
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &gatewayRouteBaseProps{
//   	routeSpec: appmesh.gatewayRouteSpec.grpc(&grpcGatewayRouteSpecOptions{
//   		routeTarget: virtualService,
//   		match: &grpcGatewayRouteMatch{
//   			hostname: appmesh.gatewayRouteHostnameMatch.endsWith(jsii.String(".example.com")),
//   		},
//   	}),
//   })
//
// Experimental.
type GatewayRouteHostnameMatch interface {
	// Returns the gateway route host name match configuration.
	// Experimental.
	Bind(scope awscdk.Construct) *GatewayRouteHostnameMatchConfig
}

// The jsii proxy struct for GatewayRouteHostnameMatch
type jsiiProxy_GatewayRouteHostnameMatch struct {
	_ byte // padding
}

// Experimental.
func NewGatewayRouteHostnameMatch_Override(g GatewayRouteHostnameMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		nil, // no parameters
		g,
	)
}

// The value of the host name with the given name must end with the specified characters.
// Experimental.
func GatewayRouteHostnameMatch_EndsWith(suffix *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		"endsWith",
		[]interface{}{suffix},
		&returns,
	)

	return returns
}

// The value of the host name must match the specified value exactly.
// Experimental.
func GatewayRouteHostnameMatch_Exactly(name *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.GatewayRouteHostnameMatch",
		"exactly",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayRouteHostnameMatch) Bind(scope awscdk.Construct) *GatewayRouteHostnameMatchConfig {
	var returns *GatewayRouteHostnameMatchConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

