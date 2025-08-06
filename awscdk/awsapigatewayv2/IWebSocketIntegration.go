package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents an Integration for an WebSocket API.
type IWebSocketIntegration interface {
	IIntegration
	// The WebSocket API associated with this integration.
	WebSocketApi() IWebSocketApi
}

// The jsii proxy for IWebSocketIntegration
type jsiiProxy_IWebSocketIntegration struct {
	jsiiProxy_IIntegration
}

func (j *jsiiProxy_IWebSocketIntegration) WebSocketApi() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"webSocketApi",
		&returns,
	)
	return returns
}

