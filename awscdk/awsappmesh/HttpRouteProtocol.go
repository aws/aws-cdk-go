package awsappmesh


// Supported :scheme options for HTTP2.
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
type HttpRouteProtocol string

const (
	// Match HTTP requests.
	HttpRouteProtocol_HTTP HttpRouteProtocol = "HTTP"
	// Match HTTPS requests.
	HttpRouteProtocol_HTTPS HttpRouteProtocol = "HTTPS"
)

