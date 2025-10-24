package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate host name matching methods.
//
// Example:
//   var gateway VirtualGateway
//   var virtualService VirtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &GatewayRouteBaseProps{
//   	RouteSpec: appmesh.GatewayRouteSpec_Grpc(&GrpcGatewayRouteSpecOptions{
//   		RouteTarget: virtualService,
//   		Match: &GrpcGatewayRouteMatch{
//   			Hostname: appmesh.GatewayRouteHostnameMatch_EndsWith(jsii.String(".example.com")),
//   		},
//   	}),
//   })
//
type GatewayRouteHostnameMatch interface {
	// Returns the gateway route host name match configuration.
	Bind(scope constructs.Construct) *GatewayRouteHostnameMatchConfig
}

// The jsii proxy struct for GatewayRouteHostnameMatch
type jsiiProxy_GatewayRouteHostnameMatch struct {
	_ byte // padding
}

func NewGatewayRouteHostnameMatch_Override(g GatewayRouteHostnameMatch) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		nil, // no parameters
		g,
	)
}

// The value of the host name with the given name must end with the specified characters.
func GatewayRouteHostnameMatch_EndsWith(suffix *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	if err := validateGatewayRouteHostnameMatch_EndsWithParameters(suffix); err != nil {
		panic(err)
	}
	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		"endsWith",
		[]interface{}{suffix},
		&returns,
	)

	return returns
}

// The value of the host name must match the specified value exactly.
func GatewayRouteHostnameMatch_Exactly(name *string) GatewayRouteHostnameMatch {
	_init_.Initialize()

	if err := validateGatewayRouteHostnameMatch_ExactlyParameters(name); err != nil {
		panic(err)
	}
	var returns GatewayRouteHostnameMatch

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteHostnameMatch",
		"exactly",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayRouteHostnameMatch) Bind(scope constructs.Construct) *GatewayRouteHostnameMatchConfig {
	if err := g.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *GatewayRouteHostnameMatchConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

