package awsappmesh


// Base interface properties for all Routes.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2-retry"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		retryPolicy: &httpRetryPolicy{
//   			// Retry if the connection failed
//   			tcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.tcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			// Retry if HTTP responds with a gateway error (502, 503, 504)
//   			httpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry five times
//   			retryAttempts: jsii.Number(5),
//   			// Use a 1 second timeout per retry
//   			retryTimeout: cdk.duration.seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
// Experimental.
type RouteBaseProps struct {
	// Protocol specific spec.
	// Experimental.
	RouteSpec RouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the route.
	// Experimental.
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
}

