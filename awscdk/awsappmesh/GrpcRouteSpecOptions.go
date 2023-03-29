package awsappmesh


// Properties specific for a GRPC Based Routes.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &GrpcRouteMatch{
//   			ServiceName: jsii.String("my-service.default.svc.cluster.local"),
//   		},
//   		Timeout: &GrpcTimeout{
//   			Idle: cdk.Duration_Seconds(jsii.Number(2)),
//   			PerRequest: cdk.Duration_*Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type GrpcRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this Route.
	Match *GrpcRouteMatch `field:"required" json:"match" yaml:"match"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// The retry policy.
	RetryPolicy *GrpcRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a grpc timeout.
	Timeout *GrpcTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

