package interfacesawsapigatewayv2


// A reference to a Route resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeReference := &RouteReference{
//   	ApiId: jsii.String("apiId"),
//   	RouteId: jsii.String("routeId"),
//   }
//
type RouteReference struct {
	// The ApiId of the Route resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The RouteId of the Route resource.
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
}

