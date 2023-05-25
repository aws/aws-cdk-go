package awscdkapigatewayv2authorizersalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha/v2/internal"
)

// Authorize WebSocket Api routes via a lambda function.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2authorizersalpha"
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   // This function handles your auth logic
//   var authHandler function
//
//   // This function handles your WebSocket requests
//   var handler function
//
//
//   authorizer := awscdkapigatewayv2authorizersalpha.NewWebSocketLambdaAuthorizer(jsii.String("Authorizer"), authHandler)
//
//   integration := awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("Integration"), handler)
//
//   apigwv2.NewWebSocketApi(this, jsii.String("WebSocketApi"), &WebSocketApiProps{
//   	ConnectRouteOptions: &WebSocketRouteOptions{
//   		Integration: *Integration,
//   		Authorizer: *Authorizer,
//   	},
//   })
//
// Experimental.
type WebSocketLambdaAuthorizer interface {
	awscdkapigatewayv2alpha.IWebSocketRouteAuthorizer
	// Bind this authorizer to a specified WebSocket route.
	// Experimental.
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig
}

// The jsii proxy struct for WebSocketLambdaAuthorizer
type jsiiProxy_WebSocketLambdaAuthorizer struct {
	internal.Type__awscdkapigatewayv2alphaIWebSocketRouteAuthorizer
}

// Experimental.
func NewWebSocketLambdaAuthorizer(id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) WebSocketLambdaAuthorizer {
	_init_.Initialize()

	if err := validateNewWebSocketLambdaAuthorizerParameters(id, handler, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketLambdaAuthorizer{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaAuthorizer_Override(w WebSocketLambdaAuthorizer, id *string, handler awslambda.IFunction, props *WebSocketLambdaAuthorizerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-authorizers-alpha.WebSocketLambdaAuthorizer",
		[]interface{}{id, handler, props},
		w,
	)
}

func (w *jsiiProxy_WebSocketLambdaAuthorizer) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.WebSocketRouteAuthorizerConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

