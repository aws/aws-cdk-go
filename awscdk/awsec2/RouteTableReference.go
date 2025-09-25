package awsec2


// A reference to a RouteTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routeTableReference := &RouteTableReference{
//   	RouteTableId: jsii.String("routeTableId"),
//   }
//
type RouteTableReference struct {
	// The RouteTableId of the RouteTable resource.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

