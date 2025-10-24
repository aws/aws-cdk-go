package awsapigatewayv2authorizers

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2authorizers/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Authorize WebSocket Api routes via a lambda function.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   // This function handles your auth logic
//   var authHandler Function
//
//   // This function handles your WebSocket requests
//   var handler Function
//
//
//   authorizer := awscdk.NewWebSocketLambdaAuthorizer(jsii.String("Authorizer"), authHandler)
//
//   integration := awscdk.NewWebSocketLambdaIntegration(jsii.String("Integration"), handler)
//
//   apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &WebSocketApiProps{
//   	ConnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: *Integration,
//   		Authorizer: *Authorizer,
//   	},
//   })
//
type WebSocketLambdaAuthorizer interface {
	awsapigatewayv2.IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketLambdaAuthorizer
type jsiiProxy_WebSocketLambdaAuthorizer struct {
	internal.Type__awsapigatewayv2IWebSocketRouteAuthorizer
}

func NewWebSocketLambdaAuthorizer(id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) WebSocketLambdaAuthorizer {
	_init_.Initialize()

	if err := validateNewWebSocketLambdaAuthorizerParameters(id, handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketLambdaAuthorizer{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

func NewWebSocketLambdaAuthorizer_Override(w WebSocketLambdaAuthorizer, id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_authorizers.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		w,
	)
}

func (w *jsiiProxy_WebSocketLambdaAuthorizer) Bind(options *awsapigatewayv2.WebSocketRouteAuthorizerBindOptions) *awsapigatewayv2.WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

