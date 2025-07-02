package awsappmesh


// Properties specific for HTTP Based GatewayRoutes.
//
// Example:
//   var gateway virtualGateway
//   var virtualService virtualService
//
//
//   gateway.addGatewayRoute(jsii.String("gateway-route-http-2"), &GatewayRouteBaseProps{
//   	RouteSpec: appmesh.GatewayRouteSpec_Http(&HttpGatewayRouteSpecOptions{
//   		RouteTarget: virtualService,
//   		Match: &HttpGatewayRouteMatch{
//   			// This rewrites the path from '/test' to '/rewrittenPath'.
//   			Path: appmesh.HttpGatewayRoutePathMatch_Exactly(jsii.String("/test"), jsii.String("/rewrittenPath")),
//   		},
//   	}),
//   })
//
type HttpGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	// Default: - no particular priority.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The VirtualService this GatewayRoute directs traffic to.
	RouteTarget IVirtualService `field:"required" json:"routeTarget" yaml:"routeTarget"`
	// The criterion for determining a request match for this GatewayRoute.
	//
	// When path match is defined, this may optionally determine the path rewrite configuration.
	// Default: - matches any path and automatically rewrites the path to '/'.
	//
	Match *HttpGatewayRouteMatch `field:"optional" json:"match" yaml:"match"`
}

