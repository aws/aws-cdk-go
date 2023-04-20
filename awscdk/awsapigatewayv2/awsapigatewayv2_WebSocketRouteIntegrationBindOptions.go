package awsapigatewayv2

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Options to the WebSocketRouteIntegration during its bind operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import monocdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var construct construct
//   var webSocketRoute webSocketRoute
//
//   webSocketRouteIntegrationBindOptions := &WebSocketRouteIntegrationBindOptions{
//   	Route: webSocketRoute,
//   	Scope: construct,
//   }
//
// Experimental.
type WebSocketRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	// Experimental.
	Route IWebSocketRoute `field:"required" json:"route" yaml:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `WebSocketRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	// Experimental.
	Scope awscdk.Construct `field:"required" json:"scope" yaml:"scope"`
}

