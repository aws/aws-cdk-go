package awscdkapigatewayv2alpha

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to the HttpRouteIntegration during its bind operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var httpRoute httpRoute
//
//   httpRouteIntegrationBindOptions := &HttpRouteIntegrationBindOptions{
//   	Route: httpRoute,
//   	Scope: construct,
//   }
//
// Experimental.
type HttpRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	// Experimental.
	Route IHttpRoute `field:"required" json:"route" yaml:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `HttpRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	// Experimental.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}

