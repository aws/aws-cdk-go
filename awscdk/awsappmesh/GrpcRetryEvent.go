package awsappmesh


// gRPC events.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
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
//   			ServiceName: jsii.String("servicename"),
//   		},
//   		RetryPolicy: &GrpcRetryPolicy{
//   			TcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			HttpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry if gRPC responds that the request was cancelled, a resource
//   			// was exhausted, or if the service is unavailable
//   			GrpcRetryEvents: []grpcRetryEvent{
//   				appmesh.*grpcRetryEvent_CANCELLED,
//   				appmesh.*grpcRetryEvent_RESOURCE_EXHAUSTED,
//   				appmesh.*grpcRetryEvent_UNAVAILABLE,
//   			},
//   			RetryAttempts: jsii.Number(5),
//   			RetryTimeout: cdk.Duration_Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type GrpcRetryEvent string

const (
	// Request was cancelled.
	// See: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	//
	GrpcRetryEvent_CANCELLED GrpcRetryEvent = "CANCELLED"
	// The deadline was exceeded.
	// See: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	//
	GrpcRetryEvent_DEADLINE_EXCEEDED GrpcRetryEvent = "DEADLINE_EXCEEDED"
	// Internal error.
	// See: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	//
	GrpcRetryEvent_INTERNAL_ERROR GrpcRetryEvent = "INTERNAL_ERROR"
	// A resource was exhausted.
	// See: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	//
	GrpcRetryEvent_RESOURCE_EXHAUSTED GrpcRetryEvent = "RESOURCE_EXHAUSTED"
	// The service is unavailable.
	// See: https://grpc.github.io/grpc/core/md_doc_statuscodes.html
	//
	GrpcRetryEvent_UNAVAILABLE GrpcRetryEvent = "UNAVAILABLE"
)

