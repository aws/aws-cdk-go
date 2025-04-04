package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Explicitly configure no authorizers on specific WebSocket API routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketNoneAuthorizer := awscdk.Aws_apigatewayv2.NewWebSocketNoneAuthorizer()
//
type WebSocketNoneAuthorizer interface {
	IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	Bind(_options *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketNoneAuthorizer
type jsiiProxy_WebSocketNoneAuthorizer struct {
	jsiiProxy_IWebSocketRouteAuthorizer
}

func NewWebSocketNoneAuthorizer() WebSocketNoneAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketNoneAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewWebSocketNoneAuthorizer_Override(w WebSocketNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketNoneAuthorizer",
		nil, // no parameters
		w,
	)
}

func (w *jsiiProxy_WebSocketNoneAuthorizer) Bind(_options *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(_options); err != nil {
		panic(err)
	}
	var returns *WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{_options},
		&returns,
	)

	return returns
}

