package awsec2


// A reference to a TransitGatewayRouteTable resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteTableReference := &TransitGatewayRouteTableReference{
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type TransitGatewayRouteTableReference struct {
	// The TransitGatewayRouteTableId of the TransitGatewayRouteTable resource.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

