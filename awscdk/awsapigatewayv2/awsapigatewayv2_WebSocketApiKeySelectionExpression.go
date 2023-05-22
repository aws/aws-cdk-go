package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the currently available API Key Selection Expressions.
//
// Example:
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"), &WebSocketApiProps{
//   	ApiKeySelectionExpression: apigwv2.WebSocketApiKeySelectionExpression_HEADER_X_API_KEY(),
//   })
//
// Experimental.
type WebSocketApiKeySelectionExpression interface {
	// The expression used by API Gateway.
	// Experimental.
	CustomApiKeySelector() *string
}

// The jsii proxy struct for WebSocketApiKeySelectionExpression
type jsiiProxy_WebSocketApiKeySelectionExpression struct {
	_ byte // padding
}

func (j *jsiiProxy_WebSocketApiKeySelectionExpression) CustomApiKeySelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customApiKeySelector",
		&returns,
	)
	return returns
}


// Experimental.
func NewWebSocketApiKeySelectionExpression(customApiKeySelector *string) WebSocketApiKeySelectionExpression {
	_init_.Initialize()

	if err := validateNewWebSocketApiKeySelectionExpressionParameters(customApiKeySelector); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketApiKeySelectionExpression{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		[]interface{}{customApiKeySelector},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketApiKeySelectionExpression_Override(w WebSocketApiKeySelectionExpression, customApiKeySelector *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		[]interface{}{customApiKeySelector},
		w,
	)
}

func WebSocketApiKeySelectionExpression_AUTHORIZER_USAGE_IDENTIFIER_KEY() WebSocketApiKeySelectionExpression {
	_init_.Initialize()
	var returns WebSocketApiKeySelectionExpression
	_jsii_.StaticGet(
		"monocdk.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		"AUTHORIZER_USAGE_IDENTIFIER_KEY",
		&returns,
	)
	return returns
}

func WebSocketApiKeySelectionExpression_HEADER_X_API_KEY() WebSocketApiKeySelectionExpression {
	_init_.Initialize()
	var returns WebSocketApiKeySelectionExpression
	_jsii_.StaticGet(
		"monocdk.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		"HEADER_X_API_KEY",
		&returns,
	)
	return returns
}

