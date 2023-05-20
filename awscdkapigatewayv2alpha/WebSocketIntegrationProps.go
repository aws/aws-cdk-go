package awscdkapigatewayv2alpha


// The integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var webSocketApi webSocketApi
//
//   webSocketIntegrationProps := &WebSocketIntegrationProps{
//   	IntegrationType: apigatewayv2_alpha.WebSocketIntegrationType_AWS_PROXY,
//   	IntegrationUri: jsii.String("integrationUri"),
//   	WebSocketApi: webSocketApi,
//   }
//
// Experimental.
type WebSocketIntegrationProps struct {
	// Integration type.
	// Experimental.
	IntegrationType WebSocketIntegrationType `field:"required" json:"integrationType" yaml:"integrationType"`
	// Integration URI.
	// Experimental.
	IntegrationUri *string `field:"required" json:"integrationUri" yaml:"integrationUri"`
	// The WebSocket API to which this integration should be bound.
	// Experimental.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
}

