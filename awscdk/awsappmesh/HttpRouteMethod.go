package awsappmesh


// Supported values for matching routes based on the HTTP request method.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type HttpRouteMethod string

const (
	// GET request.
	HttpRouteMethod_GET HttpRouteMethod = "GET"
	// HEAD request.
	HttpRouteMethod_HEAD HttpRouteMethod = "HEAD"
	// POST request.
	HttpRouteMethod_POST HttpRouteMethod = "POST"
	// PUT request.
	HttpRouteMethod_PUT HttpRouteMethod = "PUT"
	// DELETE request.
	HttpRouteMethod_DELETE HttpRouteMethod = "DELETE"
	// CONNECT request.
	HttpRouteMethod_CONNECT HttpRouteMethod = "CONNECT"
	// OPTIONS request.
	HttpRouteMethod_OPTIONS HttpRouteMethod = "OPTIONS"
	// TRACE request.
	HttpRouteMethod_TRACE HttpRouteMethod = "TRACE"
	// PATCH request.
	HttpRouteMethod_PATCH HttpRouteMethod = "PATCH"
)

