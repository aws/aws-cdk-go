package awsec2


// A reference to a RouteServerEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeServerEndpointReference := &RouteServerEndpointReference{
//   	RouteServerEndpointArn: jsii.String("routeServerEndpointArn"),
//   	RouteServerEndpointId: jsii.String("routeServerEndpointId"),
//   }
//
type RouteServerEndpointReference struct {
	// The ARN of the RouteServerEndpoint resource.
	RouteServerEndpointArn *string `field:"required" json:"routeServerEndpointArn" yaml:"routeServerEndpointArn"`
	// The Id of the RouteServerEndpoint resource.
	RouteServerEndpointId *string `field:"required" json:"routeServerEndpointId" yaml:"routeServerEndpointId"`
}

