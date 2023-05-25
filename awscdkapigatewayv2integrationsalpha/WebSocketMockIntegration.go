package awscdkapigatewayv2integrationsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha/v2"
	"github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha/v2/internal"
)

// Mock WebSocket Integration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_integrations_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2integrationsalpha"
//
//   webSocketMockIntegration := apigatewayv2_integrations_alpha.NewWebSocketMockIntegration(jsii.String("id"))
//
// Experimental.
type WebSocketMockIntegration interface {
	awscdkapigatewayv2alpha.WebSocketRouteIntegration
	// Bind this integration to the route.
	// Experimental.
	Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig
}

// The jsii proxy struct for WebSocketMockIntegration
type jsiiProxy_WebSocketMockIntegration struct {
	internal.Type__awscdkapigatewayv2alphaWebSocketRouteIntegration
}

// Experimental.
func NewWebSocketMockIntegration(id *string) WebSocketMockIntegration {
	_init_.Initialize()

	if err := validateNewWebSocketMockIntegrationParameters(id); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebSocketMockIntegration{}

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketMockIntegration",
		[]interface{}{id},
		&j,
	)

	return &j
}

// Experimental.
func NewWebSocketMockIntegration_Override(w WebSocketMockIntegration, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-apigatewayv2-integrations-alpha.WebSocketMockIntegration",
		[]interface{}{id},
		w,
	)
}

func (w *jsiiProxy_WebSocketMockIntegration) Bind(options *awscdkapigatewayv2alpha.WebSocketRouteIntegrationBindOptions) *awscdkapigatewayv2alpha.WebSocketRouteIntegrationConfig {
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

