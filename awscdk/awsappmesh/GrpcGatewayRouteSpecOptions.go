package awsappmesh


// Properties specific for a gRPC GatewayRoute.
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
type GrpcGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	// Default: - no particular priority.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this GatewayRoute.
	Match *GrpcGatewayRouteMatch `field:"required" json:"match" yaml:"match"`
	// The VirtualService this GatewayRoute directs traffic to.
	RouteTarget IVirtualService `field:"required" json:"routeTarget" yaml:"routeTarget"`
}

