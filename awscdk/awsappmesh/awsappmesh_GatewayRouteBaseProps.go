package awsappmesh


// Basic configuration properties for a GatewayRoute.
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
//   			hostname: appmesh.gatewayRouteHostnameMatch.exactly(jsii.String("example.com")),
//   			// This disables the default rewrite to virtual service name and retain original request.
//   			rewriteRequestHostname: jsii.Boolean(false),
//   		},
//   	}),
//   })
//
// Experimental.
type GatewayRouteBaseProps struct {
	// What protocol the route uses.
	// Experimental.
	RouteSpec GatewayRouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the GatewayRoute.
	// Experimental.
	GatewayRouteName *string `field:"optional" json:"gatewayRouteName" yaml:"gatewayRouteName"`
}

