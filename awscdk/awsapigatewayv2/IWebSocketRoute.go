package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Route for an WebSocket API.
type IWebSocketRoute interface {
	IRoute
	// The key to this route.
	RouteKey() *string
	// The WebSocket API associated with this route.
	WebSocketApi() IWebSocketApi
}

// The jsii proxy for IWebSocketRoute
type jsiiProxy_IWebSocketRoute struct {
	jsiiProxy_IRoute
}

func (j *jsiiProxy_IWebSocketRoute) RouteKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebSocketRoute) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}

