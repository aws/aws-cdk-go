package awsapigatewayv2


// A reference to a RouteResponse resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeResponseReference := &RouteResponseReference{
//   	ApiId: jsii.String("apiId"),
//   	RouteId: jsii.String("routeId"),
//   	RouteResponseId: jsii.String("routeResponseId"),
//   }
//
type RouteResponseReference struct {
	// The ApiId of the RouteResponse resource.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The RouteId of the RouteResponse resource.
	RouteId *string `field:"required" json:"routeId" yaml:"routeId"`
	// The RouteResponseId of the RouteResponse resource.
	RouteResponseId *string `field:"required" json:"routeResponseId" yaml:"routeResponseId"`
}

