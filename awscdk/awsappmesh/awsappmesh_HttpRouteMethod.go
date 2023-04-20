package awsappmesh


// Supported values for matching routes based on the HTTP request method.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_Exactly(jsii.String("/exact")),
//   			Method: appmesh.HttpRouteMethod_POST,
//   			Protocol: appmesh.HttpRouteProtocol_HTTPS,
//   			Headers: []headerMatch{
//   				appmesh.*headerMatch_ValueIs(jsii.String("Content-Type"), jsii.String("application/json")),
//   				appmesh.*headerMatch_ValueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
//   			},
//   			QueryParameters: []queryParameterMatch{
//   				appmesh.*queryParameterMatch_ValueIs(jsii.String("query-field"), jsii.String("value")),
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type HttpRouteMethod string

const (
	// GET request.
	// Experimental.
	HttpRouteMethod_GET HttpRouteMethod = "GET"
	// HEAD request.
	// Experimental.
	HttpRouteMethod_HEAD HttpRouteMethod = "HEAD"
	// POST request.
	// Experimental.
	HttpRouteMethod_POST HttpRouteMethod = "POST"
	// PUT request.
	// Experimental.
	HttpRouteMethod_PUT HttpRouteMethod = "PUT"
	// DELETE request.
	// Experimental.
	HttpRouteMethod_DELETE HttpRouteMethod = "DELETE"
	// CONNECT request.
	// Experimental.
	HttpRouteMethod_CONNECT HttpRouteMethod = "CONNECT"
	// OPTIONS request.
	// Experimental.
	HttpRouteMethod_OPTIONS HttpRouteMethod = "OPTIONS"
	// TRACE request.
	// Experimental.
	HttpRouteMethod_TRACE HttpRouteMethod = "TRACE"
	// PATCH request.
	// Experimental.
	HttpRouteMethod_PATCH HttpRouteMethod = "PATCH"
)

