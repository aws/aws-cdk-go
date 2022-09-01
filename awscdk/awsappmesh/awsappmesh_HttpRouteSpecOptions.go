package awsappmesh


// Properties specific for HTTP Based Routes.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
type HttpRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// The criterion for determining a request match for this Route.
	Match *HttpRouteMatch `field:"optional" json:"match" yaml:"match"`
	// The retry policy.
	RetryPolicy *HttpRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a http timeout.
	Timeout *HttpTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

