package awsappmesh

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
)

// Used to generate specs with different protocols for a RouteSpec.
//
// Example:
//   var router virtualRouter
//   var node virtualNode
//
//
//   router.addRoute(jsii.String("route-http2-retry"), &routeBaseProps{
//   	routeSpec: appmesh.routeSpec.http2(&httpRouteSpecOptions{
//   		weightedTargets: []weightedTarget{
//   			&weightedTarget{
//   				virtualNode: node,
//   			},
//   		},
//   		retryPolicy: &httpRetryPolicy{
//   			// Retry if the connection failed
//   			tcpRetryEvents: []cONNECTION_ERROR{
//   				appmesh.tcpRetryEvent_*cONNECTION_ERROR,
//   			},
//   			// Retry if HTTP responds with a gateway error (502, 503, 504)
//   			httpRetryEvents: []httpRetryEvent{
//   				appmesh.*httpRetryEvent_GATEWAY_ERROR,
//   			},
//   			// Retry five times
//   			retryAttempts: jsii.Number(5),
//   			// Use a 1 second timeout per retry
//   			retryTimeout: cdk.duration.seconds(jsii.Number(1)),
//   		},
//   	}),
//   })
//
// Experimental.
type RouteSpec interface {
	// Called when the RouteSpec type is initialized.
	//
	// Can be used to enforce
	// mutual exclusivity with future properties.
	// Experimental.
	Bind(scope awscdk.Construct) *RouteSpecConfig
}

// The jsii proxy struct for RouteSpec
type jsiiProxy_RouteSpec struct {
	_ byte // padding
}

// Experimental.
func NewRouteSpec_Override(r RouteSpec) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_appmesh.RouteSpec",
		nil, // no parameters
		r,
	)
}

// Creates a GRPC Based RouteSpec.
// Experimental.
func RouteSpec_Grpc(options *GrpcRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"grpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP Based RouteSpec.
// Experimental.
func RouteSpec_Http(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"http",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates an HTTP2 Based RouteSpec.
// Experimental.
func RouteSpec_Http2(options *HttpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"http2",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Creates a TCP Based RouteSpec.
// Experimental.
func RouteSpec_Tcp(options *TcpRouteSpecOptions) RouteSpec {
	_init_.Initialize()

	var returns RouteSpec

	_jsii_.StaticInvoke(
		"monocdk.aws_appmesh.RouteSpec",
		"tcp",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteSpec) Bind(scope awscdk.Construct) *RouteSpecConfig {
	var returns *RouteSpecConfig

	_jsii_.Invoke(
		r,
		"bind",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

