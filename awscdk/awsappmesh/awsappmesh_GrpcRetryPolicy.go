package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// gRPC retry policy.
//
// Example:
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
// Experimental.
type GrpcRetryPolicy struct {
	// The maximum number of retry attempts.
	// Experimental.
	RetryAttempts *float64 `field:"required" json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout for each retry attempt.
	// Experimental.
	RetryTimeout awscdk.Duration `field:"required" json:"retryTimeout" yaml:"retryTimeout"`
	// Specify HTTP events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Experimental.
	HttpRetryEvents *[]HttpRetryEvent `field:"optional" json:"httpRetryEvents" yaml:"httpRetryEvents"`
	// TCP events on which to retry.
	//
	// The event occurs before any processing of a
	// request has started and is encountered when the upstream is temporarily or
	// permanently unavailable. You must specify at least one value for at least
	// one types of retry events.
	// Experimental.
	TcpRetryEvents *[]TcpRetryEvent `field:"optional" json:"tcpRetryEvents" yaml:"tcpRetryEvents"`
	// gRPC events on which to retry.
	//
	// You must specify at least one value
	// for at least one types of retry events.
	// Experimental.
	GrpcRetryEvents *[]GrpcRetryEvent `field:"optional" json:"grpcRetryEvents" yaml:"grpcRetryEvents"`
}

