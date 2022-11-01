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
//   webSocketIntegrationProps := &webSocketIntegrationProps{
//   	integrationType: awscdk.Aws_apigatewayv2.webSocketIntegrationType_AWS_PROXY,
//   	integrationUri: jsii.String("integrationUri"),
//   	webSocketApi: webSocketApi,
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

