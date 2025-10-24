package awsapigatewayv2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Options to the HttpRouteIntegration during its bind operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct Construct
//   var httpRoute HttpRoute
//
//   httpRouteIntegrationBindOptions := &HttpRouteIntegrationBindOptions{
//   	Route: httpRoute,
//   	Scope: construct,
//   }
//
type HttpRouteIntegrationBindOptions struct {
	// The route to which this is being bound.
	Route IHttpRoute `field:"required" json:"route" yaml:"route"`
	// The current scope in which the bind is occurring.
	//
	// If the `HttpRouteIntegration` being bound creates additional constructs,
	// this will be used as their parent scope.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}

