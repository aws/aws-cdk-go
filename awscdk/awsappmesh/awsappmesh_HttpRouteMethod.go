package awsappmesh


// Supported values for matching routes based on the HTTP request method.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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

