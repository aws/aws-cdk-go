package awsappmesh


// The criterion for determining a request match for this GatewayRoute.
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
type HttpGatewayRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the gateway route to match.
	Headers *[]HeaderMatch `field:"optional" json:"headers" yaml:"headers"`
	// The gateway route host name to be matched on.
	Hostname GatewayRouteHostnameMatch `field:"optional" json:"hostname" yaml:"hostname"`
	// The method to match on.
	Method HttpRouteMethod `field:"optional" json:"method" yaml:"method"`
	// Specify how to match requests based on the 'path' part of their URL.
	Path HttpGatewayRoutePathMatch `field:"optional" json:"path" yaml:"path"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	QueryParameters *[]QueryParameterMatch `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	RewriteRequestHostname *bool `field:"optional" json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
}

