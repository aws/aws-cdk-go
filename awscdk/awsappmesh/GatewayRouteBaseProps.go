package awsappmesh


// Basic configuration properties for a GatewayRoute.
//
// Example:
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
type GatewayRouteBaseProps struct {
	// What protocol the route uses.
	RouteSpec GatewayRouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the GatewayRoute.
	// Default: - an automatically generated name.
	//
	GatewayRouteName *string `field:"optional" json:"gatewayRouteName" yaml:"gatewayRouteName"`
}

