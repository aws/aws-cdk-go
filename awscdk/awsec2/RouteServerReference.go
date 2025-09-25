package awsec2


// A reference to a RouteServer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeServerReference := &RouteServerReference{
//   	RouteServerArn: jsii.String("routeServerArn"),
//   	RouteServerId: jsii.String("routeServerId"),
//   }
//
type RouteServerReference struct {
	// The ARN of the RouteServer resource.
	RouteServerArn *string `field:"required" json:"routeServerArn" yaml:"routeServerArn"`
	// The Id of the RouteServer resource.
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
}

