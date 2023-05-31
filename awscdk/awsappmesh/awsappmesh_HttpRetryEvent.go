package awsappmesh


// HTTP events on which to retry.
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
//   			RetryTimeout: cdk.Duration_Seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
// Experimental.
type HttpRetryEvent string

const (
	// HTTP status codes 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, and 511.
	// Experimental.
	HttpRetryEvent_SERVER_ERROR HttpRetryEvent = "SERVER_ERROR"
	// HTTP status codes 502, 503, and 504.
	// Experimental.
	HttpRetryEvent_GATEWAY_ERROR HttpRetryEvent = "GATEWAY_ERROR"
	// HTTP status code 409.
	// Experimental.
	HttpRetryEvent_CLIENT_ERROR HttpRetryEvent = "CLIENT_ERROR"
	// Retry on refused stream.
	// Experimental.
	HttpRetryEvent_STREAM_ERROR HttpRetryEvent = "STREAM_ERROR"
)

