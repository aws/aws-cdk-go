package interfacesawsec2


// A reference to a RouteServerPropagation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeServerPropagationReference := &RouteServerPropagationReference{
//   	RouteServerId: jsii.String("routeServerId"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
type RouteServerPropagationReference struct {
	// The RouteServerId of the RouteServerPropagation resource.
	RouteServerId *string `field:"required" json:"routeServerId" yaml:"routeServerId"`
	// The RouteTableId of the RouteServerPropagation resource.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

