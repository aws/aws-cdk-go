package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// gRPC retry policy.
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
type GrpcRetryPolicy struct {
	// The maximum number of retry attempts.
	RetryAttempts *float64 `field:"required" json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout for each retry attempt.
	RetryTimeout awscdk.Duration `field:"required" json:"retryTimeout" yaml:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Default: - no retries for http events.
	//
	HttpRetryEvents *[]HttpRetryEvent `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	// Default: - no retries for tcp events.
	//
	TcpRetryEvents *[]TcpRetryEvent `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
	// gRPC events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Default: - no retries for gRPC events.
	//
	GrpcRetryEvents *[]GrpcRetryEvent `field:"optional" json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
}

