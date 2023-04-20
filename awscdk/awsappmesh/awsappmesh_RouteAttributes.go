package awsappmesh


// Interface with properties ncecessary to import a reusable Route.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var virtualRouter virtualRouter
//
//   routeAttributes := &RouteAttributes{
//   	RouteName: jsii.String("routeName"),
//   	VirtualRouter: virtualRouter,
//   }
//
// Experimental.
type RouteAttributes struct {
	// The name of the Route.
	// Experimental.
	RouteName *string `field:"required" json:"routeName" yaml:"routeName"`
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter IVirtualRouter `field:"required" json:"virtualRouter" yaml:"virtualRouter"`
}

