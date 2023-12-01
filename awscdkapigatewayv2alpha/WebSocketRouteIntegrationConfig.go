package awscdkapigatewayv2alpha


// Config returned back as a result of the bind.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   webSocketRouteIntegrationConfig := &WebSocketRouteIntegrationConfig{
//   	Type: apigatewayv2_alpha.WebSocketIntegrationType_AWS_PROXY,
//   	Uri: jsii.String("uri"),
//   }
//
// Deprecated.
type WebSocketRouteIntegrationConfig struct {
	// Integration type.
	// Deprecated.
	Type WebSocketIntegrationType `field:"required" json:"type" yaml:"type"`
	// Integration URI.
	// Deprecated.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

