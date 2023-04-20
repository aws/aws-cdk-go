package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// gRPC retry policy.
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
//   			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type GrpcRetryPolicy struct {
	// The maximum number of retry attempts.
	RetryAttempts *float64 `field:"required" json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout for each retry attempt.
	RetryTimeout awscdk.Duration `field:"required" json:"retryTimeout" yaml:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	HttpRetryEvents *[]HttpRetryEvent `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	TcpRetryEvents *[]TcpRetryEvent `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
	// gRPC events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	GrpcRetryEvents *[]GrpcRetryEvent `field:"optional" json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
}

