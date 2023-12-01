package awsapigatewayv2


// The integration properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var webSocketApi webSocketApi
//
//   webSocketIntegrationProps := &WebSocketIntegrationProps{
//   	IntegrationType: awscdk.Aws_apigatewayv2.WebSocketIntegrationType_AWS_PROXY,
//   	IntegrationUri: jsii.String("integrationUri"),
//   	WebSocketApi: webSocketApi,
//   }
//
type WebSocketIntegrationProps struct {
	// Integration type.
	IntegrationType WebSocketIntegrationType `field:"required" json:"integrationType" yaml:"integrationType"`
	// Integration URI.
	IntegrationUri *string `field:"required" json:"integrationUri" yaml:"integrationUri"`
	// The WebSocket API to which this integration should be bound.
	WebSocketApi IWebSocketApi `field:"required" json:"webSocketApi" yaml:"webSocketApi"`
}

