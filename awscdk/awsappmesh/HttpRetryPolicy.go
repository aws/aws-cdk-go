package awsappmesh

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// HTTP retry policy.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2-retry"), &RouteBaseProps{
//   	RouteSpec: appmesh.RouteSpec_Http2(&HttpRouteSpecOptions{
//   		WeightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				VirtualNode: node,
//   			},
//   		},
//   		RetryPolicy: &HttpRetryPolicy{
//   			// Retry if the connection failed
//   			TcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.TcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			// Retry if HTTP responds with a gateway error (502, 503, 504)
//   			HttpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry five times
//   			RetryAttempts: jsii.Number(5),
//   			// Use a 1 second timeout per retry
//   			RetryTimeout: awscdk.Duration_Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
type HttpRetryPolicy struct {
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
}

