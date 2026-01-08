package interfacesawsappmesh


// A reference to a Route resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeReference := &RouteReference{
//   	RouteArn: jsii.String("routeArn"),
//   }
//
type RouteReference struct {
	// The Arn of the Route resource.
	RouteArn *string `field:"required" json:"routeArn" yaml:"routeArn"`
}

