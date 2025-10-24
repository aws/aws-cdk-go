package awsappmesh


// Properties specific for a GRPC Based Routes.
//
// Example:
//   var router VirtualRouter
//   var node VirtualNode
//
//
//   router.addRoute(jsii.String("route-grpc-retry"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
//   		WeightedTargets: []WeightedTarget{
//   			&WeightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &GrpcRouteMatch{
//   			// When method name is specified, service name must be also specified.
//   			MethodName: jsii.String("methodname"),
//   			ServiceName: jsii.String("servicename"),
//   			Metadata: []HeaderMatch{
//   				appmesh.HeaderMatch_ValueStartsWith(jsii.String("Content-Type"), jsii.String("application/")),
//   				appmesh.HeaderMatch_ValueDoesNotStartWith(jsii.String("Content-Type"), jsii.String("text/")),
//   			},
//   		},
//   	}),
//   })
//
type GrpcRouteSpecOptions struct {
	// The priority for the route.
	//
	// When a Virtual Router has multiple routes, route match is performed in the
	// order of specified value, where 0 is the highest priority, and first matched route is selected.
	// Default: - no particular priority.
	//
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// The criterion for determining a request match for this Route.
	Match *GrpcRouteMatch `field:"required" json:"match" yaml:"match"`
	// List of targets that traffic is routed to when a request matches the route.
	WeightedTargets *[]*WeightedTarget `field:"required" json:"weightedTargets" yaml:"weightedTargets"`
	// The retry policy.
	// Default: - no retry policy.
	//
	RetryPolicy *GrpcRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// An object that represents a grpc timeout.
	// Default: - None.
	//
	Timeout *GrpcTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}

