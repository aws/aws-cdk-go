package awsapigatewayv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents the currently available API Key Selection Expressions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketApiKeySelectionExpression := awscdk.Aws_apigatewayv2.NewWebSocketApiKeySelectionExpression(jsii.String("customApiKeySelector"))
//
type WebSocketApiKeySelectionExpression interface {
	// The expression used by API Gateway.
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


func NewWebSocketApiKeySelectionExpression(customApiKeySelector *string) WebSocketApiKeySelectionExpression {
	_init_.Initialize()

	if err := validateNewWebSocketApiKeySelectionExpressionParameters(customApiKeySelector); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketApiKeySelectionExpression{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		[]interface{}{customApiKeySelector},
		&j,
	)

	return &j
}

func NewWebSocketApiKeySelectionExpression_Override(w WebSocketApiKeySelectionExpression, customApiKeySelector *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		[]interface{}{customApiKeySelector},
		w,
	)
}

func WebSocketApiKeySelectionExpression_AUTHORIZER_USAGE_IDENTIFIER_KEY() WebSocketApiKeySelectionExpression {
	_init_.Initialize()
	var returns WebSocketApiKeySelectionExpression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		"AUTHORIZER_USAGE_IDENTIFIER_KEY",
		&returns,
	)
	return returns
}

func WebSocketApiKeySelectionExpression_HEADER_X_API_KEY() WebSocketApiKeySelectionExpression {
	_init_.Initialize()
	var returns WebSocketApiKeySelectionExpression
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_apigatewayv2.WebSocketApiKeySelectionExpression",
		"HEADER_X_API_KEY",
		&returns,
	)
	return returns
}

