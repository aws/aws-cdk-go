package awsec2


// A reference to a GatewayRouteTableAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   gatewayRouteTableAssociationReference := &GatewayRouteTableAssociationReference{
//   	GatewayId: jsii.String("gatewayId"),
//   }
//
type GatewayRouteTableAssociationReference struct {
	// The GatewayId of the GatewayRouteTableAssociation resource.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
}

