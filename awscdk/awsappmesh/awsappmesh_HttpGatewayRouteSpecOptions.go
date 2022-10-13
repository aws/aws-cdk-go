package awsappmesh


// Properties specific for HTTP Based GatewayRoutes.
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
type HttpGatewayRouteSpecOptions struct {
	// The priority for the gateway route.
	//
	// When a Virtual Gateway has multiple gateway routes, gateway route match
	// is performed in the order of specified value, where 0 is the highest priority,
	// and first matched gateway route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The VirtualService this GatewayRoute directs traffic to.
	RouteTarget IVirtualService `field:"required" json:"routeTarget" yaml:"routeTarget"`
	// The criterion for determining a request match for this GatewayRoute.
	//
	// When path match is defined, this may optionally determine the path rewrite configuration.
	Match *HttpGatewayRouteMatch `field:"optional" json:"match" yaml:"match"`
}

