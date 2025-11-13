package interfacesawsrefactorspaces


// A reference to a Route resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeReference := &RouteReference{
//   	ApplicationIdentifier: jsii.String("applicationIdentifier"),
//   	EnvironmentIdentifier: jsii.String("environmentIdentifier"),
//   	RouteArn: jsii.String("routeArn"),
//   	RouteIdentifier: jsii.String("routeIdentifier"),
//   }
//
type RouteReference struct {
	// The ApplicationIdentifier of the Route resource.
	ApplicationIdentifier *string `field:"required" json:"applicationIdentifier" yaml:"applicationIdentifier"`
	// The EnvironmentIdentifier of the Route resource.
	EnvironmentIdentifier *string `field:"required" json:"environmentIdentifier" yaml:"environmentIdentifier"`
	// The ARN of the Route resource.
	RouteArn *string `field:"required" json:"routeArn" yaml:"routeArn"`
	// The RouteIdentifier of the Route resource.
	RouteIdentifier *string `field:"required" json:"routeIdentifier" yaml:"routeIdentifier"`
}

