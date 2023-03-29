// Integrations for AWS APIGateway V2
package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// Lambda WebSocket Integration.
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
type WebSocketLambdaIntegration interface {
	awscdkapigatewayv2alpha.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketLambdaIntegration
type jsiiProxy_WebSocketLambdaIntegration struct {
	internal.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration
}

// Experimental.
func NewWebSocketLambdaIntegration(id *string, handler awslambda.IFunction) WebSocketLambdaIntegration {
	_init_.Initialize()

	if err := validateNewWebSocketLambdaIntegrationParameters(id, handler); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketLambdaIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaIntegration_Override(w WebSocketLambdaIntegration, id *string, handler awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		w,
	)
}

func (w *jsiiProxy_WebSocketLambdaIntegration) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

