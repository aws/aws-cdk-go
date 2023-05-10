// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Explicitly configure no authorizers on specific WebSocket API routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   webSocketNoneAuthorizer := apigatewayv2_alpha.NewWebSocketNoneAuthorizer()
//
// Experimental.
type WebSocketNoneAuthorizer interface {
	IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
	Bind(_options *WebSocketRouteAuthorizerBindOptions) *WebSocketRouteAuthorizerConfig
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
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketNoneAuthorizer",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketNoneAuthorizer_Override(w WebSocketNoneAuthorizer) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketNoneAuthorizer",
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

