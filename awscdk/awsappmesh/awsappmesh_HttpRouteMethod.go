package awsappmesh


// Supported values for matching routes based on the HTTP request method.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &httpRouteMatch{
//   			path: appmesh.httpRoutePathMatch.exactly(jsii.String("/exact")),
//   			method: appmesh.httpRouteMethod_POST,
//   			protocol: appmesh.httpRouteProtocol_HTTPS,
//   			headers: []headerMatch{
//   				appmesh.*headerMatch.valueIs(jsii.String("Content-Type"), jsii.String("application/json")),
//   				appmesh.*headerMatch.valueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
//   			},
//   			queryParameters: []queryParameterMatch{
//   				appmesh.*queryParameterMatch.valueIs(jsii.String("query-field"), jsii.String("value")),
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

