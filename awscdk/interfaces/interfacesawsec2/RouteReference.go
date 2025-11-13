package interfacesawsec2


// A reference to a Route resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeReference := &RouteReference{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
type RouteReference struct {
	// The CidrBlock of the Route resource.
	CidrBlock *string `field:"required" json:"cidrBlock" yaml:"cidrBlock"`
	// The RouteTableId of the Route resource.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

