package awsappmesh


// The criterion for determining a request match for this Route.
//
// At least one match type must be selected.
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
type GrpcRouteMatch struct {
	// Create metadata based gRPC route match.
	//
	// All specified metadata must match for the route to match.
	Metadata *[]HeaderMatch `field:"optional" json:"metadata" yaml:"metadata"`
	// The method name to match from the request.
	//
	// If the method name is specified, service name must be also provided.
	MethodName *string `field:"optional" json:"methodName" yaml:"methodName"`
	// Create service name based gRPC route match.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
}

