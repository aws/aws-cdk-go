// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The interface that various route integration classes will inherit.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   var messageHandler function
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//   webSocketApi.AddRoute(jsii.String("sendmessage"), &WebSocketRouteOptions{
//   	Integration: awscdkapigatewayv2integrationsalpha.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketRouteIntegration interface {
	// Bind this integration to the route.
	// Experimental.
	Bind(options *WebSocketRouteIntegrationBindOptions) *WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketRouteIntegration
type jsiiProxy_WebSocketRouteIntegration struct {
	_ byte // padding
}

// Initialize an integration for a route on websocket api.
// Experimental.
func NewWebSocketRouteIntegration_Override(w WebSocketRouteIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-alpha.WebSocketRouteIntegration",
		[]interface{}{id},
		w,
	)
}

func (w *jsiiProxy_WebSocketRouteIntegration) Bind(options *WebSocketRouteIntegrationBindOptions) *WebSocketRouteIntegrationConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

