// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Input to the bind() operation, that binds an authorizer to a route.
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
//   httpRouteAuthorizerBindOptions := &HttpRouteAuthorizerBindOptions{
//   	Route: httpRoute,
//   	Scope: construct,
//   }
//
// Experimental.
type HttpRouteAuthorizerBindOptions struct {
	// The route to which the authorizer is being bound.
	// Experimental.
	Route IHttpRoute `field:"required" json:"route" yaml:"route"`
	// The scope for any constructs created as part of the bind.
	// Experimental.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}

