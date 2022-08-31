package awsapigatewayv2integrations

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations/internal"
)

// Mock WebSocket Integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   webSocketMockIntegration := awscdk.Aws_apigatewayv2_integrations.NewWebSocketMockIntegration(jsii.String("id"))
//
// Experimental.
type WebSocketMockIntegration interface {
	awsapigatewayv2.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketMockIntegration
type jsiiProxy_WebSocketMockIntegration struct {
	internal.Type__awsapigatewayv2WebSocketRouteIntegration
}

// Experimental.
func NewWebSocketMockIntegration(id *string) WebSocketMockIntegration {
	_init_.Initialize()

	j := jsiiProxy_WebSocketMockIntegration{}

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketMockIntegration_Override(w WebSocketMockIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigatewayv2_integrations.WebSocketMockIntegration",
		[]interface{}{id},
		w,
	)
}

func (w *jsiiProxy_WebSocketMockIntegration) Bind(options *awsapigatewayv2.WebSocketRouteIntegrationBindOptions) *awsapigatewayv2.WebSocketRouteIntegrationConfig {
	var returns *awsapigatewayv2.WebSocketRouteIntegrationConfig

	_jsii_.Invoke(
		w,
		"bind",
		[]interface{}{options},
		&returns,
	)

	return returns
}

