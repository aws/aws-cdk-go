package awsappmesh


// Base interface properties for all Routes.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2-retry"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		RetryPolicy: &HttpRetryPolicy{
//   			// Retry if the connection failed
//   			TcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			// Retry if HTTP responds with a gateway error (502, 503, 504)
//   			HttpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry five times
//   			RetryAttempts: jsii.Number(5),
//   			// Use a 1 second timeout per retry
//   			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type RouteBaseProps struct {
	// Protocol specific spec.
	RouteSpec RouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the route.
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
}

