package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
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
// Experimental.
type WebSocketNoneAuthorizer interface {
	IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
	Bind(_arg *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketNoneAuthorizer
type jsiiProxy_WebSocketNoneAuthorizer struct {
	jsiiProxy_IWebSocketRouteAuthorizer
}

// Experimental.
func NewWebSocketNoneAuthorizer() WebSocketNoneAuthorizer {
	_init_.Initialize()

	j := jsiiProxy_WebSocketNoneAuthorizer{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.WebSocketNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketNoneAuthorizer_Override(w WebSocketNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.WebSocketNoneAuthorizer",
		nil, // no parameters
		w,
	)
}

func (w *jsiiProxy_WebSocketNoneAuthorizer) Bind(_arg *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(_arg); err != nil {
		panic(err)
	}
	var returns *WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{_arg},
		&returns,
	)

	return returns
}

