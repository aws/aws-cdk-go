package awsapigatewayv2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to the WebSocketRouteIntegration during its bind operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct Construct
//   var webSocketRoute WebSocketRoute
//
//   webSocketRouteIntegrationBindOptions := &WebSocketRouteIntegrationBindOptions{
//   	Route: webSocketRoute,
//   	Scope: construct,
//   }
//
type WebSocketRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	Route IWebSocketRoute `field:"required" json:"route" yaml:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `WebSocketRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}

