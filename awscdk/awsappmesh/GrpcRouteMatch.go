package awsappmesh


// The criterion for determining a request match for this Route.
//
// At least one match type must be selected.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-grpc-retry"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Grpc(&GrpcRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		Match: &GrpcRouteMatch{
//   			// When method name is specified, service name must be also specified.
//   			MethodName: jsii.String("methodname"),
//   			ServiceName: jsii.String("servicename"),
//   			Metadata: []headerMatch{
//   				appmesh.*headerMatch_ValueStartsWith(jsii.String("Content-Type"), jsii.String("application/")),
//   				appmesh.*headerMatch_ValueDoesNotStartWith(jsii.String("Content-Type"), jsii.String("text/")),
//   			},
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

