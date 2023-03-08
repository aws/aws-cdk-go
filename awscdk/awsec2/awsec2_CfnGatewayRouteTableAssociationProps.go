package awsec2


// Properties for defining a `CfnGatewayRouteTableAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnGatewayRouteTableAssociationProps := &cfnGatewayRouteTableAssociationProps{
//   	gatewayId: jsii.String("gatewayId"),
//   	routeTableId: jsii.String("routeTableId"),
//   }
//
type CfnGatewayRouteTableAssociationProps struct {
	// The ID of the gateway.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// The ID of the route table.
	RouteTableId *string `field:"required" json:"routeTableId" yaml:"routeTableId"`
}

