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
//   routeAttributes := &routeAttributes{
//   	routeName: jsii.String("routeName"),
//   	virtualRouter: virtualRouter,
//   }
//
type RouteAttributes struct {
	// The name of the Route.
	RouteName *string `field:"required" json:"routeName" yaml:"routeName"`
	// The VirtualRouter the Route belongs to.
	VirtualRouter IVirtualRouter `field:"required" json:"virtualRouter" yaml:"virtualRouter"`
}

