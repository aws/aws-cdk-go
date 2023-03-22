// The CDK Construct Library for AWS::APIGatewayv2
package awscdkapigatewayv2alpha


// Options used when configuring multiple routes, at once.
//
// The options here are the ones that would be configured for all being set up.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import apigatewayv2_alpha "github.com/aws/aws-cdk-go/awscdkapigatewayv2alpha"
//
//   var httpRouteIntegration httpRouteIntegration
//
//   batchHttpRouteOptions := &BatchHttpRouteOptions{
//   	Integration: httpRouteIntegration,
//   }
//
// Experimental.
type BatchHttpRouteOptions struct {
	// The integration to be configured on this route.
	// Experimental.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
}

