package awsappmesh


// Properties specific for a GRPC Based Routes.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &grpcRouteMatch{
//   			serviceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		timeout: &grpcTimeout{
//   			idle: cdk.duration.seconds(jsii.Number(2)),
//   			perRequest: cdk.*duration.seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
// Experimental.
type GrpcRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this Route.
	// Experimental.
	Match *GrpcRouteMatch `field:"required" json:"match" yaml:"match"`
	// List of targets that traffic is routed to when a request matches the route.
	// Experimental.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// The retry policy.
	// Experimental.
	RetryPolicy *GrpcRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a grpc timeout.
	// Experimental.
	Timeout *GrpcTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

