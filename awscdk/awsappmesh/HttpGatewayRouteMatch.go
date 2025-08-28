package awsappmesh


// The criterion for determining a request match for this GatewayRoute.
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
type HttpGatewayRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the gateway route to match.
	// Default: - do not match on headers.
	//
	Headers *[]HeaderMatch `field:"optional" json:"headers" yaml:"headers"`
	// The gateway route host name to be matched on.
	// Default: - do not match on host name.
	//
	Hostname GatewayRouteHostnameMatch `field:"optional" json:"hostname" yaml:"hostname"`
	// The method to match on.
	// Default: - do not match on method.
	//
	Method HttpRouteMethod `field:"optional" json:"method" yaml:"method"`
	// Specify how to match requests based on the 'path' part of their URL.
	// Default: - matches requests with any path.
	//
	Path HttpGatewayRoutePathMatch `field:"optional" json:"path" yaml:"path"`
	// The port number to match on.
	// Default: - no default port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	// Default: - do not match on query parameters.
	//
	QueryParameters *[]QueryParameterMatch `field:"optional" json:"queryParameters" yaml:"queryParameters"`
	// When `true`, rewrites the original request received at the Virtual Gateway to the destination Virtual Service name.
	//
	// When `false`, retains the original hostname from the request.
	// Default: true.
	//
	RewriteRequestHostname *bool `field:"optional" json:"rewriteRequestHostname" yaml:"rewriteRequestHostname"`
}

