package awsec2


// A reference to a LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localGatewayRouteTableVirtualInterfaceGroupAssociationReference := &LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference{
//   	LocalGatewayRouteTableVirtualInterfaceGroupAssociationId: jsii.String("localGatewayRouteTableVirtualInterfaceGroupAssociationId"),
//   }
//
type LocalGatewayRouteTableVirtualInterfaceGroupAssociationReference struct {
	// The LocalGatewayRouteTableVirtualInterfaceGroupAssociationId of the LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
	LocalGatewayRouteTableVirtualInterfaceGroupAssociationId *string `field:"required" json:"localGatewayRouteTableVirtualInterfaceGroupAssociationId" yaml:"localGatewayRouteTableVirtualInterfaceGroupAssociationId"`
}

