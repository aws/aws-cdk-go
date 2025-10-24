package awsappmesh


// The criterion for determining a request match for this Route.
//
// Example:
//   var router VirtualRouter
//   var node VirtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http(&HttpRouteSpecOptions{
//   		WeightedTargets: []WeightedTarget{
//   			&WeightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   			&WeightedTarget{
//   				VirtualNode: node,
//   				Weight: jsii.Number(50),
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_StartsWith(jsii.String("/path-to-app")),
//   		},
//   	}),
//   })
//
type HttpRouteMatch struct {
	// Specifies the client request headers to match on.
	//
	// All specified headers
	// must match for the route to match.
	// Default: - do not match on headers.
	//
	Headers *[]HeaderMatch `field:"optional" json:"headers" yaml:"headers"`
	// The HTTP client request method to match on.
	// Default: - do not match on request method.
	//
	Method HttpRouteMethod `field:"optional" json:"method" yaml:"method"`
	// Specifies how is the request matched based on the path part of its URL.
	// Default: - matches requests with all paths.
	//
	Path HttpRoutePathMatch `field:"optional" json:"path" yaml:"path"`
	// The port to match from the request.
	// Default: - do not match on port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The client request protocol to match on.
	//
	// Applicable only for HTTP2 routes.
	// Default: - do not match on HTTP2 request protocol.
	//
	Protocol HttpRouteProtocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The query parameters to match on.
	//
	// All specified query parameters must match for the route to match.
	// Default: - do not match on query parameters.
	//
	QueryParameters *[]QueryParameterMatch `field:"optional" json:"queryParameters" yaml:"queryParameters"`
}

