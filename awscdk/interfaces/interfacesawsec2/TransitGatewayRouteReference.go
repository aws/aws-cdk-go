package interfacesawsec2


// A reference to a TransitGatewayRoute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   transitGatewayRouteReference := &TransitGatewayRouteReference{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	TransitGatewayRouteTableId: jsii.String("transitGatewayRouteTableId"),
//   }
//
type TransitGatewayRouteReference struct {
	// The DestinationCidrBlock of the TransitGatewayRoute resource.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The TransitGatewayRouteTableId of the TransitGatewayRoute resource.
	TransitGatewayRouteTableId *string `field:"required" json:"transitGatewayRouteTableId" yaml:"transitGatewayRouteTableId"`
}

