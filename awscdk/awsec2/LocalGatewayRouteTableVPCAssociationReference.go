package awsec2


// A reference to a LocalGatewayRouteTableVPCAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayRouteTableVPCAssociationReference := &LocalGatewayRouteTableVPCAssociationReference{
//   	LocalGatewayRouteTableVpcAssociationId: jsii.String("localGatewayRouteTableVpcAssociationId"),
//   }
//
type LocalGatewayRouteTableVPCAssociationReference struct {
	// The LocalGatewayRouteTableVpcAssociationId of the LocalGatewayRouteTableVPCAssociation resource.
	LocalGatewayRouteTableVpcAssociationId *string `field:"required" json:"localGatewayRouteTableVpcAssociationId" yaml:"localGatewayRouteTableVpcAssociationId"`
}

