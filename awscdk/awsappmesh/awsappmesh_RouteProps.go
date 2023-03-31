package awsappmesh


// Properties to define new Routes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var mesh mesh
//   var routeSpec routeSpec
//   var virtualRouter virtualRouter
//
//   routeProps := &routeProps{
//   	mesh: mesh,
//   	routeSpec: routeSpec,
//   	virtualRouter: virtualRouter,
//
//   	// the properties below are optional
//   	routeName: jsii.String("routeName"),
//   }
//
// Experimental.
type RouteProps struct {
	// Protocol specific spec.
	// Experimental.
	RouteSpec RouteSpec `field:"required" json:"routeSpec" yaml:"routeSpec"`
	// The name of the route.
	// Experimental.
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
	// The service mesh to define the route in.
	// Experimental.
	Mesh IMesh `field:"required" json:"mesh" yaml:"mesh"`
	// The VirtualRouter the Route belongs to.
	// Experimental.
	VirtualRouter IVirtualRouter `field:"required" json:"virtualRouter" yaml:"virtualRouter"`
}

