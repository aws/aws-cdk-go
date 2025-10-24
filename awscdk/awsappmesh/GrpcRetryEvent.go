package awsappmesh


// gRPC events.
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
//   			ServiceName: jsii.String("servicename"),
//   		},
//   		RetryPolicy: &GrpcRetryPolicy{
//   			TcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			HttpRetryEvents: []HttpRetryEvent{
//   				appmesh.HttpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry if gRPC responds that the request was cancelled, a resource
//   			// was exhausted, or if the service is unavailable
//   			GrpcRetryEvents: []GrpcRetryEvent{
//   				appmesh.GrpcRetryEvent_CANCELLED,
//   				appmesh.GrpcRetryEvent_RESOURCE_EXHAUSTED,
//   				appmesh.GrpcRetryEvent_UNAVAILABLE,
//   			},
//   			RetryAttempts: jsii.Number(5),
//   			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
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

