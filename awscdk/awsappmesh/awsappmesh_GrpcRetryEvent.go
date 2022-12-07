package awsappmesh


// gRPC events.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-grpc-retry"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.grpc(&grpcRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		match: &grpcRouteMatch{
//   			serviceName: jsii.String("servicename"),
//   		},
//   		retryPolicy: &grpcRetryPolicy{
//   			tcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.tcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			httpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry if gRPC responds that the request was cancelled, a resource
//   			// was exhausted, or if the service is unavailable
//   			grpcRetryEvents: []grpcRetryEvent{
//   				appmesh.*grpcRetryEvent_CANCELLED,
//   				appmesh.*grpcRetryEvent_RESOURCE_EXHAUSTED,
//   				appmesh.*grpcRetryEvent_UNAVAILABLE,
//   			},
//   			retryAttempts: jsii.Number(5),
//   			retryTimeout: cdk.duration.seconds(jsii.Number(1)),
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

