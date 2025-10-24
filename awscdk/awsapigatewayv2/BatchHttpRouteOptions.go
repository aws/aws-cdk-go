package awsapigatewayv2


// Options used when configuring multiple routes, at once.
//
// The options here are the ones that would be configured for all being set up.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var httpRouteIntegration HttpRouteIntegration
//
//   batchHttpRouteOptions := &BatchHttpRouteOptions{
//   	Integration: httpRouteIntegration,
//   }
//
type BatchHttpRouteOptions struct {
	// The integration to be configured on this route.
	Integration HttpRouteIntegration `field:"required" json:"integration" yaml:"integration"`
}

