package awscdkapigatewayv2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the WebSocketStage.
// Experimental.
type IWebSocketStage interface {
	IStage
	// The API this stage is associated to.
	// Experimental.
	Api() IWebSocketApi
	// The callback URL to this stage.
	//
	// You can use the callback URL to send messages to the client from the backend system.
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-basic-concept.html
	// https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-how-to-call-websocket-api-connections.html
	// Experimental.
	CallbackUrl() *string
}

// The jsii proxy for IWebSocketStage
type jsiiProxy_IWebSocketStage struct {
	jsiiProxy_IStage
}

func (j *jsiiProxy_IWebSocketStage) Api() IWebSocketApi {
	var returns IWebSocketApi
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebSocketStage) CallbackUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callbackUrl",
		&returns,
	)
	return returns
}

