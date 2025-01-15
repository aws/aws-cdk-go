package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations/internal"
)

// Mock WebSocket Integration.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   webSocketApi := apigwv2.NewWebSocketApi(this, jsii.String("mywsapi"))
//   apigwv2.NewWebSocketStage(this, jsii.String("mystage"), &WebSocketStageProps{
//   	WebSocketApi: WebSocketApi,
//   	StageName: jsii.String("dev"),
//   	AutoDeploy: jsii.Boolean(true),
//   })
//
//   webSocketApi.AddRoute(jsii.String("sendMessage"), &WebSocketRouteOptions{
//   	Integration: awscdk.NewWebSocketMockIntegration(jsii.String("DefaultIntegration"), &WebSocketMockIntegrationProps{
//   		RequestTemplates: map[string]*string{
//   			"application/json": JSON.stringify(map[string]*f64{
//   				"statusCode": jsii.Number(200),
//   			}),
//   		},
//   		TemplateSelectionExpression: jsii.String("\\$default"),
//   	}),
//   	ReturnResponse: jsii.Boolean(true),
//   })
//
type WebSocketMockIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketMockIntegration
type jsiiProxy_WebSocketMockIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

func NewWebSocketMockIntegration(id *string, props *WebSocketMockIntegrationProps) WebSocketMockIntegration {
	_init_.Initialize()

	if err := validateNewWebSocketMockIntegrationParameters(id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketMockIntegration{}

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id, props},
		&j,
	)

	return &j
}

func NewWebSocketMockIntegration_Override(w WebSocketMockIntegration, id *string, props *WebSocketMockIntegrationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id, props},
		w,
	)
}

func (w *jsiiProxy_WebSocketMockIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
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

