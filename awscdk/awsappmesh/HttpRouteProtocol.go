package awsappmesh


// Supported :scheme options for HTTP2.
//
// Example:
//   var router VirtualRouter
//   var node VirtualNode
//
//
//   router.addRoute(jsii.String("route-http2"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
//   		WeightedTargets: []WeightedTarget{
//   			&WeightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &HttpRouteMatch{
//   			Path: appmesh.HttpRoutePathMatch_Exactly(jsii.String("/exact")),
//   			Method: appmesh.HttpRouteMethod_POST,
//   			Protocol: appmesh.HttpRouteProtocol_HTTPS,
//   			Headers: []HeaderMatch{
//   				appmesh.HeaderMatch_ValueIs(jsii.String("Content-Type"), jsii.String("application/json")),
//   				appmesh.HeaderMatch_ValueIsNot(jsii.String("Content-Type"), jsii.String("application/json")),
//   			},
//   			QueryParameters: []QueryParameterMatch{
//   				appmesh.QueryParameterMatch_ValueIs(jsii.String("query-field"), jsii.String("value")),
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

