package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations/internal"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
)

// Lambda WebSocket Integration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
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
//   	Integration: awscdk.NewWebSocketLambdaIntegration(jsii.String("SendMessageIntegration"), messageHandler),
//   })
//
// Experimental.
type WebSocketLambdaIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketLambdaIntegration
type jsiiProxy_WebSocketLambdaIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

// Experimental.
func NewWebSocketLambdaIntegration(id *string, handler awslambda.IFunction) WebSocketLambdaIntegration {
	_init_.Initialize()

	if err := validateNewWebSocketLambdaIntegrationParameters(id, handler); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketLambdaIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketLambdaIntegration_Override(w WebSocketLambdaIntegration, id *string, handler awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketLambdaIntegration",
		[]interface{}{id, handler},
		w,
	)
}

func (w *jsiiProxy_WebSocketLambdaIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	if err := w.validateBindParameters(options); err != nil {
		panic(err)
	}
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

