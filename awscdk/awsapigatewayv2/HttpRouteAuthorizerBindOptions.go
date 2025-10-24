package awsapigatewayv2

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Input to the bind() operation, that binds an authorizer to a route.
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
//   httpRouteAuthorizerBindOptions := &HttpRouteAuthorizerBindOptions{
//   	Route: httpRoute,
//   	Scope: construct,
//   }
//
type HttpRouteAuthorizerBindOptions struct {
	// The route to which the authorizer is being bound.
	Route IHttpRoute `field:"required" json:"route" yaml:"route"`
	// The scope for any constructs created as part of the bind.
	Scope constructs.Construct `field:"required" json:"scope" yaml:"scope"`
}

