package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Used to generate specs with different protocols for a GatewayRoute.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-grpc"), &GatewayRouteBaseProps{
//   	RouteSpec: appmesh.GatewayRouteSpec_Grpc(&GrpcGatewayRouteSpecOptions{
//   		RouteTarget: virtualService,
//   		Match: &GrpcGatewayRouteMatch{
//   			Hostname: appmesh.GatewayRouteHostnameMatch_Exactly(jsii.String("example.com")),
//   			// This disables the default rewrite to virtual service name and retain original request.
//   			RewriteRequestHostname: jsii.Boolean(false),
//   		},
//   	}),
//   })
//
type GatewayRouteSpec interface {
	// Called when the GatewayRouteSpec type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	Bind(scope constructs.Construct) *GatewayRouteSpecConfig
}

// The jsii proxy struct for GatewayRouteSpec
type jsiiProxy_GatewayRouteSpec struct {
	_ byte // padding
}

func NewGatewayRouteSpec_Override(g GatewayRouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		nil, // no parameters
		g,
	)
}

// Creates an gRPC Based GatewayRoute.
func GatewayRouteSpec_Grpc(options *GrpcGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	if err := validateGatewayRouteSpec_GrpcParameters(options); err != nil {
		panic(err)
	}
	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based GatewayRoute.
func GatewayRouteSpec_Http(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	if err := validateGatewayRouteSpec_HttpParameters(options); err != nil {
		panic(err)
	}
	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based GatewayRoute.
func GatewayRouteSpec_Http2(options *HttpGatewayRouteSpecOptions) GatewayRouteSpec {
	_init_.Initialize()

	if err := validateGatewayRouteSpec_Http2Parameters(options); err != nil {
		panic(err)
	}
	var returns GatewayRouteSpec

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_appmesh.GatewayRouteSpec",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GatewayRouteSpec) Bind(scope constructs.Construct) *GatewayRouteSpecConfig {
	if err := g.validateBindParameters(scope); err != nil {
		panic(err)
	}
	var returns *GatewayRouteSpecConfig

	_jsii_.Invoke(
		g,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

